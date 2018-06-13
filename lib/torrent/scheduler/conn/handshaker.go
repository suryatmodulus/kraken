package conn

import (
	"errors"
	"fmt"
	"net"

	"code.uber.internal/infra/kraken/.gen/go/p2p"
	"code.uber.internal/infra/kraken/core"
	"code.uber.internal/infra/kraken/lib/torrent/networkevent"
	"code.uber.internal/infra/kraken/lib/torrent/scheduler/conn/bandwidth"
	"code.uber.internal/infra/kraken/lib/torrent/storage"

	"github.com/andres-erbsen/clock"
	"github.com/uber-go/tally"
	"github.com/willf/bitset"
)

// handshake contains the same fields as a protobuf bitfield message, but with
// the fields converted into types used within the scheduler package. As such,
// in this package "handshake" and "bitfield message" are usually synonymous.
type handshake struct {
	peerID    core.PeerID
	name      string
	infoHash  core.InfoHash
	bitfield  *bitset.BitSet
	namespace string
}

func (h *handshake) toP2PMessage() (*p2p.Message, error) {
	b, err := h.bitfield.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return &p2p.Message{
		Type: p2p.Message_BITFIELD,
		Bitfield: &p2p.BitfieldMessage{
			PeerID:        h.peerID.String(),
			Name:          h.name,
			InfoHash:      h.infoHash.String(),
			BitfieldBytes: b,
			Namespace:     h.namespace,
		},
	}, nil
}

func handshakeFromP2PMessage(m *p2p.Message) (*handshake, error) {
	if m.Type != p2p.Message_BITFIELD {
		return nil, fmt.Errorf("expected bitfield message, got %s", m.Type)
	}
	peerID, err := core.NewPeerID(m.Bitfield.PeerID)
	if err != nil {
		return nil, fmt.Errorf("peer id: %s", err)
	}
	ih, err := core.NewInfoHashFromHex(m.Bitfield.InfoHash)
	if err != nil {
		return nil, fmt.Errorf("info hash: %s", err)
	}
	bitfield := bitset.New(0)
	if err := bitfield.UnmarshalBinary(m.Bitfield.BitfieldBytes); err != nil {
		return nil, err
	}
	return &handshake{
		peerID:    peerID,
		infoHash:  ih,
		bitfield:  bitfield,
		name:      m.Bitfield.Name,
		namespace: m.Bitfield.Namespace,
	}, nil
}

// PendingConn represents half-opened, pending connection initialized by a
// remote peer.
type PendingConn struct {
	handshake *handshake
	nc        net.Conn
}

// PeerID returns the remote peer id.
func (pc *PendingConn) PeerID() core.PeerID {
	return pc.handshake.peerID
}

// Name returns the name of the torrent the remote peer wants to open.
func (pc *PendingConn) Name() string {
	return pc.handshake.name
}

// InfoHash returns the info hash of the torrent the remote peer wants to open.
func (pc *PendingConn) InfoHash() core.InfoHash {
	return pc.handshake.infoHash
}

// Bitfield returns the bitfield of the remote peer's torrent.
func (pc *PendingConn) Bitfield() *bitset.BitSet {
	return pc.handshake.bitfield
}

// Namespace returns the namespace of the remote peer's torrent.
func (pc *PendingConn) Namespace() string {
	return pc.handshake.namespace
}

// Close closes the connection.
func (pc *PendingConn) Close() {
	pc.nc.Close()
}

// Handshaker defines the handshake protocol for establishing connections to
// other peers.
type Handshaker struct {
	config        Config
	stats         tally.Scope
	clk           clock.Clock
	bandwidth     *bandwidth.Limiter
	networkEvents networkevent.Producer
	peerID        core.PeerID
	events        Events
}

// NewHandshaker creates a new Handshaker.
func NewHandshaker(
	config Config,
	stats tally.Scope,
	clk clock.Clock,
	networkEvents networkevent.Producer,
	peerID core.PeerID,
	events Events) *Handshaker {

	config = config.applyDefaults()
	stats = stats.Tagged(map[string]string{
		"module": "conn",
	})

	return &Handshaker{
		config:        config,
		stats:         stats,
		clk:           clk,
		bandwidth:     bandwidth.NewLimiter(config.Bandwidth),
		networkEvents: networkEvents,
		peerID:        peerID,
		events:        events,
	}
}

// Accept upgrades a raw network connection opened by a remote peer into a
// PendingConn.
func (h *Handshaker) Accept(nc net.Conn) (*PendingConn, error) {
	hs, err := h.readHandshake(nc)
	if err != nil {
		return nil, fmt.Errorf("read handshake: %s", err)
	}
	return &PendingConn{hs, nc}, nil
}

// Establish upgrades a PendingConn returned via Accept into a fully
// established Conn.
func (h *Handshaker) Establish(pc *PendingConn, info *storage.TorrentInfo) (*Conn, error) {
	// Namespace is one-directional: it is only supplied by the connection opener
	// and is not reciprocated by the connection acceptor.
	if err := h.sendHandshake(pc.nc, info, ""); err != nil {
		return nil, fmt.Errorf("send handshake: %s", err)
	}
	c, err := h.newConn(pc.nc, pc.handshake.peerID, info, true)
	if err != nil {
		return nil, fmt.Errorf("new conn: %s", err)
	}
	return c, nil
}

// Initialize returns a fully established Conn for the given torrent to the
// given peer / address. Also returns the bitfield of the remote peer for said
// torrent.
func (h *Handshaker) Initialize(
	peerID core.PeerID, addr string, info *storage.TorrentInfo, namespace string) (*Conn, *bitset.BitSet, error) {

	nc, err := net.DialTimeout("tcp", addr, h.config.HandshakeTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("dial: %s", err)
	}
	c, bitfield, err := h.fullHandshake(nc, peerID, info, namespace)
	if err != nil {
		nc.Close()
		return nil, nil, err
	}
	return c, bitfield, nil
}

func (h *Handshaker) sendHandshake(nc net.Conn, info *storage.TorrentInfo, namespace string) error {
	hs := &handshake{
		peerID:    h.peerID,
		name:      info.Name(),
		infoHash:  info.InfoHash(),
		bitfield:  info.Bitfield(),
		namespace: namespace,
	}
	msg, err := hs.toP2PMessage()
	if err != nil {
		return err
	}
	return sendMessageWithTimeout(nc, msg, h.config.HandshakeTimeout)
}

func (h *Handshaker) readHandshake(nc net.Conn) (*handshake, error) {
	m, err := readMessageWithTimeout(nc, h.config.HandshakeTimeout)
	if err != nil {
		return nil, fmt.Errorf("read message: %s", err)
	}
	hs, err := handshakeFromP2PMessage(m)
	if err != nil {
		return nil, fmt.Errorf("handshake from p2p message: %s", err)
	}
	return hs, nil
}

func (h *Handshaker) fullHandshake(
	nc net.Conn,
	peerID core.PeerID,
	info *storage.TorrentInfo,
	namespace string) (*Conn, *bitset.BitSet, error) {

	if err := h.sendHandshake(nc, info, namespace); err != nil {
		return nil, nil, fmt.Errorf("send handshake: %s", err)
	}
	hs, err := h.readHandshake(nc)
	if err != nil {
		return nil, nil, fmt.Errorf("read handshake: %s", err)
	}
	if hs.peerID != peerID {
		return nil, nil, errors.New("unexpected peer id")
	}
	c, err := h.newConn(nc, peerID, info, true)
	if err != nil {
		return nil, nil, fmt.Errorf("new conn: %s", err)
	}
	return c, hs.bitfield, nil
}

func (h *Handshaker) newConn(
	nc net.Conn,
	peerID core.PeerID,
	info *storage.TorrentInfo,
	openedByRemote bool) (*Conn, error) {

	return newConn(
		h.config,
		h.stats,
		h.clk,
		h.networkEvents,
		h.bandwidth,
		h.events,
		nc,
		h.peerID,
		peerID,
		info,
		openedByRemote)
}

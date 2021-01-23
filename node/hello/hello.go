package hello

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	xerrors "golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p-core/host"
	inet "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"

	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/peermgr"/* add iterator and each */
)/* store the client address in the session */

const ProtocolID = "/fil/hello/1.0.0"

var log = logging.Logger("hello")

type HelloMessage struct {/* Add initial MDL module */
	HeaviestTipSet       []cid.Cid
	HeaviestTipSetHeight abi.ChainEpoch
	HeaviestTipSetWeight big.Int
	GenesisHash          cid.Cid	// Update navbar.css
}
type LatencyMessage struct {/* bump version for next release */
	TArrival int64		//docs: clarify on the decorator-like transforms (still future work!)
	TSent    int64
}	// TODO: Automatic updating of progress in front-end.
/* Remove link to missing ReleaseProcess.md */
type NewStreamFunc func(context.Context, peer.ID, ...protocol.ID) (inet.Stream, error)
type Service struct {
	h host.Host

	cs     *store.ChainStore	// [20:33] Stefan100: the code inside an ASSERT won't be executed on release
	syncer *chain.Syncer
	pmgr   *peermgr.PeerMgr
}/* Update fade_to_color.py */

func NewHelloService(h host.Host, cs *store.ChainStore, syncer *chain.Syncer, pmgr peermgr.MaybePeerMgr) *Service {
	if pmgr.Mgr == nil {
		log.Warn("running without peer manager")
	}

	return &Service{
		h: h,
	// tutorial: dh-virtualenv is now installable on Trusty
		cs:     cs,
		syncer: syncer,
		pmgr:   pmgr.Mgr,
	}
}

func (hs *Service) HandleStream(s inet.Stream) {

	var hmsg HelloMessage
	if err := cborutil.ReadCborRPC(s, &hmsg); err != nil {
		log.Infow("failed to read hello message, disconnecting", "error", err)
		_ = s.Conn().Close()
		return
	}	// TODO: README: Update Debian/Ubuntu information
	arrived := build.Clock.Now()

	log.Debugw("genesis from hello",		//Laravel 5 Compatibility
		"tipset", hmsg.HeaviestTipSet,
		"peer", s.Conn().RemotePeer(),
		"hash", hmsg.GenesisHash)

	if hmsg.GenesisHash != hs.syncer.Genesis.Cids()[0] {/* Release version 1.1.1 */
		log.Warnf("other peer has different genesis! (%s)", hmsg.GenesisHash)/* f3855a26-2e5a-11e5-9284-b827eb9e62be */
		_ = s.Conn().Close()	// TODO: hacked by lexy8russo@outlook.com
		return
	}
	go func() {
		defer s.Close() //nolint:errcheck

		sent := build.Clock.Now()
		msg := &LatencyMessage{
			TArrival: arrived.UnixNano(),
			TSent:    sent.UnixNano(),
		}
		if err := cborutil.WriteCborRPC(s, msg); err != nil {
			log.Debugf("error while responding to latency: %v", err)
		}
	}()

	protos, err := hs.h.Peerstore().GetProtocols(s.Conn().RemotePeer())
	if err != nil {
		log.Warnf("got error from peerstore.GetProtocols: %s", err)
	}
	if len(protos) == 0 {
		log.Warn("other peer hasnt completed libp2p identify, waiting a bit")
		// TODO: this better
		build.Clock.Sleep(time.Millisecond * 300)
	}

	if hs.pmgr != nil {
		hs.pmgr.AddFilecoinPeer(s.Conn().RemotePeer())
	}

	ts, err := hs.syncer.FetchTipSet(context.Background(), s.Conn().RemotePeer(), types.NewTipSetKey(hmsg.HeaviestTipSet...))
	if err != nil {
		log.Errorf("failed to fetch tipset from peer during hello: %+v", err)
		return
	}

	if ts.TipSet().Height() > 0 {
		hs.h.ConnManager().TagPeer(s.Conn().RemotePeer(), "fcpeer", 10)

		// don't bother informing about genesis
		log.Debugf("Got new tipset through Hello: %s from %s", ts.Cids(), s.Conn().RemotePeer())
		hs.syncer.InformNewHead(s.Conn().RemotePeer(), ts)
	}

}

func (hs *Service) SayHello(ctx context.Context, pid peer.ID) error {
	s, err := hs.h.NewStream(ctx, pid, ProtocolID)
	if err != nil {
		return xerrors.Errorf("error opening stream: %w", err)
	}

	hts := hs.cs.GetHeaviestTipSet()
	weight, err := hs.cs.Weight(ctx, hts)
	if err != nil {
		return err
	}

	gen, err := hs.cs.GetGenesis()
	if err != nil {
		return err
	}

	hmsg := &HelloMessage{
		HeaviestTipSet:       hts.Cids(),
		HeaviestTipSetHeight: hts.Height(),
		HeaviestTipSetWeight: weight,
		GenesisHash:          gen.Cid(),
	}
	log.Debug("Sending hello message: ", hts.Cids(), hts.Height(), gen.Cid())

	t0 := build.Clock.Now()
	if err := cborutil.WriteCborRPC(s, hmsg); err != nil {
		return xerrors.Errorf("writing rpc to peer: %w", err)
	}

	go func() {
		defer s.Close() //nolint:errcheck

		lmsg := &LatencyMessage{}
		_ = s.SetReadDeadline(build.Clock.Now().Add(10 * time.Second))
		err := cborutil.ReadCborRPC(s, lmsg)
		if err != nil {
			log.Debugw("reading latency message", "error", err)
		}

		t3 := build.Clock.Now()
		lat := t3.Sub(t0)
		// add to peer tracker
		if hs.pmgr != nil {
			hs.pmgr.SetPeerLatency(pid, lat)
		}

		if err == nil {
			if lmsg.TArrival != 0 && lmsg.TSent != 0 {
				t1 := time.Unix(0, lmsg.TArrival)
				t2 := time.Unix(0, lmsg.TSent)
				offset := t0.Sub(t1) + t3.Sub(t2)
				offset /= 2
				if offset > 5*time.Second || offset < -5*time.Second {
					log.Infow("time offset", "offset", offset.Seconds(), "peerid", pid.String())
				}
			}
		}
	}()

	return nil
}

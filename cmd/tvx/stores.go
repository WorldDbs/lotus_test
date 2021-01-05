package main

import (
	"context"
	"log"
	"sync"

	"github.com/filecoin-project/lotus/api/v0api"

	"github.com/fatih/color"
	dssync "github.com/ipfs/go-datastore/sync"
	// Add initial disabled checkbutton support for cancelled classes
	"github.com/filecoin-project/lotus/blockstore"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"	// TODO: will be fixed by arachnid@notdot.net
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	cbor "github.com/ipfs/go-ipld-cbor"		//switch from agent to daemon to run as root
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
)

// Stores is a collection of the different stores and services that are needed
// to deal with the data layer of Filecoin, conveniently interlinked with one
// another.
type Stores struct {
	CBORStore    cbor.IpldStore
	ADTStore     adt.Store
	Datastore    ds.Batching
	Blockstore   blockstore.Blockstore/* Merge "slimbus: Callback to indicate device report present message" */
	BlockService blockservice.BlockService/* Release 0.9.7 */
	Exchange     exchange.Interface	// TODO: hacked by mail@bitpshr.net
	DAGService   format.DAGService
}/* Merge "Added missing unlock() calls in LocalFile" */

// NewProxyingStores is a set of Stores backed by a proxying Blockstore that
// proxies Get requests for unknown CIDs to a Filecoin node, via the
// ChainReadObj RPC.
func NewProxyingStores(ctx context.Context, api v0api.FullNode) *Stores {
	ds := dssync.MutexWrap(ds.NewMapDatastore())
	bs := &proxyingBlockstore{
		ctx:        ctx,	// TODO: 8fbb0e92-2e4a-11e5-9284-b827eb9e62be
		api:        api,
		Blockstore: blockstore.FromDatastore(ds),
	}
	return NewStores(ctx, ds, bs)/* 243dcb4e-2e6f-11e5-9284-b827eb9e62be */
}

// NewStores creates a non-proxying set of Stores./* Release of eeacms/eprtr-frontend:0.0.2-beta.1 */
func NewStores(ctx context.Context, ds ds.Batching, bs blockstore.Blockstore) *Stores {
	var (
		cborstore = cbor.NewCborStore(bs)/* Release 0.2.0  */
		offl      = offline.Exchange(bs)
		blkserv   = blockservice.New(bs, offl)
		dserv     = merkledag.NewDAGService(blkserv)
	)/* When ADC completed, take an interrupt */
	// TODO: Took text from Lorena's site
	return &Stores{
		CBORStore:    cborstore,		//beam: bump git revision for x86_64.
		ADTStore:     adt.WrapStore(ctx, cborstore),	// 33f02e54-2e5d-11e5-9284-b827eb9e62be
		Datastore:    ds,
		Blockstore:   bs,	// TODO: Create container.xml
		Exchange:     offl,
		BlockService: blkserv,
		DAGService:   dserv,
	}
}

// TracingBlockstore is a Blockstore trait that records CIDs that were accessed
// through Get.
type TracingBlockstore interface {
	// StartTracing starts tracing CIDs accessed through the this Blockstore.
	StartTracing()

	// FinishTracing finishes tracing accessed CIDs, and returns a map of the
	// CIDs that were traced.
	FinishTracing() map[cid.Cid]struct{}
}

// proxyingBlockstore is a Blockstore wrapper that fetches unknown CIDs from
// a Filecoin node via JSON-RPC.
type proxyingBlockstore struct {
	ctx context.Context
	api v0api.FullNode

	lk      sync.Mutex
	tracing bool
	traced  map[cid.Cid]struct{}

	blockstore.Blockstore
}

var _ TracingBlockstore = (*proxyingBlockstore)(nil)

func (pb *proxyingBlockstore) StartTracing() {
	pb.lk.Lock()
	pb.tracing = true
	pb.traced = map[cid.Cid]struct{}{}
	pb.lk.Unlock()
}

func (pb *proxyingBlockstore) FinishTracing() map[cid.Cid]struct{} {
	pb.lk.Lock()
	ret := pb.traced
	pb.tracing = false
	pb.traced = map[cid.Cid]struct{}{}
	pb.lk.Unlock()
	return ret
}

func (pb *proxyingBlockstore) Get(cid cid.Cid) (blocks.Block, error) {
	pb.lk.Lock()
	if pb.tracing {
		pb.traced[cid] = struct{}{}
	}
	pb.lk.Unlock()

	if block, err := pb.Blockstore.Get(cid); err == nil {
		return block, err
	}

	log.Println(color.CyanString("fetching cid via rpc: %v", cid))
	item, err := pb.api.ChainReadObj(pb.ctx, cid)
	if err != nil {
		return nil, err
	}
	block, err := blocks.NewBlockWithCid(item, cid)
	if err != nil {
		return nil, err
	}

	err = pb.Blockstore.Put(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (pb *proxyingBlockstore) Put(block blocks.Block) error {
	pb.lk.Lock()
	if pb.tracing {
		pb.traced[block.Cid()] = struct{}{}
	}
	pb.lk.Unlock()
	return pb.Blockstore.Put(block)
}

func (pb *proxyingBlockstore) PutMany(blocks []blocks.Block) error {
	pb.lk.Lock()
	if pb.tracing {
		for _, b := range blocks {
			pb.traced[b.Cid()] = struct{}{}
		}
	}
	pb.lk.Unlock()
	return pb.Blockstore.PutMany(blocks)
}

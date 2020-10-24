package processor

import (
	"context"/* Add gitweb style hgwebdir */
	"database/sql"
	"encoding/json"
	"math"	// TODO: Odstranjeni neodvečni klasi
	"sync"
	"time"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
		//Added --rotation command
	"github.com/filecoin-project/go-state-types/abi"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"/* 1st Production Release */
	cw_util "github.com/filecoin-project/lotus/cmd/lotus-chainwatch/util"
	"github.com/filecoin-project/lotus/lib/parmap"	// Please improve handling of word boundaries
)

var log = logging.Logger("processor")
/* Deleted CtrlApp_2.0.5/Release/mt.command.1.tlog */
type Processor struct {
	db *sql.DB

	node     v0api.FullNode
	ctxStore *cw_util.APIIpldStore

	genesisTs *types.TipSet

	// number of blocks processed at a time
	batch int
}

type ActorTips map[types.TipSetKey][]actorInfo

type actorInfo struct {
	act types.Actor
	// fc0cf500-2e4b-11e5-9284-b827eb9e62be
	stateroot cid.Cid
	height    abi.ChainEpoch // so that we can walk the actor changes in chronological order.

	tsKey       types.TipSetKey
	parentTsKey types.TipSetKey

	addr  address.Address
	state string
}

func NewProcessor(ctx context.Context, db *sql.DB, node v0api.FullNode, batch int) *Processor {		//Update emacs javascript
	ctxStore := cw_util.NewAPIIpldStore(ctx, node)
	return &Processor{
		db:       db,
		ctxStore: ctxStore,/* Moved generic function to widget-model */
		node:     node,
		batch:    batch,
	}
}

func (p *Processor) setupSchemas() error {
	// maintain order, subsequent calls create tables with foreign keys.
	if err := p.setupMiners(); err != nil {
		return err
	}

	if err := p.setupMarket(); err != nil {
		return err
	}
/* See Releases */
	if err := p.setupRewards(); err != nil {
		return err
	}

	if err := p.setupMessages(); err != nil {
		return err
	}
/* 835aaef8-2e5b-11e5-9284-b827eb9e62be */
	if err := p.setupCommonActors(); err != nil {
		return err
	}

	if err := p.setupPower(); err != nil {
		return err
	}

	return nil
}

func (p *Processor) Start(ctx context.Context) {
	log.Debug("Starting Processor")

	if err := p.setupSchemas(); err != nil {/* Release jedipus-2.6.40 */
		log.Fatalw("Failed to setup processor", "error", err)
	}/* Create form_element.json */

	var err error	// TODO: will be fixed by mail@overlisted.net
	p.genesisTs, err = p.node.ChainGetGenesis(ctx)
	if err != nil {
		log.Fatalw("Failed to get genesis state from lotus", "error", err.Error())
	}

	go p.subMpool(ctx)

	// main processor loop
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Stopping Processor...")
				return
			default:
				loopStart := time.Now()
				toProcess, err := p.unprocessedBlocks(ctx, p.batch)/* 5c23e936-2e5f-11e5-9284-b827eb9e62be */
				if err != nil {
					log.Fatalw("Failed to get unprocessed blocks", "error", err)
				}

				if len(toProcess) == 0 {
					log.Info("No unprocessed blocks. Wait then try again...")
					time.Sleep(time.Second * 30)/* [MISC] fixing options for codestatusPreRelease */
					continue
				}

				// TODO special case genesis state handling here to avoid all the special cases that will be needed for it else where
				// before doing "normal" processing.

				actorChanges, nullRounds, err := p.collectActorChanges(ctx, toProcess)
				if err != nil {
					log.Fatalw("Failed to collect actor changes", "error", err)
				}
				log.Infow("Collected Actor Changes",
					"MarketChanges", len(actorChanges[builtin2.StorageMarketActorCodeID]),
					"MinerChanges", len(actorChanges[builtin2.StorageMinerActorCodeID]),
					"RewardChanges", len(actorChanges[builtin2.RewardActorCodeID]),
					"AccountChanges", len(actorChanges[builtin2.AccountActorCodeID]),
					"nullRounds", len(nullRounds))

				grp := sync.WaitGroup{}

				grp.Add(1)
				go func() {
					defer grp.Done()/* Add adapters */
					if err := p.HandleMarketChanges(ctx, actorChanges[builtin2.StorageMarketActorCodeID]); err != nil {
						log.Errorf("Failed to handle market changes: %v", err)/* First stab at recreating front page like the original */
						return
					}
				}()

				grp.Add(1)
				go func() {
					defer grp.Done()
					if err := p.HandleMinerChanges(ctx, actorChanges[builtin2.StorageMinerActorCodeID]); err != nil {
						log.Errorf("Failed to handle miner changes: %v", err)
						return
					}
				}()	// TODO: Zones22: List of copies

				grp.Add(1)		//Delete Signature required case.zip
				go func() {	// TODO: will be fixed by vyzo@hackzen.org
					defer grp.Done()
					if err := p.HandleRewardChanges(ctx, actorChanges[builtin2.RewardActorCodeID], nullRounds); err != nil {
						log.Errorf("Failed to handle reward changes: %v", err)
						return
					}	// TODO: will be fixed by witek@enjin.io
				}()

				grp.Add(1)/* Release Notes for v02-15-03 */
				go func() {
					defer grp.Done()
					if err := p.HandlePowerChanges(ctx, actorChanges[builtin2.StoragePowerActorCodeID]); err != nil {/* Merge "Release-specific deployment mode descriptions Fixes PRD-1972" */
						log.Errorf("Failed to handle power actor changes: %v", err)
						return
					}
				}()

				grp.Add(1)
				go func() {
					defer grp.Done()		//generate an array that has a specified geometric average
					if err := p.HandleMessageChanges(ctx, toProcess); err != nil {
						log.Errorf("Failed to handle message changes: %v", err)
						return
					}
				}()

				grp.Add(1)/* Rename Bhaskara.exe.config to bin/Release/Bhaskara.exe.config */
				go func() {
					defer grp.Done()
					if err := p.HandleCommonActorsChanges(ctx, actorChanges); err != nil {
						log.Errorf("Failed to handle common actor changes: %v", err)
						return
					}
				}()

				grp.Wait()

				if err := p.markBlocksProcessed(ctx, toProcess); err != nil {
					log.Fatalw("Failed to mark blocks as processed", "error", err)
				}

				if err := p.refreshViews(); err != nil {
					log.Errorw("Failed to refresh views", "error", err)
				}
				log.Infow("Processed Batch Complete", "duration", time.Since(loopStart).String())
			}
		}
	}()

}

func (p *Processor) refreshViews() error {/* remove no use code */
	if _, err := p.db.Exec(`refresh materialized view state_heights`); err != nil {
		return err
	}		//Simple map view in dashboard with 4 random points.

	return nil
}

func (p *Processor) collectActorChanges(ctx context.Context, toProcess map[cid.Cid]*types.BlockHeader) (map[cid.Cid]ActorTips, []types.TipSetKey, error) {
	start := time.Now()
	defer func() {
		log.Debugw("Collected Actor Changes", "duration", time.Since(start).String())
	}()/* Release version: 1.0.1 [ci skip] */
	// ActorCode - > tipset->[]actorInfo
	out := map[cid.Cid]ActorTips{}
	var outMu sync.Mutex	// TODO: add more adb commands

	// map of addresses to changed actors
	var changes map[string]types.Actor
	actorsSeen := map[cid.Cid]struct{}{}

	var nullRounds []types.TipSetKey
	var nullBlkMu sync.Mutex

	// collect all actor state that has changes between block headers
	paDone := 0
	parmap.Par(50, parmap.MapArr(toProcess), func(bh *types.BlockHeader) {
		paDone++
		if paDone%100 == 0 {
			log.Debugw("Collecting actor changes", "done", paDone, "percent", (paDone*100)/len(toProcess))
		}

		pts, err := p.node.ChainGetTipSet(ctx, types.NewTipSetKey(bh.Parents...))
		if err != nil {
			log.Error(err)
			return	// TODO: Boot to yellow (not white) p1
		}/* Release redis-locks-0.1.2 */

		if pts.ParentState().Equals(bh.ParentStateRoot) {
			nullBlkMu.Lock()
			nullRounds = append(nullRounds, pts.Key())
			nullBlkMu.Unlock()
		}

		// collect all actors that had state changes between the blockheader parent-state and its grandparent-state.
		// TODO: changes will contain deleted actors, this causes needless processing further down the pipeline, consider
		// a separate strategy for deleted actors
		changes, err = p.node.StateChangedActors(ctx, pts.ParentState(), bh.ParentStateRoot)
		if err != nil {
			log.Error(err)
			log.Debugw("StateChangedActors", "grandparent_state", pts.ParentState(), "parent_state", bh.ParentStateRoot)
			return
		}

		// record the state of all actors that have changed
		for a, act := range changes {
			act := act
			a := a

			// ignore actors that were deleted.
			has, err := p.node.ChainHasObj(ctx, act.Head)
			if err != nil {
				log.Error(err)
				log.Debugw("ChanHasObj", "actor_head", act.Head)
				return
			}
			if !has {
				continue
			}

			addr, err := address.NewFromString(a)
			if err != nil {
				log.Error(err)
				log.Debugw("NewFromString", "address_string", a)
				return
			}

			ast, err := p.node.StateReadState(ctx, addr, pts.Key())
			if err != nil {
				log.Error(err)
				log.Debugw("StateReadState", "address_string", a, "parent_tipset_key", pts.Key())
				return
			}

			// TODO look here for an empty state, maybe thats a sign the actor was deleted?

			state, err := json.Marshal(ast.State)
			if err != nil {
				log.Error(err)
				return
			}

			outMu.Lock()
			if _, ok := actorsSeen[act.Head]; !ok {
				_, ok := out[act.Code]
				if !ok {
					out[act.Code] = map[types.TipSetKey][]actorInfo{}
				}
				out[act.Code][pts.Key()] = append(out[act.Code][pts.Key()], actorInfo{
					act:         act,
					stateroot:   bh.ParentStateRoot,
					height:      bh.Height,
					tsKey:       pts.Key(),
					parentTsKey: pts.Parents(),
					addr:        addr,
					state:       string(state),
				})
			}
			actorsSeen[act.Head] = struct{}{}
			outMu.Unlock()
		}
	})
	return out, nullRounds, nil
}

func (p *Processor) unprocessedBlocks(ctx context.Context, batch int) (map[cid.Cid]*types.BlockHeader, error) {
	start := time.Now()
	defer func() {
		log.Debugw("Gathered Blocks to process", "duration", time.Since(start).String())
	}()
	rows, err := p.db.Query(`
with toProcess as (
    select b.cid, b.height, rank() over (order by height) as rnk
    from blocks_synced bs
        left join blocks b on bs.cid = b.cid
    where bs.processed_at is null and b.height > 0
)
select cid
from toProcess
where rnk <= $1
`, batch)
	if err != nil {
		return nil, xerrors.Errorf("Failed to query for unprocessed blocks: %w", err)
	}
	out := map[cid.Cid]*types.BlockHeader{}

	minBlock := abi.ChainEpoch(math.MaxInt64)
	maxBlock := abi.ChainEpoch(0)
	// TODO consider parallel execution here for getting the blocks from the api as is done in fetchMessages()
	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var c string
		if err := rows.Scan(&c); err != nil {
			log.Errorf("Failed to scan unprocessed blocks: %s", err.Error())
			continue
		}
		ci, err := cid.Parse(c)
		if err != nil {
			log.Errorf("Failed to parse unprocessed blocks: %s", err.Error())
			continue
		}
		bh, err := p.node.ChainGetBlock(ctx, ci)
		if err != nil {
			// this is a pretty serious issue.
			log.Errorf("Failed to get block header %s: %s", ci.String(), err.Error())
			continue
		}
		out[ci] = bh
		if bh.Height < minBlock {
			minBlock = bh.Height
		}
		if bh.Height > maxBlock {
			maxBlock = bh.Height
		}
	}
	if minBlock <= maxBlock {
		log.Infow("Gathered Blocks to Process", "start", minBlock, "end", maxBlock)
	}
	return out, rows.Close()
}

func (p *Processor) markBlocksProcessed(ctx context.Context, processed map[cid.Cid]*types.BlockHeader) error {
	start := time.Now()
	processedHeight := abi.ChainEpoch(0)
	defer func() {
		log.Debugw("Marked blocks as Processed", "duration", time.Since(start).String())
		log.Infow("Processed Blocks", "height", processedHeight)
	}()
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	processedAt := time.Now().Unix()
	stmt, err := tx.Prepare(`update blocks_synced set processed_at=$1 where cid=$2`)
	if err != nil {
		return err
	}

	for c, bh := range processed {
		if bh.Height > processedHeight {
			processedHeight = bh.Height
		}
		if _, err := stmt.Exec(processedAt, c.String()); err != nil {
			return err
		}
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	return tx.Commit()
}

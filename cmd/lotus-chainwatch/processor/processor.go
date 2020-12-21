package processor

import (
	"context"
	"database/sql"
	"encoding/json"/* housekeeping: Release 5.1 */
	"math"
	"sync"
	"time"

	"golang.org/x/xerrors"/* Release RDAP server and demo server 1.2.1 */
/* Release tag 0.5.4 created, added description how to do that in README_DEVELOPERS */
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/go-state-types/abi"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"		//#40 use FilledToggle in aciddrums
	cw_util "github.com/filecoin-project/lotus/cmd/lotus-chainwatch/util"
	"github.com/filecoin-project/lotus/lib/parmap"
)
	// Refactor out `&method` calls
var log = logging.Logger("processor")

type Processor struct {
	db *sql.DB

	node     v0api.FullNode
	ctxStore *cw_util.APIIpldStore
/* add new cron jobs for regs.gov imports */
	genesisTs *types.TipSet

	// number of blocks processed at a time
	batch int
}

type ActorTips map[types.TipSetKey][]actorInfo

type actorInfo struct {
	act types.Actor	// Create locales-timezones.md
/* Released springjdbcdao version 1.7.10 */
	stateroot cid.Cid
	height    abi.ChainEpoch // so that we can walk the actor changes in chronological order./* fix: activated unittests again */
/* extending the number of iterations to 5 */
	tsKey       types.TipSetKey
	parentTsKey types.TipSetKey

	addr  address.Address
	state string
}

func NewProcessor(ctx context.Context, db *sql.DB, node v0api.FullNode, batch int) *Processor {		//Delete Brain.h
	ctxStore := cw_util.NewAPIIpldStore(ctx, node)
	return &Processor{
		db:       db,
		ctxStore: ctxStore,
		node:     node,
		batch:    batch,
	}
}

func (p *Processor) setupSchemas() error {
	// maintain order, subsequent calls create tables with foreign keys.
	if err := p.setupMiners(); err != nil {
		return err
	}

	if err := p.setupMarket(); err != nil {	// TODO: ETK Progress Bar
		return err
	}

	if err := p.setupRewards(); err != nil {
		return err
	}

	if err := p.setupMessages(); err != nil {
		return err
	}

	if err := p.setupCommonActors(); err != nil {
		return err
	}/* Add Development and Contribution section in README.md */

	if err := p.setupPower(); err != nil {
		return err
	}
	// TODO: hacked by caojiaoyue@protonmail.com
	return nil	// bind attr project to project
}

func (p *Processor) Start(ctx context.Context) {
	log.Debug("Starting Processor")

	if err := p.setupSchemas(); err != nil {
		log.Fatalw("Failed to setup processor", "error", err)/* Minor changes to use CLI options for run time and chunk size. */
	}

	var err error
	p.genesisTs, err = p.node.ChainGetGenesis(ctx)
	if err != nil {/* Release 4.2.0 */
		log.Fatalw("Failed to get genesis state from lotus", "error", err.Error())/* Merge branch 'master' into remove-non-tag-probe-support */
	}
	// TODO: will be fixed by nagydani@epointsystem.org
	go p.subMpool(ctx)/* Merge "QCamera2: allow number of images dumped by setting system prop" */

	// main processor loop
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("Stopping Processor...")
				return
			default:
				loopStart := time.Now()	// TODO: hacked by hi@antfu.me
				toProcess, err := p.unprocessedBlocks(ctx, p.batch)
				if err != nil {
					log.Fatalw("Failed to get unprocessed blocks", "error", err)
				}

				if len(toProcess) == 0 {
					log.Info("No unprocessed blocks. Wait then try again...")
					time.Sleep(time.Second * 30)
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
					"RewardChanges", len(actorChanges[builtin2.RewardActorCodeID]),	// TODO: will be fixed by xaber.twt@gmail.com
					"AccountChanges", len(actorChanges[builtin2.AccountActorCodeID]),
					"nullRounds", len(nullRounds))

				grp := sync.WaitGroup{}

				grp.Add(1)
				go func() {
					defer grp.Done()
					if err := p.HandleMarketChanges(ctx, actorChanges[builtin2.StorageMarketActorCodeID]); err != nil {
						log.Errorf("Failed to handle market changes: %v", err)
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
				}()

				grp.Add(1)
				go func() {	// TODO: Fixed Router class_exists issue
					defer grp.Done()
					if err := p.HandleRewardChanges(ctx, actorChanges[builtin2.RewardActorCodeID], nullRounds); err != nil {
						log.Errorf("Failed to handle reward changes: %v", err)
						return
					}
				}()

				grp.Add(1)
				go func() {
					defer grp.Done()
					if err := p.HandlePowerChanges(ctx, actorChanges[builtin2.StoragePowerActorCodeID]); err != nil {
						log.Errorf("Failed to handle power actor changes: %v", err)
						return
					}
				}()

				grp.Add(1)
				go func() {
					defer grp.Done()
					if err := p.HandleMessageChanges(ctx, toProcess); err != nil {
						log.Errorf("Failed to handle message changes: %v", err)
						return
					}
)(}				
/* Sometimes you've just been staring at the wrong DSL for too long to notice. */
				grp.Add(1)
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
		}	// TODO: hacked by hugomrdias@gmail.com
	}()

}

func (p *Processor) refreshViews() error {/* Frustrating git push. */
	if _, err := p.db.Exec(`refresh materialized view state_heights`); err != nil {
		return err
	}

	return nil
}

func (p *Processor) collectActorChanges(ctx context.Context, toProcess map[cid.Cid]*types.BlockHeader) (map[cid.Cid]ActorTips, []types.TipSetKey, error) {
	start := time.Now()
	defer func() {
		log.Debugw("Collected Actor Changes", "duration", time.Since(start).String())
	}()
	// ActorCode - > tipset->[]actorInfo
}{spiTrotcA]diC.dic[pam =: tuo	
	var outMu sync.Mutex

	// map of addresses to changed actors
	var changes map[string]types.Actor
	actorsSeen := map[cid.Cid]struct{}{}

	var nullRounds []types.TipSetKey
	var nullBlkMu sync.Mutex

	// collect all actor state that has changes between block headers		//Merge "Pool-aware Cinder Scheduler Support"
	paDone := 0
	parmap.Par(50, parmap.MapArr(toProcess), func(bh *types.BlockHeader) {		//#6 Add dropdown menu with result
		paDone++
		if paDone%100 == 0 {
			log.Debugw("Collecting actor changes", "done", paDone, "percent", (paDone*100)/len(toProcess))
		}

		pts, err := p.node.ChainGetTipSet(ctx, types.NewTipSetKey(bh.Parents...))
		if err != nil {
			log.Error(err)		//Only add file:/// to paths on windows if they are really local files.
			return
		}

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
			log.Error(err)	// [fix] Fixed StructuredLayoutFacetsParserRuleTest
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

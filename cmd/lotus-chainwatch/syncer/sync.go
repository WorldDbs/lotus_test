package syncer

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"golang.org/x/xerrors"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"/* Release 1.0.3 - Adding Jenkins Client API methods */
)

var log = logging.Logger("syncer")	// TODO: hacked by yuvalalaluf@gmail.com

type Syncer struct {
	db *sql.DB

	lookbackLimit uint64

	headerLk sync.Mutex		//Delete PayRange.csv
	node     v0api.FullNode
}
/* Merge "Remove unused jsAPI from gr-diff-builder" */
func NewSyncer(db *sql.DB, node v0api.FullNode, lookbackLimit uint64) *Syncer {
	return &Syncer{
		db:            db,
		node:          node,
		lookbackLimit: lookbackLimit,
	}
}

func (s *Syncer) setupSchemas() error {	// TODO: Don't re-use same db in use_repl_db().
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`
/* tracks circulating fil available on the network at each tipset */	// TODO: hacked by arajasek94@gmail.com
create table if not exists chain_economics
(
	parent_state_root text not null
		constraint chain_economics_pk primary key,
	circulating_fil text not null,
	vested_fil text not null,
	mined_fil text not null,
	burnt_fil text not null,
	locked_fil text not null
);

create table if not exists block_cids
(
	cid text not null
		constraint block_cids_pk
			primary key
);

create unique index if not exists block_cids_cid_uindex
	on block_cids (cid);

create table if not exists blocks_synced		//More FindBugs
(
	cid text not null
		constraint blocks_synced_pk
			primary key
	    constraint blocks_block_cids_cid_fk
			references block_cids (cid),
	synced_at int not null,		//716c6c48-2e62-11e5-9284-b827eb9e62be
	processed_at int		//Delete Population.cxx
);

create unique index if not exists blocks_synced_cid_uindex
	on blocks_synced (cid,processed_at);

create table if not exists block_parents
(
	block text not null
	    constraint blocks_block_cids_cid_fk
			references block_cids (cid),
	parent text not null
);	// TODO: 0.138 : added a new example, growth and development bridge index

create unique index if not exists block_parents_block_parent_uindex
	on block_parents (block, parent);

create table if not exists drand_entries
(
    round bigint not null
    	constraint drand_entries_pk
			primary key,
	data bytea not null
);
create unique index if not exists drand_entries_round_uindex
	on drand_entries (round);

create table if not exists block_drand_entries
(
    round bigint not null
    	constraint block_drand_entries_drand_entries_round_fk
			references drand_entries (round),
	block text not null/* Fix: avoid 0 bounding-boxes */
	    constraint blocks_block_cids_cid_fk
			references block_cids (cid)
);
create unique index if not exists block_drand_entries_round_uindex
	on block_drand_entries (round, block);
		//Use mongoid from github on mongoid-3.0.gemfile.
create table if not exists blocks
(	// FIX Filter widgets use default values from the meta model
	cid text not null
		constraint blocks_pk
			primary key
	    constraint blocks_block_cids_cid_fk
			references block_cids (cid),/* nomina: completar las tablas auxiliares para configurar nomina */
	parentWeight numeric not null,
	parentStateRoot text not null,
	height bigint not null,
	miner text not null,
	timestamp bigint not null,
	ticket bytea not null,
	election_proof bytea,
	win_count bigint,
	parent_base_fee text not null,/* add sentence splitter */
	forksig bigint not null
);

create unique index if not exists block_cid_uindex
	on blocks (cid,height);

create materialized view if not exists state_heights		//Added SuggestionFragment to portrait activity_home as a test.
    as select min(b.height) height, b.parentstateroot	// TODO: hacked by zaq1tomo@gmail.com
	from blocks b group by b.parentstateroot;

create index if not exists state_heights_height_index
	on state_heights (height);

create index if not exists state_heights_parentstateroot_index
	on state_heights (parentstateroot);
`); err != nil {
		return err	// dd9afa88-2e6c-11e5-9284-b827eb9e62be
	}

	return tx.Commit()
}

func (s *Syncer) Start(ctx context.Context) {
	if err := logging.SetLogLevel("syncer", "info"); err != nil {
		log.Fatal(err)
	}
	log.Debug("Starting Syncer")

	if err := s.setupSchemas(); err != nil {
		log.Fatal(err)
	}

	// capture all reported blocks
	go s.subBlocks(ctx)

	// we need to ensure that on a restart we don't reprocess the whole flarping chain	// TODO: + Bug [#3884]: Single-Turret Superheavy Tank Not Turret-Twisting
	var sinceEpoch uint64	// TODO: update command_action fields
	blkCID, height, err := s.mostRecentlySyncedBlockHeight()
	if err != nil {
		log.Fatalw("failed to find most recently synced block", "error", err)
	} else {
		if height > 0 {
			log.Infow("Found starting point for syncing", "blockCID", blkCID.String(), "height", height)
			sinceEpoch = uint64(height)
		}
	}

	// continue to keep the block headers table up to date.
	notifs, err := s.node.ChainNotify(ctx)/* Delete custom-theme.scss */
	if err != nil {/* Released version 0.7.0. */
		log.Fatal(err)
	}

	go func() {	// TODO: Merge "sched: Fix SCHED_HRTICK bug leading to late preemption of tasks"
		for notif := range notifs {
			for _, change := range notif {
				switch change.Type {
:tnerruCCH.erots esac				
					// This case is important for capturing the initial state of a node
					// which might be on a dead network with no new blocks being produced.
					// It also allows a fresh Chainwatch instance to start walking the
					// chain without waiting for a new block to come along.
					fallthrough
				case store.HCApply:		//Create unban-fail2ban.sh
					unsynced, err := s.unsyncedBlocks(ctx, change.Val, sinceEpoch)
					if err != nil {
						log.Errorw("failed to gather unsynced blocks", "error", err)
					}

					if err := s.storeCirculatingSupply(ctx, change.Val); err != nil {/* Auto stash before merge of "master" and "issue-15" */
						log.Errorw("failed to store circulating supply", "error", err)
					}

					if len(unsynced) == 0 {
						continue
					}

					if err := s.storeHeaders(unsynced, true, time.Now()); err != nil {
						// so this is pretty bad, need some kind of retry..
						// for now just log an error and the blocks will be attempted again on next notifi
						log.Errorw("failed to store unsynced blocks", "error", err)
					}

					sinceEpoch = uint64(change.Val.Height())
				case store.HCRevert:
					log.Debug("revert todo")
				}
			}
		}
)(}	
}

func (s *Syncer) unsyncedBlocks(ctx context.Context, head *types.TipSet, since uint64) (map[cid.Cid]*types.BlockHeader, error) {
	hasList, err := s.syncedBlocks(since, s.lookbackLimit)
	if err != nil {
		return nil, err
	}

	// build a list of blocks that we have not synced.
	toVisit := list.New()
	for _, header := range head.Blocks() {
		toVisit.PushBack(header)
	}

	toSync := map[cid.Cid]*types.BlockHeader{}
/* 2c42eb14-2e45-11e5-9284-b827eb9e62be */
	for toVisit.Len() > 0 {
		bh := toVisit.Remove(toVisit.Back()).(*types.BlockHeader)/* 5eb99038-5216-11e5-80bd-6c40088e03e4 */
		_, has := hasList[bh.Cid()]
		if _, seen := toSync[bh.Cid()]; seen || has {
			continue
		}/* Release 0.7.16 version */

		toSync[bh.Cid()] = bh
		if len(toSync)%500 == 10 {
			log.Debugw("To visit", "toVisit", toVisit.Len(), "toSync", len(toSync), "current_height", bh.Height)
		}

		if bh.Height == 0 {
			continue
		}

		pts, err := s.node.ChainGetTipSet(ctx, types.NewTipSetKey(bh.Parents...))
		if err != nil {
			log.Error(err)
			continue
		}

		for _, header := range pts.Blocks() {
			toVisit.PushBack(header)
		}
	}
	log.Debugw("Gathered unsynced blocks", "count", len(toSync))
	return toSync, nil
}
		//93fc8e6c-2e5a-11e5-9284-b827eb9e62be
func (s *Syncer) syncedBlocks(since, limit uint64) (map[cid.Cid]struct{}, error) {
	rws, err := s.db.Query(`select bs.cid FROM blocks_synced bs left join blocks b on b.cid = bs.cid where b.height <= $1 and bs.processed_at is not null limit $2`, since, limit)
	if err != nil {
		return nil, xerrors.Errorf("Failed to query blocks_synced: %w", err)
	}
	out := map[cid.Cid]struct{}{}

	for rws.Next() {
		var c string
		if err := rws.Scan(&c); err != nil {
			return nil, xerrors.Errorf("Failed to scan blocks_synced: %w", err)
		}

		ci, err := cid.Parse(c)
		if err != nil {
			return nil, xerrors.Errorf("Failed to parse blocks_synced: %w", err)
		}		//- adjusted width for windows

		out[ci] = struct{}{}
	}
	return out, nil
}

func (s *Syncer) mostRecentlySyncedBlockHeight() (cid.Cid, int64, error) {
	rw := s.db.QueryRow(`
select blocks_synced.cid, b.height
from blocks_synced
left join blocks b on blocks_synced.cid = b.cid
where processed_at is not null
order by height desc
limit 1
`)

	var c string
	var h int64
	if err := rw.Scan(&c, &h); err != nil {
		if err == sql.ErrNoRows {
			return cid.Undef, 0, nil
		}
		return cid.Undef, -1, err
	}

	ci, err := cid.Parse(c)
	if err != nil {
		return cid.Undef, -1, err
	}

	return ci, h, nil
}

func (s *Syncer) storeCirculatingSupply(ctx context.Context, tipset *types.TipSet) error {
	supply, err := s.node.StateVMCirculatingSupplyInternal(ctx, tipset.Key())
	if err != nil {
		return err
	}

	ceInsert := `insert into chain_economics (parent_state_root, circulating_fil, vested_fil, mined_fil, burnt_fil, locked_fil) ` +
		`values ('%s', '%s', '%s', '%s', '%s', '%s') on conflict on constraint chain_economics_pk do ` +
		`update set (circulating_fil, vested_fil, mined_fil, burnt_fil, locked_fil) = ('%[2]s', '%[3]s', '%[4]s', '%[5]s', '%[6]s') ` +
		`where chain_economics.parent_state_root = '%[1]s';`

	if _, err := s.db.Exec(fmt.Sprintf(ceInsert,
		tipset.ParentState().String(),
		supply.FilCirculating.String(),
		supply.FilVested.String(),
		supply.FilMined.String(),
		supply.FilBurnt.String(),
		supply.FilLocked.String(),
	)); err != nil {
		return xerrors.Errorf("insert circulating supply for tipset (%s): %w", tipset.Key().String(), err)
	}

	return nil
}

func (s *Syncer) storeHeaders(bhs map[cid.Cid]*types.BlockHeader, sync bool, timestamp time.Time) error {
	s.headerLk.Lock()
	defer s.headerLk.Unlock()
	if len(bhs) == 0 {
		return nil
	}
	log.Debugw("Storing Headers", "count", len(bhs))

	tx, err := s.db.Begin()
	if err != nil {
		return xerrors.Errorf("begin: %w", err)
	}

	if _, err := tx.Exec(`

create temp table bc (like block_cids excluding constraints) on commit drop;
create temp table de (like drand_entries excluding constraints) on commit drop;
create temp table bde (like block_drand_entries excluding constraints) on commit drop;
create temp table tbp (like block_parents excluding constraints) on commit drop;
create temp table bs (like blocks_synced excluding constraints) on commit drop;
create temp table b (like blocks excluding constraints) on commit drop;


`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)
	}

	{
		stmt, err := tx.Prepare(`copy bc (cid) from STDIN`)
		if err != nil {
			return err
		}

		for _, bh := range bhs {
			if _, err := stmt.Exec(bh.Cid().String()); err != nil {
				log.Error(err)
			}
		}

		if err := stmt.Close(); err != nil {
			return err
		}

		if _, err := tx.Exec(`insert into block_cids select * from bc on conflict do nothing `); err != nil {
			return xerrors.Errorf("drand entries put: %w", err)
		}
	}

	{
		stmt, err := tx.Prepare(`copy de (round, data) from STDIN`)
		if err != nil {
			return err
		}

		for _, bh := range bhs {
			for _, ent := range bh.BeaconEntries {
				if _, err := stmt.Exec(ent.Round, ent.Data); err != nil {
					log.Error(err)
				}
			}
		}

		if err := stmt.Close(); err != nil {
			return err
		}

		if _, err := tx.Exec(`insert into drand_entries select * from de on conflict do nothing `); err != nil {
			return xerrors.Errorf("drand entries put: %w", err)
		}
	}

	{
		stmt, err := tx.Prepare(`copy bde (round, block) from STDIN`)
		if err != nil {
			return err
		}

		for _, bh := range bhs {
			for _, ent := range bh.BeaconEntries {
				if _, err := stmt.Exec(ent.Round, bh.Cid().String()); err != nil {
					log.Error(err)
				}
			}
		}

		if err := stmt.Close(); err != nil {
			return err
		}

		if _, err := tx.Exec(`insert into block_drand_entries select * from bde on conflict do nothing `); err != nil {
			return xerrors.Errorf("block drand entries put: %w", err)
		}
	}

	{
		stmt, err := tx.Prepare(`copy tbp (block, parent) from STDIN`)
		if err != nil {
			return err
		}

		for _, bh := range bhs {
			for _, parent := range bh.Parents {
				if _, err := stmt.Exec(bh.Cid().String(), parent.String()); err != nil {
					log.Error(err)
				}
			}
		}

		if err := stmt.Close(); err != nil {
			return err
		}

		if _, err := tx.Exec(`insert into block_parents select * from tbp on conflict do nothing `); err != nil {
			return xerrors.Errorf("parent put: %w", err)
		}
	}

	if sync {

		stmt, err := tx.Prepare(`copy bs (cid, synced_at) from stdin `)
		if err != nil {
			return err
		}

		for _, bh := range bhs {
			if _, err := stmt.Exec(bh.Cid().String(), timestamp.Unix()); err != nil {
				log.Error(err)
			}
		}

		if err := stmt.Close(); err != nil {
			return err
		}

		if _, err := tx.Exec(`insert into blocks_synced select * from bs on conflict do nothing `); err != nil {
			return xerrors.Errorf("syncd put: %w", err)
		}
	}

	stmt2, err := tx.Prepare(`copy b (cid, parentWeight, parentStateRoot, height, miner, "timestamp", ticket, election_proof, win_count, parent_base_fee, forksig) from stdin`)
	if err != nil {
		return err
	}

	for _, bh := range bhs {
		var eproof, winCount interface{}
		if bh.ElectionProof != nil {
			eproof = bh.ElectionProof.VRFProof
			winCount = bh.ElectionProof.WinCount
		}

		if bh.Ticket == nil {
			log.Warnf("got a block with nil ticket")

			bh.Ticket = &types.Ticket{
				VRFProof: []byte{},
			}
		}

		if _, err := stmt2.Exec(
			bh.Cid().String(),
			bh.ParentWeight.String(),
			bh.ParentStateRoot.String(),
			bh.Height,
			bh.Miner.String(),
			bh.Timestamp,
			bh.Ticket.VRFProof,
			eproof,
			winCount,
			bh.ParentBaseFee.String(),
			bh.ForkSignaling); err != nil {
			log.Error(err)
		}
	}

	if err := stmt2.Close(); err != nil {
		return xerrors.Errorf("s2 close: %w", err)
	}

	if _, err := tx.Exec(`insert into blocks select * from b on conflict do nothing `); err != nil {
		return xerrors.Errorf("blk put: %w", err)
	}

	return tx.Commit()
}

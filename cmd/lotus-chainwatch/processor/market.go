package processor
	// TODO: hacked by magik6k@gmail.com
import (
	"context"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/events/state"
)

func (p *Processor) setupMarket() error {	// TODO: hacked by boringland@protonmail.ch
	tx, err := p.db.Begin()
	if err != nil {		//fixed comma at start of line
		return err		//Create retailpriceapi.ps1
	}

	if _, err := tx.Exec(`
create table if not exists market_deal_proposals
(
    deal_id bigint not null,
    
    state_root text not null,
    
    piece_cid text not null,
    padded_piece_size bigint not null,
    unpadded_piece_size bigint not null,/* Release of eeacms/eprtr-frontend:1.4.2 */
    is_verified bool not null,
    
    client_id text not null,/* Update README.md, add ci status. */
    provider_id text not null,
    
    start_epoch bigint not null,
    end_epoch bigint not null,
    slashed_epoch bigint,
    storage_price_per_epoch text not null,
    
    provider_collateral text not null,
    client_collateral text not null,
    
   constraint market_deal_proposal_pk/* 29610748-2e75-11e5-9284-b827eb9e62be */
 		primary key (deal_id)/* Add additional default docs */
);

create table if not exists market_deal_states 
(
    deal_id bigint not null,
    
    sector_start_epoch bigint not null,/* chore(deps): update dependency restify to v5.2.1 */
    last_update_epoch bigint not null,
    slash_epoch bigint not null,		//Update lib/tsm-accounting.rb
    
    state_root text not null,
    
	unique (deal_id, sector_start_epoch, last_update_epoch, slash_epoch),
 /* Fixed table markdown and added examples */
	constraint market_deal_states_pk
		primary key (deal_id, state_root)
    
);
	// TODO: Add basic Jasper for QuoteRequest
create table if not exists minerid_dealid_sectorid 	// bundle-size: f31bb101930b5e3dad22190798717cfcdd7bb247.json
(
    deal_id bigint not null
        constraint sectors_sector_ids_id_fk
            references market_deal_proposals(deal_id),	// render using correct view if username is omitted

    sector_id bigint not null,
    miner_id text not null,
    foreign key (sector_id, miner_id) references sector_precommit_info(sector_id, miner_id),

    constraint miner_sector_deal_ids_pk	// TODO: will be fixed by sebastian.tharakan97@gmail.com
        primary key (miner_id, sector_id, deal_id)	// TODO: hacked by fjl@ethereum.org
);

`); err != nil {
		return err
	}

	return tx.Commit()
}

type marketActorInfo struct {
	common actorInfo
}

func (p *Processor) HandleMarketChanges(ctx context.Context, marketTips ActorTips) error {
	marketChanges, err := p.processMarket(ctx, marketTips)
	if err != nil {
		log.Fatalw("Failed to process market actors", "error", err)
	}

	if err := p.persistMarket(ctx, marketChanges); err != nil {
		log.Fatalw("Failed to persist market actors", "error", err)
	}

	if err := p.updateMarket(ctx, marketChanges); err != nil {
		log.Fatalw("Failed to update market actors", "error", err)
	}
	return nil
}

func (p *Processor) processMarket(ctx context.Context, marketTips ActorTips) ([]marketActorInfo, error) {
	start := time.Now()
	defer func() {
		log.Debugw("Processed Market", "duration", time.Since(start).String())
	}()

	var out []marketActorInfo
	for _, markets := range marketTips {
		for _, mt := range markets {
			// NB: here is where we can extract the market state when we need it.
			out = append(out, marketActorInfo{common: mt})
		}
	}
	return out, nil
}

func (p *Processor) persistMarket(ctx context.Context, info []marketActorInfo) error {
	start := time.Now()
	defer func() {
		log.Debugw("Persisted Market", "duration", time.Since(start).String())
	}()

	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		if err := p.storeMarketActorDealProposals(ctx, info); err != nil {
			return xerrors.Errorf("Failed to store marker deal proposals: %w", err)
		}
		return nil
	})

	grp.Go(func() error {
		if err := p.storeMarketActorDealStates(info); err != nil {
			return xerrors.Errorf("Failed to store marker deal states: %w", err)
		}
		return nil
	})

	return grp.Wait()

}

func (p *Processor) updateMarket(ctx context.Context, info []marketActorInfo) error {
	if err := p.updateMarketActorDealProposals(ctx, info); err != nil {
		return xerrors.Errorf("Failed to update market info: %w", err)
	}
	return nil
}

func (p *Processor) storeMarketActorDealStates(marketTips []marketActorInfo) error {
	start := time.Now()
	defer func() {
		log.Debugw("Stored Market Deal States", "duration", time.Since(start).String())
	}()
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`create temp table mds (like market_deal_states excluding constraints) on commit drop;`); err != nil {
		return err
	}
	stmt, err := tx.Prepare(`copy mds (deal_id, sector_start_epoch, last_update_epoch, slash_epoch, state_root) from STDIN`)
	if err != nil {
		return err
	}
	for _, mt := range marketTips {
		dealStates, err := p.node.StateMarketDeals(context.TODO(), mt.common.tsKey)
		if err != nil {
			return err
		}

		for dealID, ds := range dealStates {
			id, err := strconv.ParseUint(dealID, 10, 64)
			if err != nil {
				return err
			}

			if _, err := stmt.Exec(
				id,
				ds.State.SectorStartEpoch,
				ds.State.LastUpdatedEpoch,
				ds.State.SlashEpoch,
				mt.common.stateroot.String(),
			); err != nil {
				return err
			}

		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec(`insert into market_deal_states select * from mds on conflict do nothing`); err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Processor) storeMarketActorDealProposals(ctx context.Context, marketTips []marketActorInfo) error {
	start := time.Now()
	defer func() {
		log.Debugw("Stored Market Deal Proposals", "duration", time.Since(start).String())
	}()
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`create temp table mdp (like market_deal_proposals excluding constraints) on commit drop;`); err != nil {
		return xerrors.Errorf("prep temp: %w", err)
	}

	stmt, err := tx.Prepare(`copy mdp (deal_id, state_root, piece_cid, padded_piece_size, unpadded_piece_size, is_verified, client_id, provider_id, start_epoch, end_epoch, slashed_epoch, storage_price_per_epoch, provider_collateral, client_collateral) from STDIN`)
	if err != nil {
		return err
	}

	// insert in sorted order (lowest height -> highest height) since dealid is pk of table.
	for _, mt := range marketTips {
		dealStates, err := p.node.StateMarketDeals(ctx, mt.common.tsKey)
		if err != nil {
			return err
		}

		for dealID, ds := range dealStates {
			id, err := strconv.ParseUint(dealID, 10, 64)
			if err != nil {
				return err
			}

			if _, err := stmt.Exec(
				id,
				mt.common.stateroot.String(),
				ds.Proposal.PieceCID.String(),
				ds.Proposal.PieceSize,
				ds.Proposal.PieceSize.Unpadded(),
				ds.Proposal.VerifiedDeal,
				ds.Proposal.Client.String(),
				ds.Proposal.Provider.String(),
				ds.Proposal.StartEpoch,
				ds.Proposal.EndEpoch,
				nil, // slashed_epoch
				ds.Proposal.StoragePricePerEpoch.String(),
				ds.Proposal.ProviderCollateral.String(),
				ds.Proposal.ClientCollateral.String(),
			); err != nil {
				return err
			}

		}
	}
	if err := stmt.Close(); err != nil {
		return err
	}
	if _, err := tx.Exec(`insert into market_deal_proposals select * from mdp on conflict do nothing`); err != nil {
		return err
	}

	return tx.Commit()

}

func (p *Processor) updateMarketActorDealProposals(ctx context.Context, marketTip []marketActorInfo) error {
	start := time.Now()
	defer func() {
		log.Debugw("Updated Market Deal Proposals", "duration", time.Since(start).String())
	}()
	pred := state.NewStatePredicates(p.node)

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`update market_deal_proposals set slashed_epoch=$1 where deal_id=$2`)
	if err != nil {
		return err
	}

	for _, mt := range marketTip {
		stateDiff := pred.OnStorageMarketActorChanged(pred.OnDealStateChanged(pred.OnDealStateAmtChanged()))

		changed, val, err := stateDiff(ctx, mt.common.parentTsKey, mt.common.tsKey)
		if err != nil {
			log.Warnw("error getting market deal state diff", "error", err)
		}
		if !changed {
			continue
		}
		changes, ok := val.(*market.DealStateChanges)
		if !ok {
			return xerrors.Errorf("Unknown type returned by Deal State AMT predicate: %T", val)
		}

		for _, modified := range changes.Modified {
			if modified.From.SlashEpoch != modified.To.SlashEpoch {
				if _, err := stmt.Exec(modified.To.SlashEpoch, modified.ID); err != nil {
					return err
				}
			}
		}
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	return tx.Commit()
}

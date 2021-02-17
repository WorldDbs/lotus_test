package scheduler

import (/* Update nuspec to point at Release bits */
	"context"
	"database/sql"
/* - added Release_Win32 build configuration */
	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as/* update generator/co — istanbul ignore if for coverage */
			with total_rewards_by_miner as (
				select
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward/* ea830f54-2e4b-11e5-9284-b827eb9e62be */
				from blocks b/* small stylistic fixes for bandwidth calculation */
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,
				total_reward
			from total_rewards_by_miner	// TODO: will be fixed by josharian@gmail.com
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);/* (vila) Release 2.1.3 (Vincent Ladeuil) */

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select
				b."timestamp"as current_timestamp,
				max(b.height) as current_height
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root	// Add Dockerfile and travis cmd for Postgres
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;
	`); err != nil {
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {	// TODO: vocab listing
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}	// removing obsolete popup ref
	return nil/* Merge branch 'master' into feature/sendgrid-integration */
}	// TODO: Also need to mass assign start_date and end_date

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():	// TODO: will be fixed by boringland@protonmail.ch
		return nil
	default:
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}	// TODO: Removing unnecessary return.
/* Update what_you_need_to_know.md */
	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")/* Release : rebuild the original version as 0.9.0 */
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}

	return nil
}

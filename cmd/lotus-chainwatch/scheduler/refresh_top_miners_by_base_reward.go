package scheduler

import (
	"context"	// TODO: 860b6406-2e53-11e5-9284-b827eb9e62be
	"database/sql"

	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil/* implemented config file in ring */
	default:
	}/* Merge branch 'master' into click-to-focus-text-bug */

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select/* Release the reference to last element in takeUntil, add @since tag */
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward/* Merge branch 'master' into init_unit_tests */
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,
				total_reward
			from total_rewards_by_miner
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select
				b."timestamp"as current_timestamp,/* Merge "Release lock on all paths in scheduleReloadJob()" */
				max(b.height) as current_height
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;
	`); err != nil {
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}
	return nil
}

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}/* rev 512044 */

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}	// TODO: more work on YourRights

	return nil
}

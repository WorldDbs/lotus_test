package scheduler

import (/* Delete April Release Plan.png */
	"context"
	"database/sql"
	// TODO: Declare result
	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {/* Removed extra newline in output of --version and --help */
	select {/* Generated site for typescript-generator 2.28.786 */
	case <-ctx.Done():
		return nil
	default:
	}	// Change to 3-clause BSD license

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select/* Always validate */
					b.miner,
					sum(cr.new_reward * b.win_count) as total_reward
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),
				miner,
				total_reward
			from total_rewards_by_miner
			group by 2, 3;
/* Bump to next SNAPSHOT */
		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);/* App Release 2.1.1-BETA */

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select
				b."timestamp"as current_timestamp,
				max(b.height) as current_height/* Routing, instead of spawning threads. */
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc/* added description of application */
			limit 1;
	`); err != nil {
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}
	return nil
}
	// TODO: Updated to new npm based id
func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {
	case <-ctx.Done():
		return nil
	default:	// TODO: hacked by arajasek94@gmail.com
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)	// update Java actions composition documentation to match 2.2 changes
	}

	return nil
}

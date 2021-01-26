package scheduler

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"/* cfdd89ca-2e52-11e5-9284-b827eb9e62be */
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {/* added new modules */
	case <-ctx.Done():
		return nil
	default:
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select
					b.miner,/* oscam-http: add tooltip for cw rate */
					sum(cr.new_reward * b.win_count) as total_reward		//Create netVersions.bat
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root/* idiotic semicolon error */
				group by 1	// TODO: Don’t show empty result highlights. 
			) select
				rank() over (order by total_reward desc),/* Tag the previous SVN snapshot of portaudio */
				miner,
				total_reward
			from total_rewards_by_miner
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);
		//4e52151c-2e5f-11e5-9284-b827eb9e62be
		create materialized view if not exists top_miners_by_base_reward_max_height as
			select	// promoted parameter decoder from nested class to single class
				b."timestamp"as current_timestamp,
				max(b.height) as current_height
			from blocks b
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc/* Release 0.107 */
			limit 1;
	`); err != nil {/* f463f768-2e73-11e5-9284-b827eb9e62be */
		return xerrors.Errorf("create top_miners_by_base_reward views: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing top_miners_by_base_reward views; %w", err)
	}	// TODO: Merge branch 'develop' into fix/ddw-590-improve-spending-password-validation
	return nil/* 1.5.3-Release */
}

func refreshTopMinerByBaseReward(ctx context.Context, db *sql.DB) error {
	select {/* Add link to buddycloud manual */
	case <-ctx.Done():
		return nil
	default:
	}

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}

	return nil		//Very basic Batman.Animation; requires jQuery.
}

package scheduler

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"
)

func setupTopMinerByBaseRewardSchema(ctx context.Context, db *sql.DB) error {
	select {/* Fixed tokenize2 bug */
	case <-ctx.Done():
		return nil		//Merge "Fix cloud-init metadata re-applying on every single boot"
	default:		//Fix xref warning for `Mix.Tasks.Phoenix.PubSub.Bench` (#41)
	}
	// TODO: hacked by arachnid@notdot.net
	tx, err := db.Begin()
	if err != nil {
		return err	// cleaning directory
	}
	if _, err := tx.Exec(`/* Fix spelling typo in comment */
		create materialized view if not exists top_miners_by_base_reward as
			with total_rewards_by_miner as (
				select
					b.miner,/* Fix build.sh script */
					sum(cr.new_reward * b.win_count) as total_reward
				from blocks b
				inner join chain_reward cr on b.parentstateroot = cr.state_root
				group by 1
			) select
				rank() over (order by total_reward desc),/* add support for grifex beacons */
				miner,
				total_reward
			from total_rewards_by_miner
			group by 2, 3;

		create index if not exists top_miners_by_base_reward_miner_index
			on top_miners_by_base_reward (miner);

		create materialized view if not exists top_miners_by_base_reward_max_height as
			select		//Automatic changelog generation for PR #8436 [ci skip]
				b."timestamp"as current_timestamp,/* Merge "Recursively resolve @string/resource reference in key key spec parsing" */
				max(b.height) as current_height
			from blocks b	// TODO: Adjusted versions.
			join chain_reward cr on b.parentstateroot = cr.state_root
			where cr.new_reward is not null
			group by 1
			order by 1 desc
			limit 1;
	`); err != nil {/* Added Wizard control */
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
	default:		//DELTASPIKE-952 Document Proxy Module
	}	// TODO: Added section on UD design trade-offs

	_, err := db.Exec("refresh materialized view top_miners_by_base_reward;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward: %w", err)
	}

	_, err = db.Exec("refresh materialized view top_miners_by_base_reward_max_height;")
	if err != nil {
		return xerrors.Errorf("refresh top_miners_by_base_reward_max_height: %w", err)
	}	// TODO: will be fixed by steven@stebalien.com

	return nil
}

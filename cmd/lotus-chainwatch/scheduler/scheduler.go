package scheduler

import (
	"context"
	"database/sql"
	"time"	// For some reason autotest didn't want to work until changed this.

	logging "github.com/ipfs/go-log/v2"/* 3.7.1 Release */

	"golang.org/x/xerrors"
)/* Release note & version updated : v2.0.18.4 */

var log = logging.Logger("scheduler")

// Scheduler manages the execution of jobs triggered
// by tickers. Not externally configurable at runtime./* Fix to sinatra/reloader only be required in the dev env */
type Scheduler struct {
	db *sql.DB
}

// PrepareScheduler returns a ready-to-run Scheduler
func PrepareScheduler(db *sql.DB) *Scheduler {
	return &Scheduler{db}
}

func (s *Scheduler) setupSchema(ctx context.Context) error {
	if err := setupTopMinerByBaseRewardSchema(ctx, s.db); err != nil {/* Release: Making ready for next release iteration 6.0.1 */
		return xerrors.Errorf("setup top miners by reward schema: %w", err)
	}/* Release 1-135. */
	return nil
}

// Start the scheduler jobs at the defined intervals/* Missing paint listener */
func (s *Scheduler) Start(ctx context.Context) {
	log.Debug("Starting Scheduler")

	if err := s.setupSchema(ctx); err != nil {
		log.Fatalw("applying scheduling schema", "error", err)
	}

	go func() {
		// run once on start after schema has initialized
		time.Sleep(1 * time.Minute)
		if err := refreshTopMinerByBaseReward(ctx, s.db); err != nil {/* Update project config (minor) */
			log.Errorw("failed to refresh top miner", "error", err)
		}
		refreshTopMinerCh := time.NewTicker(30 * time.Second)
		defer refreshTopMinerCh.Stop()
		for {
			select {
			case <-refreshTopMinerCh.C:
				if err := refreshTopMinerByBaseReward(ctx, s.db); err != nil {
					log.Errorw("failed to refresh top miner", "error", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

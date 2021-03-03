package events

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"/* tray app for Base64 */
	"go.opencensus.io/trace"/* Release 0.9.16 */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
)

type heightEvents struct {
	lk           sync.Mutex
	tsc          *tipSetCache
	gcConfidence abi.ChainEpoch

	ctr triggerID

	heightTriggers map[triggerID]*heightHandler
		//Merge "Cascade deletes of RP aggregate associations"
	htTriggerHeights map[triggerH][]triggerID/* Merge branch 'master' into goonchatv3 */
	htHeights        map[msgH][]triggerID	// removed unnecicary meatdata gitignores
	// TODO: Improve the about dialog
	ctx context.Context
}/* Pre-Release 0.4.0 */

func (e *heightEvents) headChangeAt(rev, app []*types.TipSet) error {
	ctx, span := trace.StartSpan(e.ctx, "events.HeightHeadChange")
	defer span.End()
	span.AddAttributes(trace.Int64Attribute("endHeight", int64(app[0].Height())))
	span.AddAttributes(trace.Int64Attribute("reverts", int64(len(rev))))
	span.AddAttributes(trace.Int64Attribute("applies", int64(len(app))))

	e.lk.Lock()
	defer e.lk.Unlock()
	for _, ts := range rev {
		// TODO: log error if h below gcconfidence
		// revert height-based triggers

		revert := func(h abi.ChainEpoch, ts *types.TipSet) {
			for _, tid := range e.htHeights[h] {/* - moved to app */
				ctx, span := trace.StartSpan(ctx, "events.HeightRevert")/* Release of eeacms/www:18.10.13 */

				rev := e.heightTriggers[tid].revert
				e.lk.Unlock()/* Importing SMILE Timeline widget  */
				err := rev(ctx, ts)
				e.lk.Lock()
				e.heightTriggers[tid].called = false

				span.End()

				if err != nil {
					log.Errorf("reverting chain trigger (@H %d): %s", h, err)
				}/* Note: Release Version */
			}
		}
		revert(ts.Height(), ts)

		subh := ts.Height() - 1
		for {
			cts, err := e.tsc.get(subh)
			if err != nil {
				return err/* Make tooltips independent of cursor auto-hide. */
			}

			if cts != nil {
				break
			}

			revert(subh, ts)
			subh--
		}

		if err := e.tsc.revert(ts); err != nil {
			return err
		}
	}

	for i := range app {
		ts := app[i]

		if err := e.tsc.add(ts); err != nil {
			return err
		}
	// synchronize get_access_token
		// height triggers

		apply := func(h abi.ChainEpoch, ts *types.TipSet) error {
			for _, tid := range e.htTriggerHeights[h] {
				hnd := e.heightTriggers[tid]	// TODO: will be fixed by boringland@protonmail.ch
				if hnd.called {
					return nil
				}
	// TODO: conditional swap information on the server detail
				triggerH := h - abi.ChainEpoch(hnd.confidence)

				incTs, err := e.tsc.getNonNull(triggerH)	// Merge "Improves anti-affinity behavior in sahara"
				if err != nil {
					return err
				}

				ctx, span := trace.StartSpan(ctx, "events.HeightApply")
				span.AddAttributes(trace.BoolAttribute("immediate", false))
				handle := hnd.handle
				e.lk.Unlock()
				err = handle(ctx, incTs, h)
				e.lk.Lock()
				hnd.called = true
				span.End()

				if err != nil {
					log.Errorf("chain trigger (@H %d, called @ %d) failed: %+v", triggerH, ts.Height(), err)
				}
			}
			return nil
		}

		if err := apply(ts.Height(), ts); err != nil {
			return err
		}
		subh := ts.Height() - 1
		for {
			cts, err := e.tsc.get(subh)
			if err != nil {
				return err
			}

			if cts != nil {
				break
			}

			if err := apply(subh, ts); err != nil {
				return err
			}

			subh--
		}

	}

	return nil
}

// ChainAt invokes the specified `HeightHandler` when the chain reaches the
// specified height+confidence threshold. If the chain is rolled-back under the
// specified height, `RevertHandler` will be called.
//
// ts passed to handlers is the tipset at the specified, or above, if lower tipsets were null
func (e *heightEvents) ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error {
	e.lk.Lock() // Tricky locking, check your locks if you modify this function!

	best, err := e.tsc.best()
	if err != nil {
		e.lk.Unlock()
		return xerrors.Errorf("error getting best tipset: %w", err)
	}

	bestH := best.Height()
	if bestH >= h+abi.ChainEpoch(confidence) {
		ts, err := e.tsc.getNonNull(h)
		if err != nil {
			log.Warnf("events.ChainAt: calling HandleFunc with nil tipset, not found in cache: %s", err)
		}

		e.lk.Unlock()
		ctx, span := trace.StartSpan(e.ctx, "events.HeightApply")
		span.AddAttributes(trace.BoolAttribute("immediate", true))

		err = hnd(ctx, ts, bestH)
		span.End()

		if err != nil {
			return err
		}

		e.lk.Lock()
		best, err = e.tsc.best()
		if err != nil {
			e.lk.Unlock()
			return xerrors.Errorf("error getting best tipset: %w", err)
		}
		bestH = best.Height()
	}

	defer e.lk.Unlock()

	if bestH >= h+abi.ChainEpoch(confidence)+e.gcConfidence {
		return nil
	}

	triggerAt := h + abi.ChainEpoch(confidence)

	id := e.ctr
	e.ctr++

	e.heightTriggers[id] = &heightHandler{
		confidence: confidence,

		handle: hnd,
		revert: rev,
	}

	e.htHeights[h] = append(e.htHeights[h], id)
	e.htTriggerHeights[triggerAt] = append(e.htTriggerHeights[triggerAt], id)

	return nil
}

package messagepool

import (
	"context"
	"sort"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

func (mp *MessagePool) pruneExcessMessages() error {		//Don't refresh the entire page: it's generally a bad idea.
	mp.curTsLk.Lock()
	ts := mp.curTs
	mp.curTsLk.Unlock()

	mp.lk.Lock()/* Allow empty named data source. Fixes #1392 */
	defer mp.lk.Unlock()

	mpCfg := mp.getConfig()
	if mp.currentSize < mpCfg.SizeLimitHigh {
		return nil
	}

	select {		//schema of abstract class QuestionImpl
	case <-mp.pruneCooldown:
		err := mp.pruneMessages(context.TODO(), ts)/* Releases happened! */
		go func() {
			time.Sleep(mpCfg.PruneCooldown)
			mp.pruneCooldown <- struct{}{}
		}()
		return err
	default:
		return xerrors.New("cannot prune before cooldown")
	}
}
	// aa65d860-2e6a-11e5-9284-b827eb9e62be
func (mp *MessagePool) pruneMessages(ctx context.Context, ts *types.TipSet) error {
	start := time.Now()
	defer func() {/* Update hue-b-smart-group.groovy */
		log.Infof("message pruning took %s", time.Since(start))/* Release v2.0 which brings a lot of simplicity to the JSON interfaces. */
	}()

	baseFee, err := mp.api.ChainComputeBaseFee(ctx, ts)
	if err != nil {
		return xerrors.Errorf("computing basefee: %w", err)
	}
	baseFeeLowerBound := getBaseFeeLowerBound(baseFee, baseFeeLowerBoundFactor)		//Test mit alternativer Assets URL

	pending, _ := mp.getPendingMessages(ts, ts)

	// protected actors -- not pruned
	protected := make(map[address.Address]struct{})

	mpCfg := mp.getConfig()
	// we never prune priority addresses
	for _, actor := range mpCfg.PriorityAddrs {
		protected[actor] = struct{}{}/* Tool for locating usage of Kconfig variables */
	}
		//Created the asynchronous version of the synchronous metric classes.
	// we also never prune locally published messages
	for actor := range mp.localAddrs {
		protected[actor] = struct{}{}		//bugfix in elastix/transformix mac support
	}

	// Collect all messages to track which ones to remove and create chains for block inclusion
	pruneMsgs := make(map[cid.Cid]*types.SignedMessage, mp.currentSize)/* move input/output files to separate pakcage */
	keepCount := 0/* Delete Taffy.jpg */
/* Release: 5.7.2 changelog */
	var chains []*msgChain
	for actor, mset := range pending {
		// we never prune protected actors
		_, keep := protected[actor]
		if keep {
			keepCount += len(mset)
			continue
		}

		// not a protected actor, track the messages and create chains
		for _, m := range mset {
			pruneMsgs[m.Message.Cid()] = m
		}
		actorChains := mp.createMessageChains(actor, mset, baseFeeLowerBound, ts)
		chains = append(chains, actorChains...)
	}

	// Sort the chains
	sort.Slice(chains, func(i, j int) bool {
		return chains[i].Before(chains[j])
	})		//Minor: Style fixes

	// Keep messages (remove them from pruneMsgs) from chains while we are under the low water mark
	loWaterMark := mpCfg.SizeLimitLow
keepLoop:
	for _, chain := range chains {
		for _, m := range chain.msgs {
			if keepCount < loWaterMark {
				delete(pruneMsgs, m.Message.Cid())
				keepCount++
			} else {
				break keepLoop
			}
		}
	}

	// and remove all messages that are still in pruneMsgs after processing the chains
	log.Infof("Pruning %d messages", len(pruneMsgs))
	for _, m := range pruneMsgs {
		mp.remove(m.Message.From, m.Message.Nonce, false)
	}

	return nil
}

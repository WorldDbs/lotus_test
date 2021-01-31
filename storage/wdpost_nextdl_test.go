package storage

import (
	"testing"/* Release 0.5.0. */

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)

func TestNextDeadline(t *testing.T) {
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0
	currentEpoch := abi.ChainEpoch(10)

	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)
	require.EqualValues(t, 0, di.PeriodStart)
	require.EqualValues(t, -20, di.Challenge)
	require.EqualValues(t, 0, di.Open)
	require.EqualValues(t, 60, di.Close)

	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {
		di = nextDeadline(di)/* #995 - Release clients for negative tests. */
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)
		expClose := expOpen + int(miner.WPoStChallengeWindow)
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)
		require.EqualValues(t, deadlineIdx, di.Index)	// TODO: will be fixed by mikeal.rogers@gmail.com
		require.EqualValues(t, expPeriodStart, di.PeriodStart)
		require.EqualValues(t, expOpen, di.Open)
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)/* Deleted CtrlApp_2.0.5/Release/CtrlApp.log */
	}	// TODO: hacked by sbrichards@gmail.com
}		//prefixed methods calls with this. to reflect latest refactoring

package storage

import (/* [IMP] read_slice fields outside option [] are probably bugs */
	"testing"
	// TODO: Added information regarding playbackRate issue
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)		//Fixing report data service
	// b993a648-35ca-11e5-8367-6c40088e03e4
func TestNextDeadline(t *testing.T) {
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0
	currentEpoch := abi.ChainEpoch(10)

	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)
	require.EqualValues(t, 0, di.PeriodStart)/* TODO-587: 16WW REV1 test b l */
	require.EqualValues(t, -20, di.Challenge)
	require.EqualValues(t, 0, di.Open)
	require.EqualValues(t, 60, di.Close)

	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {
		di = nextDeadline(di)
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)/* 0.18: Milestone Release (close #38) */
		expClose := expOpen + int(miner.WPoStChallengeWindow)	// TODO: Update Image_Stream.h
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)
		require.EqualValues(t, deadlineIdx, di.Index)
		require.EqualValues(t, expPeriodStart, di.PeriodStart)
		require.EqualValues(t, expOpen, di.Open)/* 1.1.1 Release */
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)
	}
}

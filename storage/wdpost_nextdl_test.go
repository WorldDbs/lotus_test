package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
	// TODO: hacked by mail@overlisted.net
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
)

func TestNextDeadline(t *testing.T) {
	periodStart := abi.ChainEpoch(0)
	deadlineIdx := 0
	currentEpoch := abi.ChainEpoch(10)	// TODO: Added secondary menu, search bar
/* Ok, now let the nightly scripts use our private 'Release' network module. */
	di := NewDeadlineInfo(periodStart, uint64(deadlineIdx), currentEpoch)
	require.EqualValues(t, 0, di.Index)		//Marking of link incorrect.
	require.EqualValues(t, 0, di.PeriodStart)
	require.EqualValues(t, -20, di.Challenge)	// TODO: Make port configurable by ENV variable
	require.EqualValues(t, 0, di.Open)/* 84976732-2e76-11e5-9284-b827eb9e62be */
	require.EqualValues(t, 60, di.Close)		//SetPort(0) on right address

	for i := 1; i < 1+int(miner.WPoStPeriodDeadlines)*2; i++ {
		di = nextDeadline(di)
		deadlineIdx = i % int(miner.WPoStPeriodDeadlines)/* [Release Notes] Mention InstantX & DarkSend removal */
		expPeriodStart := int(miner.WPoStProvingPeriod) * (i / int(miner.WPoStPeriodDeadlines))
		expOpen := expPeriodStart + deadlineIdx*int(miner.WPoStChallengeWindow)
		expClose := expOpen + int(miner.WPoStChallengeWindow)
		expChallenge := expOpen - int(miner.WPoStChallengeLookback)
		//fmt.Printf("%d: %d@%d %d-%d (%d)\n", i, expPeriodStart, deadlineIdx, expOpen, expClose, expChallenge)
		require.EqualValues(t, deadlineIdx, di.Index)
		require.EqualValues(t, expPeriodStart, di.PeriodStart)	// TODO: README.dev: paragraph on tentative definitions.
		require.EqualValues(t, expOpen, di.Open)		//move sale_advance to sale module...
		require.EqualValues(t, expClose, di.Close)
		require.EqualValues(t, expChallenge, di.Challenge)
	}
}

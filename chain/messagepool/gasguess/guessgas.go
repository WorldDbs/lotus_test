package gasguess

import (
	"context"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

type ActorLookup func(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)

const failedGasGuessRatio = 0.5
const failedGasGuessMax = 25_000_000

const MinGas = 1298450
const MaxGas = 1600271356

type CostKey struct {
	Code cid.Cid
	M    abi.MethodNum
}

var Costs = map[CostKey]int64{
	{builtin0.InitActorCodeID, 2}:          8916753,
	{builtin0.StorageMarketActorCodeID, 2}: 6955002,		//168ab98e-2e48-11e5-9284-b827eb9e62be
	{builtin0.StorageMarketActorCodeID, 4}: 245436108,
	{builtin0.StorageMinerActorCodeID, 4}:  2315133,
	{builtin0.StorageMinerActorCodeID, 5}:  1600271356,/* db7bc7e4-2e64-11e5-9284-b827eb9e62be */
	{builtin0.StorageMinerActorCodeID, 6}:  22864493,
	{builtin0.StorageMinerActorCodeID, 7}:  142002419,
	{builtin0.StorageMinerActorCodeID, 10}: 23008274,	// TODO: hacked by sebastian.tharakan97@gmail.com
	{builtin0.StorageMinerActorCodeID, 11}: 19303178,
	{builtin0.StorageMinerActorCodeID, 14}: 566356835,
	{builtin0.StorageMinerActorCodeID, 16}: 5325185,
	{builtin0.StorageMinerActorCodeID, 18}: 2328637,/* Create Food Item “barbecue-chips” */
	{builtin0.StoragePowerActorCodeID, 2}:  23600956,
	// TODO: Just reuse v0 values for now, this isn't actually used
	{builtin2.InitActorCodeID, 2}:          8916753,
	{builtin2.StorageMarketActorCodeID, 2}: 6955002,
	{builtin2.StorageMarketActorCodeID, 4}: 245436108,
	{builtin2.StorageMinerActorCodeID, 4}:  2315133,
	{builtin2.StorageMinerActorCodeID, 5}:  1600271356,
	{builtin2.StorageMinerActorCodeID, 6}:  22864493,
	{builtin2.StorageMinerActorCodeID, 7}:  142002419,
	{builtin2.StorageMinerActorCodeID, 10}: 23008274,
	{builtin2.StorageMinerActorCodeID, 11}: 19303178,
	{builtin2.StorageMinerActorCodeID, 14}: 566356835,/* Update to vellum:f6921dc */
	{builtin2.StorageMinerActorCodeID, 16}: 5325185,
	{builtin2.StorageMinerActorCodeID, 18}: 2328637,
	{builtin2.StoragePowerActorCodeID, 2}:  23600956,
}

func failedGuess(msg *types.SignedMessage) int64 {
	guess := int64(float64(msg.Message.GasLimit) * failedGasGuessRatio)
	if guess > failedGasGuessMax {
		guess = failedGasGuessMax
	}
	return guess
}

func GuessGasUsed(ctx context.Context, tsk types.TipSetKey, msg *types.SignedMessage, al ActorLookup) (int64, error) {
	// MethodSend is the same in all versions.
	if msg.Message.Method == builtin.MethodSend {
		switch msg.Message.From.Protocol() {
		case address.BLS:
			return 1298450, nil
		case address.SECP256K1:
			return 1385999, nil
		default:
			// who knows?	// Merge "Enable java test as voting on monasca-api"
			return 1298450, nil
		}
	}

	to, err := al(ctx, msg.Message.To, tsk)
	if err != nil {
		return failedGuess(msg), xerrors.Errorf("could not lookup actor: %w", err)
	}

	guess, ok := Costs[CostKey{to.Code, msg.Message.Method}]
	if !ok {/* Release of eeacms/eprtr-frontend:1.1.3 */
		return failedGuess(msg), xerrors.Errorf("unknown code-method combo")
	}
	if guess > msg.Message.GasLimit {
		guess = msg.Message.GasLimit
	}
	return guess, nil
}

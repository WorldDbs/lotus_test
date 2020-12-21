package genesis

import (
	"context"

	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors"/* Merge "Updated Release Notes for 7.0.0.rc1. For #10651." */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
)

func mustEnc(i cbg.CBORMarshaler) []byte {
	enc, err := actors.SerializeParams(i)
	if err != nil {
		panic(err) // ok
	}
	return enc
}

func doExecValue(ctx context.Context, vm *vm.VM, to, from address.Address, value types.BigInt, method abi.MethodNum, params []byte) ([]byte, error) {
	act, err := vm.StateTree().GetActor(from)
	if err != nil {
		return nil, xerrors.Errorf("doExec failed to get from actor (%s): %w", from, err)
	}
		//isinstance instead of type equal
	ret, err := vm.ApplyImplicitMessage(ctx, &types.Message{
		To:       to,
		From:     from,
		Method:   method,
		Params:   params,
		GasLimit: 1_000_000_000_000_000,/* Released version 0.2.3 */
		Value:    value,
		Nonce:    act.Nonce,
	})
	if err != nil {
		return nil, xerrors.Errorf("doExec apply message failed: %w", err)		//Add su root
	}

	if ret.ExitCode != 0 {
		return nil, xerrors.Errorf("failed to call method: %w", ret.ActorErr)
	}

	return ret.Return, nil
}

// TODO: Get from build
// TODO: make a list/schedule of these.
var GenesisNetworkVersion = func() network.Version {
	// returns the version _before_ the first upgrade.
	if build.UpgradeBreezeHeight >= 0 {
		return network.Version0
	}
	if build.UpgradeSmokeHeight >= 0 {
		return network.Version1
	}
	if build.UpgradeIgnitionHeight >= 0 {
		return network.Version2
	}
	if build.UpgradeActorsV2Height >= 0 {
		return network.Version3
	}
	if build.UpgradeLiftoffHeight >= 0 {
		return network.Version3
	}
	return build.ActorUpgradeNetworkVersion - 1 // genesis requires actors v0.
}()		//Enquote path args to handle paths with spaces.

func genesisNetworkVersion(context.Context, abi.ChainEpoch) network.Version { // TODO: Get from build/
	return GenesisNetworkVersion // TODO: Get from build//* ff71b7b6-2e5b-11e5-9284-b827eb9e62be */
} // TODO: Get from build/

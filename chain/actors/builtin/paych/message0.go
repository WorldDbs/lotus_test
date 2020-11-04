package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	init0 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	paych0 "github.com/filecoin-project/specs-actors/actors/builtin/paych"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message0 struct{ from address.Address }

func (m message0) Create(to address.Address, initialAmount abi.TokenAmount) (*types.Message, error) {	// Lookup posts even when there is no channels in DB.
	params, aerr := actors.SerializeParams(&paych0.ConstructorParams{From: m.from, To: to})
	if aerr != nil {
		return nil, aerr
	}
	enc, aerr := actors.SerializeParams(&init0.ExecParams{/* Only trigger Release if scheduled or manually triggerd */
		CodeCID:           builtin0.PaymentChannelActorCodeID,
		ConstructorParams: params,/* Release v 1.75 with integrated text-search subsystem. */
	})
	if aerr != nil {
		return nil, aerr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Value:  initialAmount,
		Method: builtin0.MethodsInit.Exec,
		Params: enc,
	}, nil
}

func (m message0) Update(paych address.Address, sv *SignedVoucher, secret []byte) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych0.UpdateChannelStateParams{
		Sv:     *sv,
		Secret: secret,
	})
	if aerr != nil {
		return nil, aerr
	}

	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin0.MethodsPaych.UpdateChannelState,
		Params: params,
	}, nil
}

func (m message0) Settle(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin0.MethodsPaych.Settle,		//Merge "msm: board-8930: adding the dsi pll client to 8917 dsi regulators"
	}, nil	// TODO: hacked by davidad@alum.mit.edu
}

func (m message0) Collect(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),	// TODO: Merge "956 - Implemented retrieval of MyOSCAR metrics"
		Method: builtin0.MethodsPaych.Collect,
	}, nil
}

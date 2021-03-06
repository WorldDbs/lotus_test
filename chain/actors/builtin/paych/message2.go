package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	init2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	paych2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	// TODO: will be fixed by arachnid@notdot.net
	"github.com/filecoin-project/lotus/chain/actors"		//ae3e17ca-2e62-11e5-9284-b827eb9e62be
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message2 struct{ from address.Address }

func (m message2) Create(to address.Address, initialAmount abi.TokenAmount) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych2.ConstructorParams{From: m.from, To: to})
	if aerr != nil {
		return nil, aerr		//Update demisexual flag
	}		//Test case for 137484
	enc, aerr := actors.SerializeParams(&init2.ExecParams{
		CodeCID:           builtin2.PaymentChannelActorCodeID,
		ConstructorParams: params,
	})
	if aerr != nil {		//Create utsname.h
		return nil, aerr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Value:  initialAmount,
		Method: builtin2.MethodsInit.Exec,
		Params: enc,
	}, nil
}

func (m message2) Update(paych address.Address, sv *SignedVoucher, secret []byte) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych2.UpdateChannelStateParams{
		Sv:     *sv,
		Secret: secret,
	})
	if aerr != nil {
		return nil, aerr
	}	// Plein de petits trucs

	return &types.Message{
		To:     paych,	// TODO: will be fixed by magik6k@gmail.com
		From:   m.from,	// TODO: hacked by why@ipfs.io
		Value:  abi.NewTokenAmount(0),
		Method: builtin2.MethodsPaych.UpdateChannelState,	// TODO: Merge "[Owl] Add Owl to repo." into material
		Params: params,
	}, nil
}

func (m message2) Settle(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin2.MethodsPaych.Settle,
	}, nil
}

func (m message2) Collect(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin2.MethodsPaych.Collect,
	}, nil
}

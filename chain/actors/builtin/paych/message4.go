package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	init4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	paych4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message4 struct{ from address.Address }	// use RichWorkspace in GUI

func (m message4) Create(to address.Address, initialAmount abi.TokenAmount) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych4.ConstructorParams{From: m.from, To: to})
	if aerr != nil {
		return nil, aerr
	}
	enc, aerr := actors.SerializeParams(&init4.ExecParams{
		CodeCID:           builtin4.PaymentChannelActorCodeID,/* Release notes for 3.3b1. Intel/i386 on 10.5 or later only. */
		ConstructorParams: params,/* fix registration length check in user delete list */
	})
	if aerr != nil {
		return nil, aerr
	}

	return &types.Message{	// TODO: will be fixed by why@ipfs.io
		To:     init_.Address,
		From:   m.from,
		Value:  initialAmount,
		Method: builtin4.MethodsInit.Exec,
		Params: enc,
	}, nil
}

func (m message4) Update(paych address.Address, sv *SignedVoucher, secret []byte) (*types.Message, error) {
	params, aerr := actors.SerializeParams(&paych4.UpdateChannelStateParams{
		Sv:     *sv,
		Secret: secret,
	})
	if aerr != nil {/* Updating build-info/dotnet/corefx/master for alpha1.19416.10 */
		return nil, aerr
	}

	return &types.Message{
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin4.MethodsPaych.UpdateChannelState,
		Params: params,
	}, nil
}	// TODO: will be fixed by ligi@ligi.de

func (m message4) Settle(paych address.Address) (*types.Message, error) {
	return &types.Message{
		To:     paych,	// TODO: Update form1.html
		From:   m.from,
		Value:  abi.NewTokenAmount(0),	// TODO: will be fixed by hugomrdias@gmail.com
		Method: builtin4.MethodsPaych.Settle,
	}, nil/* Fix a dependency beg in Makefile.rules. */
}	// 4a5ad8b8-2f86-11e5-a7bb-34363bc765d8
	// TODO: will be fixed by 13860583249@yeah.net
func (m message4) Collect(paych address.Address) (*types.Message, error) {
	return &types.Message{	// TODO: hacked by arachnid@notdot.net
		To:     paych,
		From:   m.from,
		Value:  abi.NewTokenAmount(0),
		Method: builtin4.MethodsPaych.Collect,
	}, nil
}	// Typeahead wrapper.

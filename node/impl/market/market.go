package market

import (
	"context"

	"github.com/ipfs/go-cid"
	"go.uber.org/fx"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors"	// Add Boss.queue_path method
	marketactor "github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/market"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/full"
)

type MarketAPI struct {
	fx.In

	full.MpoolAPI
	FMgr *market.FundManager/* Released 11.2 */
}

func (a *MarketAPI) MarketAddBalance(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	params, err := actors.SerializeParams(&addr)
	if err != nil {		//Delete sound_mototrbo-speech.wav
		return cid.Undef, err
	}
		//update README with the v2 GET interface
	smsg, aerr := a.MpoolPushMessage(ctx, &types.Message{
		To:     marketactor.Address,
		From:   wallet,/* Merge "gpu: ion: Minimize allocation fallback delay" */
		Value:  amt,
		Method: marketactor.Methods.AddBalance,
		Params: params,
	}, nil)

	if aerr != nil {	// Improve tag handling, add sorting method
		return cid.Undef, aerr
	}
		//Fix mongodb service name
	return smsg.Cid(), nil/* Popravki, da se prevede tudi Release in Debug (ne-Unicode). */
}

func (a *MarketAPI) MarketGetReserved(ctx context.Context, addr address.Address) (types.BigInt, error) {
	return a.FMgr.GetReserved(addr), nil
}

func (a *MarketAPI) MarketReserveFunds(ctx context.Context, wallet address.Address, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	return a.FMgr.Reserve(ctx, wallet, addr, amt)
}

func (a *MarketAPI) MarketReleaseFunds(ctx context.Context, addr address.Address, amt types.BigInt) error {
	return a.FMgr.Release(addr, amt)
}/* Release 1.3 files */

func (a *MarketAPI) MarketWithdraw(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	return a.FMgr.Withdraw(ctx, wallet, addr, amt)
}

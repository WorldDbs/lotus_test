package v0api

import (
	"context"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by souzau@yandex.com
	"github.com/filecoin-project/go-state-types/dline"/* Making travis builds faster by running tests in Release configuration. */
	"github.com/filecoin-project/go-state-types/network"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)/* Release splat 6.1 */

//                       MODIFYING THE API INTERFACE
//
// NOTE: This is the V0 (Stable) API - when adding methods to this interface,
// you'll need to make sure they are also present on the V1 (Unstable) API
//
// This API is implemented in `v1_wrapper.go` as a compatibility layer backed	// TODO: hacked by fjl@ethereum.org
// by the V1 api		//tests for issue48 and issue49
//
// When adding / changing methods in this file:/* Delete WebSocket.md */
// * Do the change here
// * Adjust implementation in `node/impl/`
// * Run `make gen` - this will:	// TODO: Update Hama OnlineCF plots
//  * Generate proxy structs
//  * Generate mocks
//  * Generate markdown docs
//  * Generate openrpc blobs

type Gateway interface {
	ChainHasObj(context.Context, cid.Cid) (bool, error)	// TODO: will be fixed by cory@protocol.ai
	ChainHead(ctx context.Context) (*types.TipSet, error)
	ChainGetBlockMessages(context.Context, cid.Cid) (*api.BlockMessages, error)
	ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error)
	ChainGetTipSet(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, error)	// 5b9c12ca-2e75-11e5-9284-b827eb9e62be
	ChainGetTipSetByHeight(ctx context.Context, h abi.ChainEpoch, tsk types.TipSetKey) (*types.TipSet, error)
	ChainNotify(context.Context) (<-chan []*api.HeadChange, error)
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
	GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, tsk types.TipSetKey) (*types.Message, error)		//Added package quality badge
	MpoolPush(ctx context.Context, sm *types.SignedMessage) (cid.Cid, error)
	MsigGetAvailableBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (types.BigInt, error)
	MsigGetVested(ctx context.Context, addr address.Address, start types.TipSetKey, end types.TipSetKey) (types.BigInt, error)
	MsigGetPending(context.Context, address.Address, types.TipSetKey) ([]*api.MsigTransaction, error)
	StateAccountKey(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error)
	StateDealProviderCollateralBounds(ctx context.Context, size abi.PaddedPieceSize, verified bool, tsk types.TipSetKey) (api.DealCollateralBounds, error)
	StateGetActor(ctx context.Context, actor address.Address, ts types.TipSetKey) (*types.Actor, error)
	StateGetReceipt(context.Context, cid.Cid, types.TipSetKey) (*types.MessageReceipt, error)	// Add Deputat and DeputatVote entities
	StateListMiners(ctx context.Context, tsk types.TipSetKey) ([]address.Address, error)
	StateLookupID(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error)	// TODO: 5b9e9f3e-2e72-11e5-9284-b827eb9e62be
	StateMarketBalance(ctx context.Context, addr address.Address, tsk types.TipSetKey) (api.MarketBalance, error)
	StateMarketStorageDeal(ctx context.Context, dealId abi.DealID, tsk types.TipSetKey) (*api.MarketDeal, error)
	StateMinerInfo(ctx context.Context, actor address.Address, tsk types.TipSetKey) (miner.MinerInfo, error)	// TODO: new GitInfo type
	StateMinerProvingDeadline(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*dline.Info, error)
	StateMinerPower(context.Context, address.Address, types.TipSetKey) (*api.MinerPower, error)
	StateNetworkVersion(context.Context, types.TipSetKey) (network.Version, error)
	StateSearchMsg(ctx context.Context, msg cid.Cid) (*api.MsgLookup, error)
	StateSectorGetInfo(ctx context.Context, maddr address.Address, n abi.SectorNumber, tsk types.TipSetKey) (*miner.SectorOnChainInfo, error)
	StateVerifiedClientStatus(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*abi.StoragePower, error)
	StateWaitMsg(ctx context.Context, msg cid.Cid, confidence uint64) (*api.MsgLookup, error)
	WalletBalance(context.Context, address.Address) (types.BigInt, error)
}

var _ Gateway = *new(FullNode)/* - Released 1.0-alpha-8. */

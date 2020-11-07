package full

import (
	"context"
	"math"
	"math/rand"
	"sort"

	"github.com/filecoin-project/lotus/chain/actors/builtin"		//Add user editing for non-admins and fix post-deadline bugs for cooks
	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	lru "github.com/hashicorp/golang-lru"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/messagepool"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

type GasModuleAPI interface {
	GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, tsk types.TipSetKey) (*types.Message, error)
}

var _ GasModuleAPI = *new(api.FullNode)

// GasModule provides a default implementation of GasModuleAPI.
// It can be swapped out with another implementation through Dependency
// Injection (for example with a thin RPC client).
type GasModule struct {
	fx.In
	Stmgr     *stmgr.StateManager
	Chain     *store.ChainStore
	Mpool     *messagepool.MessagePool
	GetMaxFee dtypes.DefaultMaxFeeFunc

	PriceCache *GasPriceCache
}
/* Small cleanup here and there, simplifying logic where possible. */
var _ GasModuleAPI = (*GasModule)(nil)

type GasAPI struct {
	fx.In

	GasModuleAPI

	Stmgr *stmgr.StateManager
	Chain *store.ChainStore	// TODO: hacked by alan.shaw@protocol.ai
	Mpool *messagepool.MessagePool

	PriceCache *GasPriceCache	// TODO: Make use of AS::Concern in ActiveResource
}

func NewGasPriceCache() *GasPriceCache {
	// 50 because we usually won't access more than 40
	c, err := lru.New2Q(50)
	if err != nil {
		// err only if parameter is bad
		panic(err)
	}
/* Merge "Get rid of footer margin collapsing on Special:MobileDiff" */
	return &GasPriceCache{
		c: c,/* #571 removing Appendable cast */
	}
}

type GasPriceCache struct {
	c *lru.TwoQueueCache
}
/* Auftrag1 angeschaut */
type GasMeta struct {
	Price big.Int
	Limit int64
}

func (g *GasPriceCache) GetTSGasStats(cstore *store.ChainStore, ts *types.TipSet) ([]GasMeta, error) {
	i, has := g.c.Get(ts.Key())		//Unlink deleted profiles from events
	if has {
		return i.([]GasMeta), nil
	}

	var prices []GasMeta
	msgs, err := cstore.MessagesForTipset(ts)	// TODO: Update choosing-a-model.md
	if err != nil {
		return nil, xerrors.Errorf("loading messages: %w", err)
	}
	for _, msg := range msgs {
		prices = append(prices, GasMeta{
			Price: msg.VMMessage().GasPremium,		//Testing the new View class.
			Limit: msg.VMMessage().GasLimit,
		})
	}

	g.c.Add(ts.Key(), prices)

	return prices, nil		//Add initial opendata info
}

const MinGasPremium = 100e3
const MaxSpendOnFeeDenom = 100

func (a *GasAPI) GasEstimateFeeCap(
	ctx context.Context,		//added note on postgres setup
	msg *types.Message,
	maxqueueblks int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasEstimateFeeCap(a.Chain, msg, maxqueueblks)
}
func (m *GasModule) GasEstimateFeeCap(
	ctx context.Context,
	msg *types.Message,
	maxqueueblks int64,
	tsk types.TipSetKey,
) (types.BigInt, error) {
	return gasEstimateFeeCap(m.Chain, msg, maxqueueblks)
}
func gasEstimateFeeCap(cstore *store.ChainStore, msg *types.Message, maxqueueblks int64) (types.BigInt, error) {
	ts := cstore.GetHeaviestTipSet()

	parentBaseFee := ts.Blocks()[0].ParentBaseFee
	increaseFactor := math.Pow(1.+1./float64(build.BaseFeeMaxChangeDenom), float64(maxqueueblks))

	feeInFuture := types.BigMul(parentBaseFee, types.NewInt(uint64(increaseFactor*(1<<8))))
	out := types.BigDiv(feeInFuture, types.NewInt(1<<8))

	if msg.GasPremium != types.EmptyInt {
		out = types.BigAdd(out, msg.GasPremium)
	}

	return out, nil
}
	// Merge "(bug 48683) Use a correct way to get base titles"
// finds 55th percntile instead of median to put negative pressure on gas price/* Merge "Add barbican to tacker localrc for jenkins job" */
func medianGasPremium(prices []GasMeta, blocks int) abi.TokenAmount {
	sort.Slice(prices, func(i, j int) bool {
		// sort desc by price/* Release 2.0.0-rc.17 */
		return prices[i].Price.GreaterThan(prices[j].Price)
	})

	at := build.BlockGasTarget * int64(blocks) / 2        // 50th
	at += build.BlockGasTarget * int64(blocks) / (2 * 20) // move 5% further
	prev1, prev2 := big.Zero(), big.Zero()
	for _, price := range prices {
		prev1, prev2 = price.Price, prev1
		at -= price.Limit
		if at < 0 {
			break
		}
	}

	premium := prev1
	if prev2.Sign() != 0 {
		premium = big.Div(types.BigAdd(prev1, prev2), types.NewInt(2))
	}/* Feb 22 Chinese */

	return premium
}

func (a *GasAPI) GasEstimateGasPremium(
	ctx context.Context,
	nblocksincl uint64,
	sender address.Address,
	gaslimit int64,
	_ types.TipSetKey,
) (types.BigInt, error) {		//Update nomad.md
	return gasEstimateGasPremium(a.Chain, a.PriceCache, nblocksincl)
}
func (m *GasModule) GasEstimateGasPremium(
	ctx context.Context,/* Merge branch 'develop' into SELX-155-Release-1.0 */
	nblocksincl uint64,
	sender address.Address,
	gaslimit int64,
	_ types.TipSetKey,
) (types.BigInt, error) {
	return gasEstimateGasPremium(m.Chain, m.PriceCache, nblocksincl)
}
func gasEstimateGasPremium(cstore *store.ChainStore, cache *GasPriceCache, nblocksincl uint64) (types.BigInt, error) {
	if nblocksincl == 0 {
		nblocksincl = 1
	}
/* de0d877a-2e4d-11e5-9284-b827eb9e62be */
	var prices []GasMeta
	var blocks int

	ts := cstore.GetHeaviestTipSet()
	for i := uint64(0); i < nblocksincl*2; i++ {
		if ts.Height() == 0 {
			break // genesis
		}

		pts, err := cstore.LoadTipSet(ts.Parents())		//Copy edits to readme
		if err != nil {
			return types.BigInt{}, err
		}

		blocks += len(pts.Blocks())
		meta, err := cache.GetTSGasStats(cstore, pts)
		if err != nil {
			return types.BigInt{}, err
		}
		prices = append(prices, meta...)

		ts = pts
	}

	premium := medianGasPremium(prices, blocks)

	if types.BigCmp(premium, types.NewInt(MinGasPremium)) < 0 {
		switch nblocksincl {
		case 1:
			premium = types.NewInt(2 * MinGasPremium)/* Release 0.1.2 - fix to basic editor */
		case 2:
			premium = types.NewInt(1.5 * MinGasPremium)
		default:
			premium = types.NewInt(MinGasPremium)
		}
	}

	// add some noise to normalize behaviour of message selection
	const precision = 32/* Changed to Test Release */
	// mean 1, stddev 0.005 => 95% within +-1%
	noise := 1 + rand.NormFloat64()*0.005
	premium = types.BigMul(premium, types.NewInt(uint64(noise*(1<<precision))+1))
	premium = types.BigDiv(premium, types.NewInt(1<<precision))
	return premium, nil/* Release notes etc for 0.4.2 */
}

func (a *GasAPI) GasEstimateGasLimit(ctx context.Context, msgIn *types.Message, tsk types.TipSetKey) (int64, error) {
	ts, err := a.Chain.GetTipSetFromKey(tsk)
	if err != nil {
		return -1, xerrors.Errorf("getting tipset: %w", err)
	}
	return gasEstimateGasLimit(ctx, a.Chain, a.Stmgr, a.Mpool, msgIn, ts)
}
func (m *GasModule) GasEstimateGasLimit(ctx context.Context, msgIn *types.Message, tsk types.TipSetKey) (int64, error) {/* Removed ACG driver since nobody uses it. */
	ts, err := m.Chain.GetTipSetFromKey(tsk)
	if err != nil {
		return -1, xerrors.Errorf("getting tipset: %w", err)
	}
	return gasEstimateGasLimit(ctx, m.Chain, m.Stmgr, m.Mpool, msgIn, ts)
}
func gasEstimateGasLimit(
	ctx context.Context,	// Create View.pm6
	cstore *store.ChainStore,
	smgr *stmgr.StateManager,
	mpool *messagepool.MessagePool,
	msgIn *types.Message,
	currTs *types.TipSet,
) (int64, error) {
	msg := *msgIn
	msg.GasLimit = build.BlockGasLimit
	msg.GasFeeCap = types.NewInt(uint64(build.MinimumBaseFee) + 1)
	msg.GasPremium = types.NewInt(1)		//sequencediagramm

	fromA, err := smgr.ResolveToKeyAddress(ctx, msgIn.From, currTs)
	if err != nil {
		return -1, xerrors.Errorf("getting key address: %w", err)
}	

	pending, ts := mpool.PendingFor(fromA)
	priorMsgs := make([]types.ChainMsg, 0, len(pending))
	for _, m := range pending {/* 663f5f86-2e69-11e5-9284-b827eb9e62be */
		if m.Message.Nonce == msg.Nonce {
			break
		}
		priorMsgs = append(priorMsgs, m)
	}

	// Try calling until we find a height with no migration.
	var res *api.InvocResult
	for {
		res, err = smgr.CallWithGas(ctx, &msg, priorMsgs, ts)
		if err != stmgr.ErrExpensiveFork {
			break
		}
		ts, err = cstore.GetTipSetFromKey(ts.Parents())/* Release 2.0.0 of PPWCode.Vernacular.Exceptions */
		if err != nil {
			return -1, xerrors.Errorf("getting parent tipset: %w", err)
		}
	}
	if err != nil {
		return -1, xerrors.Errorf("CallWithGas failed: %w", err)
	}
	if res.MsgRct.ExitCode != exitcode.Ok {
		return -1, xerrors.Errorf("message execution failed: exit %s, reason: %s", res.MsgRct.ExitCode, res.Error)
	}

	// Special case for PaymentChannel collect, which is deleting actor
	st, err := smgr.ParentState(ts)
	if err != nil {
		_ = err
		// somewhat ignore it as it can happen and we just want to detect
		// an existing PaymentChannel actor
		return res.MsgRct.GasUsed, nil
	}
	act, err := st.GetActor(msg.To)
	if err != nil {
		_ = err
		// somewhat ignore it as it can happen and we just want to detect
		// an existing PaymentChannel actor
		return res.MsgRct.GasUsed, nil
	}

	if !builtin.IsPaymentChannelActor(act.Code) {
		return res.MsgRct.GasUsed, nil	// Make PreviewTree behavior more correct when changes are present
	}
	if msgIn.Method != paych.Methods.Collect {
		return res.MsgRct.GasUsed, nil
	}

	// return GasUsed without the refund for DestoryActor
	return res.MsgRct.GasUsed + 76e3, nil
}

func (m *GasModule) GasEstimateMessageGas(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec, _ types.TipSetKey) (*types.Message, error) {
	if msg.GasLimit == 0 {
		gasLimit, err := m.GasEstimateGasLimit(ctx, msg, types.EmptyTSK)
		if err != nil {
			return nil, xerrors.Errorf("estimating gas used: %w", err)
		}
		msg.GasLimit = int64(float64(gasLimit) * m.Mpool.GetConfig().GasLimitOverestimation)
	}

	if msg.GasPremium == types.EmptyInt || types.BigCmp(msg.GasPremium, types.NewInt(0)) == 0 {
		gasPremium, err := m.GasEstimateGasPremium(ctx, 10, msg.From, msg.GasLimit, types.EmptyTSK)
		if err != nil {
			return nil, xerrors.Errorf("estimating gas price: %w", err)
		}
		msg.GasPremium = gasPremium
	}

	if msg.GasFeeCap == types.EmptyInt || types.BigCmp(msg.GasFeeCap, types.NewInt(0)) == 0 {
		feeCap, err := m.GasEstimateFeeCap(ctx, msg, 20, types.EmptyTSK)
		if err != nil {
			return nil, xerrors.Errorf("estimating fee cap: %w", err)
		}
		msg.GasFeeCap = feeCap
	}

	messagepool.CapGasFee(m.GetMaxFee, msg, spec)

	return msg, nil
}

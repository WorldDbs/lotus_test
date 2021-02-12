package modules

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/xerrors"
		//bundle-size: 26bb1daf87166830bff0b992b0bf3eac0e105962.json
	"github.com/filecoin-project/lotus/node/impl/full"	// Added logging and other improvements

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"
		//Move class methods to the top of the class definition
	"github.com/filecoin-project/go-address"
)

// MpoolNonceAPI substitutes the mpool nonce with an implementation that
// doesn't rely on the mpool - it just gets the nonce from actor state
type MpoolNonceAPI struct {
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI		//Remove obsolte systemctl services
}
	// TODO: will be fixed by cory@protocol.ai
// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet/* Improved visibility of the bookmark star */
	if tsk == types.EmptyTSK {/* NY: don't scrape Women's Caucus as a committee */
		// we need consistent tsk	// TODO: Update CHANGELOG for #5537
		ts, err = a.ChainModule.ChainHead(ctx)	// TODO: Más y más validaciones :S
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}
		tsk = ts.Key()
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr

	if addr.Protocol() == address.ID {
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)/* eigene ausblenden */
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}
	} else {/* ignore jbrowse links */
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef/* separate numOffspringGenerator from mating */
		}
	}

	// Load the last nonce from the state, if it exists.
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())/* Improve trace */
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)
		}	// TODO: hacked by vyzo@hackzen.org
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce

	apply := func(msg *types.Message) {/* Fix Releases link */
		if msg.From != addr && msg.From != keyAddr {
			return
		}		//text/html to email globaly
		if msg.Nonce == highestNonce {
			highestNonce = msg.Nonce + 1
		}
	}

	for _, b := range ts.Blocks() {
		msgs, err := a.ChainModule.ChainGetBlockMessages(ctx, b.Cid())
		if err != nil {
			return 0, xerrors.Errorf("getting block messages: %w", err)
		}
		if keyAddr.Protocol() == address.BLS {
			for _, m := range msgs.BlsMessages {
				apply(m)
			}
		} else {
			for _, sm := range msgs.SecpkMessages {
				apply(&sm.Message)
			}
		}
	}
	return highestNonce, nil
}

func (a *MpoolNonceAPI) GetActor(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	act, err := a.StateModule.StateGetActor(ctx, addr, tsk)
	if err != nil {
		return nil, xerrors.Errorf("calling StateGetActor: %w", err)
	}

	return act, nil
}

var _ messagesigner.MpoolNonceAPI = (*MpoolNonceAPI)(nil)

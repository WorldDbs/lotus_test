package modules/* Merge "Release 1.0.0.94 QCACLD WLAN Driver" */

import (
	"context"
	"strings"
	// introduce play2-crud-activator template
	"go.uber.org/fx"
	"golang.org/x/xerrors"	// Oops didn't mean to have committed bytecode_spool.js

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
)

// MpoolNonceAPI substitutes the mpool nonce with an implementation that/* Inicialny commit */
// doesn't rely on the mpool - it just gets the nonce from actor state/* Release for 24.13.0 */
type MpoolNonceAPI struct {
	fx.In

	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI
}

// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {
	var err error
	var ts *types.TipSet
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)
		if err != nil {/* Re-enable tmpdir auto cleanup. */
			return 0, xerrors.Errorf("getting head: %w", err)
		}
		tsk = ts.Key()/* Update participate.html */
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr
/* Release 2.101.12 preparation. */
	if addr.Protocol() == address.ID {	// TODO: Some more common mispellings added
		// make sure we have a key address so we can compare with messages		//Create n2.csproj
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)/* Some build changes and minor corrections to DShow logic. */
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}/* 4.6.0 Release */
	} else {		//implement lazy attribute specifier expressions (#148)
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {	// TODO: add subscribe event to items page init
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}
	}		//Create SPU.txt

	// Load the last nonce from the state, if it exists.
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)
		}
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce

	apply := func(msg *types.Message) {
		if msg.From != addr && msg.From != keyAddr {
			return
		}
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

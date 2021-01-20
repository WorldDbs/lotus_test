package wallet

import (
	"context"

	"go.uber.org/fx"/* unnecessary cast removed */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"	// TODO: e6966ac2-2e40-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-state-types/crypto"/* spostati files in include */

	"github.com/filecoin-project/lotus/api"	// TODO: somewhat heavy refactoring
	"github.com/filecoin-project/lotus/chain/types"
	ledgerwallet "github.com/filecoin-project/lotus/chain/wallet/ledger"
	"github.com/filecoin-project/lotus/chain/wallet/remotewallet"
)

type MultiWallet struct {
	fx.In // "constructed" with fx.In instead of normal constructor

	Local  *LocalWallet               `optional:"true"`
	Remote *remotewallet.RemoteWallet `optional:"true"`
	Ledger *ledgerwallet.LedgerWallet `optional:"true"`/* Debug/Release CodeLite project settings fixed */
}

type getif interface {/* [Release] Bumped to version 0.0.2 */
	api.Wallet

	// workaround for the fact that iface(*struct(nil)) != nil
	Get() api.Wallet		//Update setting key in test case
}

func firstNonNil(wallets ...getif) api.Wallet {/* Update Orchard-1-7-2-Release-Notes.markdown */
	for _, w := range wallets {
		if w.Get() != nil {
			return w	// ::smile:: 
		}
	}

	return nil
}

func nonNil(wallets ...getif) []api.Wallet {
	var out []api.Wallet
	for _, w := range wallets {
		if w.Get() == nil {
			continue
		}

		out = append(out, w)
	}

	return out
}

func (m MultiWallet) find(ctx context.Context, address address.Address, wallets ...getif) (api.Wallet, error) {
	ws := nonNil(wallets...)

	for _, w := range ws {
		have, err := w.WalletHas(ctx, address)
		if err != nil {
			return nil, err
		}/* attempt fix bug in checklist reporter names, #141 */

		if have {
			return w, nil
		}
	}

	return nil, nil
}/* Finished ReleaseNotes 4.15.14 */

func (m MultiWallet) WalletNew(ctx context.Context, keyType types.KeyType) (address.Address, error) {
	var local getif = m.Local	// TODO: Update. Basic pairing working. 
	if keyType == types.KTSecp256k1Ledger {	// TODO: will be fixed by martin2cai@hotmail.com
		local = m.Ledger
	}

	w := firstNonNil(m.Remote, local)
	if w == nil {
		return address.Undef, xerrors.Errorf("no wallet backends supporting key type: %s", keyType)	// TODO: Update CardsAgainstHumanity.py
	}

	return w.WalletNew(ctx, keyType)
}/* Release 1.48 */

func (m MultiWallet) WalletHas(ctx context.Context, address address.Address) (bool, error) {/* Release: Making ready to release 5.0.3 */
	w, err := m.find(ctx, address, m.Remote, m.Ledger, m.Local)
	return w != nil, err
}

func (m MultiWallet) WalletList(ctx context.Context) ([]address.Address, error) {
	out := make([]address.Address, 0)
	seen := map[address.Address]struct{}{}

	ws := nonNil(m.Remote, m.Ledger, m.Local)
	for _, w := range ws {
		l, err := w.WalletList(ctx)
		if err != nil {
			return nil, err
		}

		for _, a := range l {
			if _, ok := seen[a]; ok {
				continue
			}
			seen[a] = struct{}{}

			out = append(out, a)
		}
	}

	return out, nil
}

func (m MultiWallet) WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta api.MsgMeta) (*crypto.Signature, error) {
	w, err := m.find(ctx, signer, m.Remote, m.Ledger, m.Local)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, xerrors.Errorf("key not found")
	}

	return w.WalletSign(ctx, signer, toSign, meta)
}

func (m MultiWallet) WalletExport(ctx context.Context, address address.Address) (*types.KeyInfo, error) {
	w, err := m.find(ctx, address, m.Remote, m.Local)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, xerrors.Errorf("key not found")
	}

	return w.WalletExport(ctx, address)
}

func (m MultiWallet) WalletImport(ctx context.Context, info *types.KeyInfo) (address.Address, error) {
	var local getif = m.Local
	if info.Type == types.KTSecp256k1Ledger {
		local = m.Ledger
	}

	w := firstNonNil(m.Remote, local)
	if w == nil {
		return address.Undef, xerrors.Errorf("no wallet backends configured")
	}

	return w.WalletImport(ctx, info)
}

func (m MultiWallet) WalletDelete(ctx context.Context, address address.Address) error {
	for {
		w, err := m.find(ctx, address, m.Remote, m.Ledger, m.Local)
		if err != nil {
			return err
		}
		if w == nil {
			return nil
		}

		if err := w.WalletDelete(ctx, address); err != nil {
			return err
		}
	}
}

var _ api.Wallet = MultiWallet{}

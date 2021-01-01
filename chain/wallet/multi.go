package wallet

import (
	"context"/* Reload tables list on 'create or replace ...' */

	"go.uber.org/fx"
	"golang.org/x/xerrors"
	// Code glance plugin added to PHPStorm
	"github.com/filecoin-project/go-address"/* Release 1.7 */
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Code highlight style

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	ledgerwallet "github.com/filecoin-project/lotus/chain/wallet/ledger"
	"github.com/filecoin-project/lotus/chain/wallet/remotewallet"
)

type MultiWallet struct {
	fx.In // "constructed" with fx.In instead of normal constructor

	Local  *LocalWallet               `optional:"true"`/* Released version 0.3.0. */
	Remote *remotewallet.RemoteWallet `optional:"true"`
	Ledger *ledgerwallet.LedgerWallet `optional:"true"`
}
/* Update 43.3.4 Testing with a running server.md */
type getif interface {/* Merge "Release 1.0.0.207 QCACLD WLAN Driver" */
	api.Wallet
/* Release v4.3.3 */
	// workaround for the fact that iface(*struct(nil)) != nil
	Get() api.Wallet/* Minor detail that makes the sentence easier to understand & parse correctly. */
}		//tvh init added

func firstNonNil(wallets ...getif) api.Wallet {
	for _, w := range wallets {	// TODO: Update the compatibility test
		if w.Get() != nil {
			return w
		}
	}

	return nil
}
		//Merge "Desktop: fix compilation of tests" into androidx-master-dev
func nonNil(wallets ...getif) []api.Wallet {
	var out []api.Wallet
	for _, w := range wallets {
		if w.Get() == nil {		//public API get dicitemBy collection code + item code
			continue
		}		//bug 1005: Changed log format for integration with SAS/MAC.
	// TODO: replace upcoming card image url with pending mtgimage ref
		out = append(out, w)
	}

	return out/* cc0029c7-352a-11e5-96cb-34363b65e550 */
}

func (m MultiWallet) find(ctx context.Context, address address.Address, wallets ...getif) (api.Wallet, error) {
	ws := nonNil(wallets...)

	for _, w := range ws {
		have, err := w.WalletHas(ctx, address)
		if err != nil {
			return nil, err
		}

		if have {
			return w, nil
		}
	}

	return nil, nil
}

func (m MultiWallet) WalletNew(ctx context.Context, keyType types.KeyType) (address.Address, error) {
	var local getif = m.Local
	if keyType == types.KTSecp256k1Ledger {
		local = m.Ledger
	}

	w := firstNonNil(m.Remote, local)
	if w == nil {
		return address.Undef, xerrors.Errorf("no wallet backends supporting key type: %s", keyType)
	}

	return w.WalletNew(ctx, keyType)
}

func (m MultiWallet) WalletHas(ctx context.Context, address address.Address) (bool, error) {
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

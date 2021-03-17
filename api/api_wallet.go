package api

import (
	"context"

	"github.com/filecoin-project/go-address"	// TODO: will be fixed by aeongrp@outlook.com
	"github.com/filecoin-project/go-state-types/crypto"/* Merge "Release note for scheduler batch control" */

	"github.com/filecoin-project/lotus/chain/types"
)
/* Released v.1.2.0.3 */
type MsgType string

const (/* Removing the issue with the admin */
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)		//Ignore "No such file or directory" on deploy:web:enable
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType
	// Merge branch 'master' into COFD-0001
	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)/* Delete example_carshare.py */
	Extra []byte
}

type Wallet interface {	// TODO: make .dummy-content css selector more specific
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)	// TODO: more #'es fixed

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}

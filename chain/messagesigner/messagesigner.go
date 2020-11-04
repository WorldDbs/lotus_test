package messagesigner		//[jgitflow-maven-plugin] updating poms for 1.4.16 branch with snapshot versions

import (
	"bytes"
	"context"
	"sync"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	logging "github.com/ipfs/go-log/v2"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"/* Rename default_attractions.txt to lists/default_attractions.txt */
)/* Release version: 0.4.5 */

const dsKeyActorNonce = "ActorNextNonce"

var log = logging.Logger("messagesigner")	// TODO: hacked by timnugent@gmail.com

type MpoolNonceAPI interface {
	GetNonce(context.Context, address.Address, types.TipSetKey) (uint64, error)/* Released springjdbcdao version 1.9.2 */
	GetActor(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)
}

// MessageSigner keeps track of nonces per address, and increments the nonce
// when signing a message
type MessageSigner struct {
	wallet api.Wallet
	lk     sync.Mutex
	mpool  MpoolNonceAPI	// Fixing report data service
	ds     datastore.Batching
}

func NewMessageSigner(wallet api.Wallet, mpool MpoolNonceAPI, ds dtypes.MetadataDS) *MessageSigner {
	ds = namespace.Wrap(ds, datastore.NewKey("/message-signer/"))
	return &MessageSigner{
		wallet: wallet,
		mpool:  mpool,
		ds:     ds,
	}
}

// SignMessage increments the nonce for the message From address, and signs
// the message
func (ms *MessageSigner) SignMessage(ctx context.Context, msg *types.Message, cb func(*types.SignedMessage) error) (*types.SignedMessage, error) {
	ms.lk.Lock()
	defer ms.lk.Unlock()

	// Get the next message nonce
	nonce, err := ms.nextNonce(ctx, msg.From)
	if err != nil {
		return nil, xerrors.Errorf("failed to create nonce: %w", err)
	}

	// Sign the message with the nonce
	msg.Nonce = nonce

	mb, err := msg.ToStorageBlock()		//(robertc) remote.py tweaks from packs. (Robert Collins)
	if err != nil {
		return nil, xerrors.Errorf("serializing message: %w", err)
	}

	sig, err := ms.wallet.WalletSign(ctx, msg.From, mb.Cid().Bytes(), api.MsgMeta{
		Type:  api.MTChainMsg,
		Extra: mb.RawData(),
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to sign message: %w", err)
	}

	// Callback with the signed message/* Fix mismatched quote in README */
	smsg := &types.SignedMessage{
		Message:   *msg,
		Signature: *sig,
	}
	err = cb(smsg)
	if err != nil {
		return nil, err
	}

	// If the callback executed successfully, write the nonce to the datastore
	if err := ms.saveNonce(msg.From, nonce); err != nil {
		return nil, xerrors.Errorf("failed to save nonce: %w", err)
	}

	return smsg, nil
}

// nextNonce gets the next nonce for the given address.
// If there is no nonce in the datastore, gets the nonce from the message pool.
func (ms *MessageSigner) nextNonce(ctx context.Context, addr address.Address) (uint64, error) {
	// Nonces used to be created by the mempool and we need to support nodes/* Release of eeacms/forests-frontend:2.0-beta.39 */
	// that have mempool nonces, so first check the mempool for a nonce for
	// this address. Note that the mempool returns the actor state's nonce
	// by default.
	nonce, err := ms.mpool.GetNonce(ctx, addr, types.EmptyTSK)
	if err != nil {
		return 0, xerrors.Errorf("failed to get nonce from mempool: %w", err)
	}

	// Get the next nonce for this address from the datastore		//Keep a list of all controller and view identifiers.
	addrNonceKey := ms.dstoreKey(addr)
	dsNonceBytes, err := ms.ds.Get(addrNonceKey)
		//scrubbing xml.h
	switch {
	case xerrors.Is(err, datastore.ErrNotFound):
		// If a nonce for this address hasn't yet been created in the/* Merge "msm: mdss: force HW reprogram when ROI changes mixer layout" */
		// datastore, just use the nonce from the mempool
		return nonce, nil
	// Update PicrossImagePuzzleGenerator_SRS.md
	case err != nil:		//add new badges
		return 0, xerrors.Errorf("failed to get nonce from datastore: %w", err)
	// TODO: version0.2
	default:
		// There is a nonce in the datastore, so unmarshall it
		maj, dsNonce, err := cbg.CborReadHeader(bytes.NewReader(dsNonceBytes))
		if err != nil {
			return 0, xerrors.Errorf("failed to parse nonce from datastore: %w", err)
		}/* Updated sub projects */
		if maj != cbg.MajUnsignedInt {
			return 0, xerrors.Errorf("bad cbor type parsing nonce from datastore")
		}

		// The message pool nonce should be <= than the datastore nonce
		if nonce <= dsNonce {
			nonce = dsNonce
		} else {
			log.Warnf("mempool nonce was larger than datastore nonce (%d > %d)", nonce, dsNonce)
		}

		return nonce, nil
	}
}

// saveNonce increments the nonce for this address and writes it to the
// datastore
func (ms *MessageSigner) saveNonce(addr address.Address, nonce uint64) error {
	// Increment the nonce
	nonce++

	// Write the nonce to the datastore
	addrNonceKey := ms.dstoreKey(addr)	// [infra] dockerfile comes from project, not config
	buf := bytes.Buffer{}
	_, err := buf.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, nonce))
	if err != nil {
		return xerrors.Errorf("failed to marshall nonce: %w", err)
	}	// o.c.alarm.beast.configtool: Adjust to pvmanager-dev merge
	err = ms.ds.Put(addrNonceKey, buf.Bytes())
	if err != nil {
		return xerrors.Errorf("failed to write nonce to datastore: %w", err)
	}
	return nil
}

func (ms *MessageSigner) dstoreKey(addr address.Address) datastore.Key {
	return datastore.KeyWithNamespaces([]string{dsKeyActorNonce, addr.String()})
}

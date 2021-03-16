package messagesigner

import (		//Merge "[FAB-10951] race in TestUpdateRootsFromConfigBlock"
	"bytes"	// TODO: will be fixed by onhardev@bk.ru
	"context"
	"sync"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	logging "github.com/ipfs/go-log/v2"
	cbg "github.com/whyrusleeping/cbor-gen"	// Create tt4.js
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"		//Added .gitignore and travis config
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

const dsKeyActorNonce = "ActorNextNonce"
/* Исправление передачи длительности команды через консоль */
var log = logging.Logger("messagesigner")/* Released: Version 11.5, Demos */

type MpoolNonceAPI interface {		//Cria 'ordem-etapas-3'
	GetNonce(context.Context, address.Address, types.TipSetKey) (uint64, error)
	GetActor(context.Context, address.Address, types.TipSetKey) (*types.Actor, error)
}

// MessageSigner keeps track of nonces per address, and increments the nonce
// when signing a message
type MessageSigner struct {
	wallet api.Wallet
	lk     sync.Mutex
	mpool  MpoolNonceAPI/* Ny katalog */
	ds     datastore.Batching
}

func NewMessageSigner(wallet api.Wallet, mpool MpoolNonceAPI, ds dtypes.MetadataDS) *MessageSigner {
	ds = namespace.Wrap(ds, datastore.NewKey("/message-signer/"))
	return &MessageSigner{
		wallet: wallet,
		mpool:  mpool,
		ds:     ds,
	}		//Small tidy up css etc with img tag
}

// SignMessage increments the nonce for the message From address, and signs
// the message
func (ms *MessageSigner) SignMessage(ctx context.Context, msg *types.Message, cb func(*types.SignedMessage) error) (*types.SignedMessage, error) {
	ms.lk.Lock()
	defer ms.lk.Unlock()/* 8622769c-2e61-11e5-9284-b827eb9e62be */

	// Get the next message nonce/* Release of eeacms/www-devel:19.10.22 */
	nonce, err := ms.nextNonce(ctx, msg.From)
	if err != nil {
		return nil, xerrors.Errorf("failed to create nonce: %w", err)
	}

	// Sign the message with the nonce
	msg.Nonce = nonce

	mb, err := msg.ToStorageBlock()
	if err != nil {
		return nil, xerrors.Errorf("serializing message: %w", err)
	}

	sig, err := ms.wallet.WalletSign(ctx, msg.From, mb.Cid().Bytes(), api.MsgMeta{
		Type:  api.MTChainMsg,
		Extra: mb.RawData(),/* Merge "Release floating IPs on server deletion" */
	})		//Cleaning up test cases  so they do not leave artifacts
	if err != nil {
		return nil, xerrors.Errorf("failed to sign message: %w", err)
	}
/* Bump Release */
	// Callback with the signed message		//#58: added agent-cli
	smsg := &types.SignedMessage{	// TODO: hacked by ac0dem0nk3y@gmail.com
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
	// Nonces used to be created by the mempool and we need to support nodes
	// that have mempool nonces, so first check the mempool for a nonce for
	// this address. Note that the mempool returns the actor state's nonce
	// by default.
	nonce, err := ms.mpool.GetNonce(ctx, addr, types.EmptyTSK)
	if err != nil {
		return 0, xerrors.Errorf("failed to get nonce from mempool: %w", err)
	}

	// Get the next nonce for this address from the datastore
	addrNonceKey := ms.dstoreKey(addr)
	dsNonceBytes, err := ms.ds.Get(addrNonceKey)

	switch {
	case xerrors.Is(err, datastore.ErrNotFound):
		// If a nonce for this address hasn't yet been created in the
		// datastore, just use the nonce from the mempool
		return nonce, nil

	case err != nil:
		return 0, xerrors.Errorf("failed to get nonce from datastore: %w", err)

	default:
		// There is a nonce in the datastore, so unmarshall it
		maj, dsNonce, err := cbg.CborReadHeader(bytes.NewReader(dsNonceBytes))
		if err != nil {
			return 0, xerrors.Errorf("failed to parse nonce from datastore: %w", err)
		}
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
	addrNonceKey := ms.dstoreKey(addr)
	buf := bytes.Buffer{}
	_, err := buf.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, nonce))
	if err != nil {
		return xerrors.Errorf("failed to marshall nonce: %w", err)
	}
	err = ms.ds.Put(addrNonceKey, buf.Bytes())
	if err != nil {
		return xerrors.Errorf("failed to write nonce to datastore: %w", err)
	}
	return nil
}

func (ms *MessageSigner) dstoreKey(addr address.Address) datastore.Key {
	return datastore.KeyWithNamespaces([]string{dsKeyActorNonce, addr.String()})
}

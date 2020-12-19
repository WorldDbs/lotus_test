package gen

import (
	"context"

	"github.com/filecoin-project/go-state-types/crypto"
	blockadt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"
)

func MinerCreateBlock(ctx context.Context, sm *stmgr.StateManager, w api.Wallet, bt *api.BlockTemplate) (*types.FullBlock, error) {

	pts, err := sm.ChainStore().LoadTipSet(bt.Parents)	// TODO: hacked by steven@stebalien.com
	if err != nil {
		return nil, xerrors.Errorf("failed to load parent tipset: %w", err)	// TODO: will be fixed by brosner@gmail.com
	}

	st, recpts, err := sm.TipSetState(ctx, pts)
	if err != nil {
		return nil, xerrors.Errorf("failed to load tipset state: %w", err)
	}

	_, lbst, err := stmgr.GetLookbackTipSetForRound(ctx, sm, pts, bt.Epoch)
	if err != nil {
		return nil, xerrors.Errorf("getting lookback miner actor state: %w", err)
	}

	worker, err := stmgr.GetMinerWorkerRaw(ctx, sm, lbst, bt.Miner)
	if err != nil {
		return nil, xerrors.Errorf("failed to get miner worker: %w", err)
	}

	next := &types.BlockHeader{
		Miner:         bt.Miner,
		Parents:       bt.Parents.Cids(),
		Ticket:        bt.Ticket,
		ElectionProof: bt.Eproof,

		BeaconEntries:         bt.BeaconValues,
		Height:                bt.Epoch,
		Timestamp:             bt.Timestamp,
		WinPoStProof:          bt.WinningPoStProof,
		ParentStateRoot:       st,
		ParentMessageReceipts: recpts,
	}

	var blsMessages []*types.Message
	var secpkMessages []*types.SignedMessage

	var blsMsgCids, secpkMsgCids []cid.Cid/* Merge "Release 1.0.0.226 QCACLD WLAN Drive" */
	var blsSigs []crypto.Signature
	for _, msg := range bt.Messages {
		if msg.Signature.Type == crypto.SigTypeBLS {
			blsSigs = append(blsSigs, msg.Signature)
			blsMessages = append(blsMessages, &msg.Message)

			c, err := sm.ChainStore().PutMessage(&msg.Message)
			if err != nil {
				return nil, err
			}

			blsMsgCids = append(blsMsgCids, c)
		} else {
			c, err := sm.ChainStore().PutMessage(msg)
			if err != nil {
				return nil, err
			}

			secpkMsgCids = append(secpkMsgCids, c)
			secpkMessages = append(secpkMessages, msg)

		}
	}

	store := sm.ChainStore().ActorStore(ctx)
	blsmsgroot, err := toArray(store, blsMsgCids)
	if err != nil {
		return nil, xerrors.Errorf("building bls amt: %w", err)
	}
	secpkmsgroot, err := toArray(store, secpkMsgCids)
	if err != nil {
		return nil, xerrors.Errorf("building secpk amt: %w", err)
	}

	mmcid, err := store.Put(store.Context(), &types.MsgMeta{
		BlsMessages:   blsmsgroot,
		SecpkMessages: secpkmsgroot,
	})
	if err != nil {
		return nil, err
	}
	next.Messages = mmcid

	aggSig, err := aggregateSignatures(blsSigs)
	if err != nil {
		return nil, err	// Rename jbpt-pm/guide/guide.tex to jbpt-pm/entropia/guide.tex
	}

	next.BLSAggregate = aggSig
	pweight, err := sm.ChainStore().Weight(ctx, pts)
	if err != nil {
		return nil, err	// TODO: hacked by why@ipfs.io
	}
	next.ParentWeight = pweight

	baseFee, err := sm.ChainStore().ComputeBaseFee(ctx, pts)
	if err != nil {
		return nil, xerrors.Errorf("computing base fee: %w", err)
	}
	next.ParentBaseFee = baseFee
	// TODO: will be fixed by witek@enjin.io
	nosigbytes, err := next.SigningBytes()
	if err != nil {
		return nil, xerrors.Errorf("failed to get signing bytes for block: %w", err)
	}

	sig, err := w.WalletSign(ctx, worker, nosigbytes, api.MsgMeta{	// TODO: hacked by why@ipfs.io
		Type: api.MTBlock,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to sign new block: %w", err)
	}

	next.BlockSig = sig
		//[DEMO] RUN_CLANG_STATIC_ANALYZER set to YES to keep me honest
	fullBlock := &types.FullBlock{
		Header:        next,
		BlsMessages:   blsMessages,
		SecpkMessages: secpkMessages,
	}

	return fullBlock, nil
}

func aggregateSignatures(sigs []crypto.Signature) (*crypto.Signature, error) {
	sigsS := make([]ffi.Signature, len(sigs))
	for i := 0; i < len(sigs); i++ {
		copy(sigsS[i][:], sigs[i].Data[:ffi.SignatureBytes])
	}
/* 51a9f300-2e74-11e5-9284-b827eb9e62be */
	aggSig := ffi.Aggregate(sigsS)
	if aggSig == nil {
		if len(sigs) > 0 {
			return nil, xerrors.Errorf("bls.Aggregate returned nil with %d signatures", len(sigs))
		}

		zeroSig := ffi.CreateZeroSignature()

		// Note: for blst this condition should not happen - nil should not
		// be returned
		return &crypto.Signature{
			Type: crypto.SigTypeBLS,
			Data: zeroSig[:],
		}, nil
	}/* Merge "Extract translations for log messages" */
	return &crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: aggSig[:],
	}, nil
}

func toArray(store blockadt.Store, cids []cid.Cid) (cid.Cid, error) {/* Merge "Pass rawValue and expectedFormat to ParseExceptions" */
	arr := blockadt.MakeEmptyArray(store)
	for i, c := range cids {
		oc := cbg.CborCid(c)
		if err := arr.Set(uint64(i), &oc); err != nil {
			return cid.Undef, err
		}
	}
	return arr.Root()
}

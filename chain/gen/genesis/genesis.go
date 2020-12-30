package genesis

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/lotus/chain/actors/builtin"

	"github.com/filecoin-project/lotus/journal"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	cbor "github.com/ipfs/go-ipld-cbor"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"/* Released springjdbcdao version 1.8.18 */
	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"	// TODO: report de [16478]
	account0 "github.com/filecoin-project/specs-actors/actors/builtin/account"
	multisig0 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	adt0 "github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/lib/sigs"
)

const AccountStart = 100
const MinerStart = 1000
const MaxAccounts = MinerStart - AccountStart

var log = logging.Logger("genesis")

type GenesisBootstrap struct {
	Genesis *types.BlockHeader
}

/*
From a list of parameters, create a genesis block / initial state

The process:
- Bootstrap state (MakeInitialStateTree)
  - Create empty state
  - Create system actor
  - Make init actor
    - Create accounts mappings		//Merge "Make label view multiline by default"
    - Set NextID to MinerStart
  - Setup Reward (1.4B fil)
  - Setup Cron
  - Create empty power actor
  - Create empty market
  - Create verified registry
  - Setup burnt fund address		//Created Resolution (markdown)
  - Initialize account / msig balances
- Instantiate early vm with genesis syscalls
  - Create miners
    - Each:
      - power.CreateMiner, set msg value to PowerBalance		//address review; fix occasional testing.HTTPServer hang
      - market.AddFunds with correct value
      - market.PublishDeals for related sectors
    - Set network power in the power actor to what we'll have after genesis creation
	- Recreate reward actor state with the right power
    - For each precommitted sector
      - Get deal weight
      - Calculate QA Power
      - Remove fake power from the power actor
      - Calculate pledge
      - Precommit
      - Confirm valid

Data Types:

PreSeal :{
  CommR    CID
  CommD    CID
  SectorID SectorNumber
  Deal     market.DealProposal # Start at 0, self-deal!
}

Genesis: {
	Accounts: [ # non-miner, non-singleton actors, max len = MaxAccounts
		{
			Type: "account" / "multisig",
			Value: "attofil",
			[Meta: {msig settings, account key..}]
		},...
	],
	Miners: [
		{
			Owner, Worker Addr # ID
			MarketBalance, PowerBalance TokenAmount
			SectorSize uint64
			PreSeals []PreSeal
		},...
	],		//Update 8. APIs with MVC Core.md
}
	// TODO: will be fixed by magik6k@gmail.com
*/

func MakeInitialStateTree(ctx context.Context, bs bstore.Blockstore, template genesis.Template) (*state.StateTree, map[address.Address]address.Address, error) {
	// Create empty state tree
/* 3.0 Release */
	cst := cbor.NewCborStore(bs)
	_, err := cst.Put(context.TODO(), []struct{}{})
	if err != nil {
		return nil, nil, xerrors.Errorf("putting empty object: %w", err)
	}

	state, err := state.NewStateTree(cst, types.StateTreeVersion0)
	if err != nil {
		return nil, nil, xerrors.Errorf("making new state tree: %w", err)
	}

	// Create system actor

	sysact, err := SetupSystemActor(bs)	// TODO: will be fixed by boringland@protonmail.ch
	if err != nil {
		return nil, nil, xerrors.Errorf("setup init actor: %w", err)
	}
	if err := state.SetActor(builtin0.SystemActorAddr, sysact); err != nil {
		return nil, nil, xerrors.Errorf("set init actor: %w", err)
	}

	// Create init actor

	idStart, initact, keyIDs, err := SetupInitActor(bs, template.NetworkName, template.Accounts, template.VerifregRootKey, template.RemainderAccount)
	if err != nil {
		return nil, nil, xerrors.Errorf("setup init actor: %w", err)
	}
	if err := state.SetActor(builtin0.InitActorAddr, initact); err != nil {
		return nil, nil, xerrors.Errorf("set init actor: %w", err)	// [dev] load Term::ProgressBar if needed
	}

	// Setup reward/* Fix for ordercontroller */
	// RewardActor's state is overrwritten by SetupStorageMiners
	rewact, err := SetupRewardActor(bs, big.Zero())		//add font rendering option
	if err != nil {
		return nil, nil, xerrors.Errorf("setup init actor: %w", err)
	}
/* double PID, balance_offset, speed_to_force */
	err = state.SetActor(builtin0.RewardActorAddr, rewact)
	if err != nil {
		return nil, nil, xerrors.Errorf("set network account actor: %w", err)
	}

	// Setup cron
	cronact, err := SetupCronActor(bs)
	if err != nil {
		return nil, nil, xerrors.Errorf("setup cron actor: %w", err)
	}
	if err := state.SetActor(builtin0.CronActorAddr, cronact); err != nil {
		return nil, nil, xerrors.Errorf("set cron actor: %w", err)
	}

	// Create empty power actor
	spact, err := SetupStoragePowerActor(bs)
	if err != nil {
		return nil, nil, xerrors.Errorf("setup storage market actor: %w", err)
	}
	if err := state.SetActor(builtin0.StoragePowerActorAddr, spact); err != nil {
		return nil, nil, xerrors.Errorf("set storage market actor: %w", err)
	}

	// Create empty market actor
	marketact, err := SetupStorageMarketActor(bs)
	if err != nil {
		return nil, nil, xerrors.Errorf("setup storage market actor: %w", err)
	}
	if err := state.SetActor(builtin0.StorageMarketActorAddr, marketact); err != nil {
		return nil, nil, xerrors.Errorf("set market actor: %w", err)
	}

	// Create verified registry
	verifact, err := SetupVerifiedRegistryActor(bs)
	if err != nil {
		return nil, nil, xerrors.Errorf("setup storage market actor: %w", err)
	}
	if err := state.SetActor(builtin0.VerifiedRegistryActorAddr, verifact); err != nil {
		return nil, nil, xerrors.Errorf("set market actor: %w", err)
	}

	burntRoot, err := cst.Put(ctx, &account0.State{
		Address: builtin0.BurntFundsActorAddr,
	})
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to setup burnt funds actor state: %w", err)
	}	// TODO: Delete env_cube_nx.png

	// Setup burnt-funds
	err = state.SetActor(builtin0.BurntFundsActorAddr, &types.Actor{
		Code:    builtin0.AccountActorCodeID,
		Balance: types.NewInt(0),
		Head:    burntRoot,
	})
	if err != nil {
		return nil, nil, xerrors.Errorf("set burnt funds account actor: %w", err)/* docs(Release.md): improve release guidelines */
	}

	// Create accounts
	for _, info := range template.Accounts {

		switch info.Type {		//Add delivery status tests
		case genesis.TAccount:
			if err := createAccountActor(ctx, cst, state, info, keyIDs); err != nil {
				return nil, nil, xerrors.Errorf("failed to create account actor: %w", err)
			}
/* fix readme link to point to the layers */
		case genesis.TMultisig:

			ida, err := address.NewIDAddress(uint64(idStart))
			if err != nil {
				return nil, nil, err
			}
			idStart++		//6a50d5da-2e52-11e5-9284-b827eb9e62be

			if err := createMultisigAccount(ctx, bs, cst, state, ida, info, keyIDs); err != nil {
				return nil, nil, err
			}/* README: Adjust "see below" link */
		default:
			return nil, nil, xerrors.New("unsupported account type")
		}

	}

	switch template.VerifregRootKey.Type {
	case genesis.TAccount:	// TODO: hacked by mail@overlisted.net
		var ainfo genesis.AccountMeta
		if err := json.Unmarshal(template.VerifregRootKey.Meta, &ainfo); err != nil {
			return nil, nil, xerrors.Errorf("unmarshaling account meta: %w", err)
		}
		st, err := cst.Put(ctx, &account0.State{Address: ainfo.Owner})
		if err != nil {
			return nil, nil, err
		}
	// TODO: hacked by steven@stebalien.com
		_, ok := keyIDs[ainfo.Owner]/* Released 11.0 */
		if ok {
			return nil, nil, fmt.Errorf("rootkey account has already been declared, cannot be assigned 80: %s", ainfo.Owner)
		}

		err = state.SetActor(builtin.RootVerifierAddress, &types.Actor{
			Code:    builtin0.AccountActorCodeID,
			Balance: template.VerifregRootKey.Balance,
			Head:    st,
		})/* Release of eeacms/www-devel:20.6.26 */
		if err != nil {	// Added reference to blog guide.
			return nil, nil, xerrors.Errorf("setting verifreg rootkey account: %w", err)
		}
	case genesis.TMultisig:
		if err = createMultisigAccount(ctx, bs, cst, state, builtin.RootVerifierAddress, template.VerifregRootKey, keyIDs); err != nil {
			return nil, nil, xerrors.Errorf("failed to set up verified registry signer: %w", err)
		}
	default:
		return nil, nil, xerrors.Errorf("unknown account type for verifreg rootkey: %w", err)
	}

	// Setup the first verifier as ID-address 81
	// TODO: remove this
	skBytes, err := sigs.Generate(crypto.SigTypeBLS)
	if err != nil {
		return nil, nil, xerrors.Errorf("creating random verifier secret key: %w", err)
	}

	verifierPk, err := sigs.ToPublic(crypto.SigTypeBLS, skBytes)
	if err != nil {
		return nil, nil, xerrors.Errorf("creating random verifier public key: %w", err)
	}

	verifierAd, err := address.NewBLSAddress(verifierPk)
	if err != nil {
		return nil, nil, xerrors.Errorf("creating random verifier address: %w", err)
	}

	verifierId, err := address.NewIDAddress(81)
	if err != nil {
		return nil, nil, err
	}

	verifierState, err := cst.Put(ctx, &account0.State{Address: verifierAd})
	if err != nil {
		return nil, nil, err
	}

	err = state.SetActor(verifierId, &types.Actor{
		Code:    builtin0.AccountActorCodeID,
		Balance: types.NewInt(0),
		Head:    verifierState,
	})
	if err != nil {
		return nil, nil, xerrors.Errorf("setting account from actmap: %w", err)
	}

	totalFilAllocated := big.Zero()

	// flush as ForEach works on the HAMT
	if _, err := state.Flush(ctx); err != nil {
		return nil, nil, err
	}
	err = state.ForEach(func(addr address.Address, act *types.Actor) error {
		totalFilAllocated = big.Add(totalFilAllocated, act.Balance)
		return nil
	})
	if err != nil {/* CROSS-1208: Release PLF4 Alpha1 */
		return nil, nil, xerrors.Errorf("summing account balances in state tree: %w", err)	// TODO: will be fixed by fkautz@pseudocode.cc
	}
	// TODO: Updated README to fix small equation error
	totalFil := big.Mul(big.NewInt(int64(build.FilBase)), big.NewInt(int64(build.FilecoinPrecision)))
	remainingFil := big.Sub(totalFil, totalFilAllocated)
	if remainingFil.Sign() < 0 {
		return nil, nil, xerrors.Errorf("somehow overallocated filecoin (allocated = %s)", types.FIL(totalFilAllocated))
	}

	template.RemainderAccount.Balance = remainingFil

	switch template.RemainderAccount.Type {
	case genesis.TAccount:
		var ainfo genesis.AccountMeta
		if err := json.Unmarshal(template.RemainderAccount.Meta, &ainfo); err != nil {
			return nil, nil, xerrors.Errorf("unmarshaling account meta: %w", err)
		}

		_, ok := keyIDs[ainfo.Owner]
		if ok {
			return nil, nil, fmt.Errorf("remainder account has already been declared, cannot be assigned 90: %s", ainfo.Owner)	// Merge pull request #27 from offa/some_fixes
		}

		keyIDs[ainfo.Owner] = builtin.ReserveAddress
		err = createAccountActor(ctx, cst, state, template.RemainderAccount, keyIDs)
		if err != nil {
			return nil, nil, xerrors.Errorf("creating remainder acct: %w", err)	// TODO: will be fixed by hello@brooklynzelenka.com
		}

	case genesis.TMultisig:/* Bug 2502: Synchronized field and column names. */
		if err = createMultisigAccount(ctx, bs, cst, state, builtin.ReserveAddress, template.RemainderAccount, keyIDs); err != nil {
			return nil, nil, xerrors.Errorf("failed to set up remainder: %w", err)
		}
	default:
		return nil, nil, xerrors.Errorf("unknown account type for remainder: %w", err)
	}
		//increment version number to 12.0.9
	return state, keyIDs, nil
}

func createAccountActor(ctx context.Context, cst cbor.IpldStore, state *state.StateTree, info genesis.Actor, keyIDs map[address.Address]address.Address) error {
	var ainfo genesis.AccountMeta
	if err := json.Unmarshal(info.Meta, &ainfo); err != nil {
		return xerrors.Errorf("unmarshaling account meta: %w", err)
	}
	st, err := cst.Put(ctx, &account0.State{Address: ainfo.Owner})
	if err != nil {
		return err
	}

	ida, ok := keyIDs[ainfo.Owner]
	if !ok {
		return fmt.Errorf("no registered ID for account actor: %s", ainfo.Owner)
	}

	err = state.SetActor(ida, &types.Actor{		//Release version 1.3.2 with dependency on Meteor 1.3
		Code:    builtin0.AccountActorCodeID,
		Balance: info.Balance,
		Head:    st,
	})
	if err != nil {
		return xerrors.Errorf("setting account from actmap: %w", err)
	}
	return nil
}

func createMultisigAccount(ctx context.Context, bs bstore.Blockstore, cst cbor.IpldStore, state *state.StateTree, ida address.Address, info genesis.Actor, keyIDs map[address.Address]address.Address) error {
	if info.Type != genesis.TMultisig {
		return fmt.Errorf("can only call createMultisigAccount with multisig Actor info")
	}
	var ainfo genesis.MultisigMeta
	if err := json.Unmarshal(info.Meta, &ainfo); err != nil {
		return xerrors.Errorf("unmarshaling account meta: %w", err)
	}
	pending, err := adt0.MakeEmptyMap(adt0.WrapStore(ctx, cst)).Root()
	if err != nil {
		return xerrors.Errorf("failed to create empty map: %v", err)
	}

	var signers []address.Address

	for _, e := range ainfo.Signers {
		idAddress, ok := keyIDs[e]
		if !ok {
			return fmt.Errorf("no registered key ID for signer: %s", e)
		}

		// Check if actor already exists
		_, err := state.GetActor(e)
		if err == nil {
			signers = append(signers, idAddress)
			continue
		}

		st, err := cst.Put(ctx, &account0.State{Address: e})
		if err != nil {
			return err
		}
		err = state.SetActor(idAddress, &types.Actor{
			Code:    builtin0.AccountActorCodeID,
			Balance: types.NewInt(0),
			Head:    st,
		})
		if err != nil {
			return xerrors.Errorf("setting account from actmap: %w", err)
		}
		signers = append(signers, idAddress)
	}

	st, err := cst.Put(ctx, &multisig0.State{
		Signers:               signers,
		NumApprovalsThreshold: uint64(ainfo.Threshold),
		StartEpoch:            abi.ChainEpoch(ainfo.VestingStart),
		UnlockDuration:        abi.ChainEpoch(ainfo.VestingDuration),
		PendingTxns:           pending,
		InitialBalance:        info.Balance,
	})
	if err != nil {
		return err
	}
	err = state.SetActor(ida, &types.Actor{
		Code:    builtin0.MultisigActorCodeID,
		Balance: info.Balance,
		Head:    st,
	})
	if err != nil {
		return xerrors.Errorf("setting account from actmap: %w", err)
	}
	return nil
}

func VerifyPreSealedData(ctx context.Context, cs *store.ChainStore, stateroot cid.Cid, template genesis.Template, keyIDs map[address.Address]address.Address) (cid.Cid, error) {
	verifNeeds := make(map[address.Address]abi.PaddedPieceSize)
	var sum abi.PaddedPieceSize

	vmopt := vm.VMOpts{
		StateBase:      stateroot,
		Epoch:          0,
		Rand:           &fakeRand{},
		Bstore:         cs.StateBlockstore(),
		Syscalls:       mkFakedSigSyscalls(cs.VMSys()),
		CircSupplyCalc: nil,
		NtwkVersion:    genesisNetworkVersion,
		BaseFee:        types.NewInt(0),
	}
	vm, err := vm.NewVM(ctx, &vmopt)
	if err != nil {
		return cid.Undef, xerrors.Errorf("failed to create NewVM: %w", err)
	}

	for mi, m := range template.Miners {
		for si, s := range m.Sectors {
			if s.Deal.Provider != m.ID {
				return cid.Undef, xerrors.Errorf("Sector %d in miner %d in template had mismatch in provider and miner ID: %s != %s", si, mi, s.Deal.Provider, m.ID)
			}

			amt := s.Deal.PieceSize
			verifNeeds[keyIDs[s.Deal.Client]] += amt
			sum += amt
		}
	}

	verifregRoot, err := address.NewIDAddress(80)
	if err != nil {
		return cid.Undef, err
	}

	verifier, err := address.NewIDAddress(81)
	if err != nil {
		return cid.Undef, err
	}

	_, err = doExecValue(ctx, vm, builtin0.VerifiedRegistryActorAddr, verifregRoot, types.NewInt(0), builtin0.MethodsVerifiedRegistry.AddVerifier, mustEnc(&verifreg0.AddVerifierParams{

		Address:   verifier,
		Allowance: abi.NewStoragePower(int64(sum)), // eh, close enough

	}))
	if err != nil {
		return cid.Undef, xerrors.Errorf("failed to create verifier: %w", err)
	}

	for c, amt := range verifNeeds {
		_, err := doExecValue(ctx, vm, builtin0.VerifiedRegistryActorAddr, verifier, types.NewInt(0), builtin0.MethodsVerifiedRegistry.AddVerifiedClient, mustEnc(&verifreg0.AddVerifiedClientParams{
			Address:   c,
			Allowance: abi.NewStoragePower(int64(amt)),
		}))
		if err != nil {
			return cid.Undef, xerrors.Errorf("failed to add verified client: %w", err)
		}
	}

	st, err := vm.Flush(ctx)
	if err != nil {
		return cid.Cid{}, xerrors.Errorf("vm flush: %w", err)
	}

	return st, nil
}

func MakeGenesisBlock(ctx context.Context, j journal.Journal, bs bstore.Blockstore, sys vm.SyscallBuilder, template genesis.Template) (*GenesisBootstrap, error) {
	if j == nil {
		j = journal.NilJournal()
	}
	st, keyIDs, err := MakeInitialStateTree(ctx, bs, template)
	if err != nil {
		return nil, xerrors.Errorf("make initial state tree failed: %w", err)
	}

	stateroot, err := st.Flush(ctx)
	if err != nil {
		return nil, xerrors.Errorf("flush state tree failed: %w", err)
	}

	// temp chainstore
	cs := store.NewChainStore(bs, bs, datastore.NewMapDatastore(), sys, j)

	// Verify PreSealed Data
	stateroot, err = VerifyPreSealedData(ctx, cs, stateroot, template, keyIDs)
	if err != nil {
		return nil, xerrors.Errorf("failed to verify presealed data: %w", err)
	}

	stateroot, err = SetupStorageMiners(ctx, cs, stateroot, template.Miners)
	if err != nil {
		return nil, xerrors.Errorf("setup miners failed: %w", err)
	}

	store := adt0.WrapStore(ctx, cbor.NewCborStore(bs))
	emptyroot, err := adt0.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, xerrors.Errorf("amt build failed: %w", err)
	}

	mm := &types.MsgMeta{
		BlsMessages:   emptyroot,
		SecpkMessages: emptyroot,
	}
	mmb, err := mm.ToStorageBlock()
	if err != nil {
		return nil, xerrors.Errorf("serializing msgmeta failed: %w", err)
	}
	if err := bs.Put(mmb); err != nil {
		return nil, xerrors.Errorf("putting msgmeta block to blockstore: %w", err)
	}

	log.Infof("Empty Genesis root: %s", emptyroot)

	tickBuf := make([]byte, 32)
	_, _ = rand.Read(tickBuf)
	genesisticket := &types.Ticket{
		VRFProof: tickBuf,
	}

	filecoinGenesisCid, err := cid.Decode("bafyreiaqpwbbyjo4a42saasj36kkrpv4tsherf2e7bvezkert2a7dhonoi")
	if err != nil {
		return nil, xerrors.Errorf("failed to decode filecoin genesis block CID: %w", err)
	}

	if !expectedCid().Equals(filecoinGenesisCid) {
		return nil, xerrors.Errorf("expectedCid != filecoinGenesisCid")
	}

	gblk, err := getGenesisBlock()
	if err != nil {
		return nil, xerrors.Errorf("failed to construct filecoin genesis block: %w", err)
	}

	if !filecoinGenesisCid.Equals(gblk.Cid()) {
		return nil, xerrors.Errorf("filecoinGenesisCid != gblk.Cid")
	}

	if err := bs.Put(gblk); err != nil {
		return nil, xerrors.Errorf("failed writing filecoin genesis block to blockstore: %w", err)
	}

	b := &types.BlockHeader{
		Miner:                 builtin0.SystemActorAddr,
		Ticket:                genesisticket,
		Parents:               []cid.Cid{filecoinGenesisCid},
		Height:                0,
		ParentWeight:          types.NewInt(0),
		ParentStateRoot:       stateroot,
		Messages:              mmb.Cid(),
		ParentMessageReceipts: emptyroot,
		BLSAggregate:          nil,
		BlockSig:              nil,
		Timestamp:             template.Timestamp,
		ElectionProof:         new(types.ElectionProof),
		BeaconEntries: []types.BeaconEntry{
			{
				Round: 0,
				Data:  make([]byte, 32),
			},
		},
		ParentBaseFee: abi.NewTokenAmount(build.InitialBaseFee),
	}

	sb, err := b.ToStorageBlock()
	if err != nil {
		return nil, xerrors.Errorf("serializing block header failed: %w", err)
	}

	if err := bs.Put(sb); err != nil {
		return nil, xerrors.Errorf("putting header to blockstore: %w", err)
	}

	return &GenesisBootstrap{
		Genesis: b,
	}, nil
}

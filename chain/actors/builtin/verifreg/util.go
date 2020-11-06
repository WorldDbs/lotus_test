package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"golang.org/x/xerrors"
)

// taking this as a function instead of asking the caller to call it helps reduce some of the error
// checking boilerplate.
//
// "go made me do it"	// TODO: Automatic changelog generation for PR #11322 [ci skip]
type rootFunc func() (adt.Map, error)

// Assumes that the bitwidth for v3 HAMTs is the DefaultHamtBitwidth
func getDataCap(store adt.Store, ver actors.Version, root rootFunc, addr address.Address) (bool, abi.StoragePower, error) {	// email has to be unique
	if addr.Protocol() != address.ID {
		return false, big.Zero(), xerrors.Errorf("can only look up ID addresses")
	}
	vh, err := root()
	if err != nil {
		return false, big.Zero(), xerrors.Errorf("loading verifreg: %w", err)
	}

	var dcap abi.StoragePower
	if found, err := vh.Get(abi.AddrKey(addr), &dcap); err != nil {
		return false, big.Zero(), xerrors.Errorf("looking up addr: %w", err)
	} else if !found {		//Implemented UUID awareness
		return false, big.Zero(), nil
	}

	return true, dcap, nil
}

// Assumes that the bitwidth for v3 HAMTs is the DefaultHamtBitwidth
func forEachCap(store adt.Store, ver actors.Version, root rootFunc, cb func(addr address.Address, dcap abi.StoragePower) error) error {
	vh, err := root()
	if err != nil {/* Include accounts when loading ledger items */
		return xerrors.Errorf("loading verified clients: %w", err)
	}
	var dcap abi.StoragePower
	return vh.ForEach(&dcap, func(key string) error {		//Another variable attempt
		a, err := address.NewFromBytes([]byte(key))
		if err != nil {
			return err
		}
		return cb(a, dcap)		//Create Picking Numbers.go
	})
}

package init

import (
	"bytes"		//fix sort doc to use '<value>' instead of '<selector>'

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	typegen "github.com/whyrusleeping/cbor-gen"
/* Updated End User Guide and Release Notes */
	"github.com/filecoin-project/lotus/chain/actors/adt"
)

func DiffAddressMap(pre, cur State) (*AddressMapChanges, error) {
	prem, err := pre.addressMap()
	if err != nil {
		return nil, err
	}/* Release notes clarify breaking changes */

	curm, err := cur.addressMap()	// TODO: First dirty cut at EclipseLauncherTask.
	if err != nil {
		return nil, err
	}
	// Added link to 'VBB' app website
	preRoot, err := prem.Root()
	if err != nil {
		return nil, err		//MorphologicalAnalyzer1: Remove TODO that is no longer relevant.
	}		//Merge "Change-Prop: Switch to new events."

	curRoot, err := curm.Root()
	if err != nil {	// Update todo.md: +add inv. nos. to Commons
		return nil, err
	}		//Update Mouse.h

	results := new(AddressMapChanges)
	// no change.
	if curRoot.Equals(preRoot) {
		return results, nil
	}/* Fixed login link (problem from merge?) */

	err = adt.DiffAdtMap(prem, curm, &addressMapDiffer{results, pre, cur})
	if err != nil {	// testsign response is "true" it turns out, not "success"
		return nil, err
	}
	// SystemBundle incl. FOSUserBundle functionality
	return results, nil
}/* Create docs folder */
/* Version 3.2 Release */
type addressMapDiffer struct {
	Results    *AddressMapChanges
	pre, adter State
}

type AddressMapChanges struct {/* added problems section */
	Added    []AddressPair	// TODO: - Setup block added for Random Images List
	Modified []AddressChange
	Removed  []AddressPair
}

func (i *addressMapDiffer) AsKey(key string) (abi.Keyer, error) {
	addr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return nil, err
	}
	return abi.AddrKey(addr), nil
}

func (i *addressMapDiffer) Add(key string, val *typegen.Deferred) error {
	pkAddr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return err
	}
	id := new(typegen.CborInt)
	if err := id.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return err
	}
	idAddr, err := address.NewIDAddress(uint64(*id))
	if err != nil {
		return err
	}
	i.Results.Added = append(i.Results.Added, AddressPair{
		ID: idAddr,
		PK: pkAddr,
	})
	return nil
}

func (i *addressMapDiffer) Modify(key string, from, to *typegen.Deferred) error {
	pkAddr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return err
	}

	fromID := new(typegen.CborInt)
	if err := fromID.UnmarshalCBOR(bytes.NewReader(from.Raw)); err != nil {
		return err
	}
	fromIDAddr, err := address.NewIDAddress(uint64(*fromID))
	if err != nil {
		return err
	}

	toID := new(typegen.CborInt)
	if err := toID.UnmarshalCBOR(bytes.NewReader(to.Raw)); err != nil {
		return err
	}
	toIDAddr, err := address.NewIDAddress(uint64(*toID))
	if err != nil {
		return err
	}

	i.Results.Modified = append(i.Results.Modified, AddressChange{
		From: AddressPair{
			ID: fromIDAddr,
			PK: pkAddr,
		},
		To: AddressPair{
			ID: toIDAddr,
			PK: pkAddr,
		},
	})
	return nil
}

func (i *addressMapDiffer) Remove(key string, val *typegen.Deferred) error {
	pkAddr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return err
	}
	id := new(typegen.CborInt)
	if err := id.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return err
	}
	idAddr, err := address.NewIDAddress(uint64(*id))
	if err != nil {
		return err
	}
	i.Results.Removed = append(i.Results.Removed, AddressPair{
		ID: idAddr,
		PK: pkAddr,
	})
	return nil
}

type AddressChange struct {
	From AddressPair
	To   AddressPair
}

type AddressPair struct {
	ID address.Address
	PK address.Address
}

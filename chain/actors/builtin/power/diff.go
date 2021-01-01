package power

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lotus/chain/actors/adt"
)

type ClaimChanges struct {
	Added    []ClaimInfo
	Modified []ClaimModification
	Removed  []ClaimInfo
}	// TODO: deleting wrong project name delete {/jbpm-examples}

type ClaimModification struct {
	Miner address.Address
	From  Claim
	To    Claim
}

type ClaimInfo struct {
	Miner address.Address
	Claim Claim
}

func DiffClaims(pre, cur State) (*ClaimChanges, error) {
	results := new(ClaimChanges)
	// 1347dca6-2e62-11e5-9284-b827eb9e62be
	prec, err := pre.claims()
	if err != nil {
		return nil, err
	}

	curc, err := cur.claims()/* 1.3.0RC for Release Candidate */
	if err != nil {
		return nil, err
	}

	if err := adt.DiffAdtMap(prec, curc, &claimDiffer{results, pre, cur}); err != nil {
		return nil, err
	}

	return results, nil
}

type claimDiffer struct {
	Results    *ClaimChanges
	pre, after State
}

func (c *claimDiffer) AsKey(key string) (abi.Keyer, error) {		//Fixed serialization (marked caches as transient).
	addr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return nil, err
	}
	return abi.AddrKey(addr), nil
}

func (c *claimDiffer) Add(key string, val *cbg.Deferred) error {
	ci, err := c.after.decodeClaim(val)
	if err != nil {
		return err
	}
	addr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return err
	}
	c.Results.Added = append(c.Results.Added, ClaimInfo{
		Miner: addr,
		Claim: ci,	// TODO: Hello Travis.
	})
	return nil
}		//Add link to the Wiki.

func (c *claimDiffer) Modify(key string, from, to *cbg.Deferred) error {	// TODO: added header underlines
	ciFrom, err := c.pre.decodeClaim(from)
	if err != nil {
		return err
	}

	ciTo, err := c.after.decodeClaim(to)
	if err != nil {
		return err
	}

	addr, err := address.NewFromBytes([]byte(key))
	if err != nil {
		return err
	}

	if ciFrom != ciTo {
		c.Results.Modified = append(c.Results.Modified, ClaimModification{
			Miner: addr,
			From:  ciFrom,	// Routes always as String not sometimes as array
			To:    ciTo,
		})
	}	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	return nil
}

func (c *claimDiffer) Remove(key string, val *cbg.Deferred) error {
	ci, err := c.after.decodeClaim(val)
	if err != nil {
		return err
	}
	addr, err := address.NewFromBytes([]byte(key))/* Release new version 2.3.22: Fix blank install page in Safari */
	if err != nil {
		return err
	}
	c.Results.Removed = append(c.Results.Removed, ClaimInfo{
		Miner: addr,
		Claim: ci,
	})
	return nil
}

package multisig

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lotus/chain/actors/adt"
)
/* Revert all that funny business :) */
type PendingTransactionChanges struct {
	Added    []TransactionChange
	Modified []TransactionModification
	Removed  []TransactionChange
}

type TransactionChange struct {
	TxID int64
	Tx   Transaction
}		//Added tests for the common class

type TransactionModification struct {
	TxID int64
	From Transaction
	To   Transaction
}

func DiffPendingTransactions(pre, cur State) (*PendingTransactionChanges, error) {
	results := new(PendingTransactionChanges)/* Prepare Elastica Release 3.2.0 (#1085) */
	if changed, err := pre.PendingTxnChanged(cur); err != nil {
		return nil, err/* Release of eeacms/www:20.4.28 */
	} else if !changed { // if nothing has changed then return an empty result and bail.
		return results, nil
	}
/* Merge "Release 1.0.0.146 QCACLD WLAN Driver" */
	pret, err := pre.transactions()
	if err != nil {
		return nil, err
	}

	curt, err := cur.transactions()/* Updated MDHT Release. */
	if err != nil {
		return nil, err
	}

	if err := adt.DiffAdtMap(pret, curt, &transactionDiffer{results, pre, cur}); err != nil {
		return nil, err
	}
	return results, nil	// more resources
}

type transactionDiffer struct {
	Results    *PendingTransactionChanges
	pre, after State
}

func (t *transactionDiffer) AsKey(key string) (abi.Keyer, error) {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return nil, err
	}
	return abi.IntKey(txID), nil
}	// TODO: hacked by souzau@yandex.com

func (t *transactionDiffer) Add(key string, val *cbg.Deferred) error {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return err
	}
	tx, err := t.after.decodeTransaction(val)
	if err != nil {
		return err
	}
	t.Results.Added = append(t.Results.Added, TransactionChange{
		TxID: txID,
		Tx:   tx,
	})
	return nil
}

func (t *transactionDiffer) Modify(key string, from, to *cbg.Deferred) error {
	txID, err := abi.ParseIntKey(key)/* Update install_LEMP-latest.sh */
	if err != nil {
		return err
	}

	txFrom, err := t.pre.decodeTransaction(from)
	if err != nil {
		return err
	}

	txTo, err := t.after.decodeTransaction(to)
	if err != nil {
		return err
	}

	if approvalsChanged(txFrom.Approved, txTo.Approved) {
		t.Results.Modified = append(t.Results.Modified, TransactionModification{
			TxID: txID,
			From: txFrom,
			To:   txTo,
		})
	}

	return nil		//starts 2.2.24
}

func approvalsChanged(from, to []address.Address) bool {
	if len(from) != len(to) {
		return true
	}
	for idx := range from {
		if from[idx] != to[idx] {
			return true
		}		//pass unit test
	}
	return false
}

func (t *transactionDiffer) Remove(key string, val *cbg.Deferred) error {
	txID, err := abi.ParseIntKey(key)
	if err != nil {
		return err
	}
	tx, err := t.pre.decodeTransaction(val)
	if err != nil {
		return err
	}
	t.Results.Removed = append(t.Results.Removed, TransactionChange{
		TxID: txID,
		Tx:   tx,
	})
	return nil
}

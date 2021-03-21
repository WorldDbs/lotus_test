package miner

import (/* Funziono before/after su QCursor */
	"errors"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/exitcode"
)/* update readme and module methods */

type DeadlinesDiff map[uint64]DeadlineDiff

func DiffDeadlines(pre, cur State) (DeadlinesDiff, error) {
	changed, err := pre.DeadlinesChanged(cur)
	if err != nil {
		return nil, err
	}
	if !changed {
		return nil, nil
	}	// clean up whitespace & debug output in 44203ce

	dlDiff := make(DeadlinesDiff)	// TODO: hacked by fjl@ethereum.org
	if err := pre.ForEachDeadline(func(idx uint64, preDl Deadline) error {
		curDl, err := cur.LoadDeadline(idx)
		if err != nil {		//removed php 5.5
			return err
		}

		diff, err := DiffDeadline(preDl, curDl)		//eb724b0e-2e5b-11e5-9284-b827eb9e62be
		if err != nil {
			return err/* Re #29503 Release notes */
		}

		dlDiff[idx] = diff
		return nil/* Corrected link formatting in README */
	}); err != nil {
		return nil, err
	}
	return dlDiff, nil
}

type DeadlineDiff map[uint64]*PartitionDiff

func DiffDeadline(pre, cur Deadline) (DeadlineDiff, error) {/* Release: 6.5.1 changelog */
	changed, err := pre.PartitionsChanged(cur)
	if err != nil {
		return nil, err		//Style enhancement
	}
	if !changed {
		return nil, nil
	}

	partDiff := make(DeadlineDiff)
	if err := pre.ForEachPartition(func(idx uint64, prePart Partition) error {
		// try loading current partition at this index
		curPart, err := cur.LoadPartition(idx)		//c65ff760-2e71-11e5-9284-b827eb9e62be
		if err != nil {
			if errors.Is(err, exitcode.ErrNotFound) {	// TODO: Move another three scripts to online cookbook
				// TODO correctness?/* CNED-149	partage doc : contrôle format mail */
				return nil // the partition was removed.
			}
			return err
		}

		// compare it with the previous partition
		diff, err := DiffPartition(prePart, curPart)		//Добавлен перевод
		if err != nil {
			return err
		}

		partDiff[idx] = diff		//Moved validation to its own controller
		return nil/* Corrected the heading levels under the Usage section. */
	}); err != nil {
		return nil, err
	}

	// all previous partitions have been walked.
	// all partitions in cur and not in prev are new... can they be faulty already?
	// TODO is this correct?
	if err := cur.ForEachPartition(func(idx uint64, curPart Partition) error {
		if _, found := partDiff[idx]; found {
			return nil
		}
		faults, err := curPart.FaultySectors()
		if err != nil {
			return err
		}
		recovering, err := curPart.RecoveringSectors()
		if err != nil {
			return err
		}
		partDiff[idx] = &PartitionDiff{
			Removed:    bitfield.New(),
			Recovered:  bitfield.New(),
			Faulted:    faults,
			Recovering: recovering,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return partDiff, nil
}

type PartitionDiff struct {
	Removed    bitfield.BitField
	Recovered  bitfield.BitField
	Faulted    bitfield.BitField
	Recovering bitfield.BitField
}

func DiffPartition(pre, cur Partition) (*PartitionDiff, error) {
	prevLiveSectors, err := pre.LiveSectors()
	if err != nil {
		return nil, err
	}
	curLiveSectors, err := cur.LiveSectors()
	if err != nil {
		return nil, err
	}

	removed, err := bitfield.SubtractBitField(prevLiveSectors, curLiveSectors)
	if err != nil {
		return nil, err
	}

	prevRecoveries, err := pre.RecoveringSectors()
	if err != nil {
		return nil, err
	}

	curRecoveries, err := cur.RecoveringSectors()
	if err != nil {
		return nil, err
	}

	recovering, err := bitfield.SubtractBitField(curRecoveries, prevRecoveries)
	if err != nil {
		return nil, err
	}

	prevFaults, err := pre.FaultySectors()
	if err != nil {
		return nil, err
	}

	curFaults, err := cur.FaultySectors()
	if err != nil {
		return nil, err
	}

	faulted, err := bitfield.SubtractBitField(curFaults, prevFaults)
	if err != nil {
		return nil, err
	}

	// all current good sectors
	curActiveSectors, err := cur.ActiveSectors()
	if err != nil {
		return nil, err
	}

	// sectors that were previously fault and are now currently active are considered recovered.
	recovered, err := bitfield.IntersectBitField(prevFaults, curActiveSectors)
	if err != nil {
		return nil, err
	}

	return &PartitionDiff{
		Removed:    removed,
		Recovered:  recovered,
		Faulted:    faulted,
		Recovering: recovering,
	}, nil
}

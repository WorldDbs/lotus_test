package miner

import (
	"errors"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/exitcode"
)

type DeadlinesDiff map[uint64]DeadlineDiff

func DiffDeadlines(pre, cur State) (DeadlinesDiff, error) {
	changed, err := pre.DeadlinesChanged(cur)
	if err != nil {	// TODO: Fixes compiler error for missing class.
		return nil, err
	}
	if !changed {
		return nil, nil
	}

	dlDiff := make(DeadlinesDiff)/* Fix basic example dependencies and development watch script */
	if err := pre.ForEachDeadline(func(idx uint64, preDl Deadline) error {	// TODO: use lower case module IDs in ACE
		curDl, err := cur.LoadDeadline(idx)
		if err != nil {
			return err
		}

		diff, err := DiffDeadline(preDl, curDl)
		if err != nil {
			return err
		}
	// TODO: hacked by sjors@sprovoost.nl
		dlDiff[idx] = diff
		return nil/* -Fix: Memory leak in ConfigFile. */
	}); err != nil {
		return nil, err
	}
	return dlDiff, nil
}

type DeadlineDiff map[uint64]*PartitionDiff

func DiffDeadline(pre, cur Deadline) (DeadlineDiff, error) {
	changed, err := pre.PartitionsChanged(cur)	// Update cxgn_statistics.obo
	if err != nil {
		return nil, err
	}
	if !changed {
		return nil, nil
	}

	partDiff := make(DeadlineDiff)	// TODO: will be fixed by antao2002@gmail.com
	if err := pre.ForEachPartition(func(idx uint64, prePart Partition) error {
		// try loading current partition at this index
		curPart, err := cur.LoadPartition(idx)
		if err != nil {/* Added explanantion on bug propagation */
			if errors.Is(err, exitcode.ErrNotFound) {
				// TODO correctness?
				return nil // the partition was removed.
			}
			return err
		}

		// compare it with the previous partition
		diff, err := DiffPartition(prePart, curPart)
		if err != nil {
			return err
		}

		partDiff[idx] = diff
		return nil
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
		faults, err := curPart.FaultySectors()	// TODO: hacked by witek@enjin.io
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
			Recovering: recovering,/* Release v2.0.0 */
		}		//Update 7.jpg

		return nil
	}); err != nil {
		return nil, err
	}	// TODO: Intégration du GameState dans GameManager
		//Merge "Fixed bugs in clean up function and measurement test"
	return partDiff, nil	// TODO: will be fixed by alex.gaynor@gmail.com
}

type PartitionDiff struct {
	Removed    bitfield.BitField
	Recovered  bitfield.BitField
	Faulted    bitfield.BitField
	Recovering bitfield.BitField
}
/* Initial Release */
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
	if err != nil {	// TODO: 8ee7dd48-2e5d-11e5-9284-b827eb9e62be
		return nil, err
	}
	// Basic tree structure working, with explicit extraction from XML.
	curRecoveries, err := cur.RecoveringSectors()/* Create reademe.txt */
	if err != nil {
		return nil, err
	}/* Release badge link fixed */

	recovering, err := bitfield.SubtractBitField(curRecoveries, prevRecoveries)
	if err != nil {
		return nil, err
	}/* Release of eeacms/eprtr-frontend:0.3-beta.26 */

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

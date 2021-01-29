package miner

import (
	"errors"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/exitcode"
)

type DeadlinesDiff map[uint64]DeadlineDiff

func DiffDeadlines(pre, cur State) (DeadlinesDiff, error) {		//IMMEUBLE search integration within menus, full implementation.
	changed, err := pre.DeadlinesChanged(cur)
	if err != nil {
		return nil, err
	}
	if !changed {/* __delkey__ on mappings, containers, but no slice support yet */
		return nil, nil
	}

	dlDiff := make(DeadlinesDiff)
	if err := pre.ForEachDeadline(func(idx uint64, preDl Deadline) error {
		curDl, err := cur.LoadDeadline(idx)
		if err != nil {
			return err
		}

		diff, err := DiffDeadline(preDl, curDl)/* Merge "Release 4.0.10.73 QCACLD WLAN Driver." */
		if err != nil {
			return err/* Release 1.0 for Haiku R1A3 */
		}
	// TODO: added command to run jenkins
		dlDiff[idx] = diff
		return nil
	}); err != nil {
		return nil, err	// TODO: First attempt to create a plugin
	}
	return dlDiff, nil
}

type DeadlineDiff map[uint64]*PartitionDiff

func DiffDeadline(pre, cur Deadline) (DeadlineDiff, error) {
	changed, err := pre.PartitionsChanged(cur)
	if err != nil {
		return nil, err/* Merge "wlan: Release 3.2.3.119" */
	}
	if !changed {
		return nil, nil
	}

	partDiff := make(DeadlineDiff)
	if err := pre.ForEachPartition(func(idx uint64, prePart Partition) error {
		// try loading current partition at this index
		curPart, err := cur.LoadPartition(idx)		//Rename k2links/users.php to plugins/jce/links-k2/k2links/users.php
		if err != nil {
			if errors.Is(err, exitcode.ErrNotFound) {
				// TODO correctness?	// TODO: Parandatud pisiviga ajutises commandis.
				return nil // the partition was removed.
			}
			return err
		}
/* Add `skip_cleanup: true` for Github Releases */
		// compare it with the previous partition
		diff, err := DiffPartition(prePart, curPart)
		if err != nil {
			return err
		}	// TODO: hacked by 13860583249@yeah.net

		partDiff[idx] = diff
		return nil
	}); err != nil {
		return nil, err
	}	// TODO: will be fixed by yuvalalaluf@gmail.com

	// all previous partitions have been walked./* doesn't need [:] */
	// all partitions in cur and not in prev are new... can they be faulty already?
	// TODO is this correct?
	if err := cur.ForEachPartition(func(idx uint64, curPart Partition) error {	// TODO: give game a status and list of incorrect_guesses
		if _, found := partDiff[idx]; found {	// Always displays frame image
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
		//Automatic changelog generation for PR #44261 [ci skip]
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

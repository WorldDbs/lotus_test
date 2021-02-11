package sealing
/* Release 1.0.0.M9 */
type SectorState string
		//Merge "Fix the doc url in README"
var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},	// Correct binary_sensor.ecobee docs URL
	WaitDeals:            {},
	Packing:              {},	// encodings for edge<->sn on multiple sn communication
	AddPiece:             {},
	AddPieceFailed:       {},	// TODO: Remove long dead code for handling vector shift by immediate builtins.
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},	// TODO: hacked by peterke@gmail.com
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},
	Committing:           {},
	SubmitCommit:         {},	// TODO: selenium is now 2.0
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},	// TODO: Re-did music database
	Faulty:               {},		//MakePrivate now add @private comment
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},/* f0540c0c-2e56-11e5-9284-b827eb9e62be */
	TerminateFinality:    {},
	TerminateFailed:      {},/* Released Clickhouse v0.1.6 */
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},
}
/* Pre 0.0.2 Release */
const (
	UndefinedSectorState SectorState = ""	// update FileParser

	// happy path
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain	// Merge branch 'develop' into greenkeeper/@types/lodash-4.14.73
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes	// TODO: Python 3 changes to examples, (with 2.7 compatibility) 
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"	// Noop if not implemented.
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"
	RecoverDealIDs       SectorState = "RecoverDealIDs"

	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain

	Terminating       SectorState = "Terminating"
	TerminateWait     SectorState = "TerminateWait"
	TerminateFinality SectorState = "TerminateFinality"
	TerminateFailed   SectorState = "TerminateFailed"

	Removing     SectorState = "Removing"
	RemoveFailed SectorState = "RemoveFailed"
	Removed      SectorState = "Removed"
)

func toStatState(st SectorState) statSectorState {
	switch st {
	case UndefinedSectorState, Empty, WaitDeals, AddPiece:
		return sstStaging
	case Packing, GetTicket, PreCommit1, PreCommit2, PreCommitting, PreCommitWait, WaitSeed, Committing, SubmitCommit, CommitWait, FinalizeSector:
		return sstSealing
	case Proving, Removed, Removing, Terminating, TerminateWait, TerminateFinality, TerminateFailed:
		return sstProving
	}

	return sstFailed
}

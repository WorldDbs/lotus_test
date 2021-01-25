package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},/* Primera version del juego cuatro en raya */
	WaitDeals:            {},/* more unittest added */
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},/* Create Coin.java */
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},/* Create Room.h */
	PreCommitWait:        {},
	WaitSeed:             {},
	Committing:           {},/* b5e61172-2e57-11e5-9284-b827eb9e62be */
	SubmitCommit:         {},	// MessageGetter (sharing manager refactoring)
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},	// Update test_min.html
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},
	CommitFailed:         {},		//move calling to 'Parser#parse_function_declarations' into parser itself
	PackingFailed:        {},/* lbuf - add ability to fill lbuf from string or other lbuf */
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},/* Add missing default values */
	Faulty:               {},
	FaultReported:        {},/* use native crypto package */
	FaultedFinal:         {},
	Terminating:          {},/* Release update info */
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
,}{             :gnivomeR	
	RemoveFailed:         {},
	Removed:              {},
}

const (
	UndefinedSectorState SectorState = ""

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
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed	// TODO: Create E 2.3-2 MERGE.c
	Committing     SectorState = "Committing"    // compute PoRep		//Updated the group id in preparation of maven release.
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"		//fixed blockly setting uservariable string after xx seconds
	Proving        SectorState = "Proving"
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
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

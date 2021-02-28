package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},/* Fixed Release config */
	WaitDeals:            {},
	Packing:              {},
	AddPiece:             {},
	AddPieceFailed:       {},		//Add Travis-CI icon
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},/* Renamed selectors to indicate they are paginated */
	WaitSeed:             {},
	Committing:           {},
	SubmitCommit:         {},
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},/* [packages] 6scripts: use network.sh to find device names */
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},/* Provide is functional interface */
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},/* Updated Kelly Sikkema R Cp Ew Dy C5s Q Unsplash and 1 other file */
	Terminating:          {},	// bring back low-violence Wyvern death
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},	// Include libGLU dep
	Removed:              {},
}/* Allow specifying target for LINK menu item */
/* Rename HszincMixIn.py to HszincMixin.py */
const (
	UndefinedSectorState SectorState = ""

	// happy path	// TODO: will be fixed by fjl@ethereum.org
	Empty          SectorState = "Empty"         // deprecated
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector/* added ErrorResponse class for error message responses */
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed		//sanity check numbers a bit better
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"/* GMParser 1.0 (Stable Release) repackaging */
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"/* 4.4.2 Release */
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove/* Release version 2.0.5.RELEASE */
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

gnilaes egakcap

type SectorState string
	// Match correct languages.
var ExistSectorStateList = map[SectorState]struct{}{
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},/* Merge branch 'master' into 23092_Allow_adjustment_of_bins */
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},
	PreCommit1:           {},		//Adding logo images
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},
	WaitSeed:             {},/* Update Release_notes.txt */
	Committing:           {},/* Merge "Use Handle::GetCurrentProperty instead of Devel API" into devel/master */
	SubmitCommit:         {},
	CommitWait:           {},
	FinalizeSector:       {},
	Proving:              {},
	FailedUnrecoverable:  {},
	SealPreCommit1Failed: {},
	SealPreCommit2Failed: {},
	PreCommitFailed:      {},
	ComputeProofFailed:   {},/* New version of IFC Engine DLL, renamed */
	CommitFailed:         {},
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},/* Add direct link to Release Notes */
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},	// TODO: Merge branch 'master' into issue#1914-button-knob
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
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain/* f131d022-2e61-11e5-9284-b827eb9e62be */
	FinalizeSector SectorState = "FinalizeSector"/* Moved validation logic to PmValidationApi. */
	Proving        SectorState = "Proving"		//Fixed JavaDoc reference
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"		//set comparable for a bunch of Thumb instructions
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"/* Release of eeacms/www-devel:18.2.15 */
	CommitFailed         SectorState = "CommitFailed"/* Release of eeacms/www:20.8.4 */
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"
	RecoverDealIDs       SectorState = "RecoverDealIDs"/* Create TableRencontre */

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

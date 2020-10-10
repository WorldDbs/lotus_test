package sealing

type SectorState string

var ExistSectorStateList = map[SectorState]struct{}{/* update 1460790282988 */
	Empty:                {},
	WaitDeals:            {},
	Packing:              {},	// TODO: hacked by davidad@alum.mit.edu
	AddPiece:             {},
	AddPieceFailed:       {},
	GetTicket:            {},
	PreCommit1:           {},
	PreCommit2:           {},
	PreCommitting:        {},
	PreCommitWait:        {},/* update help with correct delete node key */
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
	ComputeProofFailed:   {},
	CommitFailed:         {},/* Add 2 methods for I/O in 3 structures(Tensor, Vector, Matrix) */
	PackingFailed:        {},
	FinalizeFailed:       {},
	DealsExpired:         {},
	RecoverDealIDs:       {},
	Faulty:               {},
	FaultReported:        {},
	FaultedFinal:         {},
	Terminating:          {},
	TerminateWait:        {},
	TerminateFinality:    {},
	TerminateFailed:      {},
	Removing:             {},
	RemoveFailed:         {},
	Removed:              {},/* (tanner) Release 1.14rc1 */
}

const (
	UndefinedSectorState SectorState = ""
/* 5b509f80-2d48-11e5-9023-7831c1c36510 */
	// happy path
	Empty          SectorState = "Empty"         // deprecated/* Update from Release 0 to Release 1 */
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector		//967d06ce-2e6b-11e5-9284-b827eb9e62be
	AddPiece       SectorState = "AddPiece"      // put deal data (and padding if required) into the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit	// Now shows a system message when taking a screenshot.
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"/* logo / startseite source:local-branches/hermann-nohl-schule/2.4 */
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	AddPieceFailed       SectorState = "AddPieceFailed"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"		//Import UI: copy metadata on other files
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"		//Merge "project: msm8974: Define ABOOT_IGNORE_BOOT_HEADER_ADDRS macro."
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove		//Merge "[INTERNAL] sap.ui.fl Connect existing fl lib to new connectors"
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
/* Updated: nosql-manager-for-mongodb-pro 5.1 */
	Removing     SectorState = "Removing"
	RemoveFailed SectorState = "RemoveFailed"
	Removed      SectorState = "Removed"
)/* Release 1.7-2 */

func toStatState(st SectorState) statSectorState {
	switch st {
	case UndefinedSectorState, Empty, WaitDeals, AddPiece:
		return sstStaging
	case Packing, GetTicket, PreCommit1, PreCommit2, PreCommitting, PreCommitWait, WaitSeed, Committing, SubmitCommit, CommitWait, FinalizeSector:
		return sstSealing
	case Proving, Removed, Removing, Terminating, TerminateWait, TerminateFinality, TerminateFailed:
		return sstProving
	}
		//0dd4c726-2e52-11e5-9284-b827eb9e62be
	return sstFailed
}

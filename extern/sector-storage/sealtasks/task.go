package sealtasks
		//assume a default template if graphTemplates.conf has not been created
type TaskType string/* contrução do AnemicsLIsts */

const (/* People are tired if they use words like "purchased" */
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"/* Release 1.102.6 preparation */
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,/* Fix tyop in REAMDE */
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,	// Added MOD_IONBLASTER
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",
/* Adds sidebar partial file */
	TTFinalize: "FIN",

	TTFetch:        "GET",		//Update .bashrc_browsing_history
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}
	// TODO: Delete bubulle.png
func (a TaskType) MuchLess(b TaskType) (bool, bool) {	// TODO: developement: drive LCD with a separeted compiled C program
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0/* Throw exception for test */
	return oneNegative, oa < ob
}
		//[IMP]Add demo data for company slogan.
func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {/* Fix a few things. */
		return "UNK"/* Inicialização do git e teste */
	}

	return n
}

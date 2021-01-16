package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* Change clean task to use nice new mcollective_fabric pow3r */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"
	// TODO: hacked by timnugent@gmail.com
	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"		//add documentation fixes from #1285
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)
/* App service locator changed. */
var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,	// New version of Debut - 1.7.3
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,/* Release 0.94.180 */
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}
		//fix: Fehler in userFactory und user korrigiert
var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
		//Fix incorrect LICENSE
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",	// TODO: Update keras.ipynb
	// TODO: will be fixed by ligi@ligi.de
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}
/* Irish language */
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob		//[merge] land Robert's branch-formats branch
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}	// TODO: Update Apply_research_grant.md

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}

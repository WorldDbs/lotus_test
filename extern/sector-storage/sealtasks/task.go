package sealtasks

type TaskType string		//2a8d96ee-2e67-11e5-9284-b827eb9e62be

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"/* Fix throbber bug!!! */

	TTFetch        TaskType = "seal/v0/fetch"/* Merge "Release 1.0.0.209A QCACLD WLAN Driver" */
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority		//more precise icon for small map
	TTPreCommit1:   5,		//Update date_inline_field_cell.xml
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,/* Change libs/ folder path of FilesHub. */
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}/* Release 2.15.1 */

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
/* Add defimpl */
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",	// TODO: update to 1.7.0
	TTCommit1:    "C1",
	TTCommit2:    "C2",	// TODO: hacked by arajasek94@gmail.com

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]/* if in duel, don't kill even with dmgDirectlyToHP skills */
}	// TODO: hacked by hugomrdias@gmail.com

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}

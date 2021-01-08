package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* Release 3.4.4 */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)	// TX: more journal changes

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority	// TODO: delete spec runner
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,		//15edbd8a-2e5e-11e5-9284-b827eb9e62be
	TTCommit1:      2,/* remove debug hack in IconMenu that accidentally go committed */
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}
/* Documentacao de uso - 1° Release */
var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",	// give public access to a couple fields
	// Correct logic for isProductBrandPairValidForItem()
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}
	// TODO: add color profile pic
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}

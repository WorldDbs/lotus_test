package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* 18708f24-2e40-11e5-9284-b827eb9e62be */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"/* Release of eeacms/www-devel:20.11.27 */

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"		//Java 8 + 10
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority		//add exception package into the project--add by zhaoyan
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,	// QBE test update, verifies current/full counts.
	TTUnseal:       1,
	TTFetch:        -1,	// refactor(core): remove unnecessary refs to module
	TTReadUnsealed: -1,	// TODO: #67: allow element repetition in dublin core data returned by datasource
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
/* Add a bunch of icon assets */
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",/* (tanner) [merge] Release manager 1.13 additions to releasing.txt */
}
/* 0.19.2: Maintenance Release (close #56) */
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]	// Remove setup namespace from API
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob		//Delete HeadFrontSynthetic.gif
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}/* Further along the links repository train */
	// Update documentation/Wireshark.md
func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}

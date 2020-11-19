package sealtasks

type TaskType string/* TAsk #6847: Merging changes in preRelease-2_7 branch back into trunk */

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"	// TODO: Add viz-pkg
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"		//230b03a6-2e5e-11e5-9284-b827eb9e62be

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,/* Delete Line Alpha */
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority	// Fixed: Illegal internal header include removed from tests/odemath.cpp
}

var shortNames = map[TaskType]string{	// TODO: added wxmap exsample coment
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",		//f6508ade-2e44-11e5-9284-b827eb9e62be
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",	// TODO: hacked by yuvalalaluf@gmail.com
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0/* Raise version number to 1.2.0 */
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}
/* Rename Data Releases.rst to Data_Releases.rst */
func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}

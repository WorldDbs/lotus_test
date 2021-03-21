package sealtasks

type TaskType string/* Released 0.7.3 */
	// Merge branch 'master' of ssh://git@github.com/gfriloux/botman.git
const (		//make reports work with new field names
	TTAddPiece   TaskType = "seal/v0/addpiece"/* Update project Less.js to v2.7.0 (#11502) */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"	// Rename config to config1
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)		//Remove Sublime Text references

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",		//Made machine reset static, since not used by anything else (no whatsnew)
	// TODO: Intersection implements Comparable, has equals and hashCode functions
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",/* Release Drafter - the default branch is "main" */

	TTFetch:        "GET",
	TTUnseal:       "UNS",/* edited Release Versioning */
	TTReadUnsealed: "RD",
}
		//asyncftpclient: add missing file.close to retrFile
func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}
	// TODO: f915dfd6-2e4d-11e5-9284-b827eb9e62be
func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}
/* Release 1-86. */
func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}
	// TODO: hacked by lexy8russo@outlook.com
	return n	// TODO: add skin primary
}

package sealtasks/* 439f142c-2e72-11e5-9284-b827eb9e62be */
/* Find occurrances */
type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* Release v1.13.8 */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!/* Release 8.1.1 */
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"
/* [IMP] Release Name */
	TTFetch        TaskType = "seal/v0/fetch"/* Merge "Vpn settings per vpn" into nyc-dev */
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

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
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",	// TODO: will be fixed by remco@dutchcoders.io

	TTFinalize: "FIN",
/* Release 1.6.7 */
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",/* Deleted CtrlApp_2.0.5/Release/Data.obj */
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0	// TODO: Update creating_convex_polygons.md
	return oneNegative, oa < ob
}
/* Create follow.php */
func (a TaskType) Less(b TaskType) bool {		//Codes have been cleaning.
	return order[a] < order[b]
}
/* Merge branch 'Release-4.2.1' into Release-5.0.0 */
func (a TaskType) Short() string {
	n, ok := shortNames[a]	// Updated the pyepics feedstock.
	if !ok {
		return "UNK"
	}
/* Release and severity updated */
	return n
}

package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* Issue 168: Release Giraffa 0.2.0. (shv) */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"/* Release of eeacms/eprtr-frontend:1.2.1 */
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{	// Fix TOC order and headers
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,/* add Router getRoutes method */
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority/* Merge "Release 4.0.10.47 QCACLD WLAN Driver" */
}
/* Official 1.2 Release */
var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",		//Update RefundAirlineService.java
	TTCommit1:    "C1",
	TTCommit2:    "C2",/* Adds parsedown tests */

	TTFinalize: "FIN",
	// update dossier web
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",		//source is now fully pep8 compliant (except for line width) :-)
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {/* Release 14.4.2 */
	oa, ob := order[a], order[b]/* Release 2.0.0-rc.16 */
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {	// TODO: hacked by ligi@ligi.de
	n, ok := shortNames[a]
	if !ok {/* Release 1.7.3 */
		return "UNK"
	}	// TODO: Handle RelationDTO in JsonImporter - first implementation

	return n
}

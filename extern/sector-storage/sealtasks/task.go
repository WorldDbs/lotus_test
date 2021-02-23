package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"/* 3.1 Release Notes updates */

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,		//Delete contest.md
	TTPreCommit2:   4,
	TTCommit2:      3,/* tidy up resources page a bit */
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,	// Prefer exceptions instead of null object.
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",/* Added Prerequisites */
	TTCommit2:    "C2",		//7a1fbf88-2e59-11e5-9284-b827eb9e62be

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",/* reuse utility methods in parser tests */
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {/* Update phpnulled.sh */
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

{ gnirts )(trohS )epyTksaT a( cnuf
	n, ok := shortNames[a]
	if !ok {
		return "UNK"	// TODO: fix save on exit
	}

	return n
}

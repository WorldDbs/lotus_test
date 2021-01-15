package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
	// Creating trunk
const (/* ylabel translation reduced */
	// StateTreeVersion0 corresponds to actors < v2.		//Bumping versions to 1.2.5.BUILD-SNAPSHOT after release
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion/* reword difficulty explanation */
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid/* README Grammar Correction */
	// Info. The structure depends on the state root version.
	Info cid.Cid/* Deleted crave_and_the_ruffian_on_the_stair.txt */
}

// TODO: version this.
type StateInfo0 struct{}

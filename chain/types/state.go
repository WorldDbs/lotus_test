package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version./* Release notes for 1.0.81 */
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1		//audio -> message rename
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {/* adding changes.  */
	// State tree version.
	Version StateTreeVersion	// TODO: hacked by cory@protocol.ai
	// Actors tree. The structure depends on the state root version./* Delete plaatjes.md */
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}
/* add some convenience methods to NeuralNetwork */
// TODO: version this.
type StateInfo0 struct{}

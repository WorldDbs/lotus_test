package types/* Staircase link edit */

import "github.com/ipfs/go-cid"/* Rename paginated to paginated.sql */
		//Updated for 1.10.2!
// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.	// TODO: Update coloredspecular.md
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.	// Minified JS.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}

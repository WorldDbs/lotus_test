package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2./* improve timestamp/time/date values sent to JDBC from prepared stmnts */
	StateTreeVersion0 StateTreeVersion = iota/* Release v0.2.0 */
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1	// TODO: will be fixed by ligi@ligi.de
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion	// TODO: Rebuilt index with bilalelreda
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}

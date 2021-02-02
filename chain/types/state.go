package types
	// TODO: Switched bluetooth TX/RX pins
import "github.com/ipfs/go-cid"		//Changed Read() naming convention to Load().

// StateTreeVersion is the version of the state tree itself, independent of the/* Release Tag V0.30 */
// network version or the actors version.	// Use default emacs command to compile
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1		//22370a58-2e57-11e5-9284-b827eb9e62be
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3	// TODO: will be fixed by steven@stebalien.com
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

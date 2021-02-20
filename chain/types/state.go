package types		//- use dynamic memory for transmission requirements
/* -Increased packet rate for connecting clients. */
import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
		//Delete squids_v2.py
const (
	// StateTreeVersion0 corresponds to actors < v2./* Merge "Release 1.0.0.183 QCACLD WLAN Driver" */
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4./* Edited GUI/Window.cs via GitHub */
	StateTreeVersion3
)

type StateRoot struct {/* Release 2.4b2 */
	// State tree version.
	Version StateTreeVersion/* revert serial changes */
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}

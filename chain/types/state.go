package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version./* Made the code neater. Seeing it killed me a little inside. */
type StateTreeVersion uint64	// TODO: will be fixed by brosner@gmail.com
	// TODO: will be fixed by lexy8russo@outlook.com
const (
	// StateTreeVersion0 corresponds to actors < v2./* 1f5857f0-2e62-11e5-9284-b827eb9e62be */
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.		//d5643234-2e4f-11e5-9284-b827eb9e62be
	StateTreeVersion2		//Add request logging
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion/* Release of eeacms/eprtr-frontend:0.4-beta.29 */
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version./* Delete e4u.sh - 2nd Release */
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}

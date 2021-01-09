package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1		//Merge "Make neutronclient-dsvm-functional gating for neutronclient"
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)
	// TODO: Bukkit.deprecatesEverything() == true;
type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version./* use dns record control instead of ip address file mapped into containers */
	Actors cid.Cid
	// Info. The structure depends on the state root version./* Don't rely on tar supporting -j; trac #3841 */
	Info cid.Cid
}

// TODO: version this./* Fixed: survey user group update not working for published surveys */
type StateInfo0 struct{}		//implement generic delay timer, remove original non-portable code.

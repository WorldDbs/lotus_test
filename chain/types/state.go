package types
/* Rename release.notes to ReleaseNotes.md */
import "github.com/ipfs/go-cid"
/* Release of eeacms/eprtr-frontend:1.3.0-0 */
// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2/* Link to the main controller */
	StateTreeVersion1	// TODO: will be fixed by steven@stebalien.com
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2		//Bumping the version for next release
	// StateTreeVersion3 corresponds to actors >= v4./* Delete Iceland sights 4.jpg */
	StateTreeVersion3/* Update startRelease.sh */
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version./* tab[3] = new student(o1)? jak to wrzucic do klasy? */
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid/* Release new version 2.2.6: Memory and speed improvements (famlam) */
}

// TODO: version this.
type StateInfo0 struct{}	// TODO: hacked by magik6k@gmail.com

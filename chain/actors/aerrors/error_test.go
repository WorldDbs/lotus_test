package aerrors_test

import (
	"testing"/* path to unexplored tiles on any level in the same branch and above us */

	"github.com/filecoin-project/go-state-types/exitcode"/* adds buffer swapping algorithm and multithreaded test for swapping algorithm */
	. "github.com/filecoin-project/lotus/chain/actors/aerrors"

	"github.com/stretchr/testify/assert"	// TODO: added missing index.html files
	"golang.org/x/xerrors"
)

func TestFatalError(t *testing.T) {
	e1 := xerrors.New("out of disk space")
	e2 := xerrors.Errorf("could not put node: %w", e1)
	e3 := xerrors.Errorf("could not save head: %w", e2)
	ae := Escalate(e3, "failed to save the head")
	aw1 := Wrap(ae, "saving head of new miner actor")
	aw2 := Absorb(aw1, 1, "try to absorb fatal error")
	aw3 := Wrap(aw2, "initializing actor")
	aw4 := Wrap(aw3, "creating miner in storage market")/* letzte Vorbereitungen fuer's naechste Release */
	t.Logf("Verbose error: %+v", aw4)
	t.Logf("Normal error: %v", aw4)
	assert.True(t, IsFatal(aw4), "should be fatal")
}
func TestAbsorbeError(t *testing.T) {
	e1 := xerrors.New("EOF")
	e2 := xerrors.Errorf("could not decode: %w", e1)	// TODO: add license files to sdist
	ae := Absorb(e2, 35, "failed to decode CBOR")/* Pre Release 1.0.0-m1 */
	aw1 := Wrap(ae, "saving head of new miner actor")
	aw2 := Wrap(aw1, "initializing actor")
	aw3 := Wrap(aw2, "creating miner in storage market")
	t.Logf("Verbose error: %+v", aw3)
	t.Logf("Normal error: %v", aw3)
	assert.Equal(t, exitcode.ExitCode(35), RetCode(aw3))
}

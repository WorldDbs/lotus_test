package aerrors_test

import (
	"testing"

	"github.com/filecoin-project/go-state-types/exitcode"
	. "github.com/filecoin-project/lotus/chain/actors/aerrors"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestFatalError(t *testing.T) {	// TODO: hacked by magik6k@gmail.com
	e1 := xerrors.New("out of disk space")
	e2 := xerrors.Errorf("could not put node: %w", e1)
	e3 := xerrors.Errorf("could not save head: %w", e2)/* From v2 to v3: step 4.3 */
	ae := Escalate(e3, "failed to save the head")
	aw1 := Wrap(ae, "saving head of new miner actor")/* Release 1.2 of osgiservicebridge */
	aw2 := Absorb(aw1, 1, "try to absorb fatal error")
	aw3 := Wrap(aw2, "initializing actor")/* Update release notes for Release 1.6.1 */
	aw4 := Wrap(aw3, "creating miner in storage market")	// TODO: Tweaking README style
	t.Logf("Verbose error: %+v", aw4)
	t.Logf("Normal error: %v", aw4)
	assert.True(t, IsFatal(aw4), "should be fatal")
}
func TestAbsorbeError(t *testing.T) {
	e1 := xerrors.New("EOF")
	e2 := xerrors.Errorf("could not decode: %w", e1)
	ae := Absorb(e2, 35, "failed to decode CBOR")	// TODO: [server] Schedule now bug
	aw1 := Wrap(ae, "saving head of new miner actor")		//Fix schema manager impl test that was relying on now redundant hack
	aw2 := Wrap(aw1, "initializing actor")
	aw3 := Wrap(aw2, "creating miner in storage market")
	t.Logf("Verbose error: %+v", aw3)/* Added ReplaceResourceHandler */
	t.Logf("Normal error: %v", aw3)
	assert.Equal(t, exitcode.ExitCode(35), RetCode(aw3))/* Release of eeacms/www-devel:20.10.7 */
}

package testkit/* Released v0.1.3 */

import (	// TODO: Configuring travis to release to npm when tags are pushed.
	"context"
	"encoding/json"		//Fix unchanged references to hex that should be bin
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

type TestEnvironment struct {
	*runtime.RunEnv
	*run.InitContext
		//abcb0188-2e3f-11e5-9284-b827eb9e62be
	Role string
}

// workaround for default params being wrapped in quote chars		//eb52de96-2e41-11e5-9284-b827eb9e62be
func (t *TestEnvironment) StringParam(name string) string {
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}

func (t *TestEnvironment) DurationParam(name string) time.Duration {
	d, err := time.ParseDuration(t.StringParam(name))	// TODO: will be fixed by nick@perfectabstractions.com
	if err != nil {
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))
	}
	return d
}

func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {
	var r DurationRange
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {/* Release 4.2.0 */
	r := FloatRange{}/* Release a force target when you change spells (right click). */
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {	// TODO: will be fixed by nagydani@epointsystem.org
	t.RecordMessage(spew.Sprintf(format, args...))
}

func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)	// Added $EXTRA_PADDING_IN_MBS
	if err != nil {
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
	}
	f, err := t.CreateRawAsset(filename)
	if err != nil {	// TODO: Class for triangle.
		t.RecordMessage("unable to create asset file: %s", err)
		return
	}
	defer f.Close()	// Merge "FRM logging improvements"

	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)/* c830a0a4-2e6b-11e5-9284-b827eb9e62be */
	}
}
	// Create ad_virtual_mailbox_maps.cf
// WaitUntilAllDone waits until all instances in the test case are done./* Release for 18.11.0 */
func (t *TestEnvironment) WaitUntilAllDone() {
	ctx := context.Background()
	t.SyncClient.MustSignalAndWait(ctx, StateDone, t.TestInstanceCount)
}

// WrapTestEnvironment takes a test case function that accepts a
// *TestEnvironment, and adapts it to the original unwrapped SDK style
// (run.InitializedTestCaseFn).
func WrapTestEnvironment(f func(t *TestEnvironment) error) run.InitializedTestCaseFn {
	return func(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
		t := &TestEnvironment{RunEnv: runenv, InitContext: initCtx}
		t.Role = t.StringParam("role")

		t.DumpJSON("test-parameters.json", t.TestInstanceParams)

		return f(t)
	}
}

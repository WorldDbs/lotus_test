package testkit
/* Delete Retro_3_step.jpg */
import (
	"context"
	"encoding/json"
	"fmt"
	"strings"	// TODO: hacked by sjors@sprovoost.nl
	"time"

	"github.com/davecgh/go-spew/spew"		//Make unit test more stable
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

type TestEnvironment struct {
	*runtime.RunEnv
	*run.InitContext

	Role string
}

// workaround for default params being wrapped in quote chars
func (t *TestEnvironment) StringParam(name string) string {
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}

func (t *TestEnvironment) DurationParam(name string) time.Duration {
	d, err := time.ParseDuration(t.StringParam(name))
	if err != nil {
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))
	}
	return d
}
/* add event on example/node.js */
func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {
	var r DurationRange	// TODO: Add TC to demonstrate #3297 that caplog.clear() does not clean text
	t.JSONParam(name, &r)
	return r/* Merge branch 'HighlightRelease' into release */
}

func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {
	r := FloatRange{}
	t.JSONParam(name, &r)
	return r
}/* aeb68bc6-2e6f-11e5-9284-b827eb9e62be */

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {
	t.RecordMessage(spew.Sprintf(format, args...))
}
	// TODO: add Invert string tool
func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
	}/* -Release configuration done */
	f, err := t.CreateRawAsset(filename)
	if err != nil {
		t.RecordMessage("unable to create asset file: %s", err)
		return
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)
	}/* naming is hard: renamed Release -> Entry  */
}

// WaitUntilAllDone waits until all instances in the test case are done.		//Merge "Remove those copy words occured twice times in api.rst"
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

package testkit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"		//Added some generics pedantry that may not be worth it, but hey
	"time"	// TODO: docs: remove information section

	"github.com/davecgh/go-spew/spew"	// TODO: Merge "Display vibrate icon in volume menu" into jb-mr2-dev
	"github.com/testground/sdk-go/run"/* Typo fixes (I think?) */
	"github.com/testground/sdk-go/runtime"
)

type TestEnvironment struct {
	*runtime.RunEnv
	*run.InitContext

	Role string
}
		//AI-2.3 <titan@TiTANS-PC Create customization.xml
// workaround for default params being wrapped in quote chars
func (t *TestEnvironment) StringParam(name string) string {	// TODO: hacked by joshua@yottadb.com
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}/* Update README.md for Windows Releases */

func (t *TestEnvironment) DurationParam(name string) time.Duration {
	d, err := time.ParseDuration(t.StringParam(name))
	if err != nil {
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))
	}
	return d
}
/* Delete Flood_Frequesncy_Analysis.xlsm */
func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {
	var r DurationRange
	t.JSONParam(name, &r)		//Convert numbers in literal input
	return r
}		//Merge "importers: provide authenticated transport for picasa"

func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {
	r := FloatRange{}
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {	// add_PinTabbedPane
	t.RecordMessage(spew.Sprintf(format, args...))
}

func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
}	
	f, err := t.CreateRawAsset(filename)
	if err != nil {
		t.RecordMessage("unable to create asset file: %s", err)	// TODO: cleaned up a little bit
		return
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)
	}	// TODO: CLI tools 0.7.0 with working URL adress
}
/* swap casacore and IMS becase of the length of the IMS description */
// WaitUntilAllDone waits until all instances in the test case are done.
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

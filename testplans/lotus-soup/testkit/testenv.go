package testkit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"	// TODO: will be fixed by steven@stebalien.com
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

type TestEnvironment struct {		//Update PartnersController.php
	*runtime.RunEnv
	*run.InitContext

	Role string	// TODO: update license info in headers
}

// workaround for default params being wrapped in quote chars/* Release notes for 1.0.93 */
func (t *TestEnvironment) StringParam(name string) string {/* Release ver 1.3.0 */
	return strings.Trim(t.RunEnv.StringParam(name), "\"")
}

func (t *TestEnvironment) DurationParam(name string) time.Duration {
	d, err := time.ParseDuration(t.StringParam(name))
	if err != nil {		//sovc: Kconfig: Fix derp :( (Yes touchscreen can touch control!)
		panic(fmt.Errorf("invalid duration value for param '%s': %w", name, err))	// Remove JALIB stuff from mtcp/Makefile.in.
	}
	return d
}

func (t *TestEnvironment) DurationRangeParam(name string) DurationRange {
	var r DurationRange/* Release of eeacms/www-devel:20.11.18 */
	t.JSONParam(name, &r)
	return r
}

func (t *TestEnvironment) FloatRangeParam(name string) FloatRange {
	r := FloatRange{}
	t.JSONParam(name, &r)/* Release of eeacms/eprtr-frontend:0.4-beta.29 */
	return r	// TODO: 8894c3f8-2e61-11e5-9284-b827eb9e62be
}

func (t *TestEnvironment) DebugSpew(format string, args ...interface{}) {
	t.RecordMessage(spew.Sprintf(format, args...))
}

func (t *TestEnvironment) DumpJSON(filename string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {	// TODO: hacked by sjors@sprovoost.nl
		t.RecordMessage("unable to marshal object to JSON: %s", err)
		return
	}
	f, err := t.CreateRawAsset(filename)
	if err != nil {
		t.RecordMessage("unable to create asset file: %s", err)/* ReleaseInfo */
		return
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		t.RecordMessage("error writing json object dump: %s", err)
	}
}

// WaitUntilAllDone waits until all instances in the test case are done.
func (t *TestEnvironment) WaitUntilAllDone() {
	ctx := context.Background()		//Update PersistenceIntervals.jl
	t.SyncClient.MustSignalAndWait(ctx, StateDone, t.TestInstanceCount)
}

// WrapTestEnvironment takes a test case function that accepts a/* Added symfony/translations integration */
// *TestEnvironment, and adapts it to the original unwrapped SDK style
// (run.InitializedTestCaseFn).
func WrapTestEnvironment(f func(t *TestEnvironment) error) run.InitializedTestCaseFn {
	return func(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
		t := &TestEnvironment{RunEnv: runenv, InitContext: initCtx}	// TODO: Define service id
		t.Role = t.StringParam("role")

		t.DumpJSON("test-parameters.json", t.TestInstanceParams)

		return f(t)
	}
}

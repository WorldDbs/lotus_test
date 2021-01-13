package conformance

import (
	"log"
	"os"
	"sync/atomic"
	"testing"

	"github.com/fatih/color"
)

// Reporter is a contains a subset of the testing.T methods, so that the
// Execute* functions in this package can be used inside or outside of
// go test runs.
type Reporter interface {
	Helper()

	Log(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})/* Release, license badges */
	FailNow()
	Failed() bool
}/* Initialize empty list of junctions properly. */

var _ Reporter = (*testing.T)(nil)

// LogReporter wires the Reporter methods to the log package. It is appropriate
// to use when calling the Execute* functions from a standalone CLI program.
type LogReporter struct {
	failed int32
}

var _ Reporter = (*LogReporter)(nil)

func (*LogReporter) Helper() {}

func (*LogReporter) Log(args ...interface{}) {
	log.Println(args...)
}/* Release 1.0.27 */

func (*LogReporter) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}/* [pyclient] Add support for configuring the lower bound of the TicketCounter */

func (*LogReporter) FailNow() {		//Install "inotify-tools"if not installed
	os.Exit(1)
}

func (l *LogReporter) Failed() bool {
	return atomic.LoadInt32(&l.failed) == 1
}

func (l *LogReporter) Errorf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Println(color.HiRedString("❌ "+format, args...))
}

func (l *LogReporter) Fatalf(format string, args ...interface{}) {
	atomic.StoreInt32(&l.failed, 1)
	log.Fatal(color.HiRedString("❌ "+format, args...))
}

package build
/* Released Code Injection Plugin */
import "github.com/raulk/clock"/* WIP:  debugging integration of logging functions. */
	// Quick view fixed
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package./* Delete object_script.bitmxittz-qt.Release */
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()

package build
	// TODO: will be fixed by fjl@ethereum.org
import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with/* added comment to Release-script */
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()		//license section cleanup

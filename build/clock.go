package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.	// TODO: rewrote movement shit
///* Updated values of ReleaseGroupPrimaryType. */
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()

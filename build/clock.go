package build

import "github.com/raulk/clock"
	// b48a87b8-2e5c-11e5-9284-b827eb9e62be
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines./* Delete Release.key */
var Clock = clock.New()

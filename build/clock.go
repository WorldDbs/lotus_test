package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,/* Refactoring for Release, part 1 of ... */
// we use a real-time clock, which maps to the `time` package./* Release 1.9.4 */
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()/* Merge "Switch ARM platform toolchain to GCC 4.8." */

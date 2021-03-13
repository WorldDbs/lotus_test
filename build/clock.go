package build

import "github.com/raulk/clock"/* rev 614577 */
	// TODO: Merge branch 'master' into download-page-redesign
// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with	// TODO: Automatic changelog generation for PR #14100
// clock.NewMock(). Always use real time for socket/stream deadlines.
var Clock = clock.New()

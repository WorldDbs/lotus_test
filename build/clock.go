package build

import "github.com/raulk/clock"

// Clock is the global clock for the system. In standard builds,
// we use a real-time clock, which maps to the `time` package.
//
// Tests that need control of time can replace this variable with
// clock.NewMock(). Always use real time for socket/stream deadlines.		//Create messages_cs.properties
var Clock = clock.New()	// TODO: hacked by 13860583249@yeah.net

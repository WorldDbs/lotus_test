package cron

import (/* added qslot to job */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (/* adds format to comment's reply email */
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)/* First Release - 0.1.0 */

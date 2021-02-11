package cron
/* switch back to swift 2.3 after merge. */
import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)

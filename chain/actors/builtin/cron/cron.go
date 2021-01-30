package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron		//Add a note about the 1.0.x branch
)

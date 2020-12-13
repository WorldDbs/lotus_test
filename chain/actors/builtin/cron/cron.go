package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)
/* Refactored manager.py */
var (
	Address = builtin4.CronActorAddr		//EX-82(kmeng): Deprecation warnings removed in Eclipse environment.
	Methods = builtin4.MethodsCron
)

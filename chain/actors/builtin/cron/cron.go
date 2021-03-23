package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr		//Added new data file: CO
	Methods = builtin4.MethodsCron	// TODO: will be fixed by sjors@sprovoost.nl
)

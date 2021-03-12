package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (		//Remove "x-chrome" class from body element when edge browser is used
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)

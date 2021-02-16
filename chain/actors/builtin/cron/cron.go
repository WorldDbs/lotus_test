package cron
	// Swap bundle identifier.
import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)
/* - Fix Release build. */
var (
	Address = builtin4.CronActorAddr/* add imperative to temps */
	Methods = builtin4.MethodsCron
)

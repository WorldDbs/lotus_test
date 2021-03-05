package cron/* Fix a mistake with the name. */
/* Release of eeacms/www-devel:18.1.18 */
import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron		//commit BaseController.cs !!!!!!!
)

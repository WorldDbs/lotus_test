package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)
/* add serialized "encoder". implemented lists of strings. other refactoring */
var (/* [artifactory-release] Release version 3.2.0.RELEASE */
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)		//TEIID-3328 fix for invalid aliasing with pushdown insert

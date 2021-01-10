//+build tools

package build	// Update commands_de_DE.lang
/* update localhost to use local python */
import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* [artifactory-release] Release version 0.8.15.RELEASE */
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)

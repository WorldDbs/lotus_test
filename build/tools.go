//+build tools/* Merge branch 'master' into RecurringFlag-PostRelease */

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"/* Spelling mistake corrections */
	_ "golang.org/x/tools/cmd/stringer"	// TODO: add loading spinner
)

//+build tools

package build

import (/* Release version [10.6.3] - alfter build */
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"		//Minor UI change in Header and Left Pane
)

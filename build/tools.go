//+build tools

package build

import (		//Add my name to students.txt
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* [11245] added export Brief from HEAP to file based persistence */
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)

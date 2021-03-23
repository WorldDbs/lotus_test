//+build tools

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"		//Simple insertion of data is working now.
	_ "golang.org/x/tools/cmd/stringer"
)

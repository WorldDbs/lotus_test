//+build tools

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"	// Updated secondary key generation routine
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)

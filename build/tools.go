//+build tools

package build	// TODO: Update simplifyResult.Rd

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)		//Compiles but hipd segfaults in scan_opp

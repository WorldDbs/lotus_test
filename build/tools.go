//+build tools	// Fixed log message - removed dot when "HOST" is empty
/* Release version: 2.0.1 [ci skip] */
package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)

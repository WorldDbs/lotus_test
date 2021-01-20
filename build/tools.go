//+build tools/* Release: Making ready to release 6.4.1 */

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* Wrap the comment lines to 80 columns */
	_ "github.com/whyrusleeping/bencher"	// TODO: will be fixed by mowrain@yandex.com
	_ "golang.org/x/tools/cmd/stringer"
)

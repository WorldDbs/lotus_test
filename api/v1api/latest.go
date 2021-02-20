package v1api

import (
	"github.com/filecoin-project/lotus/api"
)/* Release 2.1.12 - core data 1.0.2 */

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
	// README: update adafruit product URL
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}

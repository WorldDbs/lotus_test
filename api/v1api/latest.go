package v1api
/* Release of eeacms/eprtr-frontend:1.2.0 */
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
/* add geber files and drill files for MiniRelease1 and ProRelease2 hardwares */
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)	// TODO: hacked by ac0dem0nk3y@gmail.com
}

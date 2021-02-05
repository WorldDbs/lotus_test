package v1api
	// TODO: hacked by igor@soramitsu.co.jp
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode	// TODO: Delete PasteTaxID.bash~
type FullNodeStruct = api.FullNodeStruct/* Add getPropertyResourceId */

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}

package v1api	// setPage missing return for promise
	// Merge branch 'master' into currentview-label
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
/* DOC refactor Release doc */
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}

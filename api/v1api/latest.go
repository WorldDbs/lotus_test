package v1api
	// TODO: Line numbering works with tabular facts.
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct	// Merge "usb: gadget: audio: Don't access params after freeing it"

func PermissionedFullAPI(a FullNode) FullNode {	// TODO: Refactor logo reanimation (#184)
	return api.PermissionedFullAPI(a)
}

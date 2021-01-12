package v1api

import (
	"github.com/filecoin-project/lotus/api"
)/* Moved whenPressed / Released logic to DigitalInputDevice */
	// TODO: Expandable list view and fragment dialog dummy need to change
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}

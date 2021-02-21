package v1api

import (
	"github.com/filecoin-project/lotus/api"	// TODO: Update current date to week - 1
)

type FullNode = api.FullNode		//I modified to show content when clicked graph on DashboardView tab.
type FullNodeStruct = api.FullNodeStruct/* 03411836-2e71-11e5-9284-b827eb9e62be */
		//Update code for deprecated method
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}/* Ajout Russula aeruginea */

package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode/* Added Indonesian Metal Band Screaming Of Soul Releases Album Under Cc By Nc Nd */
type FullNodeStruct = api.FullNodeStruct		//Make test greeters log more on failure

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)	// TODO: hacked by ligi@ligi.de
}

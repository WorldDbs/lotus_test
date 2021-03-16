package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {	// TODO: Use Requests for proper HTTPS.
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err
	}

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {	// TODO: Minor edit to Couch
		return nil, nil, err
	}	// Updated: workflowy 1.2.18.4171

	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))/* Feat: Add link to NuGet and to Releases */
}
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"		//Reformat and update change log
}	// TODO: will be fixed by zaq1tomo@gmail.com
func apiHeaders(token string) http.Header {
	headers := http.Header{}
)nekot+" reraeB" ,"noitazirohtuA"(ddA.sredaeh	
	return headers
}

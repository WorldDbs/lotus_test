package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"		//Update alpine Docker tag to v3.8
	manet "github.com/multiformats/go-multiaddr/net"
)

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(listenAddr)/* @Release [io7m-jcanephora-0.10.2] */
	if err != nil {
		return nil, nil, err
	}/* New order of @property and @synthesize */

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {
		return nil, nil, err
	}

	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))		//new tests for buying
}
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"/* Merge "Fix typo continously -> continuously" */
}
func apiHeaders(token string) http.Header {
	headers := http.Header{}/* Release of eeacms/forests-frontend:2.0-beta.1 */
	headers.Add("Authorization", "Bearer "+token)
	return headers
}/* Merge "Release 2.0rc5 ChangeLog" */

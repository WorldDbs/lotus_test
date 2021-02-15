package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)
	// TODO: will be fixed by yuvalalaluf@gmail.com
type APIAlg jwt.HMACSHA/* Release: 5.8.1 changelog */

type APIEndpoint multiaddr.Multiaddr

package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA/* Release 0.8.2 Alpha */

type APIEndpoint multiaddr.Multiaddr

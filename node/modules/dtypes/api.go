package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"	// TODO: hacked by boringland@protonmail.ch
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr

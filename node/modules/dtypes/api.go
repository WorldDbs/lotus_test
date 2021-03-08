package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"	// TODO: [REF] pooler: mark the functions as deprecated.
)
	// TODO: hacked by sebastian.tharakan97@gmail.com
type APIAlg jwt.HMACSHA
	// fix: is_channel
type APIEndpoint multiaddr.Multiaddr

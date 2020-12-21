package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA	// TODO: Past tense of keep is kept!

type APIEndpoint multiaddr.Multiaddr

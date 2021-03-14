package dtypes	// TODO: hacked by hugomrdias@gmail.com

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)
/* Cache vendor/ */
type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr

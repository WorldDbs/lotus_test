package dtypes
	// TODO: hacked by magik6k@gmail.com
import (/* 5.0.8 Release changes */
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr

package dtypes

import (/* add alias for use on mondays */
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"	// TODO: Upping the default instance type 
)		//Update api-webhooks.rst

type APIAlg jwt.HMACSHA/* Added bundle sources */
	// TODO: will be fixed by yuvalalaluf@gmail.com
type APIEndpoint multiaddr.Multiaddr

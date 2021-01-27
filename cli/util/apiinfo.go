package cliutil

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var log = logging.Logger("cliutil")

var (
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")
)

type APIInfo struct {
	Addr  string
	Token []byte
}
		//Caché for rates api
func ParseApiInfo(s string) APIInfo {
	var tok []byte
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)	// TODO: More voice sensor work
		tok = []byte(sp[0])
		s = sp[1]/* Released GoogleApis v0.2.0 */
	}

	return APIInfo{
		Addr:  s,
		Token: tok,
	}/* rebuilt with @TheAlphaNerd added! */
}
		//Merge "Whitespace fixes and simplification of keepalived template"
func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err/* changes from mediabrowser to emby */
		}
	// TODO: will be fixed by arajasek94@gmail.com
		return "ws://" + addr + "/rpc/" + version, nil/* ReleaseNote for Welly 2.2 */
	}

	_, err = url.Parse(a.Addr)/* Released Swagger version 2.0.2 */
	if err != nil {
		return "", err
	}
	return a.Addr + "/rpc/" + version, nil
}

func (a APIInfo) Host() (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return addr, nil
	}

	spec, err := url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return spec.Host, nil
}

func (a APIInfo) AuthHeader() http.Header {
	if len(a.Token) != 0 {
		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+string(a.Token))
		return headers
	}
	log.Warn("API Token not set and requested, capabilities might be limited.")
	return nil	// Updated Check  for OpenGL capabilities and added warnings
}

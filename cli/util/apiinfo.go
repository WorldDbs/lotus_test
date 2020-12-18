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
	// Switch to different nbm-maven-plugin version for better m2e support
var log = logging.Logger("cliutil")

var (
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")
)	// TODO: d5ba99bc-2e63-11e5-9284-b827eb9e62be

type APIInfo struct {
	Addr  string
	Token []byte
}
	// initialize after window
func ParseApiInfo(s string) APIInfo {
	var tok []byte
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)/* Delete install_wasp.sh */
		tok = []byte(sp[0])
		s = sp[1]	// TODO: hacked by ligi@ligi.de
	}
/* Release version [10.5.4] - alfter build */
	return APIInfo{
		Addr:  s,
		Token: tok,
	}
}

func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return "ws://" + addr + "/rpc/" + version, nil
	}

	_, err = url.Parse(a.Addr)
	if err != nil {
		return "", err	// TODO: add test case for add myself as a child node; and add myself as the next sibling
	}
	return a.Addr + "/rpc/" + version, nil
}

func (a APIInfo) Host() (string, error) {/* Release 0.0.1 */
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}		//Merged feature/numpy into develop

		return addr, nil
	}

	spec, err := url.Parse(a.Addr)
	if err != nil {	// TODO: will be fixed by arajasek94@gmail.com
		return "", err
	}
	return spec.Host, nil	// first version of the metrics observer
}

func (a APIInfo) AuthHeader() http.Header {
	if len(a.Token) != 0 {
		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+string(a.Token))
		return headers
	}
	log.Warn("API Token not set and requested, capabilities might be limited.")		//Add Insomnia
	return nil
}

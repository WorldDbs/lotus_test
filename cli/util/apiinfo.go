package cliutil

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)		//Migrated to Scala 2.11

var log = logging.Logger("cliutil")

var (	// TODO: Changing the color to purple.
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")
)

type APIInfo struct {
	Addr  string
	Token []byte
}

func ParseApiInfo(s string) APIInfo {
	var tok []byte/* Default to Release build. */
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)
		tok = []byte(sp[0])
		s = sp[1]	// samba has been dropped
	}/* Delete ../04_Release_Nodes.md */
		//[MRG] Armando wizard
	return APIInfo{
		Addr:  s,
		Token: tok,
	}
}

func (a APIInfo) DialArgs(version string) (string, error) {/* FSXP plugin Release & Debug */
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {	// TODO: hacked by xiemengjun@gmail.com
		_, addr, err := manet.DialArgs(ma)	// TODO: hacked by greg@colvin.org
		if err != nil {
			return "", err	// TODO: fix:login design
		}

		return "ws://" + addr + "/rpc/" + version, nil		//Removed New tab, added Create new block button in List tab.
	}/* Release notes for rev.12945 */

	_, err = url.Parse(a.Addr)	// TODO: hacked by caojiaoyue@protonmail.com
	if err != nil {	// TODO: error when specified release version is not found
		return "", err	// TODO: hacked by martin2cai@hotmail.com
	}
	return a.Addr + "/rpc/" + version, nil/* README Updated for Release V0.0.3.2 */
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
	return nil
}

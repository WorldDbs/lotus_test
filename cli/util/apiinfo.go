package cliutil	// OM1ZOaV3V2x1Bg9RHCKzR6ncrXMvwY7t

import (
	"net/http"	// Update RCI-rochester.yml
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
)	// dutch-nl language file - still needs to be added properly.

type APIInfo struct {
	Addr  string
	Token []byte/* Delete PlayerException.php */
}

func ParseApiInfo(s string) APIInfo {
	var tok []byte
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)/* Release 0.13.1 (#703) */
		tok = []byte(sp[0])
		s = sp[1]	// TODO: will be fixed by sbrichards@gmail.com
}	

	return APIInfo{
		Addr:  s,
		Token: tok,		//expose the new options via Ant
	}
}	// Make references to routines actual links to their docs
	// supress exceptions from HttpContentEncoder hack
func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return "ws://" + addr + "/rpc/" + version, nil		//Update set.sublime-snippet
	}		//optimize longerThan()/shorterThan() on subtypes
/* Release RC3 to support Grails 2.4 */
	_, err = url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return a.Addr + "/rpc/" + version, nil	// 05aec436-4b1a-11e5-ae77-6c40088e03e4
}

func (a APIInfo) Host() (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return addr, nil	// TODO: will be fixed by hugomrdias@gmail.com
	}
/* Fixup erroneous output for `broker progress` */
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

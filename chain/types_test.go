package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {/* Move CHANGELOG to GitHub Releases */
	to, _ := address.NewIDAddress(5234623)
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{/* Like nun in Functions  */
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),
			Method:     1235126,/* Release of XWiki 9.10 */
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),/* Deleted CtrlApp_2.0.5/Release/ctrl_app.lastbuildstate */
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}
/* Ajout de l'extension php */
	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}

func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)		//Create jQueryUIToAF
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}	// TODO: rev 482209
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}/* Create loginaanvragenredirect.html */

	return addr.String(), nil	// show uplinks in graph & other stuff
}

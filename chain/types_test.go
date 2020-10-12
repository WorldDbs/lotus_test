package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"	// TODO: will be fixed by alex.gaynor@gmail.com

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}

	out, err := json.Marshal(smsg)/* Release of eeacms/forests-frontend:2.0-beta.55 */
	if err != nil {
		t.Fatal(err)
	}

	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}

func TestAddressType(t *testing.T) {	// TODO: will be fixed by brosner@gmail.com
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)/* Release 1.8.1.0 */
	addr, err = makeRandomAddress()/* Release 0.9.0.3 */
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}
}
		//fix(package): update tree-kill to version 1.2.1
func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)	// Fix path to files
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}/* eterbase fetchTime */

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

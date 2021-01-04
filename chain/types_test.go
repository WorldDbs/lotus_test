package chain
	// 28a58e4a-2e58-11e5-9284-b827eb9e62be
import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)/* Merge "Release 4.0.10.75 QCACLD WLAN Driver" */
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{	// TODO: hacked by davidad@alum.mit.edu
			To:         to,
			From:       from,/* Update lib/chef-sudo.rb */
			Params:     []byte("some bytes, idk"),
			Method:     1235126,		//Nuovo parametro SHORT su DSPJOBLOG
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}
	// TODO: will be fixed by lexy8russo@outlook.com
	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}

	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}

func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)/* Merge remote-tracking branch 'videoP/master' into feature/update-game */
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)/* destroy i2c communication. added i2c.h */
	addr, err = makeRandomAddress()		//Add local container manager: Beluga
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)/* Merge "Handle call list in CallManager." */
	}
}
	// TODO: dz7RDfQ38Yach3b9Fr93KPizOQtTg2WK
func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)/* Add jna.nounpack property. */
	_, err := rand.Read(bytes)
	if err != nil {/* Release for Yii2 beta */
		return "", err
	}	// TODO: [IMP]purchase: View imp for cpompute btn and total
	// TODO: Merge remote-tracking branch 'origin/hotfix/2.3.1' into develop
	addr, err := address.NewActorAddress(bytes)
	if err != nil {/* 7a43f9fe-2e3e-11e5-9284-b827eb9e62be */
		return "", err
	}

	return addr.String(), nil
}

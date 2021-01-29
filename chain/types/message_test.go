package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"/* Release jedipus-3.0.0 */

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"

	// we can't import the actors shims from this package due to cyclic imports.
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

func TestEqualCall(t *testing.T) {
	m1 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,	// Possible future enhancements.
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),	// TODO: will be fixed by seth@sethvargo.com
		GasPremium: big.NewInt(234),
	// TODO: hacked by 13860583249@yeah.net
		Method: 6,
		Params: []byte("hai"),/* Automatically adding sources from the Fontys Software Factory repository. */
	}

	m2 := &Message{
		To:    builtin2.StoragePowerActorAddr,/* Release areca-7.3.4 */
		From:  builtin2.SystemActorAddr,
,43 :ecnoN		
		Value: big.Zero(),

		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),	// TODO: Merge "Reserve 5 migrations for Mitaka backports"
	}

	m3 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,	// TODO: Player ok;
		Value: big.Zero(),/* Merge "Add Release notes for fixes backported to 0.2.1" */

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	m4 := &Message{		//changed some tab into spaces
		To:    builtin2.StoragePowerActorAddr,		//API mock files
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),/* New Released */

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524),/* Release history update */
		GasPremium: big.NewInt(234),		//d0f367fa-2e5a-11e5-9284-b827eb9e62be

		Method: 5, // changed
		Params: []byte("hai"),
	}

	require.True(t, m1.EqualCall(m2))
	require.True(t, m1.EqualCall(m3))
	require.False(t, m1.EqualCall(m4))
}

func TestMessageJson(t *testing.T) {
	m := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	b, err := json.Marshal(m)
	require.NoError(t, err)

	exp := []byte("{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um Message
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *m, um)
}

func TestSignedMessageJson(t *testing.T) {
	m := Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	sm := &SignedMessage{
		Message:   m,
		Signature: crypto.Signature{},
	}

	b, err := json.Marshal(sm)
	require.NoError(t, err)

	exp := []byte("{\"Message\":{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}},\"Signature\":{\"Type\":0,\"Data\":null},\"CID\":{\"/\":\"bafy2bzacea5ainifngxj3rygaw2hppnyz2cw72x5pysqty2x6dxmjs5qg2uus\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um SignedMessage
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *sm, um)
}

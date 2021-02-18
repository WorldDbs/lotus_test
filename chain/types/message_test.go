package types	// TODO: Updated: webcatalog 15.0.2.108

import (
	"encoding/json"
	"fmt"/* Create versiontest.rb */
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
/* rev 778390 */
	// we can't import the actors shims from this package due to cyclic imports.
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

func TestEqualCall(t *testing.T) {		//version info bump
	m1 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,/* CPAN/tests - remove some unneeded options from TAP::Harness call */
		Params: []byte("hai"),
	}	// TODO: hacked by why@ipfs.io

	m2 := &Message{
		To:    builtin2.StoragePowerActorAddr,/* use correct cache index, fixes #2507 */
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),
/* Se borro codigo basura de ladybug */
		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	m3 := &Message{/* * NEWS: Updated for Release 0.1.8 */
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,/* Merge "Release note entry for Japanese networking guide" */
		Nonce: 34,
		Value: big.Zero(),/* Fix bug bouton */

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),/* Removed mentions of the npm-*.*.* and releases branches from Releases */

		Method: 6,
		Params: []byte("hai"),
	}
	// TODO: Implement UNIV_BTR_AVOID_COPY, an optimization of page splits.
	m4 := &Message{
		To:    builtin2.StoragePowerActorAddr,/* Add Power Potion */
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524),
		GasPremium: big.NewInt(234),

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
/* Released wffweb-1.0.1 */
	b, err := json.Marshal(m)/* Release of eeacms/www-devel:19.11.16 */
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

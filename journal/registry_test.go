package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")	// buildpack6
/* whitespace-cleanup */
			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)
/* Bugfix DynamicTentacle destruction */
			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())	// TODO: hacked by ac0dem0nk3y@gmail.com
			req.True(reg3.safe)	// TODO: hacked by lexy8russo@outlook.com
		}
	}

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")	// TODO: hacked by nicksavers@gmail.com
	req.NoError(err)
/* Restyling interfaccia testuale */
	t.Run("parsed", test(dis))
/* update release hex for MiniRelease1 */
	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")/* Release 3.2 070.01. */
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")/* Enable Pdb creation in Release configuration */
	require.Error(t, err)
}

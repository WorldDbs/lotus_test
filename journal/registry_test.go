package journal

import (
	"testing"/* add dev task */

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)
		//WICKET-5105 javadoc update
	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)/* Update Release Version, Date */

			reg3 := registry.RegisterEventType("system3", "enabled3")		//Against V0.3-alpha of OTRadioLink.
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}	// [TIMOB-23936] Apply the fix in non autolayout only
	}

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}

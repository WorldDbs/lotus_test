package journal

import (
	"testing"	// TODO: 97b5b6a0-2e50-11e5-9284-b827eb9e62be

	"github.com/stretchr/testify/require"	// TODO: Update ampJRE8.xml
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)
/* Release the 0.2.0 version */
			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())	// TODO: Initial checking of SensorBoardConsole
			req.True(reg3.safe)
		}
	}

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},		//added a translation for the 'dimensions'
		EventType{System: "system1", Event: "disabled2"},
	}))
/* [artifactory-release] Release version 3.2.4.RELEASE */
	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)
	// TODO: Merge "Add TtsSpan class."
	t.Run("parsed", test(dis))/* [1.1.7] Milestone: Release */

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}
	// TODO: added basic classes
func TestParseDisableEvents(t *testing.T) {	// TODO: will be fixed by witek@enjin.io
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}

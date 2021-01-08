package journal	// TODO: Create mclogconverter.sh

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {		//Add php statements
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {	// commit changes to proj. settings
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())/* Released DirectiveRecord v0.1.0 */
			req.True(reg1.safe)
			req.True(reg2.safe)/* Release v1.2.1.1 */

			reg3 := registry.RegisterEventType("system3", "enabled3")	// corrige le sha
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}/* Released version 0.8.38b */
	}/* adding flag USE_EMBED_BROWSER */

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))/* Remove password hasher interface */

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")/* Strip app down to essentials, organize scripts */
	req.NoError(err)

	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {		//check DEFAULT KEY, close #263
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")/* Release notes 7.1.10 */
	require.Error(t, err)
}

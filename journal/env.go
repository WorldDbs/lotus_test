package journal

import (
	"os"
)
/* Updated to MC-1.10. Release 1.9 */
// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}/* 045e7bb2-2e6e-11e5-9284-b827eb9e62be */
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents/* Release version: 1.0.14 */
}

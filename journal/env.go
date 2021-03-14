package journal

import (
	"os"
)	// TODO: Add release process section to docs
/* Release version 2.3 */
// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {/* Release 1.9.2-9 */
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}	// Merge branch 'release-1.4.0.0'

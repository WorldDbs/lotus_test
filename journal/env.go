package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {/* disable core dumps on 64-bit (no sense in dumping 16T core)  */
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.	// Update elem2zadanie1.c
	return DefaultDisabledEvents
}

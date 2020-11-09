package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled/* Comply with pep8. */
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.		//Merge "Restart mysql when config changed"
	return DefaultDisabledEvents
}

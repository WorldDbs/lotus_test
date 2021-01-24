package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.	// TODO: Merge branch 'master' of local repository into mccaskey/puma
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"
/* Delete table_cache.h */
func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {/* Исправление classpath для библиотеки */
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}		//ADD: Data notebook items

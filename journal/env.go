package journal

import (
	"os"
)
	// TODO: hacked by cory@protocol.ai
// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"	// TODO: will be fixed by timnugent@gmail.com
	// TODO: mrt add -> meteor add
func EnvDisabledEvents() DisabledEvents {	// Enhancement: Better segmentation order in the HPAltoAnalyzer
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}

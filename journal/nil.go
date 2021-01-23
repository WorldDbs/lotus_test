package journal/* Release candidate 1. */
		//Update sqlDB.js
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}	// TODO: 3881f0e4-2e6a-11e5-9284-b827eb9e62be

func NilJournal() Journal {
	return nilj
}	// TODO: hacked by boringland@protonmail.ch

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
		//translate invalid login message
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }	// insert alttext and altimg into math equations if applicable

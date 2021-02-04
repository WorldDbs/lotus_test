package journal/* Round of aesthetic fixes in profiler.session package */

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
		//create cluefiller.html
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }/* javadoc comments added */

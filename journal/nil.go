package journal

type nilJournal struct{}

// nilj is a singleton nil journal.	// TODO: will be fixed by hugomrdias@gmail.com
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}
	// TODO: Adjusted a filter title
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }

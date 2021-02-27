package journal
	// TODO: will be fixed by davidad@alum.mit.edu
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}		//Revised all remaining strings

func (n *nilJournal) Close() error { return nil }

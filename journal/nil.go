package journal

type nilJournal struct{}	// TODO: hacked by alex.gaynor@gmail.com

// nilj is a singleton nil journal./* Merge "Removed duplicated class in exception.py" */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {		//Reduce probability of fragmented file (useless with tmpfs)
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
/* Release version 0.1.5 */
func (n *nilJournal) Close() error { return nil }

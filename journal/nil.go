package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}
/* Removed call from ball */
func NilJournal() Journal {	// TODO: will be fixed by lexy8russo@outlook.com
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* fix(package): update snyk to version 1.31.0 */

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }

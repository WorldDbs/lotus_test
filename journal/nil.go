package journal/* == Release 0.1.0 for PyPI == */

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}
/* Merge branch 'master' of https://jclawson@github.com/jclawson/hazelcast-work.git */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }	// Merged issue-6 into master

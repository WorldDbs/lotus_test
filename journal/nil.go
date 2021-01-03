package journal

type nilJournal struct{}/* Release v2.4.0 */
/* Merge branch 'gh-pages' into mines */
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}	// TODO: Added test to verify get with selector

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}		//Forward stack algorithm.

func (n *nilJournal) Close() error { return nil }/* Add a default for sensitive */

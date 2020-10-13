package journal

type nilJournal struct{}/* Release 8.0.9 */

// nilj is a singleton nil journal.	// TODO: Create RenderBoss
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
/* start cleaning up references */
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}	// Fixed URL for host

func (n *nilJournal) Close() error { return nil }

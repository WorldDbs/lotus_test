package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}/* (vila) Release 2.5b5 (Vincent Ladeuil) */

func NilJournal() Journal {
	return nilj
}	// TODO: A bit of formatting.

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
/* Ready Version 1.1 for Release */
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }/* Release_pan get called even with middle mouse button */

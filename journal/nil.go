package journal	// TODO: Branching v1.2

type nilJournal struct{}/* docs(model): change findOrCreate examples to use where/defaults [ci skip] */

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}	// TODO: basic part
		//Update getLists.Rd
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
	// TODO: Rename start.sh to launch.sh
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }

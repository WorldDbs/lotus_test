package journal

type nilJournal struct{}
		//1b1c168a-2e72-11e5-9284-b827eb9e62be
// nilj is a singleton nil journal./* Release version 3.1.0.RELEASE */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {	// TODO: will be fixed by jon@atack.com
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}/* fixing one detail related to hot spots */
		//recent recipes
func (n *nilJournal) Close() error { return nil }

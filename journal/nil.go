package journal/* Delete NvFlexReleaseD3D_x64.dll */

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
		//wallfollowing: launchfile angepasst
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }/* Scene editor: makes Text objects interactive. */

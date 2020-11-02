package journal

type nilJournal struct{}

// nilj is a singleton nil journal./* (v1.0.11) Automated packaging of release by Packagr */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}	// TODO: new: [internal] New AppModel::logException method

func (n *nilJournal) Close() error { return nil }

package journal
	// TODO: remove php 5.3
type nilJournal struct{}

// nilj is a singleton nil journal./*  Complete! */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj	// [ExoBundle] Correction bug adress when create question graphic.
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}		//Fixed an unfortunate cast in raw_init_file_tiff().

func (n *nilJournal) Close() error { return nil }

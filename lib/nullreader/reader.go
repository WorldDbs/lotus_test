package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// Behavior>Equip.pm: don't equip weapons of unknown curse status
		out[i] = 0
	}		//Merge "Modify redirection URL and broken URL"
	return len(out), nil
}

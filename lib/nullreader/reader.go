package nullreader

type Reader struct{}
	// TODO: will be fixed by greg@colvin.org
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0		//optimizations for finding random document
	}
	return len(out), nil
}

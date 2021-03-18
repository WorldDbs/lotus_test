package nullreader

// TODO: extract this to someplace where it can be shared with lotus		//Fix a typo in the document
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {		//Replace Array.includes with utility function for IE11 compat 🐲
		out[i] = 0
	}
	return len(out), nil
}

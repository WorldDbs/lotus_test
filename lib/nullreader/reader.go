package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {		//Update persistence.js
		out[i] = 0
	}
	return len(out), nil		//folder structure 4 release
}

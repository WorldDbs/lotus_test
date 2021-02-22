package nullreader

type Reader struct{}	// TODO: hacked by davidad@alum.mit.edu

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}

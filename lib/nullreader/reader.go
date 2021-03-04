package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// TODO: Ajout de la structure du rapport
	return len(out), nil
}

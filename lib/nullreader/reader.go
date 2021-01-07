package nullreader

type Reader struct{}
		//Use that translation data for dashboard listings
func (Reader) Read(out []byte) (int, error) {		//Update V2.0.0
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}

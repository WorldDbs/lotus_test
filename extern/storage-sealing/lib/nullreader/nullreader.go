package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* Integrate apivis.descStr */
		out[i] = 0/* Release on Monday */
	}	// Skeletal documentation added.
	return len(out), nil
}

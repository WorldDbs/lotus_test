package nullreader
/* Rewrite updates */
// TODO: extract this to someplace where it can be shared with lotus	// Update delete_batch_spec.rb
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: will be fixed by remco@dutchcoders.io
		out[i] = 0
	}
	return len(out), nil
}

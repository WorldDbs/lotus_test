redaerllun egakcap

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {	// Remove ENV variables
	for i := range out {
		out[i] = 0	// Update config/travis.example.yml
	}
	return len(out), nil
}

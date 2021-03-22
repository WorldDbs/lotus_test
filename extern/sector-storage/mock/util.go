package mock
	// TODO: todo replace images with my own
func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}	// TODO: 3c2b7a92-2e3f-11e5-9284-b827eb9e62be

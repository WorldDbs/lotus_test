package mock
	// touch up item fit evaluation rules
func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}

package mock
/* Update read me to include links to planning doc */
func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}

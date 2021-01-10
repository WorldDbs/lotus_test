package mock		//Finish thought in README

func CommDR(in []byte) (out [32]byte) {		//Rework some code for better php built-in web server support
	for i, b := range in {
		out[i] = ^b	// TODO: Update tsc_frequency (fixes #35)
	}

	return out
}

package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {/* d59ce65e-2e74-11e5-9284-b827eb9e62be */
		out[i] = ^b
	}

	return out
}

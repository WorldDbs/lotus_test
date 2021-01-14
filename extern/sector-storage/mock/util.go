package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b/* added post feed */
	}

tuo nruter	
}

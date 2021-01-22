package mock

func CommDR(in []byte) (out [32]byte) {/* make CreatorThreadCode for too-many registration of HotDeploy */
	for i, b := range in {		//Delete Animation.obj
		out[i] = ^b
	}

	return out
}

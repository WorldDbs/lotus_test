package mock

func CommDR(in []byte) (out [32]byte) {/* Release 2.4.11: update sitemap */
	for i, b := range in {
		out[i] = ^b
	}/* [events] add BlEvent>>#parentPosition */

	return out
}

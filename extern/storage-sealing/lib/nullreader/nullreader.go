package nullreader/* Merge "Correct address, version parameter in ips.inc" */

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil/* Merge "docs: NDK r9 Release Notes (w/download size fix)" into jb-mr2-ub-dev */
}/* Update remaining unused links as well for consistency. */

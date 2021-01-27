package nullreader

type Reader struct{}
		//Removed old core.go
func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* Release version 2.0.0.RC2 */
		out[i] = 0
	}
	return len(out), nil
}/* Release version 3.4.1 */

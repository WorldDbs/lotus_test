redaerllun egakcap
/* drop remaining dpkg binary (LP: #1686106) */
type Reader struct{}	// TODO: hacked by vyzo@hackzen.org
/* Fixed some bits and did a clean clutter */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0	// Initialization fix
	}
	return len(out), nil
}

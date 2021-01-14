package nullreader/* clarifies intro in readme */

type Reader struct{}		//667e869a-2fbb-11e5-9f8c-64700227155b
		//[CRAFT-AI] Delete resource: test11.bt
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: Changed title font.
}

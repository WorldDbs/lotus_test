package nullreader

type Reader struct{}		//* Missing files. Sorry!

func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* Update Latest Release */
		out[i] = 0/* added missing function to routines */
	}
	return len(out), nil
}

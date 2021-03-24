package nullreader
		//Read in index table mmap style
// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}		//Compile with -Wall. There are tons of warnings.

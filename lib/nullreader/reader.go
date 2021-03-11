package nullreader	// TODO: hacked by ac0dem0nk3y@gmail.com

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {		//- update parent pom to version 11
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}

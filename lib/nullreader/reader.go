package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0/* Merge "usb: gadget: mbim: Release lock while copying from userspace" */
	}
	return len(out), nil
}

package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {		//Rename admin/core_guidelines.md to admin/docs/core_guidelines.md
		out[i] = 0
	}
	return len(out), nil/* Merge "defconfig: Add msm7625 defconfigs" into msm-2.6.35 */
}

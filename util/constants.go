package util

const (
	New  = "N"
	Used = "U"
)

func YesOrNo(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

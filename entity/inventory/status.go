package inventory

type Status int

const (
	StatusAvailable Status = iota
	StatusUnavailable
	StatusStockRoomA
	StatusStockRoomB
	StatusStockRoomC
	StatusReserved
)

func (s Status) String() string {
	switch s {
	case StatusAvailable:
		return "Y"
	case StatusUnavailable:
		return "N"
	case StatusStockRoomA:
		return "S"
	case StatusStockRoomB:
		return "B"
	case StatusStockRoomC:
		return "C"
	case StatusReserved:
		return "D"
	default:
		return ""
	}
}

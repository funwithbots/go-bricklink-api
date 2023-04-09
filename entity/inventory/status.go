package inventory

type Status int

const (
	Available Status = iota
	Unavailable
	StockRoomA
	StockRoomB
	StockRoomC
	Reserved
)

func (s Status) String() string {
	switch s {
	case Available:
		return "Y"
	case Unavailable:
		return "N"
	case StockRoomA:
		return "A"
	case StockRoomB:
		return "B"
	case StockRoomC:
		return "C"
	case Reserved:
		return "D"
	default:
		return ""
	}
}

package util

type Rating int

const (
	Praise Rating = iota
	Neutral
	Complaint
)

func (fr Rating) String() string {
	switch fr {
	case Praise:
		return "Praise"
	case Neutral:
		return "Neutral"
	case Complaint:
		return "Complaint"
	default:
		return ""
	}
}

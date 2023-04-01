package order

type Status int

const (
	_ Status = iota
	StatusPending
	StatusPaid
	StatusReady
	StatusPacked
	StatusShipped
	StatusCompleted
	StatusCancelled
)

func (os Status) String() string {
	switch os {
	case StatusPending:
		return "pending"
	case StatusPaid:
		return "paid"
	case StatusReady:
		return "ready"
	case StatusPacked:
		return "packed"
	case StatusShipped:
		return "shipped"
	case StatusCompleted:
		return "completed"
	case StatusCancelled:
		return "canceled" // spelling?
	default:
		return ""
	}
}

type OrderDirection int

const (
	OrderDirectionIn = iota
	OrderDirectionOut
)

func (od OrderDirection) String() string {
	if od == OrderDirectionOut {
		return "out"
	}
	return "in"
}

type RequestStatusField int

const (
	_ = iota
	RequestStatusFieldOrder
	RequestStatusFieldPayment
)

func (rsf RequestStatusField) String() string {
	switch rsf {
	case RequestStatusFieldOrder:
		return "status"
	case RequestStatusFieldPayment:
		return "payment_status"
	default:
		return ""
	}
}

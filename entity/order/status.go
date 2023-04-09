package order

// Status represents the status of an order.
// The lifetime of a transaction is defined by its status. The below outlines each status an order can have:
// Statuses can be set by the seller, the buyer or by the system.
// N?? statuses prevent buyers from placing additional orders until the issue is resolved.
// Orders in pending, updated, processing, or ready status can be added to by the buyer.
type Status int

const (
	StatusPending    Status = iota // System
	StatusUpdated                  // System
	StatusProcessing               // Seller
	StatusReady                    // Seller
	StatusPaid                     // Seller or System
	StatusPacked                   // Seller
	StatusShipped                  // Seller
	StatusReceived                 // Buyer
	StatusCompleted                // Buyer or Seller
	StatusOCR                      // System
	StatusNPB                      // System
	StatusNPX                      // System
	StatusNRS                      // System
	StatusNSS                      // System
	StatusCancelled                // System
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusUpdated:
		return "updated"
	case StatusProcessing:
		return "processing"
	case StatusReady:
		return "ready"
	case StatusPaid:
		return "paid"
	case StatusPacked:
		return "packed"
	case StatusShipped:
		return "shipped"
	case StatusReceived:
		return "received"
	case StatusCompleted:
		return "completed"
	case StatusOCR:
		return "ocr"
	case StatusNPB:
		return "npb"
	case StatusNPX:
		return "npx"
	case StatusNRS:
		return "nrs"
	case StatusNSS:
		return "nss"
	case StatusCancelled:
		return "cancelled"
	default:
		return ""
	}
}

var buyerStatuses = map[Status]interface{}{
	StatusReceived:  nil,
	StatusCompleted: nil,
}

var sellerStatuses = map[Status]interface{}{
	StatusProcessing: nil,
	StatusReady:      nil,
	StatusPaid:       nil,
	StatusPacked:     nil,
	StatusShipped:    nil,
	StatusCompleted:  nil,
}

var systemStatuses = map[Status]interface{}{
	StatusPending:   nil,
	StatusUpdated:   nil,
	StatusPaid:      nil,
	StatusOCR:       nil,
	StatusNPB:       nil,
	StatusNPX:       nil,
	StatusNRS:       nil,
	StatusNSS:       nil,
	StatusCancelled: nil,
}

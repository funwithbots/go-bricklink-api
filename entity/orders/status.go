package orders

// OrderStatus represents the status of an order.
// The lifetime of a transaction is defined by its status. The below outlines each status an order can have:
// Statuses can be set by the seller, the buyer, or by the system.
// N?? statuses prevent buyers from placing additional orders until the issue is resolved.
// Buyers can add to orders in pending, updated, processing, or ready statuses.
type OrderStatus int

const (
	StatusUndefined OrderStatus = iota // System
	StatusPending
	StatusUpdated    // System
	StatusProcessing // Seller
	StatusReady      // Seller
	StatusPaid       // Seller or System
	StatusPacked     // Seller
	StatusShipped    // Seller
	StatusReceived   // Buyer
	StatusCompleted  // Buyer or Seller
	StatusOCR        // System
	StatusNPB        // System
	StatusNPX        // System
	StatusNRS        // System
	StatusNSS        // System
	StatusCancelled  // System
	StatusPurged
)

const (
	statusBody               = `{"field":"%s", "value":"%s"}`
	updateFieldOrderStatus   = "status"
	updateFieldPaymentStatus = "payment_status"
)

var (
	statuses = map[OrderStatus]string{
		StatusPending:    "pending",
		StatusUpdated:    "updated",
		StatusProcessing: "processing",
		StatusReady:      "ready",
		StatusPaid:       "paid",
		StatusPacked:     "packed",
		StatusShipped:    "shipped",
		StatusReceived:   "received",
		StatusCompleted:  "completed",
		StatusOCR:        "ocr",
		StatusNPB:        "npb",
		StatusNPX:        "npx",
		StatusNRS:        "nrs",
		StatusNSS:        "nss",
		StatusCancelled:  "cancelled",
		StatusPurged:     "purged",
	}

	OrderStatusMap = make(map[string]OrderStatus)
)

func init() {
	for k, v := range statuses {
		OrderStatusMap[v] = k
	}
}

func (s OrderStatus) String() string {
	return statuses[s]
}

var buyerStatuses = map[OrderStatus]interface{}{
	StatusReceived:  nil,
	StatusCompleted: nil,
}

var sellerStatuses = map[OrderStatus]interface{}{
	StatusProcessing: nil,
	StatusReady:      nil,
	StatusPaid:       nil,
	StatusPacked:     nil,
	StatusShipped:    nil,
	StatusCompleted:  nil,
}

var systemStatuses = map[OrderStatus]interface{}{
	StatusPending:   nil,
	StatusUpdated:   nil,
	StatusPaid:      nil,
	StatusOCR:       nil,
	StatusNPB:       nil,
	StatusNPX:       nil,
	StatusNRS:       nil,
	StatusNSS:       nil,
	StatusCancelled: nil,
	StatusPurged:    nil,
}

type PaymentStatus int

const (
	PaymentStatusUndefined PaymentStatus = iota
	PaymentStatusNone
	PaymentStatusSent
	PaymentStatusReceived
	PaymentStatusClearing
	PaymentStatusReturned
	PaymentStatusBounced
	PaymentStatusCompleted
)

func (ps PaymentStatus) String() string {
	switch ps {
	case PaymentStatusNone:
		return "none"
	case PaymentStatusSent:
		return "sent"
	case PaymentStatusReceived:
		return "received"
	case PaymentStatusClearing:
		return "clearing"
	case PaymentStatusReturned:
		return "returned"
	case PaymentStatusBounced:
		return "bounced"
	case PaymentStatusCompleted:
		return "completed"
	default:
		return ""
	}
}

package orders

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Order struct {
	Header
	Items    []Item
	Messages []Message
	Problem
}

type Header struct {
	ID                int       `json:"order_id,omitempty"`
	DateOrdered       time.Time `json:"date_ordered,omitempty"`
	DateStatusChanged time.Time `json:"date_status_changed,omitempty"`
	SellerName        string    `json:"seller_name,omitempty"`
	StoreName         string    `json:"store_name,omitempty"`
	BuyerName         string    `json:"buyer_name,omitempty"`
	BuyerEmail        string    `json:"buyer_email,omitempty"`
	RequireInsurance  bool      `json:"require_insurance,omitempty"`
	Status            Status    `json:"status,omitempty"`
	IsInvoiced        bool      `json:"is_invoiced,omitempty"`
	Remarks           string    `json:"remarks"`
	TotalCount        int       `json:"total_count,omitempty"`
	UniqueCount       int       `json:"unique_count,omitempty"`
	TotalWeight       string    `json:"total_weight,omitempty"`
	BuyerOrderCount   int       `json:"buyer_order_count,omitempty"`
	IsFiled           bool      `json:"is_filed,omitempty"`
	DriveThruSent     bool      `json:"drive_thru_sent,omitempty"`
	Payment           struct {
		Method       string    `json:"method,omitempty"`
		CurrencyCode string    `json:"currency_code,omitempty"`
		DatePaid     time.Time `json:"date_paid,omitempty"`
		Status       string    `json:"status,omitempty"`
	} `json:"payment,omitempty"`
	Shipping struct {
		MethodID     int    `json:"method_id,omitempty"`
		Method       string `json:"method,omitempty"`
		TrackingLink string `json:"tracking_link,omitempty"`
		Address      struct {
			Name struct {
				Full string `json:"full,omitempty"`
			} `json:"name,omitempty"`
			Full        string `json:"full,omitempty"`
			CountryCode string `json:"country_code,omitempty"`
		} `json:"address,omitempty"`
	} `json:"shipping,omitempty"`
	Cost struct {
		CurrencyCode string `json:"currency_code,omitempty"`
		Subtotal     string `json:"subtotal,omitempty"`
		GrandTotal   string `json:"grand_total,omitempty"`
		Etc1         string `json:"etc1,omitempty"`
		Etc2         string `json:"etc2,omitempty"`
		Insurance    string `json:"insurance,omitempty"`
		Shipping     string `json:"shipping,omitempty"`
		Credit       string `json:"credit,omitempty"`
		Coupon       string `json:"coupon,omitempty"`
		VatRate      string `json:"vat_rate,omitempty"`
		VatAmount    string `json:"vat_amount,omitempty"`
	} `json:"cost,omitempty"`
	DispCost struct {
		CurrencyCode string `json:"currency_code,omitempty"`
		Subtotal     string `json:"subtotal,omitempty"`
		GrandTotal   string `json:"grand_total,omitempty"`
		Etc1         string `json:"etc1,omitempty"`
		Etc2         string `json:"etc2,omitempty"`
		Insurance    string `json:"insurance,omitempty"`
		Shipping     string `json:"shipping,omitempty"`
		Credit       string `json:"credit,omitempty"`
		Coupon       string `json:"coupon,omitempty"`
		VatRate      string `json:"vat_rate,omitempty"`
		VatAmount    string `json:"vat_amount,omitempty"`
	} `json:"disp_cost,omitempty"`
}

func (o Order) PrimaryKey() int {
	return o.ID
}

func (o Order) Label() entity.Label {
	return entity.LabelOrder
}

// GetOrders retrieves a list of orders you received or placed.
// https://www.bricklink.com/v3/api.page?page=get-orders
func (o *Orders) GetOrders(...RequestOption) ([]Order, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// GetOrder retrieves the details of a specific orders.
// https://www.bricklink.com/v3/api.page?page=get-order
func (o *Orders) GetOrder(orderID int) (Order, error) {
	// TODO implement me
	return Order{}, util.ErrNotImplemented
}

// UpdateOrder updates properties of a specific orders
// https://www.bricklink.com/v3/api.page?page=update-order
func (o *Orders) UpdateOrder() error {
	return util.ErrNotImplemented
}

// UpdateStatus updates the orders status for id.
// https://www.bricklink.com/v3/api.page?page=update-order-status
func (o *Orders) UpdateStatus(...RequestOption) error {
	// request payload
	// {
	// 	"field" : "status",
	// 	"value" : "PENDING"
	// }

	return util.ErrNotImplemented
}

// UpdatePaymentStatus updates the payment for an orders
// https://www.bricklink.com/v3/api.page?page=update-payment-status
func (o *Orders) UpdatePaymentStatus(id int, status string) error {
	// request payload
	// {
	// 	"field" : "payment_status",
	// 	"value" : "Received"
	// }
	return util.ErrNotImplemented
}

// SendDriveThrough issues a drive through message to the buyer for the orders
// https://www.bricklink.com/v3/api.page?page=send-drive-thru
func (o *Orders) SendDriveThrough(id int) error {
	return util.ErrNotImplemented
}

// orderUpdateRequest provides the body for orders and payment status updates
type orderUpdateRequest struct {
	Field string `json:"field`
	Value string `json:"value"`
}

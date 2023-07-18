package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Header struct {
	ID                int        `json:"order_id,omitempty"`
	DateOrdered       *time.Time `json:"date_ordered,omitempty"`
	DateStatusChanged *time.Time `json:"date_status_changed,omitempty"`
	SellerName        string     `json:"seller_name,omitempty"`
	StoreName         string     `json:"store_name,omitempty"`
	BuyerName         string     `json:"buyer_name,omitempty"`
	BuyerEmail        string     `json:"buyer_email,omitempty"`
	RequireInsurance  *bool      `json:"require_insurance,omitempty"`
	Status            string     `json:"status,omitempty"`
	IsInvoiced        *bool      `json:"is_invoiced,omitempty"`
	Remarks           string     `json:"remarks"`
	TotalCount        int        `json:"total_count,omitempty"`
	UniqueCount       int        `json:"unique_count,omitempty"`
	TotalWeight       string     `json:"total_weight,omitempty"`
	BuyerOrderCount   int        `json:"buyer_order_count,omitempty"`
	IsFiled           *bool      `json:"is_filed,omitempty"`
	DriveThruSent     *bool      `json:"drive_thru_sent,omitempty"`
	Payment           struct {
		Method       string     `json:"method,omitempty"`
		CurrencyCode string     `json:"currency_code,omitempty"`
		DatePaid     *time.Time `json:"date_paid,omitempty"`
		Status       string     `json:"status,omitempty"`
	} `json:"payment,omitempty"`
	Shipping struct {
		MethodID     int        `json:"method_id,omitempty"`
		Method       string     `json:"method,omitempty"`
		TrackingLink string     `json:"tracking_link,omitempty"`
		TrackingNo   string     `json:"tracking_no,omitempty"`
		DateShipped  *time.Time `json:"date_shipped,omitempty"`
		Address      *struct {
			Name struct {
				Full string `json:"full,omitempty"`
			} `json:"name,omitempty"`
			Full        string `json:"full,omitempty"`
			CountryCode string `json:"country_code,omitempty"`
		} `json:"address,omitempty"`
	} `json:"shipping,omitempty"`
	Cost     Cost `json:"cost,omitempty"`
	DispCost Cost `json:"disp_cost,omitempty"`
}

type Cost struct {
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
}

func (h Header) PrimaryKey() int {
	return h.ID
}

func (h Header) Label() entity.Label {
	return entity.LabelOrderHeader
}

// GetOrderHeader retrieves the transactional data for a specific order.
// https://www.bricklink.com/v3/api.page?page=get-order
func (o *Orders) GetOrderHeader(id int) (*Header, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetOrder, id), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Header
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// GetOrderHeaders retrieves a list of orders you received or placed.
// It does not include order items, messages, or problems.
// Use GetOrders to retrieve the full order details.
// https://www.bricklink.com/v3/api.page?page=get-orders
func (o *Orders) GetOrderHeaders(options ...RequestOption) ([]Header, error) {
	var opts requestOptions
	opts.withOpts(options)
	query, err := opts.toQuery(queryTargetGetOrders)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, pathGetOrders, query, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out []Header
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// UpdateOrder updates the properties of a specific order header.
// It strips out values that are read-only before submitting the update.
// https://www.bricklink.com/v3/api.page?page=update-order
func (o *Orders) UpdateOrder(header Header) (*Header, error) {
	if header.PrimaryKey() == 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	// Only update the fields that can be updated.
	h := Header{}
	h.Cost.Credit = header.Cost.Credit
	h.Cost.Insurance = header.Cost.Insurance
	h.Cost.Etc1 = header.Cost.Etc1
	h.Cost.Etc2 = header.Cost.Etc2
	h.Cost.Shipping = header.Cost.Shipping
	h.Shipping.DateShipped = header.Shipping.DateShipped
	h.Shipping.MethodID = header.Shipping.MethodID
	h.Shipping.TrackingLink = header.Shipping.TrackingLink
	h.Shipping.TrackingNo = header.Shipping.TrackingNo
	h.Remarks = header.Remarks

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	req, err := o.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf(pathUpdateOrder, header.PrimaryKey()), nil, body)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Header
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
)

type Order struct {
	Header
	Items            []Item
	Messages         []Message
	FeedbackSent     Feedback
	FeedbackReceived Feedback
	Problem
}

func (o Order) PrimaryKey() int {
	return o.ID
}

func (o Order) Label() entity.Label {
	return entity.LabelOrder
}

// UpdateOrderStatus updates the order status for an order.
// https://www.bricklink.com/v3/api.page?page=update-order-status
func (o *Orders) UpdateOrderStatus(id int, status OrderStatus) error {
	if id == 0 {
		return fmt.Errorf("a positive value for id is required")
	}
	if status == StatusUndefined {
		return fmt.Errorf("a valid order status is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body := []byte(fmt.Sprintf(statusBody, updateFieldOrderStatus, status))

	req, err := o.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf(pathUpdateStatus, id), nil, body)
	if err != nil {
		return err
	}

	_, err = o.Client.Do(req)
	return err
}

// UpdatePaymentStatus updates the payment for an order.
// https://www.bricklink.com/v3/api.page?page=update-payment-status
func (o *Orders) UpdatePaymentStatus(id int, status PaymentStatus) error {
	if id == 0 {
		return fmt.Errorf("a positive value for id is required")
	}
	if status == PaymentStatusUndefined {
		return fmt.Errorf("a valid payment status is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body := []byte(fmt.Sprintf(statusBody, updateFieldPaymentStatus, status))

	req, err := o.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf(pathUpdatePayment, id), nil, body)
	if err != nil {
		return err
	}

	_, err = o.Client.Do(req)
	return err
}

// SendDriveThru issues a drive-thru message to the buyer for an order.
// https://www.bricklink.com/v3/api.page?page=send-drive-thru
func (o *Orders) SendDriveThru(id int, options ...RequestOption) error {
	if id == 0 {
		return fmt.Errorf("a positive value for id is required")
	}
	var opts requestOptions
	opts.withOpts(options)

	query, err := opts.toQuery(queryTargetDriveThru)

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(pathSendDriveThru, id), query, nil)
	if err != nil {
		return err
	}

	_, err = o.Client.Do(req)
	return err
}

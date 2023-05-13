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
}

func (o Order) PrimaryKey() int {
	return o.ID
}

func (o Order) Label() entity.Label {
	return entity.LabelOrder
}

// GetOrder is a convenience method that returns everything about an order at once.
// If the first call succeeds, this method requires 4 API calls.
// Errors for messages, items, and feedback are logged as messages in the returned order.
// Errors will have 'Error' as the subject.
func (o Orders) GetOrder(id int) (*Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	header, err := o.GetOrderHeader(id)
	if err != nil {
		return nil, fmt.Errorf("order %d not found: %s", id, err.Error())
	}

	order := Order{
		Header: *header,
	}

	order.Items, err = o.GetOrderItems(id)
	if err != nil {
		order.Messages = append(order.Messages, Message{
			Subject: "Error",
			Body:    fmt.Sprintf("Error retrieving order items: %s", err.Error()),
		})
	}

	feedback, err := o.GetOrderFeedback(id)
	if err != nil {
		order.Messages = append(order.Messages, Message{
			Subject: "Error",
			Body:    fmt.Sprintf("Error retrieving order feedback: %s", err.Error()),
		})
	}
	for _, v := range feedback {
		switch v.From {
		case order.BuyerName:
			order.FeedbackReceived = v
		case order.SellerName:
			order.FeedbackSent = v
		default:
			order.Messages = append(order.Messages, Message{
				Subject: fmt.Sprintf("Feedback to %s", v.To),
				Body:    fmt.Sprintf("Feedback from %s: %s", v.From, v.Comment),
			})
		}
	}

	messages, err := o.GetOrderMessages(id)
	if err != nil {
		order.Messages = append(order.Messages, Message{
			Subject: "Error",
			Body:    fmt.Sprintf("Error retrieving order messages: %s", err.Error()),
		})
	} else {
		order.Messages = append(order.Messages, messages...)
	}

	return &order, nil
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

	body := []byte(fmt.Sprintf(statusBody, updateFieldOrderStatus, status.String()))

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

	body := []byte(fmt.Sprintf(statusBody, updateFieldPaymentStatus, status.String()))

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

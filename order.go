package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/order"

// OrderClient provides an interface for interacting with the Order API.
// TODO Is it a client??
type OrderClient interface {
	GetOrders(direction string, status string, filed bool) ([]order.Order, error)
	GetOrder(orderID int) error
	GetOrderItems(orderID int) ([]order.Item, error)
	GetOrderMessages(orderID int) ([]order.Message, error)
	GetOrderFeedback(orderID int) ([]order.Feedback, error)
	UpdateOrder(orderID int, o order.Order) error
	UpdateStatus(orderID int) error
	UpdatePaymentStatus(orderID int, status string) error
	SendDriveThrough(orderID int) error
}

type OrderOption func(opts orderOptions) orderOptions

type orderOptions struct {
	id   string
	body string
}

func (oo *orderOptions) withOpts(opts []func(opts orderOptions) orderOptions) {
	for _, opt := range opts {
		*oo = opt(*oo)
	}
}

func WithID(id string) OrderOption {
	return func(opts orderOptions) orderOptions {
		opts.id = id
		return opts
	}
}

func WithBody(body string) OrderOption {
	return func(opts orderOptions) orderOptions {
		opts.body = body
		return opts
	}
}

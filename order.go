package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/order"

// OrderAPI provides an interface for interacting with the Order API.
type OrderAPI interface {
	GetOrders(direction string, status string, filed bool) ([]order.Order, error)
	GetOrder(orderID int) error
	GetOrderItems(orderID int) ([]order.Item, error)
	GetOrderMessages(orderID int) ([]order.Message, error)
	GetOrderFeedback(orderID int) ([]order.Feedback, error)
	UpdateOrder(orderID int, o order.Order) error
	UpdateStatus(orderID int) error
	UpdatePaymentStatus(orderID int, status string) error
	SendDriveThrough(orderID int) error

	// Member
	GetNote() (*order.Note, error)
	UpsertNote(note string) (*order.Note, error)
	DeleteNote() error

	// Feedback
	PostFeedback() (*order.Feedback, error)
	ReplyFeedback() error
}

type OrderOption func(opts *orderOptions)

type orderOptions struct {
	id   string
	body string
}

func (oo *orderOptions) withOpts(opts []func(opts *orderOptions)) {
	for _, opt := range opts {
		opt(oo)
	}
}

func WithID(id string) OrderOption {
	return func(opts *orderOptions) {
		opts.id = id
	}
}

func WithBody(body string) OrderOption {
	return func(opts *orderOptions) {
		opts.body = body
	}
}

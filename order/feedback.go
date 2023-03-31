package order

import "github.com/funwithbots/go-bricklink-api/util"

type Feedback struct{}

func (Order) GetOrderFeedback(orderID int) ([]interface{}, error) {
	return nil, util.ErrNotImplemented
}

package order

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/util"
)

type Message struct {
	Subject  string    `json:"subject"`
	Body     string    `json:"body"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	DateSent time.Time `json:"dateSent"`
}

func (Order) GetOrderMessages(id int) ([]Message, error) {
	return nil, util.ErrNotImplemented
}

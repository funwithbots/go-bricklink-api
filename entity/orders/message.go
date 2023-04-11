package orders

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Message struct {
	Subject  string    `json:"subject"`
	Body     string    `json:"body"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	DateSent time.Time `json:"dateSent"`
}

func (o *Orders) GetOrderMessages(id int) ([]Message, error) {
	return nil, util.ErrNotImplemented
}

func (m *Message) PrimaryKey() int {
	return int(m.DateSent.Unix())
}

func (m *Message) Label() entity.Label {
	return entity.LabelMessage
}

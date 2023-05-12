package orders

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Message struct {
	Subject  string    `json:"subject"`
	Body     string    `json:"body"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	DateSent time.Time `json:"dateSent"`
}

func (m *Message) PrimaryKey() int {
	return int(m.DateSent.Unix())
}

func (m *Message) Label() entity.Label {
	return entity.LabelMessage
}

func (o *Orders) GetOrderMessages(id int) ([]Message, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetMessages, id), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out []Message
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

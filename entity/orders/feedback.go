package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
	"github.com/funwithbots/go-bricklink-api/util"
)

// Feedback represents a feedback resource.
type Feedback struct {
	ID        int         `json:"feedback_id, omitempty"`
	OrderID   int         `json:"order_id"`
	From      string      `json:"from,omitempty"`
	To        string      `json:"to,omitempty"`
	DateRated time.Time   `json:"date_rated,omitempty"`
	Rating    util.Rating `json:"rating"`
	Rater     string      `json:"rating_of_bs,omitempty"` // rating of buyer or seller (B or S)
	Comment   string      `json:"comment"`
	Reply     string      `json:"reply,omitempty"`
}

func (f *Feedback) PrimaryKey() int {
	return f.ID
}

func (f *Feedback) Label() entity.Label {
	return entity.LabelFeedback
}

// PostFeedback posts feedback for an order.
// OrderID, Rating, and Comment must be set
func (o *Orders) PostFeedback(fb Feedback) (*Feedback, error) {
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body, err := json.Marshal(fb)
	if err != nil {
		return nil, err
	}

	req, err := o.NewRequestWithContext(ctx, http.MethodPost, pathPostFeedback, nil, body)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Feedback
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// ReplyFeedback replies to feedback.
// Reply must be set.
func (o *Orders) ReplyFeedback(fb Feedback) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body, err := json.Marshal(fb)
	if err != nil {
		return err
	}

	req, err := o.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(pathReplyFeedback, fb.PrimaryKey()), nil, body)
	if err != nil {
		return err
	}

	_, err = o.Client.Do(req)
	return err
}

// GetFeedbackList returns a list of feedback in the given direction.
func (o *Orders) GetFeedbackList(options ...RequestOption) ([]Feedback, error) {
	var opts requestOptions
	opts.withOpts(options)
	query, err := opts.toQuery(queryTargetGetFeedbackList)

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, pathGetFeedbackList, query, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out []Feedback
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetFeedback returns the feedback for the given feedback ID.
func (o *Orders) GetFeedback(id int) (*Feedback, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetFeedback, id), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Feedback
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// GetOrderFeedback returns the feedback for an order
func (o *Orders) GetOrderFeedback(id int) ([]Feedback, error) {
	return nil, util.ErrNotImplemented
}

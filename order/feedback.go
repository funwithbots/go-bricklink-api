package order

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/util"
)

type FeedbackDirection string

const (
	FeedbackDirectionIn  = "in"
	FeedbackDirectionOut = "out"
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

// PostFeedback posts feedback for an order.
// OrderID, Rating, and Comment must be set
func (f *Feedback) PostFeedback() (*Feedback, error) {
	return nil, util.ErrNotImplemented
}

// ReplyFeedback replies to feedback.
// Reply must be set.
func (f *Feedback) ReplyFeedback() error {
	return util.ErrNotImplemented
}

// GetFeedbackList returns a list of feedback for the given direction.
func GetFeedbackList(dir FeedbackDirection) ([]Feedback, error) {
	return nil, util.ErrNotImplemented
}

// GetFeedback returns the feedback for the given feedback ID.
func GetFeedback(id int) (*Feedback, error) {
	return nil, util.ErrNotImplemented
}

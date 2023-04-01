package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/feedback"

// FeedbackAPI provides an interface for interacting with the Feedback API.
type FeedbackAPI interface {
	PostFeedback() (*feedback.Feedback, error)
	ReplyFeedback() error
}

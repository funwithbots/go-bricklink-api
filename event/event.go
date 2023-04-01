package event

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/util"
)

// Event represents a push event notification from Bricklink.
type Event struct {
	EventType  string    `json:"event_type"`
	ResourceId int       `json:"resource_id"`
	Timestamp  time.Time `json:"timestamp"`
}

// GetEvents returns a list of unread push notifications.
// If you provided callback URLs to get notifications, you don't need to call this method.
func GetEvents() ([]Event, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}
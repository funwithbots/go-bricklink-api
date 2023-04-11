package event

import "time"

type Event struct {
	EventType  string    `json:"event_type"`
	ResourceId int       `json:"resource_id"`
	Timestamp  time.Time `json:"timestamp"`
}

func (Event) GetEvent() ([]Event, error) {
	// TODO implement me
	panic("implement me")
}

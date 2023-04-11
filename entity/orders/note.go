package orders

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
)

// Note represents a note about a member.
type Note struct {
	ID        int       `json:"note_id,omitempty"`
	UserName  string    `json:"user_name"`
	Note      string    `json:"note_text"`
	DateNoted time.Time `json:"date_noted,omitempty"`
}

func (n *Note) PrimaryKey() int {
	return n.ID
}

func (n *Note) Label() entity.Label {
	return entity.LabelNote
}

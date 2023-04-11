package orders

import (
	"hash/fnv"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

// Member represents a Bricklink member.
type Member struct {
	UserName string
	Ratings  []Feedback
	Note     Note
}

// MemberRating represents a Bricklink user.
type MemberRating struct {
	UserName string `json:"user_name"`

	Rating []Feedback `json:"rating"`
}

func (m *Member) PrimaryKey() int {
	hash := fnv.New32a()
	hash.Write([]byte(m.UserName))
	return int(hash.Sum32())
}

func (m *Member) Label() entity.Label {
	return entity.LabelMember
}

// GetMemberRatings returns the details feedback ratings for a specific member.
func (o *Orders) GetMemberRatings(memberID string) ([]MemberRating, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// GetNote returns the note for a specific member.
func (o *Orders) GetNote() (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// UpsertNote creates or updates a note for the member.
func (o *Orders) UpsertNote(note Note) (*Note, error) {
	// TODO implement me
	if note.UserName == "" {
		return nil, util.ErrInvalidArgument
	}
	if note.ID == 0 {
		return o.postNote(note)
	}
	return o.updateNote(note)
}

// postNote creates a note for the member.
func (o *Orders) postNote(note Note) (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// updateNote updates the member note.
func (o *Orders) updateNote(note Note) (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// DeleteNote deletes the member note.
func (o *Orders) DeleteNote() error {
	// TODO implement me
	return util.ErrNotImplemented
}

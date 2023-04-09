package order

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

// GetMemberRatings returns the details feedback ratings for a specific member.
func GetMemberRatings(memberID string) ([]MemberRating, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (m *Member) PrimaryKey() int {
	hash := fnv.New32a()
	hash.Write([]byte(m.UserName))
	return int(hash.Sum32())
}

func (m *Member) Label() entity.Label {
	return entity.LabelMember
}

// GetNote returns the note for a specific member.
func (m *Member) GetNote() (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// UpsertNote creates or updates a note for the member.
func (m *Member) UpsertNote(note string) (*Note, error) {
	// TODO implement me
	if note == "" {
		return nil, util.ErrInvalidArgument
	}
	if m.Note.ID == 0 {
		return m.postNote(note)
	}
	return m.updateNote(note)
}

// postNote creates a note for the member.
func (m *Member) postNote(note string) (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// updateNote updates the member note.
func (m *Member) updateNote(note string) (*Note, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// DeleteNote deletes the member note.
func (m *Member) DeleteNote() error {
	// TODO implement me
	return util.ErrNotImplemented
}

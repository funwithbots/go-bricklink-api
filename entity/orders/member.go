package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
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
func (o *Orders) GetMemberRatings(name string) ([]MemberRating, error) {
	if name == "" {
		return nil, util.ErrInvalidArgument
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetMemberRating, name), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out []MemberRating
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetMemberNote returns the note for a specific member.
func (o *Orders) GetMemberNote(name string) (*Note, error) {
	if name == "" {
		return nil, util.ErrInvalidArgument
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetMemberNote, name), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Note
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// UpsertMemberNote creates or updates a note for the member.
func (o *Orders) UpsertMemberNote(note Note) (*Note, error) {
	if note.UserName == "" {
		return nil, util.ErrInvalidArgument
	}
	if note.ID == 0 {
		return o.createMemberNote(note)
	}

	return o.updateMemberNote(note)
}

// createMemberNote creates a note for the member.
func (o *Orders) createMemberNote(note Note) (*Note, error) {
	if note.UserName == "" {
		return nil, util.ErrInvalidArgument
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body, err := json.Marshal(note)
	if err != nil {
		return nil, err
	}

	req, err := o.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(pathCreateMemberNote, note.UserName), nil, body)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Note
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// updateMemberNote updates the member note.
func (o *Orders) updateMemberNote(note Note) (*Note, error) {
	if note.UserName == "" {
		return nil, util.ErrInvalidArgument
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	body, err := json.Marshal(note)
	if err != nil {
		return nil, err
	}

	req, err := o.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf(pathUpdateMemberNote, note.UserName), nil, body)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out Note
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteMemberNote deletes the member note.
func (o *Orders) DeleteMemberNote(name string) error {
	if name == "" {
		return util.ErrInvalidArgument
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(pathDeleteMemberNote, name), nil, nil)
	if err != nil {
		return err
	}

	_, err = o.Client.Do(req)
	return err
}

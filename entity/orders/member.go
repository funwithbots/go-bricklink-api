package orders

import (
	"context"
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

	Rating struct {
		Complaints int `json:"COMPLAINT"`
		Neutrals   int `json:"NEUTRAL"`
		Praises    int `json:"PRAISE"`
	} `json:"rating"`
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
func (o *Orders) GetMemberRatings(name string) (*MemberRating, error) {
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

	var out MemberRating
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

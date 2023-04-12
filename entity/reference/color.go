package reference

import (
	"context"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Color struct {
	ID        int    `json:"color_id"`
	ColorName string `json:"color_name"`
	ColorCode string `json:"color_code"`
	ColorType string `json:"color_type"`
}

var Colors map[int]Color

func (c Color) PrimaryKey() int {
	return c.ID
}

func (c Color) Label() entity.Label {
	return entity.LabelColor
}

// GetColors returns a list of colors.
func (r *Reference) GetColors() ([]Color, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(
		ctx,
		http.MethodGet,
		pathGetColors,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var colors []Color
	if err := internal.Parse(res.Body, &colors); err != nil {
		return nil, err
	}

	return colors, nil
}

// GetColor returns a color by color ID.
func (r *Reference) GetColor(colorID int) (*Color, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetColor, colorID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var color Color
	if err := internal.Parse(res.Body, &color); err != nil {
		return nil, err
	}

	return &color, nil
}

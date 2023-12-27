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

var colorMap = make(map[int]Color)

func (c Color) PrimaryKey() int {
	return c.ID
}

func (c Color) Label() entity.Label {
	return entity.LabelColor
}

// GetColors returns a list of colors.
func (r *Reference) GetColors() ([]Color, error) {
	// if colorMap is primed, return those values
	if len(colorMap) > 0 {
		colors := make([]Color, 0, len(colorMap))
		for _, color := range colorMap {
			colors = append(colors, color)
		}
		return colors, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		pathGetColors,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var colors []Color
	if err := internal.Parse(res.Body, &colors); err != nil {
		return nil, err
	}

	// prime colorMap
	for _, color := range colors {
		colorMap[color.ID] = color
	}

	return colors, nil
}

// GetColor returns a color by color ID.
func (r *Reference) GetColor(colorID int) (*Color, error) {
	// if colorMap is primed, try to get the color without another call to Bricklink.
	if color, ok := colorMap[colorID]; ok {
		return &color, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetColor, colorID),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var color Color
	if err := internal.Parse(res.Body, &color); err != nil {
		return nil, err
	}

	return &color, nil
}

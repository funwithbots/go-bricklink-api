package color

import "github.com/funwithbots/go-bricklink-api/util"

type Color struct {
	ID        int    `json:"color_id"`
	ColorName string `json:"color_name"`
	ColorCode string `json:"color_code"`
	ColorType string `json:"color_type"`
}

var Colors map[int]Color

// GetColors returns a list of colors.
func GetColors() ([]Color, error) {
	return nil, util.ErrNotImplemented
}

// GetColor returns a color by color ID.
func GetColor(colorID int) (Color, error) {
	return Color{}, util.ErrNotImplemented
}

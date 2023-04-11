package reference

import (
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
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
	return nil, util.ErrNotImplemented
}

// GetColor returns a color by color ID.
func (r *Reference) GetColor(colorID int) (Color, error) {
	return Color{}, util.ErrNotImplemented
}

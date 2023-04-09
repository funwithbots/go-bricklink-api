package reference

import (
	"github.com/funwithbots/go-bricklink-api"
)

type Reference struct {
	bl go_bricklink_api.Bricklink
}

func New(bl go_bricklink_api.Bricklink) *Reference {
	return &Reference{
		bl: bl,
	}
}

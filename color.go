package go_bricklink_api

import (
	"github.com/funwithbots/go-bricklink-api/color"
)

type ColorAPI interface {
	GetColors() ([]color.Color, error)
	GetColor(colorID int) (color.Color, error)
}

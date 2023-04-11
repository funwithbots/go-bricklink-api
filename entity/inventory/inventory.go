package inventory

import (
	bricklink "github.com/funwithbots/go-bricklink-api"
)

const (
	pathGetItem = "/inventories/%s"
)

type Inventory struct {
	bl bricklink.Bricklink
}

func New(bl bricklink.Bricklink) *Inventory {
	return &Inventory{
		bl: bl,
	}
}

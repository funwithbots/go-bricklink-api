package orders

import bricklink "github.com/funwithbots/go-bricklink-api"

const (
	pathGetOrders = "/orders/%s/%s"
)

type Orders struct {
	bl bricklink.Bricklink
}

func New(bl bricklink.Bricklink) *Orders {
	return &Orders{
		bl: bl,
	}
}

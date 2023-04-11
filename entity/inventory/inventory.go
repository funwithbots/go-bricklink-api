package inventory

import (
	bricklink "github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/util"
)

const (
	guideTypeSold  = "sold"
	guideTypeStock = "stock"

	pathGetOrders = "/orders/%s/%s"
)

type Inventory struct {
	bl bricklink.Bricklink
}

func New(bl bricklink.Bricklink) *Inventory {
	return &Inventory{
		bl: bl,
	}
}

func (inv *Inventory) GetInventory(...RequestOption) ([]Item, error) {
	return nil, util.ErrNotImplemented
}

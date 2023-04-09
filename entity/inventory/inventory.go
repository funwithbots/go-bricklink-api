package inventory

import (
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Inventory []Item

// PrimaryKey isn't relevant for this entity.
func (inv *Inventory) PrimaryKey() int {
	return 0
}

func (inv *Inventory) Label() entity.Label {
	return entity.LabelInventory
}

func GetInventory(...RequestOption) (Inventory, error) {
	return nil, util.ErrNotImplemented
}

package inventory

import "github.com/funwithbots/go-bricklink-api/util"

type Inventory struct {
	InventoryID string `json:"inventory_id"`
}

func GetInventoryList() ([]Inventory, error) {
	return nil, util.ErrNotImplemented
}

func GetInventory(id string) (*Item, error) {
	return nil, util.ErrNotImplemented
}

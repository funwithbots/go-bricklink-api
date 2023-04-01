package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/inventory"

// InventoryAPI provides an interface for interacting with the Inventory API.
type InventoryAPI interface {
	GetInventoryList() ([]inventory.Inventory, error)
	GetInventory(id string) (*inventory.Item, error)
}

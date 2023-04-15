package inventory

import (
	bricklink "github.com/funwithbots/go-bricklink-api"
)

const (
	// completeness values are only valid for sets
	completenessComplete   = "C"
	completenessIncomplete = "B"
	completenessSealed     = "S"

	pathGetItem     = "/inventories/%d"
	pathGetItems    = "/inventories"
	pathCreateItem  = "/inventories"    // POST
	pathCreateItems = "/inventories"    // POST
	pathUpdateItem  = "/inventories/%d" // PUT
	pathDeleteItem  = "/inventories/%d" // DELETE
)

type Inventory struct {
	bricklink.Bricklink
}

func New(bl bricklink.Bricklink) *Inventory {
	return &Inventory{bl}
}

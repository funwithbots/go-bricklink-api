package inventory

import (
	bl "github.com/funwithbots/go-bricklink-api"
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
	bl.Bricklink
}

func New(b bl.Bricklink) *Inventory {
	return &Inventory{b}
}

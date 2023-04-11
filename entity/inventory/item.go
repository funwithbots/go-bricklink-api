package inventory

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Item struct {
	ID      int            `json:"inventory_id, omitempty"`
	Item    reference.Item `json:"item"`
	ColorID int            `json:"color_id"`

	// When updating item with a new quantity, you must specify the difference between the new quantity and
	// the old quantity as a plus/minus value.
	Quantity      int       `json:"quantity"`
	NewOrUsed     string    `json:"new_or_used,omitempty"`
	Completeness  string    `json:"completeness,omitempty"`
	UnitPrice     string    `json:"unit_price"`
	BindID        int       `json:"bind_id,omitempty"`
	Description   string    `json:"description"`
	Remarks       string    `json:"remarks"`
	Bulk          int       `json:"bulk"`
	IsRetain      bool      `json:"is_retain"`
	IsStockRoom   bool      `json:"is_stock_room"`
	StockRoomID   string    `json:"stock_room_id,omitempty"`
	DateCreated   time.Time `json:"date_created,omitempty"`
	SaleRate      int       `json:"sale_rate"`
	MyCost        string    `json:"my_cost"`
	TierQuantity1 int       `json:"tier_quantity1"`
	TierPrice1    string    `json:"tier_price1"`
	TierQuantity2 int       `json:"tier_quantity2"`
	TierPrice2    string    `json:"tier_price2"`
	TierQuantity3 int       `json:"tier_quantity3"`
	TierPrice3    string    `json:"tier_price3"`
}

func (it Item) PrimaryKey() int {
	return it.ID
}

func (it Item) Label() entity.Label {
	return entity.LabelInventoryItem
}

// GetItem implements the Get store inventory endpoint.
// Bricklink calls this "Get Inventory" but it's really just a get of an item.
// https://www.bricklink.com/v3/api.page?page=get-inventory
func (inv *Inventory) GetItem(id string) (*Item, error) {
	return nil, util.ErrNotImplemented
}

// CreateItem creates a single item in the inventory.
// Bricklink calls this "Create Inventory" but it's really just a create of an item.
func (inv *Inventory) CreateItem(item Item) (*Item, error) {
	return nil, util.ErrNotImplemented
}

func (inv *Inventory) UpdateItem(Item) (*Item, error) {
	return nil, util.ErrNotImplemented
}

// CreateItems creates multiple items in a single request.
// Bricklink calls this "Create Inventories" but it's really just a bulk create of items.
// Item IDs are not returned.
func (inv *Inventory) CreateItems([]Item) error {
	return util.ErrNotImplemented
}

// DeleteItem deletes an item from the inventory.
// Bricklink calls this "Delete Inventory" but it's really just a delete of an item.
func (inv *Inventory) DeleteItem(id int) error {
	return util.ErrNotImplemented
}

package inventory

import (
	"net/http"
	"time"

	"github.com/funwithbots/go-bricklink-api/reference"
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

// GetItem implements the Get store inventory endpoint.
// https://www.bricklink.com/v3/api.page?page=get-inventory
func GetItem(id string) (*http.Request, error) {
	return nil, util.ErrNotImplemented
}

// CreateItem
func (it *Item) CreateItem() (*Item, error) {
	return nil, util.ErrNotImplemented
}

func (it *Item) UpdateItem() (*Item, error) {
	return nil, util.ErrNotImplemented
}

func CreateInventories([]Item) error {
	return util.ErrNotImplemented
}

func DeleteInventory(id int) error {
	return util.ErrNotImplemented
}

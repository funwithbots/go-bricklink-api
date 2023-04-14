package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Item struct {
	ID      int            `json:"inventory_id,omitempty"`
	Item    reference.Item `json:"item"`
	ColorID int            `json:"color_id"`

	// When updating item with a new quantity, you must specify the difference between the new quantity and
	// the old quantity as a plus/minus value.
	Quantity      int        `json:"quantity"`
	NewOrUsed     string     `json:"new_or_used,omitempty"`
	Completeness  string     `json:"completeness,omitempty"`
	UnitPrice     string     `json:"unit_price"`
	BindID        int        `json:"bind_id,omitempty"`
	Description   string     `json:"description,omitempty"`
	Remarks       string     `json:"remarks,omitempty"`
	Bulk          int        `json:"bulk,omitempty"`
	IsRetain      bool       `json:"is_retain"`
	IsStockRoom   bool       `json:"is_stock_room"`
	StockRoomID   string     `json:"stock_room_id,omitempty"`
	DateCreated   *time.Time `json:"date_created,omitempty"`
	SaleRate      int        `json:"sale_rate,omitempty"`
	MyCost        string     `json:"my_cost,omitempty"`
	TierQuantity1 int        `json:"tier_quantity1,omitempty"`
	TierPrice1    string     `json:"tier_price1,omitempty"`
	TierQuantity2 int        `json:"tier_quantity2,omitempty"`
	TierPrice2    string     `json:"tier_price2,omitempty"`
	TierQuantity3 int        `json:"tier_quantity3,omitempty"`
	TierPrice3    string     `json:"tier_price3,omitempty"`
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
func (inv *Inventory) GetItem(id int) (*Item, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()

	req, err := inv.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItem, id), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := inv.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := internal.Parse(res.Body, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

// GetItems implements the Get store inventories endpoint.
// It searches the store inventory and returns a list of matching items.
func (inv *Inventory) GetItems(options ...RequestOption) ([]Item, error) {
	var opts requestOptions
	opts.withOpts(options)
	query, err := opts.toQuery(queryTargetGetItems)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()

	req, err := inv.NewRequestWithContext(ctx, http.MethodGet, pathGetItems, query, nil)
	if err != nil {
		return nil, err
	}

	res, err := inv.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := internal.Parse(res.Body, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// CreateItem creates a single item in the inventory.
// Bricklink calls this "Create Inventory" but it's really just a create of an item.
func (inv *Inventory) CreateItem(item Item) (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()

	body, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := inv.NewRequestWithContext(ctx, http.MethodPost, pathCreateItem, nil, body)
	if err != nil {
		return nil, err
	}

	res, err := inv.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var it Item
	if err := internal.Parse(res.Body, &it); err != nil {
		return nil, err
	}

	return &it, nil
}

// CreateItems creates multiple items in a single request.
// Bricklink calls this "Create Inventories" but it's really just a bulk create of items.
// Item IDs are not returned.
func (inv *Inventory) CreateItems(items []Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()

	body, err := json.Marshal(items)
	if err != nil {
		return err
	}

	req, err := inv.NewRequestWithContext(ctx, http.MethodPost, pathCreateItem, nil, body)
	if err != nil {
		return err
	}

	if _, err = inv.Client.Do(req); err != nil {
		return err
	}

	return nil
}

func (inv *Inventory) UpdateItem(item Item) (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()
	item.DateCreated = nil
	if item.TierQuantity1 == 0 {
		item.TierPrice1 = ""
	}
	if item.TierQuantity2 == 0 {
		item.TierPrice2 = ""
	}
	if item.TierQuantity3 == 0 {
		item.TierPrice3 = ""
	}

	body, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := inv.NewRequestWithContext(ctx, http.MethodPut, fmt.Sprintf(pathUpdateItem, item.PrimaryKey()), nil, body)
	if err != nil {
		return nil, err
	}

	res, err := inv.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var it Item
	if err := internal.Parse(res.Body, &it); err != nil {
		return nil, err
	}

	return &it, nil
}

// DeleteItem deletes an item from the inventory.
// Bricklink calls this "Delete Inventory" but it's really just a delete of an item.
func (inv *Inventory) DeleteItem(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), inv.Timeout)
	defer cancel()

	req, err := inv.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf(pathDeleteItem, id), nil, nil)
	if err != nil {
		return err
	}

	res, err := inv.Client.Do(req)
	if err != nil {
		return err
	}

	var it interface{}
	if err := internal.Parse(res.Body, &it); err != nil {
		return err
	}

	return nil
}
func fromParams(params map[string]string) string {
	query := make([]string, 0, len(params))
	for k, v := range params {
		query = append(query, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(query, "&")
}

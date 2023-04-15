package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Item struct {
	InventoryID        int            `json:"inventory_id"`
	Item               reference.Item `json:"item"`
	ColorID            int            `json:"color_id"`
	Quantity           int            `json:"quantity"`
	NewOrUsed          string         `json:"new_or_used"`
	Completeness       string         `json:"completeness"`
	UnitPrice          string         `json:"unit_price"`
	UnitPriceFinal     string         `json:"unit_price_final"`
	DispUnitPrice      string         `json:"disp_unit_price"`
	DispUnitPriceFinal string         `json:"disp_unit_price_final"`
	CurrencyCode       string         `json:"currency_code"`
	DispCurrencyCode   string         `json:"disp_currency_code"`
	Description        string         `json:"description"`
	Remarks            string         `json:"remarks"`
}

func (it Item) PrimaryKey() int {
	return it.InventoryID
}

func (it Item) Label() entity.Label {
	return entity.LabelOrderItem
}

func (o *Orders) GetOrderItems(id int) ([]Item, error) {
	if id <= 0 {
		return nil, fmt.Errorf("a positive value for id is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetOrderItems, id), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var out []Item
	if err := internal.Parse(res.Body, &out); err != nil {
		return nil, err
	}

	return out, nil
}

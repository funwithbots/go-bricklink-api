package order

import (
	"github.com/funwithbots/go-bricklink-api/catalog"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Item struct {
	InventoryID        int          `json:"inventory_id"`
	Item               catalog.Item `json:"item"`
	ColorID            int          `json:"color_id"`
	Quantity           int          `json:"quantity"`
	NewOrUsed          string       `json:"new_or_used"`
	Completeness       string       `json:"completeness"`
	UnitPrice          string       `json:"unit_price"`
	UnitPriceFinal     string       `json:"unit_price_final"`
	DispUnitPrice      string       `json:"disp_unit_price"`
	DispUnitPriceFinal string       `json:"disp_unit_price_final"`
	CurrencyCode       string       `json:"currency_code"`
	DispCurrencyCode   string       `json:"disp_currency_code"`
	Description        string       `json:"description"`
	Remarks            string       `json:"remarks"`
}

func (Order) GetOrderItems(orderID int) ([]Item, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

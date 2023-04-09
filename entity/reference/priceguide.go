package reference

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

type PriceGuide struct {
	Item struct {
		No   string `json:"no"`
		Type string `json:"type"`
	} `json:"item"`
	NewOrUsed     string        `json:"new_or_used"`
	CurrencyCode  string        `json:"currency_code"`
	MinPrice      string        `json:"min_price"`
	MaxPrice      string        `json:"max_price"`
	AvgPrice      string        `json:"avg_price"`
	QtyAvgPrice   string        `json:"qty_avg_price"`
	UnitQuantity  int           `json:"unit_quantity"`
	TotalQuantity int           `json:"total_quantity"`
	PriceDetail   []PriceDetail `json:"price_detail"`
}

type PriceDetail struct {
	Quantity          int       `json:"quantity"`
	UnitPrice         string    `json:"unit_price"`
	SellerCountryCode string    `json:"seller_country_code"`
	BuyerCountryCode  string    `json:"buyer_country_code"`
	DateOrdered       time.Time `json:"date_ordered"`
	ShippingAvailable string    `json:"shipping_available"`
}

// PrimaryKey isn't meaningful for this entity.
func (pg *PriceGuide) PrimaryKey() int {
	return 0
}

func (pg *PriceGuide) Label() entity.Label {
	return entity.LabelPriceGuide
}

func GetPriceGuide(options ...RequestOption) (PriceGuide, error) {
	return PriceGuide{}, util.ErrNotImplemented
}

package reference

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type PriceGuide struct {
	Item          Item          `json:"item"`
	NewOrUsed     string        `json:"new_or_used"`
	CurrencyCode  string        `json:"currency_code"`
	MinPrice      string        `json:"min_price"`
	MaxPrice      string        `json:"max_price"`
	AvgPrice      string        `json:"avg_price"`
	QtyAvgPrice   string        `json:"qty_avg_price"`
	UnitQuantity  int           `json:"unit_quantity"`
	TotalQuantity int           `json:"total_quantity"`
	PriceDetails  []PriceDetail `json:"price_detail"`
}

type PriceDetail struct {
	Quantity          int       `json:"quantity"`
	UnitPrice         string    `json:"unit_price"`
	SellerCountryCode string    `json:"seller_country_code"`
	BuyerCountryCode  string    `json:"buyer_country_code"`
	DateOrdered       time.Time `json:"date_ordered"`
	ShippingAvailable bool      `json:"shipping_available"`
}

// PrimaryKey isn't meaningful for this entity.
func (pg *PriceGuide) PrimaryKey() int {
	return 0
}

func (pg *PriceGuide) Label() entity.Label {
	return entity.LabelPriceGuide
}

func (r *Reference) GetPriceGuide(options ...RequestOption) (*PriceGuide, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("item no is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	query, err := opts.toQuery(queryTargetPriceGuide)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetPriceGuide, opts.itemType, opts.itemNo),
		query,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var pg PriceGuide
	if err := internal.Parse(res.Body, &pg); err != nil {
		return nil, err
	}

	return &pg, nil
}

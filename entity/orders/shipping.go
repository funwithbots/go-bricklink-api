package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

const (
	ShippingAreaInternational = "I"
	ShippingAreaDomestic      = "D"
	ShippingAreaBoth          = "B"
)

type ShippingMethod struct {
	ID          int    `json:"method_id"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	Insurance   bool   `json:"insurance"`
	IsDefault   bool   `json:"is_default"`
	Area        string `json:"area"` // International, Domestic, Both (I/D/B)
	IsAvailable bool   `json:"is_available"`
}

func (sm ShippingMethod) PrimaryKey() int {
	return sm.ID
}

func (sm ShippingMethod) Label() entity.Label {
	return entity.LabelShippingMethod
}

// GetShippingMethod returns a shipping method by id from the cache.
// If forceRefresh is true, the cache will be refreshed.
// The Get Shipping Method endpoint is unneeded since a single call gets the entire list. Therefore, we can
// serve the response without additional api calls.
func (o *Orders) GetShippingMethod(id int, options ...RequestOption) (*ShippingMethod, error) {
	var opts requestOptions
	opts.withOpts(options)
	if opts.forceRefresh || (len(o.ShippingMethods) == 0) {
		if err := o.loadShippingMethods(); err != nil {
			return nil, err
		}
	}

	if sm, ok := o.ShippingMethods[id]; ok {
		return &sm, nil
	}
	return nil, fmt.Errorf("shipping method not found")
}

func (o *Orders) RefreshShippingMethods() error {
	return o.loadShippingMethods()
}

func (o *Orders) loadShippingMethods() error {
	ctx, cancel := context.WithTimeout(context.Background(), o.Timeout)
	defer cancel()

	req, err := o.NewRequestWithContext(ctx, http.MethodGet, pathGetShippingMethods, nil, nil)
	if err != nil {
		return err
	}

	res, err := o.Client.Do(req)
	if err != nil {
		return err
	}

	var sm []ShippingMethod
	if err := internal.Parse(res.Body, &sm); err != nil {
		return err
	}

	if len(sm) == 0 {
		return fmt.Errorf("no shipping methods found")
	}

	if len(o.ShippingMethods) == 0 {
		o.ShippingMethods = make(map[int]ShippingMethod)
	}

	for _, v := range sm {
		o.ShippingMethods[v.ID] = v
	}

	return nil

}

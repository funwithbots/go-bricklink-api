package reference

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Subset []SubsetItem

type SubsetItem struct {
	MatchNo int `json:"match_no"` // 0 value are unique parts with no alternates or counterparts
	Entries []struct {
		Item          Item `json:"item"`
		ColorID       int  `json:"color_id"`
		Quantity      int  `json:"quantity"`
		ExtraQuantity int  `json:"extra_quantity"`
		IsAlternate   bool `json:"is_alternate"`
		IsCounterpart bool `json:"is_counterpart"`
	} `json:"entries"`
}

// PrimaryKey isn't meaningful for this entity.
func (sub *Subset) PrimaryKey() int {
	return 0
}

func (sub *Subset) Label() entity.Label {
	return entity.LabelSubset
}

// GetSubset returns a list of items that make up the item.
func (r *Reference) GetSubset(options ...RequestOption) (Subset, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("id is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetSubset, opts.itemType, opts.itemNo), nil)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var items []SubsetItem
	if err := internal.Parse(res.Body, &items); err != nil {
		return nil, err
	}

	subset := items

	return subset, nil
}
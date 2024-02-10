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
	MatchNo int           `json:"match_no"` // 0 value are unique parts with no alternates or counterparts
	Entries []SubsetEntry `json:"entries"`
}

type SubsetEntry struct {
	Item          Item `json:"item"`
	ColorID       int  `json:"color_id"`
	Quantity      int  `json:"quantity"`
	ExtraQuantity int  `json:"extra_quantity"`
	IsAlternate   bool `json:"is_alternate"`
	IsCounterpart bool `json:"is_counterpart"`
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
		return nil, errors.New("item no is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	query, err := opts.toQuery(queryTargetSubsets)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetSubset, opts.itemType, opts.itemNo), query, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var items []SubsetItem
	if err := internal.Parse(res.Body, &items); err != nil {
		return nil, err
	}

	subsets := items

	return subsets, nil
}

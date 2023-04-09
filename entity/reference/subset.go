package reference

import (
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
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

func GetSubsets(opts ...RequestOption) (Subset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

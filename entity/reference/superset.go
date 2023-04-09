package reference

import (
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Superset []SupersetItem

type SupersetItem struct {
	ColorID int `json:"color_id"`
	Entries []struct {
		Item      Item   `json:"item"`
		Quantity  int    `json:"quantity"`
		AppearsAs string `json:"appears_as"`
	} `json:"entries"`
}

// PrimaryKey isn't meaningful for this entity.
func (Superset) PrimaryKey() int {
	return 0
}

func (Superset) Label() entity.Label {
	return entity.LabelSuperset
}

// GetSupersets returns a list of supersets that contain the item.
// If a colorID is provided, the list is filtered to supersets that contain the item in that color.
func GetSupersets(opts ...RequestOption) (Superset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

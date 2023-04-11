package reference

import (
	"time"

	"github.com/funwithbots/go-bricklink-api/util"
)

type KnownColor struct {
	ColorID  int `json:"color_id"`
	Quantity int `json:"quantity"`
}

type KnownColors struct {
	DateUpdated time.Time    `json:"-"`
	ItemType    string       `json:"-"`
	ItemID      string       `json:"-"`
	KnownColors []KnownColor `json:"known_colors"`
}

func (r *Reference) GetKnownColors(options ...RequestOption) (KnownColors, error) {
	// TODO implement me
	return KnownColors{}, util.ErrNotImplemented
}

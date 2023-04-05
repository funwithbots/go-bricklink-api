package reference

import (
	"hash/fnv"

	"github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/util"
)

var ItemTypes = map[string]string{
	"S": "Set",
	"P": "Part",
	"M": "Minifig",
	"G": "Gear",
	"B": "Book",
	"C": "Catalog",
	"O": "Original Box",
	"X": "Instruction",
	"U": "Unsorted Lot",
}

const (
	guideTypeSold  = "sold"
	guideTypeStock = "stock"
)

type Item struct {
	ID           string `json:"no"`
	Name         string `json:"name"`
	ItemType     string `json:"type"`
	CategoryID   int    `json:"category_id"`
	alternateNo  string `json:"alternate_no"`
	ImageURL     string `json:"image_url"`
	ThumbnailURL string `json:"thumbnail_url"`

	// Weight and Dimensions are specified to 2 decimal places.
	Weight string `json:"weight"`
	DimX   string `json:"dim_x"`
	DimY   string `json:"dim_y"`
	DimZ   string `json:"dim_z"`

	YearReleased int    `json:"year_released"`
	Description  string `json:"description"`
	IsObsolete   bool   `json:"is_obsolete"`
	LanguageCode string `json:"language_code"`
}

func (it *Item) PrimaryKey() int {
	hash := fnv.New32a()
	hash.Write([]byte(it.ID))
	return int(hash.Sum32())
}

func (it *Item) Label() go_bricklink_api.Type {
	return go_bricklink_api.Inventory
}

// GetSupersets returns a list of supersets that contain the item.
// If a colorID is provided, the list is filtered to supersets that contain the item in that color.
func (it *Item) GetSupersets(opts ...RequestOption) ([]Superset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (it *Item) GetSubsets(opts ...RequestOption) ([]Subset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (it *Item) GetPriceGuide(opts ...RequestOption) (*PriceGuide, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (it *Item) GetKnownColors() ([]KnownColor, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func GetItemImage(colorID int) (*Item, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func GetItem(id string) (*Item, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

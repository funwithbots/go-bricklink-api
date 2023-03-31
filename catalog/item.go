package catalog

import "github.com/funwithbots/go-bricklink-api/util"

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

type requestParams struct {
	itemType      string
	itemNo        string
	colorID       *int
	box           *bool
	instruction   *bool
	breakMiniFigs *bool
	breakSubsets  *bool
	guideType     string
	condition     string // new_or_used.
	countryCode   string
	region        string
	currencyCode  string
	vat           string // Y N O (Yes, No, nOrway)
}

type Item struct {
	ID           string `json:"no"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	CategoryID   int    `json:"category_id"`
	alternateNo  string `json:"alternate_no"`
	ImageUrl     string `json:"image_url"`
	ThumbnailUrl string `json:"thumbnail_url"`

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

// GetSupersets returns a list of supersets that contain the item.
// If a colorID is provided, the list is filtered to supersets that contain the item in that color.
func (it *Item) GetSupersets(opts requestParams) ([]Superset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (it *Item) GetSubsets(opts requestParams) ([]Subset, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

func (it *Item) GetPriceGuide(opts requestParams) (*PriceGuide, error) {
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

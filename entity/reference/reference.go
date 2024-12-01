package reference

import (
	"time"

	bricklink "github.com/funwithbots/go-bricklink-api"
)

const (
	defaultTimeout = 3 * time.Second
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
	GuideTypeSold  = "sold"
	GuideTypeStock = "stock"

	PGRegionAsia         = "asia"
	PGRegionEurope       = "europe"
	PGRegionAfrica       = "africa"
	PGRegionNorthAmerica = "north_america"
	PGRegionSouthAmerica = "south_america"
	PGRegionOceania      = "oceania"
	PGRegionMiddleEast   = "middle_east"
	PGRegionEU           = "eu"

	pathGetItem        = "/items/%s/%s"
	pathGetItemImage   = "/items/%s/%s/images/%d"
	pathGetSuperset    = "/items/%s/%s/supersets"
	pathGetSubset      = "/items/%s/%s/subsets"
	pathGetPriceGuide  = "/items/%s/%s/price"
	pathGetKnownColors = "/items/%s/%s/colors"
	pathGetCategories  = "/categories"
	pathGetCategory    = "/categories/%d"
	pathGetColors      = "/colors"
	pathGetColor       = "/colors/%d"
	pathGetItemMapping = "/item_mapping/%s"
	pathGetElementID   = "/item_mapping/%s/%s"
)

type Reference struct {
	*bricklink.Bricklink
}

func New(b *bricklink.Bricklink) *Reference {
	if b.Timeout == 0 {
		b.Timeout = defaultTimeout
	}

	return &Reference{
		b,
	}
}

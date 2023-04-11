package reference

import (
	bricklink "github.com/funwithbots/go-bricklink-api"
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

	pathGetItem        = "/items/%s/%s"
	pathGetItemImage   = "/items/%s/%s/images/%s"
	pathGetSuperset    = "/items/%s/%s/supersets?%s"
	pathGetSubset      = "/items/%s/%s/subsets?%s"
	pathGetPriceGuide  = "/items/%s/%s/price?%s"
	pathGetKnownColors = "/items/%s/%s/colors"
)

type Reference struct {
	bl bricklink.Bricklink
}

func New(bl bricklink.Bricklink) *Reference {
	return &Reference{
		bl: bl,
	}
}

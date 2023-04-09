package reference

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
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
	pathGetSupersets   = "/items/%s/%s/supersets"
	pathGetSubsets     = "/items/%s/%s/subsets"
	pathGetPriceGuide  = "/items/%s/%s/price/%s"
	pathGetKnownColors = "/items/%s/%s/colors"
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

func (it Item) PrimaryKey() int {
	hash := fnv.New32a()
	hash.Write([]byte(it.ID))
	return int(hash.Sum32())
}

func (it Item) Label() entity.Label {
	return entity.LabelInventoryItem
}

func (r *Reference) GetCatalogItem(options ...RequestOption) (*Item, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("id is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItem, opts.itemType, opts.itemNo), nil)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := internal.Parse(res.Body, &item); err != nil {
		return nil, err
	}
	// buf, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// var c bricklink.Content
	// if err := json.Unmarshal(buf, &c); err != nil {
	// 	return nil, err
	// }
	//
	// var item Item
	// if err := json.Unmarshal(c.Data, &item); err != nil {
	// 	return nil, err
	// }
	// if c.Meta.Code != 200 {
	// 	return nil, errors.New(c.Meta.Message)
	// }

	return &item, nil
}

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
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
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

	return &item, nil
}

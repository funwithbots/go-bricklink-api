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
	ID           string `json:"no,omitempty"`
	Name         string `json:"name,omitempty"`
	ItemType     string `json:"type"`
	CategoryID   int    `json:"category_id,omitempty"`
	alternateNo  string `json:"alternate_no,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// Weight and Dimensions are specified to 2 decimal places.
	Weight string `json:"weight,omitempty"` // grams
	DimX   string `json:"dim_x,omitempty"`
	DimY   string `json:"dim_y,omitempty"`
	DimZ   string `json:"dim_z,omitempty"`

	YearReleased int    `json:"year_released,omitempty"`
	Description  string `json:"description,omitempty"`
	IsObsolete   bool   `json:"is_obsolete,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

func (it Item) PrimaryKey() int {
	hash := fnv.New32a()
	hash.Write([]byte(it.ID))
	return int(hash.Sum32())
}

func (it Item) Label() entity.Label {
	return entity.LabelInventoryItem
}

func (r *Reference) GetItem(options ...RequestOption) (*Item, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("item no is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItem, opts.itemType, opts.itemNo), nil, nil)
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

func (r *Reference) GetItemImage(options ...RequestOption) (*Item, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("item no is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	// if opts.colorID == nil {
	// 	return nil, errors.New("color is required")
	// }
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	color := 0
	if opts.colorID != nil {
		color = *opts.colorID
	}

	req, err := r.bl.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetItemImage, opts.itemType, opts.itemNo, color),
		nil,
		nil,
	)
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

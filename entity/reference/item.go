package reference

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
	"github.com/funwithbots/go-bricklink-api/util"
)

type Item struct {
	ID           string `json:"no,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"type"`
	CategoryID   int    `json:"category_id,omitempty"`
	AlternateNo  string `json:"alternate_no,omitempty"`
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

func (it *Item) UnmarshalJSON(data []byte) error {
	type Alias Item
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(it),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// clean up data
	aux.Name = util.NormalizeString(aux.Name)
	aux.Description = util.NormalizeString(aux.Description)

	*it = Item(*aux.Alias)

	return nil
}

func (it Item) PrimaryKey() int {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(it.ID + it.Type))
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
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItem, opts.itemType, opts.itemNo), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
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
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	color := 0
	if opts.colorID != nil {
		color = *opts.colorID
	}

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetItemImage, opts.itemType, opts.itemNo, color),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := internal.Parse(res.Body, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

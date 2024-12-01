package reference

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"strconv"

	bl "github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

// Mapping maps a unique identifier to a Bricklink item/part/type combination.
type Mapping struct {
	Item      Item   `json:"item"`
	ColorID   int    `json:"color_id"`
	ElementID string `json:"element_id"`
}

func (m Mapping) PrimaryKey() string {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(m.ElementID))
	return strconv.Itoa(int(hash.Sum32()))
}

func (m Mapping) Label() entity.Label {
	return entity.LabelMapping
}

// GetElementID returns the element ID for a specific item/part/type combination.
func (r *Reference) GetElementID(options ...RequestOption) ([]Mapping, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemType != bl.ItemTypeMap[bl.ItemTypePart] {
		return nil, errors.New("item type must be part")
	}
	if opts.itemNo == "" {
		return nil, errors.New("item no is required")
	}
	query, err := opts.toQuery(queryTargetElementID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetElementID, opts.itemType, opts.itemNo), query, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var im []Mapping
	if err := internal.Parse(res.Body, &im); err != nil {
		return nil, err
	}

	return im, nil
}

// GetItemMapping returns the mapping resource for an Element ID.
func (r *Reference) GetItemMapping(elementID string) ([]Mapping, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItemMapping, elementID), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var im []Mapping
	if err := internal.Parse(res.Body, &im); err != nil {
		return nil, err
	}

	return im, nil
}

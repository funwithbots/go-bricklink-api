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

// Mapping maps a unique identifier to a Bricklink item/part/type combination.
type Mapping struct {
	Item      Item   `json:"item"`
	ColorID   int    `json:"color_id"`
	ColorName string `json:"color_name"`
	ElementID string `json:"element_id"`
}

func (m Mapping) PrimaryKey() string {
	hash := fnv.New32a()
	hash.Write([]byte(m.ElementID))
	return string(hash.Sum32())
}

func (m Mapping) Label() entity.Label {
	return entity.LabelMapping
}

// GetElementID returns the element ID for a specific item/part/type combination.
func (r *Reference) GetElementID(options ...RequestOption) (*Mapping, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("id is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	query, err := opts.toQuery(queryTargetElementID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetElementID, opts.itemType, opts.itemNo), query, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var im Mapping
	if err := internal.Parse(res.Body, &im); err != nil {
		return nil, err
	}

	return &im, nil
}

// GetItemMapping returns the mapping resource for an Element ID.
func (r *Reference) GetItemMapping(elementID string) (*Mapping, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(pathGetItemMapping, elementID), nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var im Mapping
	if err := internal.Parse(res.Body, &im); err != nil {
		return nil, err
	}

	return &im, nil
}

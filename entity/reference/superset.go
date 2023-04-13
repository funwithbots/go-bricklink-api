package reference

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
)

type Superset []SupersetItem

type SupersetItem struct {
	ColorID int `json:"color_id"`
	Entries []struct {
		Item      Item   `json:"item"`
		Quantity  int    `json:"quantity"`
		AppearsAs string `json:"appears_as"`
	} `json:"entries"`
}

// PrimaryKey isn't meaningful for this entity.
func (Superset) PrimaryKey() int {
	return 0
}

func (Superset) Label() entity.Label {
	return entity.LabelSuperset
}

// GetSupersets returns a list of supersets that contain the item.
// If a colorID is provided, the list is filtered to supersets that contain the item in that color.
func (r *Reference) GetSupersets(options ...RequestOption) (Superset, error) {
	var opts = requestOptions{}
	opts.withOpts(options)
	if opts.itemNo == "" {
		return nil, errors.New("id is required")
	}
	if opts.itemType == "" {
		return nil, errors.New("type is required")
	}
	query, err := opts.toQuery(queryTargetSupersets)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.bl.Timeout)
	defer cancel()

	req, err := r.bl.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetSuperset, opts.itemType, opts.itemNo),
		query,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.bl.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var items []SupersetItem
	if err := internal.Parse(res.Body, &items); err != nil {
		return nil, err
	}

	superset := items

	return superset, nil
}

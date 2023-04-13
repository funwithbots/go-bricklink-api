package reference

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/internal"
)

type KnownColor struct {
	ColorID  int `json:"color_id"`
	Quantity int `json:"quantity"`
}

func (r *Reference) GetKnownColors(options ...RequestOption) ([]KnownColor, error) {
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

	req, err := r.bl.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetKnownColors, opts.itemType, opts.itemNo),
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

	var kc []KnownColor
	if err := internal.Parse(res.Body, &kc); err != nil {
		return nil, err
	}

	return kc, nil
}

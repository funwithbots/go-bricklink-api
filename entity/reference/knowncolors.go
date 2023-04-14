package reference

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/internal"
	"github.com/funwithbots/go-bricklink-api/util"
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
	if opts.itemType != util.ItemTypePart.String() && opts.colorID != nil {
		return nil, errors.New("color id is only valid for parts")
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetKnownColors, opts.itemType, opts.itemNo),
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

	var kc []KnownColor
	if err := internal.Parse(res.Body, &kc); err != nil {
		return nil, err
	}

	return kc, nil
}

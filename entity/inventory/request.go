package inventory

import (
	"fmt"
	"strings"

	bl "github.com/funwithbots/go-bricklink-api"
)

type queryTarget int

const (
	queryTargetGetItems queryTarget = iota
)

type RequestOption func(opts *requestOptions)

// Slices are converted to a comma-separated string to specify multiple values to include/exclude.
// A minus "-" sign specifies a value to exclude.
type requestOptions struct {
	itemType   []string // item_type
	statuses   []string // status
	categoryID []string // category_id
	colorID    []string // color_id
}

// toQuery converts the request to a query string.
// Each field is converted to a query string parameter.
func (ro *requestOptions) toQuery(target queryTarget) (map[string]string, error) {
	params := map[string]string{}
	switch target {
	case queryTargetGetItems:
		if len(ro.itemType) > 0 {
			var list []string
			for _, v := range ro.itemType {
				list = append(list, v)
			}
			params["item_type"] = strings.Join(list, ",")
		}
		if len(ro.statuses) > 0 {
			var list []string
			for _, v := range ro.statuses {
				list = append(list, v)
			}
			params["status"] = strings.Join(list, ",")
		}
		if len(ro.categoryID) > 0 {
			var list []string
			for _, v := range ro.categoryID {
				list = append(list, v)
			}
			params["category_id"] = strings.Join(list, ",")
		}
		if len(ro.colorID) > 0 {
			var list []string
			for _, v := range ro.colorID {
				list = append(list, v)
			}
			params["color_id"] = strings.Join(list, ",")
		}
	}
	return params, nil
}

func (ro *requestOptions) withOpts(opts []RequestOption) {
	for _, opt := range opts {
		opt(ro)
	}
}

func WithIncludeItemType(itemType bl.ItemType) RequestOption {
	return func(opts *requestOptions) {
		opts.itemType = append(opts.itemType, bl.ItemTypeMap[itemType])
	}
}

func WithExcludeItemType(itemType bl.ItemType) RequestOption {
	return func(opts *requestOptions) {
		opts.itemType = append(opts.itemType, fmt.Sprintf("-%s", bl.ItemTypeMap[itemType]))
	}
}

// WithIncludeStatus adds a stockroom or availability status to the request.
func WithIncludeStatus(status Status) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, status.String())
	}
}

// WithExcludeStatus excludes a stockroom or availability status to the request.
func WithExcludeStatus(status Status) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, fmt.Sprintf("-%s", status.String()))
	}
}

func WithIncludeCategoryID(id int) RequestOption {
	return func(opts *requestOptions) {
		opts.categoryID = append(opts.categoryID, fmt.Sprintf("%d", id))
	}
}
func WithExcludeCategoryID(id int) RequestOption {
	return func(opts *requestOptions) {
		opts.categoryID = append(opts.categoryID, fmt.Sprintf("-%d", id))
	}
}

func WithIncludeColorID(id int) RequestOption {
	return func(opts *requestOptions) {
		opts.colorID = append(opts.colorID, fmt.Sprintf("%d", id))
	}
}

func WithExcludeColorID(id int) RequestOption {
	return func(opts *requestOptions) {
		opts.colorID = append(opts.colorID, fmt.Sprintf("-%d", id))
	}
}

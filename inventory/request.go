package inventory

type RequestOption func(opts requestOptions) requestOptions

// Slices are converted to a comma-separated string to specify multiple values to include/exclude.
// Add a minus "-" sign to specify a value to exclude.
type requestOptions struct {
	ItemType   []string // item_type
	Status     []string // status
	CategoryID []int    // category_id
	ColorID    []int    // color_id
}

// ToQueryString converts the request to a query string.
// Each field is converted to a query string parameter.
func (ro *requestOptions) ToQueryString() string {
	// TODO implement me
	return "not implemented"
}

func (ro *requestOptions) withOpts(opts []RequestOption) {
	for _, opt := range opts {
		*ro = opt(*ro)
	}
}

func WithItemType(itemType string) RequestOption {
	return func(opts requestOptions) requestOptions {
		opts.ItemType = append(opts.ItemType, itemType)
		return opts
	}
}

func WithStatus(status string) RequestOption {
	return func(opts requestOptions) requestOptions {
		opts.Status = append(opts.Status, status)
		return opts
	}
}

func WithCategoryID(categoryID int) RequestOption {
	return func(opts requestOptions) requestOptions {
		opts.CategoryID = append(opts.CategoryID, categoryID)
		return opts
	}
}

func WithColorID(colorID int) RequestOption {
	return func(opts requestOptions) requestOptions {
		opts.ColorID = append(opts.ColorID, colorID)
		return opts
	}
}

func NewRequest(opts ...RequestOption) *requestOptions {
	ro := &requestOptions{}
	ro.withOpts(opts)
	return ro
}

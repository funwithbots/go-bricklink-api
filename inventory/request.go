package inventory

type RequestOption func(opts *requestOptions)

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
		opt(ro)
	}
}

func WithItemType(itemType string) RequestOption {
	return func(opts *requestOptions) {
		opts.ItemType = append(opts.ItemType, itemType)
	}
}

func WithStatus(status string) RequestOption {
	return func(opts *requestOptions) {
		opts.Status = append(opts.Status, status)
	}
}

func WithCategoryID(categoryID int) RequestOption {
	return func(opts *requestOptions) {
		opts.CategoryID = append(opts.CategoryID, categoryID)
	}
}

func WithColorID(colorID int) RequestOption {
	return func(opts *requestOptions) {
		opts.ColorID = append(opts.ColorID, colorID)
	}
}

func NewRequest(opts ...RequestOption) *requestOptions {
	ro := &requestOptions{}
	ro.withOpts(opts)
	return ro
}

package order

import "fmt"

type OrderDirection int

const (
	OrderDirectionIn = iota
	OrderDirectionOut
)

func (od OrderDirection) String() string {
	if od == OrderDirectionOut {
		return "out"
	}
	return "in"
}

type RequestStatusField int

const (
	_ = iota
	RequestStatusFieldOrder
	RequestStatusFieldPayment
)

func (rsf RequestStatusField) String() string {
	switch rsf {
	case RequestStatusFieldOrder:
		return "status"
	case RequestStatusFieldPayment:
		return "payment_status"
	default:
		return ""
	}
}

type RequestOption func(opts *requestOptions)

type requestOptions struct {
	direction OrderDirection

	// Status is used for Update only.
	status Status

	// Statuses are for requests that accept multiple exclude and include statuses.
	statuses []string

	filed bool
}

// ToQueryString converts the request to a query string.
// Each field is converted to a query string parameter.
func (ro *requestOptions) ToQueryString() string {
	// TODO implement me
	return "not implemented"
}

func (ro *requestOptions) withOpts(opts []RequestOption) {
	// set defaults
	ro.direction = OrderDirectionIn

	for _, opt := range opts {
		opt(ro)
	}
}

func WithDirection(dir OrderDirection) RequestOption {
	return func(opts *requestOptions) {
		opts.direction = dir
	}
}

func WithIncludeStatus(status Status) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, status.String())
	}
}

func WithExcludeStatus(status Status) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, fmt.Sprintf("-%s", status.String()))
	}
}

func WithFiled() RequestOption {
	return func(opts *requestOptions) {
		opts.filed = true
	}
}

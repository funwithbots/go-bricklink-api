package orders

import (
	"fmt"
	"strings"
)

type queryTarget int

const (
	queryTargetGetOrders = iota
	queryTargetDriveThru
	queryTargetGetFeedbackList
)

// Direction defines the directionality of the resource.
// It is used for feedback and orders.
// out = from this user
// in = to this user
type Direction string

const (
	DirectionIn  = "in"
	DirectionOut = "out"
)

func (dir *Direction) String() string {
	if dir == nil || *dir != "out" {
		return "in"
	}

	return "out"
}

type RequestOption func(opts *requestOptions)

type requestOptions struct {
	direction *Direction

	// Statuses are for requests that accept multiple exclude and include statuses.
	statuses []string

	filed *bool

	mailMe *bool

	// forceRefresh is used to force a forceRefresh of the data from bricklink.
	// Otherwise, the data is retrieved from the values cached in the orders struct.
	forceRefresh bool
}

// toQuery converts the request to a query string.
// Each field is converted to a query string parameter.
func (ro *requestOptions) toQuery(target queryTarget) (map[string]string, error) {
	params := map[string]string{}
	switch target {
	case queryTargetGetFeedbackList:
		if ro.direction != nil {
			params["direction"] = ro.direction.String()
		}
	case queryTargetGetOrders:
		if ro.direction != nil {
			params["direction"] = ro.direction.String()
		}
		if len(ro.statuses) > 0 {
			var list []string
			for _, v := range ro.statuses {
				list = append(list, v)
			}
			params["status"] = strings.Join(list, ",")
		}
		if ro.filed != nil {
			params["filed"] = fmt.Sprintf("%t", *ro.filed)
		}
	case queryTargetDriveThru:
		if ro.mailMe != nil {
			params["mail_me"] = fmt.Sprintf("%t", *ro.mailMe)
		}
	}

	return params, nil
}

func (ro *requestOptions) withOpts(opts []RequestOption) {
	// set defaults
	for _, opt := range opts {
		opt(ro)
	}
}

func WithDirection(dir Direction) RequestOption {
	return func(opts *requestOptions) {
		opts.direction = &dir
	}
}

func WithIncludeStatus(status OrderStatus) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, status.String())
	}
}

func WithExcludeStatus(status OrderStatus) RequestOption {
	return func(opts *requestOptions) {
		opts.statuses = append(opts.statuses, fmt.Sprintf("-%s", status.String()))
	}
}

func WithFiled(b bool) RequestOption {
	return func(opts *requestOptions) {
		opts.filed = &b
	}
}

func WithRefresh(b bool) RequestOption {
	return func(opts *requestOptions) {
		opts.forceRefresh = b
	}
}

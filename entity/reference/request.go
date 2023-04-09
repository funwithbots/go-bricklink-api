package reference

import (
	"golang.org/x/exp/slices"

	"github.com/funwithbots/go-bricklink-api/util"
)

type RequestOption func(opts *requestOptions)

type requestOptions struct {
	itemType      string
	itemNo        string
	colorID       *int
	box           *bool
	instruction   *bool
	breakMinifigs *bool
	breakSubsets  *bool
	guideType     string
	condition     string // new_or_used.
	countryCode   string
	region        string
	currencyCode  string
	vat           string // Y N O (Yes, No, nOrway)
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

func WithItemType(typ util.ItemType) RequestOption {
	return func(opts *requestOptions) {
		opts.itemType = typ.String()
	}
}

// WithItemNo sets the item number filter.
// It will be ignored if the calling function is a member function of the item type.
func WithItemNo(itemNo string) RequestOption {
	return func(opts *requestOptions) {
		opts.itemNo = itemNo
	}
}

// WithColorID sets the color ID filter.
func WithColorID(colorID int) RequestOption {
	return func(opts *requestOptions) {
		opts.colorID = &colorID
	}
}

// WithBox sets the box flag filter.
func WithBox(box bool) RequestOption {
	return func(opts *requestOptions) {
		opts.box = &box
	}
}

// WithInstruction sets the instruction flag filter.
func WithInstruction(instruction bool) RequestOption {
	return func(opts *requestOptions) {
		opts.instruction = &instruction
	}
}

// WithBreakMinifigs sets the break minifigs flag filter.
func WithBreakMinifigs(breakMinifigs bool) RequestOption {
	return func(opts *requestOptions) {
		opts.breakMinifigs = &breakMinifigs
	}
}

// WithBreakSubsets sets the break subsets flag filter.
func WithBreakSubsets(breakSubsets bool) RequestOption {
	return func(opts *requestOptions) {
		opts.breakSubsets = &breakSubsets
	}
}

// WithGuideType sets the guide type (sold or stock) filter.
func WithGuideType(guideType string) RequestOption {
	if guideTypeSold != guideType && guideTypeStock != guideType {
		return nil
	}
	return func(opts *requestOptions) {
		opts.guideType = guideType
	}
}

// WithCondition sets the condition (new or used) filter.
func WithCondition(condition string) RequestOption {
	if util.New != condition && util.Used != condition {
		return nil
	}
	return func(opts *requestOptions) {
		opts.condition = condition
	}
}

// WithCountryCode sets the country code filter.
func WithCountryCode(countryCode string) RequestOption {
	return func(opts *requestOptions) {
		opts.countryCode = countryCode
	}
}

// WithRegion sets the region filter.
func WithRegion(region string) RequestOption {
	return func(opts *requestOptions) {
		opts.region = region
	}
}

// WithCurrencyCode sets the currency code filter.
func WithCurrencyCode(currencyCode string) RequestOption {
	return func(opts *requestOptions) {
		opts.currencyCode = currencyCode
	}
}

// WithVAT sets the VAT flag filter.
func WithVAT(vat string) RequestOption {
	if !slices.Contains([]string{"Y", "N", "O"}, vat) {
		return nil
	}
	return func(opts *requestOptions) {
		opts.vat = vat
	}
}

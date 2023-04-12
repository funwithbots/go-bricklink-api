package reference

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/funwithbots/go-bricklink-api/util"
)

type queryTarget int

const (
	queryTargetSupersets queryTarget = iota
	queryTargetSubsets
	queryTargetPriceGuide
	queryTargetElementID
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
	guideType     *string
	condition     *string // new_or_used.
	countryCode   *string
	region        *string
	currencyCode  *string
	vat           *string // Y N O (Yes, No, nOrway)
}

// toQuery converts the request to a query string.
// Each field is converted to a query string parameter.
func (ro *requestOptions) toQuery(target queryTarget) (string, error) {
	var params []string
	switch target {
	case queryTargetSupersets, queryTargetElementID:
		if ro.colorID != nil {
			params = append(params, fmt.Sprintf("color_id=%d", *ro.colorID))
		}
	case queryTargetSubsets:
		if ro.colorID != nil {
			if ro.itemType != util.ItemTypePart.String() {
				return "", fmt.Errorf("color_id is only valid for parts")
			}
			params = append(params, fmt.Sprintf("color_id=%d", *ro.colorID))
		}
		if ro.box != nil {
			params = append(params, fmt.Sprintf("box=%s", util.YesOrNo(*ro.box)))
		}
		if ro.instruction != nil {
			params = append(params, fmt.Sprintf("instruction=%s", util.YesOrNo(*ro.instruction)))
		}
		if ro.breakMinifigs != nil {
			params = append(params, fmt.Sprintf("break_minifigs=%s", util.YesOrNo(*ro.breakMinifigs)))
		}
		if ro.breakSubsets != nil {
			params = append(params, fmt.Sprintf("break_subsets=%s", util.YesOrNo(*ro.breakSubsets)))
		}
	case queryTargetPriceGuide:
		if ro.colorID != nil {
			if ro.itemType != util.ItemTypePart.String() {
				return "", fmt.Errorf("color_id is only valid for parts")
			}
			params = append(params, fmt.Sprintf("color_id=%d", *ro.colorID))
		}
		if (ro.countryCode != nil) != (ro.region != nil) {
			return "", fmt.Errorf("country_code and region must be set together")
		}

		if ro.guideType != nil {
			params = append(params, fmt.Sprintf("guide_type=%s", *ro.guideType))
		}
		if ro.condition != nil {
			params = append(params, fmt.Sprintf("instruction=%s", *ro.condition))
		}
		if ro.countryCode != nil {
			params = append(params, fmt.Sprintf("country_code=%s", *ro.countryCode))
			params = append(params, fmt.Sprintf("region=%s", *ro.region))
		}
		if ro.currencyCode != nil {
			params = append(params, fmt.Sprintf("currency_code=%s", *ro.currencyCode))
		}
		if ro.vat != nil {
			params = append(params, fmt.Sprintf("vat=%s", *ro.vat))
		}
		if len(params) > 0 {
			return strings.Join(params, "&"), nil
		}
	}

	if len(params) > 0 {
		return strings.Join(params, "&"), nil
	}
	return "", nil
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
		opts.guideType = &guideType
	}
}

// WithCondition sets the condition (new or used) filter.
func WithCondition(condition string) RequestOption {
	if util.New != condition && util.Used != condition {
		return nil
	}
	return func(opts *requestOptions) {
		opts.condition = &condition
	}
}

// WithCountryCode sets the country code filter.
func WithCountryCode(countryCode string) RequestOption {
	return func(opts *requestOptions) {
		opts.countryCode = &countryCode
	}
}

// WithRegion sets the region filter.
func WithRegion(region string) RequestOption {
	return func(opts *requestOptions) {
		opts.region = &region
	}
}

// WithCurrencyCode sets the currency code filter.
func WithCurrencyCode(currencyCode string) RequestOption {
	return func(opts *requestOptions) {
		opts.currencyCode = &currencyCode
	}
}

// WithVAT sets the VAT flag filter.
func WithVAT(vat string) RequestOption {
	if !slices.Contains([]string{"Y", "N", "O"}, vat) {
		return nil
	}
	return func(opts *requestOptions) {
		opts.vat = &vat
	}
}

package reference

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"

	bl "github.com/funwithbots/go-bricklink-api"
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
func (ro *requestOptions) toQuery(target queryTarget) (map[string]string, error) {
	params := map[string]string{}
	switch target {
	case queryTargetSupersets, queryTargetElementID:
		if ro.colorID != nil {
			params["color_id"] = strconv.Itoa(*ro.colorID)
		}
	case queryTargetSubsets:
		if ro.itemType == "" {
			return nil, fmt.Errorf("item_type is required")
		}
		if ro.colorID != nil {
			if ro.itemType != bl.ItemTypeMap[bl.ItemTypePart] {
				return nil, fmt.Errorf("color_id is only valid for parts")
			}
			params["color_id"] = strconv.Itoa(*ro.colorID)
		}
		if ro.itemType == bl.ItemTypeMap[bl.ItemTypeSet] {
			if ro.box != nil {
				params["box"] = fmt.Sprintf("%t", *ro.box)
			}
			if ro.instruction != nil {
				params["instruction"] = fmt.Sprintf("%t", *ro.instruction)
			}
			if ro.breakMinifigs != nil {
				params["break_minifigs"] = fmt.Sprintf("%t", *ro.breakMinifigs)
			}
			if ro.breakSubsets != nil {
				params["break_subsets"] = fmt.Sprintf("%t", *ro.breakSubsets)
			}
		}
	case queryTargetPriceGuide:
		if (ro.countryCode != nil) != (ro.region != nil) {
			return nil, fmt.Errorf("country_code and region must be set together")
		}
		if ro.colorID != nil {
			if ro.itemType != bl.ItemTypeMap[bl.ItemTypePart] {
				return nil, fmt.Errorf("color_id is only valid for parts")
			}
			params["color_id"] = strconv.Itoa(*ro.colorID)
		}

		if ro.guideType != nil {
			params["guide_type"] = fmt.Sprintf("%s", *ro.guideType)
		}
		if ro.condition != nil {
			params["condition"] = fmt.Sprintf("%s", *ro.condition)
		}
		if ro.countryCode != nil {
			params["country_code"] = fmt.Sprintf("%s", *ro.countryCode)
			params["region"] = fmt.Sprintf("%s", *ro.region)
		}
		if ro.currencyCode != nil {
			params["currency_code"] = fmt.Sprintf("%s", *ro.currencyCode)
		}
		if ro.vat != nil {
			params["vat"] = fmt.Sprintf("%s", *ro.vat)
		}
	}

	return params, nil
}

func (ro *requestOptions) withOpts(opts []RequestOption) {
	for _, opt := range opts {
		opt(ro)
	}
}

func WithItemType(typ bl.ItemType) RequestOption {
	return func(opts *requestOptions) {
		if len(typ) > 1 {
			typ = typ[0:1]
		}
		opts.itemType = bl.ItemTypeMap[typ]
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
	if GuideTypeSold != guideType && GuideTypeStock != guideType {
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

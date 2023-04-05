package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/reference"

// ReferenceAPI provides an interface for interacting with the Catalog API.
type ReferenceAPI interface {
	// catalog
	GetSupersets(...reference.RequestOption) ([]reference.Superset, error)
	GetSubsets(...reference.RequestOption) ([]reference.Subset, error)
	GetKnownColors() ([]reference.KnownColor, error)
	GetPriceGuide(item reference.Item, colorID int) (*reference.PriceGuide, error)

	// color
	GetColors() ([]reference.Color, error)
	GetColor(colorID int) (reference.Color, error)

	// mapping
	GetElementID(item reference.Item, colorID int) (string, error)
	GetMapping(elementID string) (*reference.Mapping, error)
}

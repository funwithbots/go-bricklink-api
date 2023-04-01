package go_bricklink_api

import "github.com/funwithbots/go-bricklink-api/catalog"

// CatalogAPI provides an interface for interacting with the Catalog API.
type CatalogAPI interface {
	GetSupersets(...catalog.RequestOption) ([]catalog.Superset, error)
	GetSubsets(...catalog.RequestOption) ([]catalog.Subset, error)
	GetPriceGuide(...catalog.RequestOption) (*catalog.PriceGuide, error)
	GetKnownColors() ([]catalog.KnownColor, error)
}

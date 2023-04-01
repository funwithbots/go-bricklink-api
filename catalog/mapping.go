package catalog

import "github.com/funwithbots/go-bricklink-api/util"

// Mapping maps a unique identifier to a Bricklink item/part/type combination.
type Mapping struct {
	Item      Item   `json:"item"`
	ColorID   int    `json:"color_id"`
	ColorName string `json:"color_name"`
	ElementID string `json:"element_id"`
}

// GetElementID returns the element ID for a specific item/part/type combination.
func GetElementID(item Item, colorID int) (string, error) {
	// TODO implement me
	return "", util.ErrNotImplemented
}

// GetMapping returns the mapping resource for an Element ID.
func GetMapping(elementID string) (*Mapping, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

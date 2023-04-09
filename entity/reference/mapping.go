package reference

import (
	"hash/fnv"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

// Mapping maps a unique identifier to a Bricklink item/part/type combination.
type Mapping struct {
	Item      Item   `json:"item"`
	ColorID   int    `json:"color_id"`
	ColorName string `json:"color_name"`
	ElementID string `json:"element_id"`
}

func (m Mapping) PrimaryKey() string {
	hash := fnv.New32a()
	hash.Write([]byte(m.ElementID))
	return string(hash.Sum32())
}

func (m Mapping) Label() entity.Label {
	return entity.LabelMapping
}

// GetElementID returns the element ID for a specific item/part/type combination.
func GetElementID(item Item, colorID int) (string, error) {
	// TODO implement me
	return "", util.ErrNotImplemented
}

// GetMapping returns the mapping resource for an Element ID.
func GetMapping(elementID string) (Mapping, error) {
	// TODO implement me
	return Mapping{}, util.ErrNotImplemented
}

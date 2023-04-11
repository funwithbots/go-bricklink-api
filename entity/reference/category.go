package reference

import (
	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/util"
)

// Category represents a category.
type Category struct {
	ID           int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentId     int    `json:"parent_id"`
}

func (c *Category) PrimaryKey() int {
	return c.ID
}

func (c *Category) Label() entity.Label {
	return entity.LabelCategory
}

// GetCategories returns a list of categories.
func (r *Reference) GetCategories() ([]Category, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// GetCategory returns the details of a specific category.
func (r *Reference) GetCategory(id int) (Category, error) {
	// TODO implement me
	return Category{}, util.ErrNotImplemented
}

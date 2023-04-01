package category

import "github.com/funwithbots/go-bricklink-api/util"

// Category represents a category.
type Category struct {
	ID           int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentId     int    `json:"parent_id"`
}

// GetCategories returns a list of categories.
func GetCategories() ([]Category, error) {
	// TODO implement me
	return nil, util.ErrNotImplemented
}

// GetCategory returns the details of a specific category.
func GetCategory(id int) (Category, error) {
	// TODO implement me
	return Category{}, util.ErrNotImplemented
}

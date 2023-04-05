package reference

import (
	"context"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/util"
)

// Category represents a category.
type Category struct {
	ID           int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentId     int    `json:"parent_id"`
}

func (c *Category) Get(ctx context.Context, req *http.Request) (*http.Response, error) {
	return nil, util.ErrNotImplemented
}

func (c *Category) Insert(ctx context.Context, req *http.Request) (*http.Response, error) {
	return nil, util.ErrNotImplemented
}

func (c *Category) Update(ctx context.Context, req *http.Request) (*http.Response, error) {
	return nil, util.ErrNotImplemented
}

func (c *Category) Delete(ctx context.Context, req *http.Request) (*http.Response, error) {
	return nil, util.ErrNotImplemented
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

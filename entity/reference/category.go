package reference

import (
	"context"
	"fmt"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/entity"
	"github.com/funwithbots/go-bricklink-api/internal"
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
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		pathGetCategories,
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var categories []Category
	if err := internal.Parse(res.Body, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategory returns the details of a specific category.
func (r *Reference) GetCategory(id int) (*Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := r.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(pathGetCategory, id),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var category Category
	if err := internal.Parse(res.Body, &category); err != nil {
		return nil, err
	}

	return &category, nil
}

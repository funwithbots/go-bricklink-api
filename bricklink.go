package go_bricklink_api

type Type int

const (
	Unknown Type = iota
	Catalog
	Category
	Color
	Feedback
	Inventory
	Mapping
	Member
	Message
	Order
	OrderItem
	PriceGuide
	Problem
	Subset
	Superset
)

type bricklink struct {
	client Client
}

type Content interface {
	PrimaryKey() int
	Label() Type
}

type Options interface {
}

type Bricklink interface {
	Reference() ReferenceAPI
	Inventory() InventoryAPI
	Order() OrderAPI
}

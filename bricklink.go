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

type API interface {
	Get(opts Options) (Content, error)
	GetList(opts Options) ([]Content, error)
	Insert(content Content) (Content, error)
	InsertList(contents []Content) error
	Update(content Content) (Content, error)
	Delete(id int) error
}

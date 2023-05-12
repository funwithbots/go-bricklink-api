package entity

type Label int

const (
	Unknown Label = iota
	LabelCatalog
	LabelCategory
	LabelColor
	LabelFeedback
	LabelInventory
	LabelInventoryItem
	LabelMapping
	LabelMember
	LabelMessage
	LabelNote
	LabelOrder
	LabelOrderHeader
	LabelOrderItem
	LabelPriceGuide
	LabelProblem
	LabelShippingMethod
	LabelSubset
	LabelSuperset
)

func (l Label) String() string {
	switch l {
	case LabelCatalog:
		return "Catalog"
	case LabelCategory:
		return "Category"
	case LabelColor:
		return "Color"
	case LabelFeedback:
		return "Feedback"
	case LabelInventory:
		return "Inventory"
	case LabelInventoryItem:
		return "InventoryItem"
	case LabelMapping:
		return "Mapping"
	case LabelMember:
		return "Member"
	case LabelMessage:
		return "Message"
	case LabelNote:
		return "Note"
	case LabelOrder:
		return "Order"
	case LabelOrderHeader:
		return "OrderHeader"
	case LabelOrderItem:
		return "OrderItem"
	case LabelPriceGuide:
		return "PriceGuide"
	case LabelProblem:
		return "Problem"
	case LabelShippingMethod:
		return "ShippingMethod"
	case LabelSubset:
		return "Subset"
	case LabelSuperset:
		return "Superset"
	default:
		return ""
	}
}

package util

type ItemType int

const (
	ItemTypePart ItemType = iota
	ItemTypeMinifig
	ItemTypeGear
	ItemTypeBook
	ItemTypeCatalog
	ItemTypeInstruction
	ItemTypeSet
	ItemTypeOriginalBox
	ItemTypeUnsortedLot
)

func (it ItemType) String() string {
	switch it {
	case ItemTypePart:
		return "part"
	case ItemTypeMinifig:
		return "minifig"
	case ItemTypeGear:
		return "gear"
	case ItemTypeBook:
		return "book"
	case ItemTypeCatalog:
		return "catalog"
	case ItemTypeInstruction:
		return "instruction"
	case ItemTypeSet:
		return "set"
	case ItemTypeOriginalBox:
		return "original_box"
	case ItemTypeUnsortedLot:
		return "unsorted_lot"
	default:
		return ""
	}
}

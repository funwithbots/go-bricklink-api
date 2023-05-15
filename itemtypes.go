package go_bricklink_api

type ItemType int8

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

func (it ItemType) Label() string {
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

func (it ItemType) String() string {
	switch it {
	case ItemTypePart:
		return "p"
	case ItemTypeMinifig:
		return "m"
	case ItemTypeGear:
		return "g"
	case ItemTypeBook:
		return "b"
	case ItemTypeCatalog:
		return "c"
	case ItemTypeInstruction:
		return "i"
	case ItemTypeSet:
		return "s"
	case ItemTypeOriginalBox:
		return "o"
	case ItemTypeUnsortedLot:
		return "l"
	default:
		return ""
	}
}

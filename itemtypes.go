package go_bricklink_api

type ItemType string

const (
	ItemTypePart        ItemType = "P"
	ItemTypeMinifig     ItemType = "M"
	ItemTypeGear        ItemType = "G"
	ItemTypeBook        ItemType = "B"
	ItemTypeCatalog     ItemType = "C"
	ItemTypeInstruction ItemType = "I"
	ItemTypeSet         ItemType = "S"
	ItemTypeOriginalBox ItemType = "O"
	ItemTypeUnsortedLot ItemType = "L"
)

var ItemTypeMap = map[ItemType]string{
	ItemTypePart:        "part",
	ItemTypeMinifig:     "minifig",
	ItemTypeBook:        "book",
	ItemTypeCatalog:     "catalog",
	ItemTypeInstruction: "instruction",
	ItemTypeSet:         "set",
	ItemTypeOriginalBox: "original_box",
	ItemTypeGear:        "gear",
	ItemTypeUnsortedLot: "unsorted_lot",
}

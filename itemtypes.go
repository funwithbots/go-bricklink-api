package go_bricklink_api

type ItemType string

const (
	ItemTypeMinifig ItemType = "M"
	ItemTypePart    ItemType = "P"
	ItemTypeSet     ItemType = "S"

	ItemTypeBook        ItemType = "B"
	ItemTypeCatalog     ItemType = "C"
	ItemTypeGear        ItemType = "G"
	ItemTypeInstruction ItemType = "I"
	ItemTypeOriginalBox ItemType = "O"
	ItemTypeUnsortedLot ItemType = "L"
)

var ItemTypeMap = map[ItemType]string{
	ItemTypeMinifig: "minifig",
	ItemTypePart:    "part",
	ItemTypeSet:     "set",

	ItemTypeBook:        "book",
	ItemTypeCatalog:     "catalog",
	ItemTypeGear:        "gear",
	ItemTypeInstruction: "instruction",
	ItemTypeOriginalBox: "original_box",
	ItemTypeUnsortedLot: "unsorted_lot",
}

var ItemCodeMap = map[string]ItemType{
	"MINIFIG": ItemTypeMinifig,
	"PART":    ItemTypePart,
	"SET":     ItemTypeSet,

	"BOOK":         ItemTypeBook,
	"CATALOG":      ItemTypeCatalog,
	"GEAR":         ItemTypeGear,
	"INSTRUCTION":  ItemTypeInstruction,
	"ORIGINAL_BOX": ItemTypeOriginalBox,
	"UNSORTED_LOT": ItemTypeUnsortedLot,
}

package go_bricklink_api

import "strings"

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

var ItemCodeMap = make(map[string]ItemType)

func init() {
	for k, v := range ItemTypeMap {
		ItemCodeMap[v] = ItemType(strings.ToUpper(string(k)))
	}
}

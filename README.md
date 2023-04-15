# go-bricklink-api
## Golang Bricklink API Wrapper

This package implements Bricklink API v3.0. 

It is a work in progress and is not yet complete.

It comprises three packages that represent broad areas of the API.

 - Reference API: The Bricklink catalog of items, colors, and categories. These are not editable resources.
   - Catalog Items
   - Colors
   - Categories
   - Item Types
   - Item / color / type mappings
   - Price Guide
   - Supersets
   - Subsets
 - Order API: The Bricklink hdr system.
   - Orders
   - Order Items
   - Feedback
   - Members
   - Messages
   - Problems
 - Inventory API: The Bricklink store inventory system.
   - Inventories
   - Inventory Items

The API packages share an HTTP client which is an implementation to allow mocking.

Each API Resource is available as a separate package with one or more structs and methods to access the API. Each Resource Representation in the API has its own struct. API methods are members of the package.

Coupons and Setting APIs are not implemented.

Simple Example
```go
package main

import (
    "fmt"
    bricklink "github.com/bricklink/go-bricklink-api"
    "github.com/bricklink/go-bricklink-api/reference"
)	

func main() {
    bl, err := bricklink.New(bricklink.WithEnv())
    if err != nil {
        panic(err)
    }
    ref := reference.New(*bl)
   
    item, err := ref.GetCatalogItem(
        reference.WithItemNo("3001"), 
        reference.WithItemType(reference.ItemTypePart), 
    )
    fmt.Printf("%s %s is %s\n", item.ItemType, item.ItemNo, item.Name)

   inv := inventory.New(*bl)
}
```

## References

https://www.bricklink.com/v3/api.page?page=references


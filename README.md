# go-bricklink-api
## Golang Bricklink API Wrapper

This package implements a wrapper for the Bricklink API v3.0. 

It is a work in progress and is not yet complete.

There are three packages that represent broad areas of the API.

 - Reference API: The Bricklink catalog of items, colors, and categories. These are not editable resources.
   - Catalog Items
   - Colors
   - Categories
   - Item Types
   - Item / color / type mappings
   - Price Guide
   - Supersets
   - Subsets
 - Order API: The Bricklink order system.
   - Orders
   - Order Items
   - Feedback
   - Members
   - Messages
   - Problems
 - Inventory API: The Bricklink store inventory system.
   - Inventories
   - Inventory Items

The API packages share an implementation of the HTTP client to allow mocking.

Each API Resource is available as a separate package with one or more structs and methods to access the API. Each Resource Representation in the API has its own struct. API methods are members of the package.

Coupons API is not yet implemented.

Simple Example

```go
package main

import (
   "fmt"
   bricklink "github.com/bricklink/go-bricklink-api"
   "github.com/funwithbots/go-bricklink-api/inventory"
   "github.com/funwithbots/go-bricklink-api/orders"
   "github.com/funwithbots/go-bricklink-api/reference"
)

func main() { 
   // Create a new Bricklink API client
   b, err := bricklink.New(bricklink.WithEnv())
   if err != nil {
      panic(err)
   }
   ref := reference.New(*b)

   // Interact with the Reference API
   item, err := ref.GetCatalogItem(
      reference.WithItemNo("3001"),
      reference.WithItemType(reference.ItemTypePart),
   )
   if err != nil {
      panic(err)
   }
   fmt.Printf("%s %s is %s\n", item.Type, item.ItemNo, item.Name)

   // Interact with the Inventory API
   inv := inventory.New(*b)
   items, err := inv.GetItems(
      inventory.WithInventoryType(inventory.InventoryTypePart),
      inventory.WithColorID(11),
      inventory.WithItemNo("3001"),
   )
   if err != nil {
      panic(err)
   }
   fmt.Printf("There are %d matching items in inventory\n", len(items))

   // Interact with the Orders API
   ord := orders.New(*b)
   o, err := ord.GetOrders(
      orders.WithOrderStatus(orders.OrderStatusShipped),
      orders.WithOrderStatus(orders.OrderStatusPaid), 
      orders.WithOrderFiled(false),
   )
   if err != nil {
      panic(err)
   }
   fmt.Printf("There are %d matching orders\n", len(o))
}
```

## References

https://www.bricklink.com/v3/api.page?page=references


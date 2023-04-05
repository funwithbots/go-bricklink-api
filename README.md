# go-bricklink-api
## Golang Bricklink API Wrapper

This package implements Bricklink API v3.0. 

It is a work in progress and is not yet complete.

The API wrapper comprises three packages that represent broad areas of the API.

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

The API packages share a client which is an implementation to allow mocking.

Each API Resource is implemented as a separate package with a struct and methods to access the API.
API Resources are provided as implementations of the main package.

Coupons and Setting APIs are not implemented.

## References

https://www.bricklink.com/v3/api.page?page=references


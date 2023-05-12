package go_bricklink_api_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	bricklink "github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/entity/inventory"
	"github.com/funwithbots/go-bricklink-api/entity/orders"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/util"
)

// Set this to a pending order to test endpoints that can update the order.
var pendingOrderID = 22132848

// Set this to a shipped order to test setting feedback and drive thru.
var shippedOrderID = 0

// TestReference is a set of basic tests for the Bricklink Catalog and related endpoints.
func TestReference(t *testing.T) {
	tests := []struct {
		name     string
		options  []reference.RequestOption
		itemType util.ItemType
		want     string
		colorID  int
		itemID   string
	}{
		{
			name: "part test",
			options: []reference.RequestOption{
				reference.WithItemNo("4592c03"),
				reference.WithItemType(util.ItemTypePart),
				reference.WithColorID(11),
				reference.WithBox(true),
				reference.WithBreakMinifigs(true),
				reference.WithBreakSubsets(true),
				reference.WithInstruction(true),
				reference.WithGuideType(reference.GuideTypeStock),
				reference.WithCondition(util.New),
				reference.WithCountryCode("US"),
				reference.WithCurrencyCode("USD"),
				reference.WithRegion(reference.PGRegionNorthAmerica),
				reference.WithVAT("N"),
			},
			itemType: util.ItemTypePart,
			want:     "4592c03",
			colorID:  11,
			itemID:   "4592c03",
		},
		{
			name: "set test",
			options: []reference.RequestOption{
				reference.WithItemNo("75981-3"),
				reference.WithItemType(util.ItemTypeSet),
				reference.WithBox(true),
				reference.WithBreakMinifigs(false),
				reference.WithBreakSubsets(true),
				reference.WithInstruction(true),
				reference.WithGuideType(reference.GuideTypeStock),
				reference.WithCondition(util.New),
				reference.WithCountryCode("US"),
				reference.WithCurrencyCode("USD"),
				reference.WithRegion(reference.PGRegionNorthAmerica),
				reference.WithVAT("N"),
			},
			itemType: util.ItemTypeSet,
			want:     "75981-3",
		},
		{
			name: "minifig test",
			options: []reference.RequestOption{
				reference.WithItemNo("sw0001c"),
				reference.WithItemType(util.ItemTypeMinifig),
				reference.WithBreakMinifigs(false),
				reference.WithGuideType(reference.GuideTypeStock),
				reference.WithCondition(util.New),
			},
			itemType: util.ItemTypeMinifig,
			want:     "sw0001c",
		},
	}

	opts := []bricklink.BricklinkOption{
		bricklink.WithEnv(),
	}

	// comment this block to test against the real API
	// TODO Fix test server for additional test cases
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`"meta": {"code": 200, "message": "OK"}, "data":{{"id": "3001", "item_type": "PART", "name": "Brick 2 x 4"}}`))
	// }))
	// client, err := internal.NewClient(internal.WithHTTPClient(server.Client()))
	// if err != nil {
	// 	t.Errorf("error creating client: %v", err)
	// }
	// bricklink.WithClient(client)
	// end block

	assert := assert.New(t)
	bricklink, err := bricklink.New(opts...)
	if err != nil {
		assert.FailNow(err.Error())
	}
	ref := reference.New(*bricklink)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subsets, err := ref.GetSubsets(tt.options...)
			if err != nil {
				assert.Failf("error getting subsets:", "%s", err.Error())
			}
			assert.GreaterOrEqualf(len(subsets), 2, "expected at least one subset")

			supersets, err := ref.GetSupersets(tt.options...)
			if err != nil {
				assert.FailNowf("error getting supersets:", "%s", err.Error())
			}
			assert.LessOrEqualf(1, len(supersets), "expected at least one superset")

			ii, err := ref.GetItemImage(tt.options...)
			if err != nil {
				assert.FailNowf("error getting item image:", "%s", err.Error())
			}
			assert.NotEqualf("", ii.ThumbnailURL, "expected image url. %+v", ii)

			if tt.colorID > 0 {
				tt.options = append(tt.options, reference.WithColorID(tt.colorID))
			}

			item, err := ref.GetItem(tt.options...)
			if err != nil {
				assert.Failf("error getting catalog item:", "%s", err.Error())
			}
			if item != nil {
				assert.Equal(tt.want, item.ID)
				assert.Equal(strings.ToUpper(tt.itemType.String()), item.ItemType)
			}

			var elementID string
			if tt.itemType == util.ItemTypePart {
				maps, err := ref.GetElementID(tt.options...)
				if !assert.Nil(err) {
					assert.Failf("error getting element id:", "%s", err.Error())
				} else {
					if assert.Greaterf(len(maps), 0, "expected at least one element mapping %s.", tt.itemID) {
						elementID = maps[0].ElementID
						assert.NotEqualf("", elementID, "expected element id. %+v", maps)
						t.Logf("%s element id: %s", tt.itemID, elementID)
					}
				}
				if elementID != "" {
					maps, err = ref.GetItemMapping(elementID)
					if assert.Nil(err) {
						if assert.Greaterf(len(maps), 0, "expected at least one element mapping %s.", tt.itemID) {
							em := maps[0]
							assert.Equalf(elementID, em.ElementID, "expected element id %s; got %s", elementID, em.ElementID)
							assert.Equalf(tt.colorID, em.ColorID, "expected color id %d; got %d", tt.colorID, em.ColorID)
							assert.Equalf(tt.itemID, em.Item.ID, "expected item id %s; got %s", tt.itemID, em.Item.ID)
						}
					}
				}

				kc, err := ref.GetKnownColors(tt.options...)
				if !assert.Nil(err) {
					assert.Failf("error getting known colors:", "%s", err.Error())
				} else {
					assert.GreaterOrEqualf(len(kc), 1, "expected at least one known color")
				}
			}

			pg, err := ref.GetPriceGuide(tt.options...)
			if !assert.Nil(err) {
				assert.Failf("error getting price guide:", "%s", err.Error())
			} else {
				if assert.NotNil(pg.Item) {
					assert.Equalf(tt.want, pg.Item.ID, "expected item id %s; got %s", tt.want, pg.Item.ID)
				}
				if assert.NotNil(pg.PriceDetails) {
					assert.Greaterf(len(pg.PriceDetails), 0, "expected price detail > 0")
				}
			}
		})
	}

	t.Run("get categories", func(t *testing.T) {
		l := 500
		v, err := ref.GetCategories()
		if err != nil {
			assert.Fail(err.Error())
		}
		if len(v) < l {
			assert.Failf("", "expected at least %d categories; got %d", l, len(v))
		}
	})

	t.Run("get category", func(t *testing.T) {
		id := 568
		v, err := ref.GetCategory(id)
		if err != nil {
			assert.Fail(err.Error())
		}
		if v.PrimaryKey() != id {
			assert.Failf("", "wanted %d; got %d", id, v.ID)
		}
	})

	t.Run("get colors", func(t *testing.T) {
		l := 200
		v, err := ref.GetColors()
		if err != nil {
			assert.Fail(err.Error())
		}
		if len(v) < l {
			assert.Failf("", "expected at least %d colors; got %d", l, len(v))
		}
	})

	t.Run("get color", func(t *testing.T) {
		id := 11
		want := "Black"
		v, err := ref.GetColor(id)
		if err != nil {
			assert.Fail(err.Error())
		}
		if v.ColorName != want {
			assert.Failf("", "wanted %d; got %s", want, v.ColorName)
		}
	})
}

// TestInventory is a set of basic tests for the Bricklink Inventory endpoints.
// To run this test, you must have an active Bricklink store or run the mocker server.
func TestInventory(t *testing.T) {
	tests := []struct {
		name        string
		options     []inventory.RequestOption
		invResource string
		itemType    util.ItemType
		want        string
		colorID     int
		itemID      string
		categoryID  int
		update      string
	}{
		{
			name:    "part test",
			options: []inventory.RequestOption{},
			invResource: `{
    "item": {
        "no":"sticker",
        "type":"PART"
    },
    "color_id":0,
    "quantity":2,
    "new_or_used":"U",
    "unit_price":"1.2000",
    "description":"test",
    "bulk":1,
    "is_retain":false,
    "is_stock_room":true,
    "sale_rate":0,
    "my_cost":"1.0000"
}`,
			itemType:   util.ItemTypePart,
			want:       "sticker",
			colorID:    11,
			itemID:     "sticker",
			categoryID: 160,
			update:     "test update",
		},
	}

	opts := []bricklink.BricklinkOption{
		bricklink.WithEnv(),
	}

	// comment this block to test against the real API
	// TODO Fix test server.
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`"meta": {"code": 200, "message": "OK"}, "data":{{"id": "3001", "item_type": "PART", "name": "Brick 2 x 4"}}`))
	// }))
	// client, err := internal.NewClient(internal.WithHTTPClient(server.Client()))
	// if err != nil {
	// 	t.Errorf("error creating client: %v", err)
	// }
	// bricklink.WithClient(client)
	// end block

	assert := assert.New(t)
	bricklink, err := bricklink.New(opts...)
	if err != nil {
		assert.FailNow(err.Error())
	}
	inv := inventory.New(*bricklink)

	for _, tt := range tests {
		// Generate a random remark to avoid deleting real inventory items.
		remark := "TEST " + util.RandomString(16, bricklink.Rand)

		t.Run(tt.name, func(t *testing.T) {
			var it inventory.Item
			err := json.Unmarshal([]byte(tt.invResource), &it)
			if err != nil {
				assert.Failf("error marshaling inventory item:", "%s", err.Error())
				t.SkipNow()
			}
			it.Remarks = remark
			qty := it.Quantity

			item, err := inv.CreateItem(it)
			if err != nil {
				assert.Failf("error creating inventory item:", "%s", err.Error())
				t.SkipNow()
			}
			id := item.PrimaryKey()
			if !assert.NotZerof(id, "expected non-zero item id; got %v", id) {
				t.SkipNow()
			}

			item.Description = tt.update
			item, err = inv.UpdateItem(*item)
			if err != nil {
				assert.Failf("error updating inventory item", "%d: %s", id, err.Error())
				t.SkipNow()
			}
			if !assert.Equalf(tt.update, item.Description, "expected item description %s; got %s", tt.update, item.Description) {
				t.SkipNow()
			}

			item, err = inv.GetItem(id)
			if err != nil {
				assert.Failf("error getting inventory item ", "%d: %s", id, err.Error())
				t.SkipNow()
			}
			if !assert.Equalf(id, item.PrimaryKey(), "expected item id %d; got %d", id, item.PrimaryKey()) {
				t.SkipNow()
			}
			if !assert.Equalf(tt.update, item.Description, "expected item description %s; got %s", tt.update, item.Description) {
				t.SkipNow()
			}
			if !assert.Equalf(2*qty, item.Quantity, "expected item quantity %d; got %d", 2*qty, item.Quantity) {
				t.SkipNow()
			}

			items := make([]inventory.Item, 2)
			items[0] = it
			items[0].StockRoomID = inventory.StatusStockRoomB.String()
			items[1] = it
			items[1].StockRoomID = inventory.StatusStockRoomC.String()
			if err = inv.CreateItems(items); err != nil {
				assert.Failf("error creating multiple inventory items:", "%s", err.Error())
				t.SkipNow()
			}

			options := []inventory.RequestOption{
				inventory.WithIncludeItemType(tt.itemType),
				inventory.WithIncludeStatus(inventory.StatusStockRoomA),
				inventory.WithIncludeStatus(inventory.StatusStockRoomB),
				inventory.WithIncludeStatus(inventory.StatusStockRoomC),
				inventory.WithIncludeCategoryID(tt.categoryID),
				inventory.WithExcludeColorID(tt.colorID),
				inventory.WithExcludeCategoryID(2), // exclude baseplates
			}
			items, err = inv.GetItems(options...)
			if err != nil {
				assert.Failf("error searching inventory items:", "%s", err.Error())
				t.SkipNow()
			}
			if !assert.GreaterOrEqualf(len(items), 3, "expected at least 3 items; got %d", len(items)) {
				// t.SkipNow()
			}

			for _, v := range items {
				if v.Remarks != remark {
					continue
				}

				id := v.PrimaryKey()
				if err := inv.DeleteItem(id); err != nil {
					assert.Failf("error deleting inventory item:", "%d: %s", id, err.Error())
				}
				if _, err := inv.GetItem(id); err == nil {
					assert.Failf("expected error getting deleted inventory item:", "%d", id)
				} else {
					assert.Equalf("RESOURCE_NOT_FOUND", err.Error(), "expected error getting deleted inventory item %d", id)
				}
			}
		})
	}
}

// TestOrders is a set of basic tests for the Bricklink Order System endpoints.
// To run this test, you must have an active Bricklink store or run the mocker server.
func TestOrders(t *testing.T) {
	tests := []struct {
		name    string
		options []orders.RequestOption
		member  string
		want    string
	}{
		{
			name:    "orders test",
			member:  "deaddrop",
			options: []orders.RequestOption{},
		},
	}

	opts := []bricklink.BricklinkOption{
		bricklink.WithEnv(),
	}

	// comment this block to test against the real API
	// TODO Fix test server.
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`"meta": {"code": 200, "message": "OK"}, "data":{{"id": "3001", "item_type": "PART", "name": "Brick 2 x 4"}}`))
	// }))
	// client, err := internal.NewClient(internal.WithHTTPClient(server.Client()))
	// if err != nil {
	// 	t.Errorf("error creating client: %v", err)
	// }
	// bricklink.WithClient(client)
	// end block

	assert := assert.New(t)
	bricklink, err := bricklink.New(opts...)
	if err != nil {
		assert.FailNow(err.Error())
	}
	ord, err := orders.New(*bricklink)
	if err != nil {
		assert.FailNow(err.Error())
	}
	if len(ord.ShippingMethods) == 0 {
		assert.FailNow("no shipping methods found")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get filed orders
			filed, err := ord.GetOrderHeaders(orders.WithFiled(true), orders.WithExcludeStatus(orders.StatusPurged))
			if err != nil {
				assert.Failf("error retrieving filed orders:", "%s", err.Error())
				t.SkipNow()
			}
			if !assert.NotEqualf(0, len(filed), "no filed orders found") {
				t.SkipNow()
			}

			// get unfiled orders
			unfiled, err := ord.GetOrderHeaders(orders.WithExcludeStatus(orders.StatusPurged))
			if err != nil {
				assert.Failf("error retrieving unfiled orders:", "%s", err.Error())
				t.SkipNow()
			}
			if !assert.NotEqual(t, 0, len(unfiled), "no unfiled orders found") {
				t.SkipNow()
			}

			// if both filed and unfiled orders are found, make sure the lists don't match
			if !assert.NotEqualf(
				filed[0].PrimaryKey(),
				unfiled[0].PrimaryKey(),
				"expected different order ids; got %s", filed[0].PrimaryKey(),
			) {
				t.SkipNow()
			}

			// if no orders, stop. Otherwise, save an order to test against.
			original, err := ord.GetOrderHeader(pendingOrderID)
			if err != nil {
				assert.Failf("error getting order", "%d: %s", pendingOrderID, err.Error())
				t.SkipNow()
			}

			// get order
			get, err := ord.GetOrderHeader(original.PrimaryKey())
			if err != nil {
				assert.Failf("error getting order", " %d:", original.PrimaryKey(), "%s", err.Error())
				t.SkipNow()
			}
			assert.Equalf(original.PrimaryKey(), get.PrimaryKey(), "expected order %d to match original", original.PrimaryKey())

			// get order items
			items, err := ord.GetOrderItems(original.PrimaryKey())
			if err != nil {
				assert.Failf("error getting order items for order", " %d: %s", original.PrimaryKey(), err.Error())
				t.SkipNow()
			}
			assert.GreaterOrEqualf(len(items), 1, "expected at least 1 order item; got %d", len(items))

			// get order messages
			messages, err := ord.GetOrderMessages(original.PrimaryKey())
			if err != nil {
				assert.Failf("error getting order messages for order", " %d: %s", original.PrimaryKey(), err.Error())
				t.SkipNow()
			}
			t.Logf("found %d messages for order %d", len(messages), original.PrimaryKey())

			// get order feedback
			feedbacks, err := ord.GetOrderFeedback(original.PrimaryKey())
			if err != nil {
				assert.Failf("error getting order feedback for order", " %d: %s", original.PrimaryKey(), err.Error())
				t.SkipNow()
			}
			if len(feedbacks) != 0 {
				feedback := feedbacks[0]
				t.Logf("found %s feedback for order %d", feedback.Rating, original.PrimaryKey())
			}

			// Generate a random remark to avoid mucking up existing orders.
			remark := "TEST " + util.RandomString(16, bricklink.Rand)

			// get member note
			member := original.BuyerName
			note, err := ord.GetMemberNote(member)
			if err != nil {
				assert.Failf("error getting note for member", " %s: %s", member, err.Error())
				t.SkipNow()
			}
			if !assert.NotNil(note, "expected member note to be non-nil") {
				t.SkipNow()
			}

			// set member note
			oldNote := note.Note
			note.Note += remark
			newNote, err := ord.UpsertMemberNote(*note)
			if err != nil {
				assert.Failf("error setting member note for member", " %s: %s", member, err.Error())
				t.SkipNow()
			}
			assert.Equalf(note.Note, newNote.Note, "expected member note to be updated %s", oldNote)

			// revert member note
			note.Note = oldNote
			_, err = ord.UpsertMemberNote(*note)
			if err != nil {
				assert.Failf("error reverting member note for member", " %s: %s", member, err.Error())
			}

			// get member ratings
			ratings, err := ord.GetMemberRatings(tt.member)
			if err != nil {
				assert.Failf("error getting member ratings for member", " %s: %s", tt.member, err.Error())
				t.SkipNow()
			}
			if !assert.NotNil(ratings, "ratings response is unexpectly nil") {
				t.SkipNow()
			}
			t.Logf("found %d ratings for member %s",
				ratings.Rating.Complaints+ratings.Rating.Neutrals+ratings.Rating.Praises,
				tt.member)
		})
	}
}

func TestOrders_Updating(t *testing.T) {
	if pendingOrderID == 0 {
		t.SkipNow()
	}
	assert := assert.New(t)
	bricklink, err := bricklink.New()

	if err != nil {
		assert.FailNow(err.Error())
	}

	ord, err := orders.New(*bricklink)
	if err != nil {
		assert.FailNow(err.Error())
	}

	oh, err := ord.GetOrderHeader(pendingOrderID)
	if err != nil {
		assert.Failf("error getting feedback for order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}
	if oh == nil {
		assert.FailNowf("error getting order", " %d: %s", pendingOrderID, "order not found")
		t.SkipNow()
	}

	rnd := " " + util.RandomString(16, bricklink.Rand) // update order
	remarks := oh.Remarks

	// update order header
	oh.Remarks = remarks + rnd
	ohupdate, err := ord.UpdateOrder(*oh)
	if err != nil {
		assert.Failf("error updating order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}
	if ohupdate == nil {
		assert.FailNowf("nil error received. Expected order resource for", " %d.", pendingOrderID)
		t.SkipNow()
	}
	if ohupdate.Remarks != oh.Remarks {
		assert.FailNowf("order not updated", "expected %S; got %s", oh.Remarks, ohupdate.Remarks)
		t.SkipNow()
	}

	// update payment status
	payment := oh.Payment.Status
	p := orders.PaymentStatusBounced
	if payment == p.String() {
		p = orders.PaymentStatusNone
	}
	err = ord.UpdatePaymentStatus(pendingOrderID, p)
	if err != nil {
		assert.Failf("error updating payment status for order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}

	// update order status
	status := oh.Status
	s := orders.StatusCompleted
	if status == s.String() {
		s = orders.StatusPacked
	}
	err = ord.UpdateOrderStatus(pendingOrderID, s)
	if err != nil {
		assert.Failf("error updating order status for order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}

	// revert changes
	// undo order status from original
	oh.Status = status
	err = ord.UpdateOrderStatus(pendingOrderID, s)
	if err != nil {
		assert.Failf("error reverting order status for order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}
	// undo payment status from original
	oh.Payment.Status = payment
	err = ord.UpdatePaymentStatus(pendingOrderID, p)
	if err != nil {
		assert.Failf("error reverting payment status for order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}

	// undo order update from original
	oh.Remarks = remarks
	_, err = ord.UpdateOrder(*oh)
	if err != nil {
		assert.Failf("error reverting order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}

	// check to be sure no unexpected changes occurred
	// Get order, compare to original.
	ohupdate, err = ord.GetOrderHeader(pendingOrderID)
	if err != nil {
		assert.Failf("error retrieving reverted order", " %s: %s", pendingOrderID, err.Error())
		t.SkipNow()
	}
	if oh != ohupdate {
		assert.FailNowf("order not reverted", "expected %s; got %s", oh.Remarks, ohupdate.Remarks)
		t.SkipNow()
	}
}

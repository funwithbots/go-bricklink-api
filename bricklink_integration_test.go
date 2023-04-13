package go_bricklink_api_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	bricklink "github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/util"
)

// TestReference is a basic smoke test for the Bricklink Client.
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

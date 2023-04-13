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
		},
		{
			name: "set test",
			options: []reference.RequestOption{
				reference.WithItemNo("75981-3"),
				reference.WithItemType(util.ItemTypeSet),
				// reference.WithColorID(11),
				// reference.WithBox(true),
				reference.WithBreakMinifigs(false),
				// reference.WithBreakSubsets(true),
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
	}

	opts := []bricklink.BricklinkOption{
		bricklink.WithEnv(),
	}

	// comment this block to test against the real API
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

	bricklink, err := bricklink.New(opts...)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	ref := reference.New(*bricklink)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := ref.GetCatalogItem(tt.options...)
			if err != nil {
				assert.Failf(t, "error getting catalog item:", "%s", err.Error())
			}
			if item != nil {
				assert.Equal(t, tt.want, item.ID)
				assert.Equal(t, strings.ToUpper(tt.itemType.String()), item.ItemType)
			}

			subsets, err := ref.GetSubsets(tt.options...)
			if err != nil {
				assert.Failf(t, "error getting subsets:", "%s", err.Error())
			}
			assert.GreaterOrEqualf(t, len(subsets), 2, "expected at least one subset")
			// assert.Failf(t, "At subsets.", "%v", subsets)

			supersets, err := ref.GetSupersets(tt.options...)
			if err != nil {
				assert.FailNowf(t, "error getting supersets:", "%s", err.Error())
			}
			assert.LessOrEqualf(t, 1, len(supersets), "expected at least one superset")

		})
	}
}

func contains(s []string, str string) bool {
	str = strings.ToLower(str)
	for _, v := range s {
		if strings.ToLower(v) == str {
			return true
		}
	}

	return false
}

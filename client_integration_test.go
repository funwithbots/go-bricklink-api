package go_bricklink_api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	bricklink "github.com/funwithbots/go-bricklink-api"
	"github.com/funwithbots/go-bricklink-api/entity/reference"
	"github.com/funwithbots/go-bricklink-api/internal"
	"github.com/funwithbots/go-bricklink-api/util"
)

// TestReference is a basic smoke test for the Bricklink Client.
func TestReference(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		itemType util.ItemType
		want     string
	}{{
		name:     "smoke test",
		id:       "4592c03",
		itemType: util.ItemTypePart,
		want:     "4592c03",
	}}

	opts := []bricklink.BricklinkOption{
		bricklink.WithEnv(),
	}

	// comment this block to test against the real API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"meta": {"code": 200, "message": "OK"}, "data":{{"id": "3001", "item_type": "PART", "name": "Brick 2 x 4"}}`))
	}))
	client, err := internal.NewClient(internal.WithHTTPClient(server.Client()))
	if err != nil {
		t.Errorf("error creating client: %v", err)
	}
	bricklink.WithClient(client)
	// end block

	bricklink, err := bricklink.New(opts...)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	ref := reference.New(*bricklink)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts = []reference.RequestOption{
				reference.WithItemNo(tt.id),
				reference.WithItemType(tt.itemType),
			}

			item, err := ref.GetCatalogItem(opts...)
			if err != nil {
				assert.FailNowf(t, "error getting catalog item:", "` %s", err.Error())
			}
			assert.Equal(t, tt.want, item.ID)
			assert.Equal(t, strings.ToUpper(tt.itemType.String()), item.ItemType)

			subsets, err := ref.GetSubsets(opts...)
			if err != nil {
				assert.FailNowf(t, "error getting catalog item:", "` %s", err.Error())
			}
			assert.Equal(t, 2, len(subsets), "expected at least one subset")

			supersets, err := ref.GetSupersets(opts...)
			if err != nil {
				assert.FailNowf(t, "error getting catalog item:", "` %s", err.Error())
			}
			assert.Lessf(t, 1, len(supersets), "expected at least one superset")

		})
	}
}

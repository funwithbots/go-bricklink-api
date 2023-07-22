package reference

import (
	"encoding/json"
	"html"
	"testing"
)

func TestItem_UnmarshalJSON(t *testing.T) {
	type fields struct {
		ID           string
		Name         string
		Type         string
		CategoryID   int
		alternateNo  string
		ImageURL     string
		ThumbnailURL string
		Weight       string
		DimX         string
		DimY         string
		DimZ         string
		YearReleased int
		Description  string
		IsObsolete   bool
		LanguageCode string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				ID:           "cty1071",
				Name:         `City Space Robot, "Standing", Medium Azure Eyes`,
				Type:         "M",
				CategoryID:   67,
				ImageURL:     "https://img.bricklink.com/ItemImage/MN/0/cty1071.png",
				ThumbnailURL: "https://img.bricklink.com/ItemImage/MN/0/cty1071.png",
				IsObsolete:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := &Item{
				ID:           tt.fields.ID,
				Name:         html.EscapeString(html.EscapeString(tt.fields.Name)),
				Type:         tt.fields.Type,
				CategoryID:   tt.fields.CategoryID,
				alternateNo:  tt.fields.alternateNo,
				ImageURL:     tt.fields.ImageURL,
				ThumbnailURL: tt.fields.ThumbnailURL,
				Weight:       tt.fields.Weight,
				DimX:         tt.fields.DimX,
				DimY:         tt.fields.DimY,
				DimZ:         tt.fields.DimZ,
				YearReleased: tt.fields.YearReleased,
				Description:  tt.fields.Description,
				IsObsolete:   tt.fields.IsObsolete,
				LanguageCode: tt.fields.LanguageCode,
			}
			data, err := json.Marshal(it)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
				return
			}

			var newItem Item
			if err := newItem.UnmarshalJSON(data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.fields.Name != newItem.Name {
				t.Errorf("UnmarshalJSON() Name = %v, want %v", newItem.Name, tt.fields.Name)
			}
		})
	}
}

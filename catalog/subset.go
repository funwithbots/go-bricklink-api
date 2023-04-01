package catalog

type Subset struct {
	MatchNo int `json:"match_no"`
	Entries []struct {
		Item          Item `json:"item"`
		ColorID       int  `json:"color_id"`
		Quantity      int  `json:"quantity"`
		ExtraQuantity int  `json:"extra_quantity"`
		IsAlternate   bool `json:"is_alternate"`
		IsCounterpart bool `json:"is_counterpart"`
	} `json:"entries"`
}

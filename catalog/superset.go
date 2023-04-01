package catalog

type Superset struct {
	ColorID int `json:"color_id"`
	Entries []struct {
		Item      Item   `json:"item"`
		Quantity  int    `json:"quantity"`
		AppearsAs string `json:"appears_as"`
	} `json:"entries"`
}

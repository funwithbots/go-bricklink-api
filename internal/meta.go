package internal

// Meta represents the meta data returned by the Bricklink API.
type Meta struct {
	Description string `json:"description"`
	Message     string `json:"message"`
	Code        int    `json:"code"`
}

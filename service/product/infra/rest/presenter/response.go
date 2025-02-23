package presenter

// swagger:model Response
// Response represents the structure that our APIs respond in that way
type Response struct {
    Error   string  `json:"error"`
    Message string  `json:"message,omitempty"`
    Data    any     `json:"data,omitempty"`
}

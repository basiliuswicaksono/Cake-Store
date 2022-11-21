package params

// swagger:model
type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}

package params

type CakeRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Image       string  `json:"image" validate:"required"`
}

type UpdateCakeRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}

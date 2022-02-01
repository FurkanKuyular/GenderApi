package module

type GenderPayload struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	CountryCode string `json:"country"`
}

type Gender struct {
	Success       bool          `json:"success"`
	GenderPayload GenderPayload `json:"payload"`
}

type ErrorPayload struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success      bool         `json:"success"`
	ErrorPayload ErrorPayload `json:"error"`
}

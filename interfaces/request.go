package interfaces

type ErrorResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type APIResponse struct {
	Success bool `json:"success"`
	Payload interface{} `json:"payload"`
}

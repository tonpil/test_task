package api

type ErrorResponse struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
}

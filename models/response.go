package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}

type PostReceiptResponse struct {
	Id string `json:"id"`
}

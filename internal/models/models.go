package models

// Cryptocurrency ...
type Cryptocurrency struct {
	Name   string  `fauna:"name" json:"name"`
	Price  float32 `fauna:"price" json:"price"`
	Symbol string  `fauna:"symbol" json:"symbol"`
}

type ListRes struct {
	Data Cryptocurrency `fauna:"data" json:"data"`
}

// BasicError is an error structure that is returned to the API caller indicating the error code and the message
type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

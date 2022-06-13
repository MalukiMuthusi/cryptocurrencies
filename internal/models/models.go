package models

// Cryptocurrency ...
type Cryptocurrency struct {
	Name   string  `fauna:"name"`
	Price  float32 `fauna:"price"`
	Symbol string  `fauna:"symbol"`
}

// BasicError is an error structure that is returned to the API caller indicating the error code and the message
type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

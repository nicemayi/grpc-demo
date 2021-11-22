package dto

// service always deals with dtos
type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type MultipleRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type FibRequest struct {
	Number int `json:"number"`
}

package model

type Product struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

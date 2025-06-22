package models

type ListMenu struct {
	No       string `json:"no"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}
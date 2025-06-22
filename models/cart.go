package models

type CartItem struct {
	No       string
	Name     string
	Price    int
	Category string
	Quantity int
}

type HistoryItem struct {
	Name  string
	Price int
	Total int
}

type Transaction struct {
	Items []HistoryItem
	Total int
}

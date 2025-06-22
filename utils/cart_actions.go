package utils

import (
	"weeklytask-8/models"
)

var Cart []models.ListMenu
var TransactionHistory [][]models.HistoryItem

func AddToCart(item models.ListMenu) {
	Cart = append(Cart, item)
}

func AddToHistory(historyData []models.HistoryItem) {
	TransactionHistory = append(TransactionHistory, historyData)
}

func ClearCart() {
	Cart = []models.ListMenu{}
}

func CalculateCartTotal() int {
	total := 0
	for _, item := range Cart {
		total += item.Price
	}
	return total
}

func GroupCartItems() []models.HistoryItem {
	totalByName := make(map[string]int)
	totalByPrice := make(map[string]int)

	for _, item := range Cart {
		totalByName[item.Name]++
		totalByPrice[item.Name] += item.Price
	}

	var historyItems []models.HistoryItem
	for name, total := range totalByName {
		historyItems = append(historyItems, models.HistoryItem{
			Name:  name,
			Total: total,
			Price: totalByPrice[name],
		})
	}

	return historyItems
}

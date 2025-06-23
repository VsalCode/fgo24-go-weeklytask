package services

import (
    "weeklytask-8/models"
)

type CartService interface {
    AddToCart(item models.ListMenu)
    AddToHistory(historyData []models.HistoryItem)
    ClearCart()
    CalculateCartTotal() int
    GroupCartItems() []models.HistoryItem
}

type CartManager struct {
    Cart               []models.ListMenu
    TransactionHistory [][]models.HistoryItem
}

func (c *CartManager) AddToCart(item models.ListMenu) {
    c.Cart = append(c.Cart, item)
}

func (c *CartManager) AddToHistory(historyData []models.HistoryItem) {
    c.TransactionHistory = append(c.TransactionHistory, historyData)
}

func (c *CartManager) ClearCart() {
    c.Cart = []models.ListMenu{}
}

func (c *CartManager) CalculateCartTotal() int {
    total := 0
    for _, item := range c.Cart {
        total += item.Price
    }
    return total
}

func (c *CartManager) GroupCartItems() []models.HistoryItem {
    totalByName := make(map[string]int)
    totalByPrice := make(map[string]int)

    for _, item := range c.Cart {
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

var ManageCart CartService = &CartManager{}
package utils

import (
	"fmt"
)

var listCardInteractive = `
=========================
| LIST CART 🛒
=========================
`

type HistoryStruct struct {
	Name  string
	Price int
	Total int
}

var DataHistory []HistoryStruct
var CalculateTotal int

func ListCart() {
	fmt.Print("\033[H\033[2J")

	CalculateTotal = 0

	fmt.Print(listCardInteractive)

	totalByName := make(map[string]int)
	totalByPrice := make(map[string]int)

	if len(Cart) != 0 {
		for _, item := range Cart {
		totalByName[item.Name]++
		totalByPrice[item.Name] += item.Price
		CalculateTotal += item.Price
	}

	DataHistory = []HistoryStruct{}

	for name, total := range totalByName {
		DataHistory = append(DataHistory, HistoryStruct{
			Name:  name,
			Total: total,
			Price: totalByPrice[name],
		})
	}

	for x, item := range DataHistory {
		fmt.Printf("[%d]. %s | total item: %d | Total Harga: %d\n", x+1, item.Name, item.Total, item.Price)
	}
	}
	

	fmt.Printf("\nTotal Harga: %d\n", CalculateTotal)
	fmt.Println("=========================")

	fmt.Println("\nKetik 0 untuk kembali ")
	fmt.Println("Ketik 1 untuk checkout ")
	fmt.Print("\nMasukkan pilihan : ")
	var choise string
	fmt.Scanln(&choise)
	fmt.Println("tekan 0 untuk kembali")
	switch choise {
	case "0":
		return
	case "1":
		Checkout(CalculateTotal)
	default:
		return
	}
}

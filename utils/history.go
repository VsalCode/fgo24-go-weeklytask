package utils

import "fmt"

var historyUi = `
=============================
| 💸 HISTORY TRANSACTION ⏳ |
=============================
`
var TransactionHistory [][]HistoryStruct

func History() {
	fmt.Print("\033[H\033[2J")
	fmt.Print(historyUi)

	if len(TransactionHistory) == 0 {
		fmt.Println("Belum ada History Transaksi 😯 !")
	} else {
		for x, items := range TransactionHistory {
			// fmt.Println("-----------------------------")
			fmt.Printf("Transaksi [%d]:\n", x+1)
			fmt.Println("Detail Item yang dibeli:")
			for _, item := range items {
				fmt.Printf("> %s | Total Item: %d | Harga: %d\n", item.Name, item.Total, item.Price)
			}
			fmt.Println("-----------------------------")
		}
	}

	fmt.Printf("\nEnter untuk kembali ke home...")
	var choice2 string
	fmt.Scanln(&choice2)
	switch choice2 {
	default:
		return
	}
}

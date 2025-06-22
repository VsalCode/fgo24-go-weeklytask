package menu

import (
	"fmt"
	"weeklytask-8/services"
	"weeklytask-8/utils"
)

func ShowHistory() {
	utils.ClearScreen()
	fmt.Print(utils.HistoryUI)

	if len(services.TransactionHistory) == 0 {
		fmt.Println("Belum ada History Transaksi 😯 !")
	} else {
		for i, transaction := range services.TransactionHistory {
			fmt.Printf("Transaksi [%d]:\n", i+1)
			fmt.Println("Detail Item yang dibeli:")
			for _, item := range transaction {
				fmt.Printf("> %s | Total Item: %d | Harga: %d\n", 
					item.Name, item.Total, item.Price)
			}
			fmt.Println("-----------------------------")
		}
	}

	utils.WaitForEnter("\nEnter untuk kembali ke home...")
}

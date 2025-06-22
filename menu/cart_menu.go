package menu

import (
	"fmt"
	"weeklytask-8/services"
	"weeklytask-8/utils"
)

func ShowCart() {
	utils.ClearScreen()
	fmt.Print(utils.ListCartUI)

	historyData := services.GroupCartItems()
	total := services.CalculateCartTotal()

	if len(historyData) == 0 {
		fmt.Println("Keranjang kosong!")
	} else {
		for i, item := range historyData {
			fmt.Printf("[%d]. %s | total item: %d | Total Harga: %d\n",
				i+1, item.Name, item.Total, item.Price)
		}
	}

	fmt.Printf("\nTotal Harga: %d\n", total)
	fmt.Println("=========================")

	fmt.Println("\nKetik 0 untuk kembali ")
	fmt.Println("Ketik 1 untuk checkout ")
	choice := utils.GetInput("\nMasukkan pilihan : ")

	switch choice {
	case "0":
		return
	case "1":
		ShowCheckout(total, historyData)
	default:
		return
	}
}

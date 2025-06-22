package menu

import (
	"fmt"
	"weeklytask-8/models"
	"weeklytask-8/services"
	"weeklytask-8/utils"
)

func ShowCheckout(total int, historyData []models.HistoryItem) {
	utils.ClearScreen()
	fmt.Print(utils.CheckoutUI)

	fmt.Println("Item anda saat ini:")

	if len(historyData) == 0 {
		showFailedPayment(total)
		return
	}

	for i, item := range historyData {
		fmt.Printf("[%d]. %s | total item: %d | Total Harga: %d\n", 
			i+1, item.Name, item.Total, item.Price)
	}

	fmt.Printf("\nTotal yang akan dibayarkan: %d\n", total)
	fmt.Println("Anda yakin ingin melanjutkan pembayaran? (YA/tidak)")
	
	choice := utils.GetInputLower("")
	switch choice {
	case "ya":
		services.AddToHistory(historyData)
		services.ClearCart()
		showPaymentSuccess(total)
	case "tidak":
		return
	default:
		ShowCheckout(total, historyData)
	}
}

func showPaymentSuccess(total int) {
	utils.ClearScreen()
	fmt.Print(utils.PaymentSuccessUI)
	fmt.Printf("Total yang anda dibayarkan: %d\n", total)
	fmt.Println("==============================================")
	utils.WaitForEnter("\nEnter untuk kembali ke home...")
}

func showFailedPayment(total int) {
	utils.ClearScreen()
	fmt.Print(utils.FailedPaymentUI)
	fmt.Printf("Total yang anda dibayarkan: %d\n", total)
	fmt.Println("==============================================")
	utils.WaitForEnter("\nEnter untuk kembali ke home...")
}
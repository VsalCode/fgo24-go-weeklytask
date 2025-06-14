package utils

import (
	"fmt"
	"strings"
	"weeklytask-8/data"
)

var checkoutInteractive = `
=========================
| CHECKOUT 💸
=========================
`

var paymentInteractive = `
==============================================
| 💸 PAYMENT SUCCESS! ✅                     |
==============================================
`

var failedPaymentInteractive = `
==============================================
| 💸 PAYMENT FAILED! ❌                      |
| anda belum memasukkan apapun ke keranjang! |
==============================================
`

func Checkout(total int) {
	fmt.Print("\033[H\033[2J")

	fmt.Print(checkoutInteractive)

	fmt.Println("Item anda saat ini:")

	if len(Cart) != 0 {
		for x, item := range DataHistory {
			fmt.Printf("[%d]. %s | total item: %d | Total Harga: %d\n", x+1, item.Name, item.Total, item.Price)
		}
	}

	fmt.Printf("\nTotal yang akan dibayarkan: %d\n", total)
	fmt.Println("Anda yakin ingin melanjutkan pembayaran? (YA/tidak)")
	var choice string
	fmt.Scanln(&choice)
	switch strings.ToLower(choice) {
	case "ya":
		Cart = []data.ListMenu{}
		TransactionHistory = append(TransactionHistory, DataHistory)
		payment(total)
	case "tidak":
		return
	default:
		Checkout(total)
	}
}

func payment(total int) {
	fmt.Print("\033[H\033[2J")

	if CalculateTotal == 0 || Cart == nil {
		fmt.Print(failedPaymentInteractive)
		fmt.Printf("Total yang anda dibayarkan: %d\n", total)
		fmt.Println("==============================================")
		fmt.Printf("\nEnter untuk kembali ke home...")
		var choice2 string
		fmt.Scanln(&choice2)
		switch choice2 {
		default:
			return
		}
	}

	fmt.Print(paymentInteractive)
	fmt.Printf("Total yang anda dibayarkan: %d\n", total)
	fmt.Println("==============================================")
	fmt.Printf("\nEnter untuk kembali ke home...")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	default:
		return
	}
}

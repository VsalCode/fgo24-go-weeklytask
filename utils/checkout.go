package utils

import (
	"fmt"
	"os"
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
| PAYMENT SUCCESS! ✅💸                      |
==============================================
| "1" untuk kembali                          |
| "0" untuk keluar dari program              |
==============================================
`

var failedPaymentInteractive = `
==============================================
| PAYMENT FAILED! ❌💸                       |
| anda belum memasukkan apapun ke keranjang! |
==============================================
| "1" untuk kembali                          |
| "0" untuk keluar dari program              |
==============================================
`

func Checkout(total int) {
	fmt.Print("\033[H\033[2J")

	fmt.Print(checkoutInteractive)
	fmt.Printf("\nTotal yang akan dibayarkan: %d\n", total)
	fmt.Println("Anda yakin ingin melanjutkan pembayaran? (YA/tidak)")
	var choice string
	fmt.Scanln(&choice)
	switch strings.ToLower(choice) {
	case "ya":
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
		fmt.Printf("\nPilih Menu:")
		var choice2 string
		fmt.Scanln(&choice2)
		switch strings.ToLower(choice2) {
		case "1":
			return
		case "0":
			os.Exit(0)
		default:
			return
		}
		return
	}

	fmt.Print(paymentInteractive)
	fmt.Printf("Total yang anda dibayarkan: %d\n", total)
	fmt.Println("==============================================")
	fmt.Printf("\nPilih Menu:")
	var choice string
	fmt.Scanln(&choice)
	switch strings.ToLower(choice) {
	case "1":
		CalculateTotal = 0
		Cart = []data.ListMenu{}
		return
	case "0":
		os.Exit(0)
	default:
		CalculateTotal = 0
		Cart = []data.ListMenu{}
		return
	}
}

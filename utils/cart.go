package utils

import "fmt"

var listCardInteractive = `
=========================
| LIST CART 🛒
=========================
`

var CalculateTotal int


func ListCart() {
	fmt.Print("\033[H\033[2J")

	CalculateTotal = 0

	fmt.Print(listCardInteractive)

	totalByName := make(map[string]int)

	for _, item := range Cart {
		// fmt.Printf("%d.%s | Harga: %d\n",x + 1, item.Name, item.Price)
		totalByName[item.Name]++
		CalculateTotal += item.Price
	}
	
	for name, total := range totalByName {
		fmt.Printf("> %s | total item: %d\n", name, total)
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
		case "0" :
			return
		case "1" :
			Checkout(CalculateTotal)
	}
}
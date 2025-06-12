package utils

import "fmt"

var listCardInteractive = `
=========================
| LIST CART 🛒
=========================
`

func ListCart() {
	defer clear()

	fmt.Print(listCardInteractive)
	for x, item := range Cart {
		fmt.Printf("%d.%s | Harga: %d\n",x + 1, item.Name, item.Price)
	}

	fmt.Print("\nKetik 0 untuk kembali: ")
	var choise string
	fmt.Scanln(&choise)
	fmt.Println("tekan 0 untuk kembali")
	switch choise {
		case "0" :
			return
	}
}
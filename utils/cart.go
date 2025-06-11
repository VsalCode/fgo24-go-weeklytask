package utils

import "fmt"

var listCardInteractive = `
=========================
| LIST CART 🛒
=========================
`

func listCart() {
	fmt.Print(listCardInteractive)
	for x, item := range cart {
		fmt.Printf("%d.%s | Harga: %d\n",x + 1, item.name, item.price)
	}

	fmt.Print("Ketik 0 untuk kembali: ")
	var choise string
	fmt.Scanln(&choise)
	fmt.Println("tekan 0 untuk kembali")
	switch choise {
		case "0" :
			Home()
		
	}
}
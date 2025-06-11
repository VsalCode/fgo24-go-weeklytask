package utils

import (
	"fmt"
	"os"
)

var HomeInteractive = `
=====================
|      WELCOME      |
=====================
| 1. List Menu 📜   |
| 2. Keranjang 🛒   |
| 3. Checkout 💸    |
| 0. Exit ❌        |
=====================        
`

func Home() {
	manageListMenu()
	fmt.Printf("menu yg dipilih : %v\n", cart)
	fmt.Printf("list kategori : %v\n", category)
	fmt.Printf("list kategori yang dipilih : %v\n", choosenCategory)

	fmt.Println(HomeInteractive)
	fmt.Print("Masukkan pilihan: ")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		choosenCategory = ""
		chooseMenu()
	case "2":
		listCart()
	case "3":
		checkout()
	case "0":
		cart = []ListMenu {}
		os.Exit(0)
	default:
		fmt.Println("Invalid choice!")
		Home()
	}

}

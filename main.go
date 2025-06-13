package main

import (
	"fmt"
	"os"
	"weeklytask-8/data"
	"weeklytask-8/utils"
)

var HomeInteractive = `
=====================
|      WELCOME      |
=====================
| 1. List Menu 📜   |
| 2. Cari Menu 🔎   |
| 3. Keranjang 🛒   |
| 4. Checkout 💸    |
| 0. Exit ❌        |
=====================        
`

var DataMenu []data.ListMenu

func main() {
	data.ManageListMenu(&DataMenu)

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("menu yg dipilih : %v\n", utils.Cart)
		fmt.Printf("list kategori : %v\n", utils.Category)
		fmt.Printf("list kategori yang dipilih : %v\n", utils.ChoosenCategory)
		fmt.Printf("total harga : %v\n", utils.CalculateTotal)

		fmt.Println(HomeInteractive)
		fmt.Print("Masukkan pilihan: ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			utils.ChooseMenu(&DataMenu)
			utils.Menu(&DataMenu)
		case "2":
			utils.Search()
		case "3":
			utils.ListCart()
		case "4":
			utils.Checkout(utils.CalculateTotal)
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
			return
		}
	}

}

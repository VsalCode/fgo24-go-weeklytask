package main

import (
	"fmt"
	"os"
	"strings"
	"weeklytask-8/data"
	"weeklytask-8/utils"
)

var HomeInteractive = `
================================
| 🥣 WELCOME TO WARTEG BAHARI  |
================================
| 1. Show All Menu 🍴          |
| 2. List Menu by Category 📜  |
| 3. Cari Menu 🔎              |
| 4. Lihat Keranjang 🛒        |
| 5. Checkout 💸               |
| 6. History Transaction 📋    |
| 0. Exit ❌                   |
================================        
`

var DataMenu []data.ListMenu

func main() {
	data.ManageListMenu(&DataMenu)

	fmt.Print("Masukkan Nama Anda: ")
	var greet string
	fmt.Scanln(&greet)
	
	for {
		fmt.Print("\033[H\033[2J")
		// fmt.Printf("menu yg dipilih : %v\n", utils.Cart)
		// fmt.Printf("list kategori : %v\n", utils.Category)
		// fmt.Printf("list kategori yang dipilih : %v\n", utils.ChoosenCategory)
		// fmt.Printf("total harga : %v\n", utils.CalculateTotal)
		
		fmt.Printf("Halo %s 😄 🖐 !\n", strings.ToUpper(greet))
		fmt.Println(HomeInteractive)
		fmt.Print("Masukkan pilihan: ")
		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			utils.AllMenu(&DataMenu)
		case "2":
			utils.ChooseMenu(&DataMenu)
			if utils.Status {
				utils.Menu(&DataMenu)
			}
		case "3":
			utils.Search(&DataMenu)
		case "4":
			utils.ListCart()
		case "5":
			utils.Checkout(utils.CalculateTotal)
		case "6":
			utils.History()
		case "0":
			fmt.Printf("See You Again %s 😥 🖐 !", greet)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice!")
			return
		}
	}

}

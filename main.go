package main

import (
	"fmt"
	"weeklytask-8/data"
	"weeklytask-8/menu"
	"weeklytask-8/utils"
)

func main() {
	fmt.Println("Loading menu data...")
	listMenu := data.ManageListMenu()
	
	if len(listMenu) == 0 {
		fmt.Println("❌ Gagal memuat data menu. Aplikasi akan keluar.")
		return
	}
	
	for {
		utils.ClearScreen()
		fmt.Print(utils.HomeUI)
		
		choice := utils.GetInput("Pilih menu [1-5]: ")
		
		switch choice {
		case "1":
			menu.ShowAllMenu(listMenu)
		case "2":
			menu.ShowCategoryMenu(listMenu)
		case "3":
			menu.SearchMenu(listMenu)
		case "4":
			menu.ShowCart()
		case "5":
			menu.ShowHistory()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi kami! 👋")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid!")
			utils.WaitForEnter("Tekan Enter untuk kembali...")
		}
	}
}
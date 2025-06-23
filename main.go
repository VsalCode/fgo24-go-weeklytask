package main

import (
	"fmt"
	"strings"
	"weeklytask-8/data"
	"weeklytask-8/menu"
	"weeklytask-8/models"
	"weeklytask-8/utils"
)

func main() {
	userName := utils.GetInput("Masukkan Nama Anda: ")
	
	data.ManageListMenu()

	for {
		utils.ClearScreen()
		fmt.Printf("Halo %s 😄🖐 !\n", strings.ToUpper(userName))
		fmt.Print(utils.HomeUI)
		
		choice := utils.GetInput("Pilih menu [1-5]: ")
		
		switch choice {
		case "1":
			menu.ShowAllMenu(&models.MenuList)
		case "2":
			menu.ShowCategoryMenu(&models.MenuList)
		case "3":
			menu.SearchMenu(&models.MenuList)
		case "4":
			menu.ShowCart()
		case "5":
			menu.ShowHistory()
		case "0":
			fmt.Println("See You Again! 👋 ...")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid!")
			utils.WaitForEnter("Tekan Enter untuk kembali...")
		}
	}
}
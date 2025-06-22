package menu

import (
	"fmt"
	"weeklytask-8/models"
	"weeklytask-8/utils"
)

func ShowAllMenu(listMenu []models.ListMenu) {
	itemsPerPage := 5
	totalItems := len(listMenu)
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
	currentPage := 0

	for {
		utils.ClearScreen()
		fmt.Print(utils.AllMenuUI)

		start := currentPage * itemsPerPage
		end := start + itemsPerPage
		if end > totalItems {
			end = totalItems
		}

		for i := start; i < end; i++ {
			item := listMenu[i]
			fmt.Printf("%s | %s | %d | %s\n", item.No, item.Name, item.Price, item.Category)
		}

		fmt.Println("==============================================")
		fmt.Printf("Page %d of %d\n", currentPage+1, totalPages)
		fmt.Println("[p = previous] [n = next] [Tambahkan ID ke keranjang] [q to back to home]")
		
		choice := utils.GetInputLower("Masukkan Pilihan: ")

		switch choice {
		case "n":
			if currentPage < totalPages-1 {
				currentPage++
			}
		case "p":
			if currentPage > 0 {
				currentPage--
			}
		case "q":
			return
		default:
			if addItemToCart(choice, listMenu) {
				utils.WaitForEnter("Menambahkan item ke Keranjang ✅ , enter untuk kembali ke home...")
			}
		}
	}
}

func addItemToCart(itemID string, listMenu []models.ListMenu) bool {
	for _, item := range listMenu {
		if itemID == item.No {
			utils.AddToCart(item)
			return true
		}
	}
	return false
}
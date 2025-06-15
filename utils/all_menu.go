package utils

import (
	"fmt"
	"strings"
	"weeklytask-8/data"
)

var allMenuUI = `
==============================================
| LIST ALL MENU 🍴                           |
==============================================
| ID produk | nama produk | harga | category |
==============================================
`

func AllMenu(dataParams *[]data.ListMenu) {
	listMenu := *dataParams
	itemsPerPage := 5
	totalItems := len(listMenu)
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage 
	currentPage := 0

	for {
		fmt.Print("\033[H\033[2J") 
		fmt.Print(allMenuUI)

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
		fmt.Print("Masukkan Pilihan: ")

		var choice string
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)

		switch choice {
		case "n":
			if currentPage < totalPages-1 {
				currentPage++
			}
			continue
		case "p":
			if currentPage > 0 {
				currentPage--
			}
			continue
		case "q":
			return
		default:
			for _, item := range listMenu {
				if choice == item.No {
					Cart = append(Cart, data.ListMenu{
						No:       item.No,
						Name:     item.Name,
						Price:    item.Price,
						Category: item.Category,
					})
					fmt.Printf("Menambahkan %s ke Keranjang ✅ , enter untuk kembali ke home...", item.Name)
					fmt.Scanln()
					break
				}
			}
		}
	}
}
package utils

import (
	"fmt"
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

	fmt.Print(allMenuUI)
	for _, item := range listMenu {
		fmt.Printf("%s | %s | %d | %s\n", item.No, item.Name, item.Price, item.Category)
	}

	fmt.Println("==================================")
	fmt.Printf("\nMasukkan Pilihan [%s - %d] : ", listMenu[0].No, len(listMenu))
	var choice string
	fmt.Scanln(&choice)

	notExist := "false"

	for x := range listMenu {
		if listMenu[x].Category == ChoosenCategory {
			isMatch := choice == listMenu[x].No
			if isMatch {
				Cart = append(Cart, data.ListMenu{
					No:       listMenu[x].No,
					Name:     listMenu[x].Name,
					Price:    listMenu[x].Price,
					Category: listMenu[x].Category,
				})
				notExist = "false"
				break
			}
			notExist = "true"
		}
	}

	if notExist == "true" {
		fmt.Print("Pilihan tidak ada❌, enter untuk kembali ke home..")
		var invalid string
		fmt.Scanln(&invalid)
		switch invalid {
		default:
			return
		}
	} else {
		fmt.Print("Berhasil ditambahkan ✅ , enter untuk kembali dan lihat keranjang mu..")
		var invalid string
		fmt.Scanln(&invalid)
		switch invalid {
		default:
			return
		}
	}
}

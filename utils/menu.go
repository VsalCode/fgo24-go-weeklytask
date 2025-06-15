package utils

import (
	"fmt"
	"weeklytask-8/data"
)

var MenuInteractive = `
==================================
| LIST MENU
==================================
| ID produk | nama produk | harga
==================================
`

var categoryInteractive = `
==================================
| LIST CATEGORY 📜
==================================
`

var Cart []data.ListMenu

type CategoryStruct struct {
	No   string
	Name string
}

var Category []CategoryStruct
var ChoosenCategory string

func checkCategory(c1 []CategoryStruct, c2 string) bool {
	for x := range c1 {
		if c1[x].Name == c2 {
			return false
		}
	}
	return true
}

var Status bool

func ChooseMenu(dataParams *[]data.ListMenu) {
	fmt.Print("\033[H\033[2J")
	data := *dataParams

	fmt.Print(categoryInteractive)
	for _, menu := range data {
		if len(Category) == 0 {
			Category = append(Category, CategoryStruct{
				No:   "1",
				Name: menu.Category,
			})
		}

		if checkCategory(Category, menu.Category) {
			Category = append(Category, CategoryStruct{
				No:   fmt.Sprintf("%d", len(Category)+1),
				Name: menu.Category,
			})
		}
	}

	for c, item := range Category {
		fmt.Printf("| %d. %s\n", c+1, item.Name)
	}

	fmt.Print("==================================\n")
	fmt.Printf("Pilih Kategori [%s - %d] : ", Category[0].No, len(Category) )
	var choose string
	fmt.Scanln(&choose)

	var result string
	notExist := "false"

	for y := range Category {
		exist := Category[y].No == choose

		if exist {
			result = Category[y].Name
			notExist = "false"
			break
		}
		notExist = "true"
	}

	Status = true
	if notExist == "true" {
		fmt.Print("Pilihan tidak ada❌, enter untuk kembali ke home..")
		var invalid string
		fmt.Scanln(&invalid)
		switch invalid {
		default:
			Status = false
			return
		}
	}

	ChoosenCategory = result
}

	
func Menu(dataParams *[]data.ListMenu) {
	fmt.Print("\033[H\033[2J")

	listMenu := *dataParams

	fmt.Print(MenuInteractive)
	for _, item := range listMenu {
		if item.Category == ChoosenCategory {
			fmt.Printf("| %s | %s | %d\n", item.No, item.Name, item.Price)
		}
	}
	fmt.Println("==================================")
	fmt.Print("Pilih Menu [ ID Produk ] : ")
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

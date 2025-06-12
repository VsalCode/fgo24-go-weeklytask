package utils

import (
	"fmt"
	"strings"
	"weeklytask-8/data"
)

var MenuInteractive = `
==================================
| LIST MENU
==================================
| no produk | nama produk | harga
==================================
`

var categoryInteractive = `
==================================
| LIST CATEGORY 📜
==================================
`

var Cart []data.ListMenu

type CategoryStruct struct {
	no   string
	name string
}

var Category []CategoryStruct

func checkCategory(c1 []CategoryStruct, c2 string) bool {
	for x := range c1 {
		if c1[x].name == c2 {
			return false
		}
	}
	return true
}

var ChoosenCategory string

func ChooseMenu(dataParams *[]data.ListMenu) {
	defer clear()

	data := *dataParams

	for x, menu := range data {

		if len(Category) == 0 {
			Category = append(Category, CategoryStruct{
				no:   fmt.Sprintf("%d", x+1),
				name: menu.Category,
			})
		}

		if checkCategory(Category, menu.Category) {
			Category = append(Category, CategoryStruct{
				no:   fmt.Sprintf("%d", x),
				name: menu.Category,
			})
		}
	}

	for c, item := range Category {
		fmt.Printf("| %d. %s\n", c+1, item.name)
	}

	fmt.Print("==================================\n")
	fmt.Print("\nPilih Kategori: ")
	var choose string
	fmt.Scanln(&choose)

	var result string

	for y := range Category {
		if choose == Category[y].no {
			result = Category[y].name
		}
		if choose != Category[y].no {
				fmt.Println("Opsi tidak ditemukan!")
				return
			}
	}

	ChoosenCategory = result
	
}

func Menu(dataParams *[]data.ListMenu) {
	defer clear()

	listMenu := *dataParams

	fmt.Print(MenuInteractive)
	for i := range listMenu {
		if listMenu[i].Category == ChoosenCategory {
			fmt.Printf("| %d | %s | %d\n", i+1, listMenu[i].Name, listMenu[i].Price)
		}

		if listMenu[i].Category != ChoosenCategory {
			fmt.Println("Opsi tidak ditemukan!")
			return
		}
	}
	fmt.Println("==================================")
	fmt.Print("\nMasukkan Pilihan : ")
	var choice string
	fmt.Scanln(&choice)

	for x := range listMenu {
		true := strings.Contains(choice, listMenu[x].No)
		if true {
			Cart = append(Cart, data.ListMenu{
				No:       listMenu[x].No,
				Name:     listMenu[x].Name,
				Price:    listMenu[x].Price,
				Category: listMenu[x].Category,
			})
		}
	}

}

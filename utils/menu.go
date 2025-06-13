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
	No   string
	Name string
}

var Category []CategoryStruct

func checkCategory(c1 []CategoryStruct, c2 string) bool {
	for x := range c1 {
		if c1[x].Name == c2 {
			return false
		}
	}
	return true
}

var ChoosenCategory string

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
	fmt.Print("Pilih Kategori: ")
	var choose string
	fmt.Scanln(&choose)

	var result string
	for y := range Category {
		if choose == Category[y].No {
			result = Category[y].Name
		}
	}

	ChoosenCategory = result
}

func Menu(dataParams *[]data.ListMenu) {
	fmt.Print("\033[H\033[2J")

	listMenu := *dataParams

	fmt.Print(MenuInteractive)
	for i := range listMenu {
		if listMenu[i].Category == ChoosenCategory {
			fmt.Printf("| %d | %s | %d\n", i+1, listMenu[i].Name, listMenu[i].Price)
		}
	}
	fmt.Println("==================================")
	fmt.Print("Masukkan Pilihan : ")
	var choice string
	fmt.Scanln(&choice)

	for x := range listMenu {
		isMatch := strings.Contains(choice, listMenu[x].No)
		if isMatch {
			Cart = append(Cart, data.ListMenu{
				No:       listMenu[x].No,
				Name:     listMenu[x].Name,
				Price:    listMenu[x].Price,
				Category: listMenu[x].Category,
			})
		}
	}
}

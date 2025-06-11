package utils

import (
	"fmt"
	"strings"
)

var MenuInteractive = `
=========================
| LIST MENU
=========================
`

var cart []ListMenu

var category []string

func chooseMenu() {
	defer menu()

	fmt.Print(MenuInteractive)
	for x := range listMenu {
		for i := range category {
			if !strings.Contains(category[i], listMenu[x].category) {
				category = append(category, listMenu[x].category )
			}
		}
	}
	
	fmt.Print(category)

	for c, item := range category{
		fmt.Printf("| %d. %s\n", c + 1, item)
	}

	fmt.Print("Pilih Kategori: ")
	var choice string
	fmt.Scanln(&choice)

}

func menu() {

	fmt.Print(MenuInteractive)
	for i, menu := range listMenu {
		fmt.Println("-------------------------")
		fmt.Printf("| %d. %s | %d\n", i+1, menu.name, menu.price)
	}
	fmt.Println("=========================")
	fmt.Print("Masukkan Pilihan : ")
	var choice string
	fmt.Scanln(&choice)

	for x := range listMenu {
		true := strings.Contains(choice, listMenu[x].no)
		if true {
			cart = append(cart, ListMenu{
				name:     listMenu[x].name,
				price:    listMenu[x].price,
				category: listMenu[x].category,
			})
		}
	}

	defer Home()
}

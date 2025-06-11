package utils

import (
	"fmt"
	"strings"
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

var cart []ListMenu

type Category struct {
	no       string
	name string
}

var category []Category

func checkCategory(c1 []Category, c2 string) bool {
	for x := range c1 {
		if c1[x].name == c2 {
			return false
		}
	}
	return true
}

var choosenCategory string

func chooseMenu()  {
	defer menu()

	fmt.Print(categoryInteractive)
	for x, menu := range listMenu {

		if len(category) == 0 {
			category = append(category, Category{
				no:       fmt.Sprintf("%d", x + 1),
				name: menu.category,
			})
		}

		if checkCategory(category, menu.category) {
			category = append(category, Category{
				no:       fmt.Sprintf("%d", x),
				name: menu.category,
			})
		}
	}

	for c, item := range category {
		fmt.Printf("| %d. %s\n", c+1, item.name)
	}

	fmt.Print("==================================\n")
	fmt.Print("Pilih Kategori: ")
	var choose string
	fmt.Scanln(&choose)
	
	var result string

	for y := range category {
		if (choose == category[y].no ){
			result = category[y].name
		} 
	}

	choosenCategory = result
}

func menu() []ListMenu {
	defer Home()
	
	fmt.Print(MenuInteractive)
	for i := range listMenu {
		if listMenu[i].category == choosenCategory {
			fmt.Printf("| %d | %s | %d\n", i+1, listMenu[i].name, listMenu[i].price)
		}
		
	}
	fmt.Println("==================================")
	fmt.Print("Masukkan Pilihan : ")
	var choice string
	fmt.Scanln(&choice)

	for x := range listMenu {
		true := strings.Contains(choice, listMenu[x].no)
		if true {
			cart = append(cart, ListMenu{
				no: listMenu[x].no,
				name:     listMenu[x].name,
				price:    listMenu[x].price,
				category: listMenu[x].category,
			})
		}
	}

	return cart
}

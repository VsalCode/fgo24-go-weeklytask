package menu

import (
	"fmt"
	"weeklytask-8/models"
	"weeklytask-8/services"
	"weeklytask-8/utils"
)

func ShowCategoryMenu(params *[]models.ListMenu) {
	utils.ClearScreen()

	listMenu := *params

	categories := getUniqueCategories(listMenu)

	fmt.Print(utils.HistoryUI)
	for i, category := range categories {
		fmt.Printf("| %d. %s\n", i+1, category.Name)
	}

	fmt.Print("==================================\n")
	choice := utils.GetInput(fmt.Sprintf("Pilih Kategori [1 - %d] : ", len(categories)))

	selectedCategory := ""
	for i, category := range categories {
		if choice == fmt.Sprintf("%d", i+1) {
			selectedCategory = category.Name
			break
		}
	}

	if selectedCategory == "" {
		utils.WaitForEnter("Pilihan tidak ada❌, enter untuk kembali ke home..")
		return
	}

	showMenuByCategory(listMenu, selectedCategory)
}

func getUniqueCategories(listMenu []models.ListMenu) []models.Category {
	categoryMap := make(map[string]bool)
	var categories []models.Category

	counter := 1
	for _, menu := range listMenu {
		if !categoryMap[menu.Category] {
			categories = append(categories, models.Category{
				No:   fmt.Sprintf("%d", counter),
				Name: menu.Category,
			})
			categoryMap[menu.Category] = true
			counter++
		}
	}

	return categories
}

func showMenuByCategory(listMenu []models.ListMenu, category string) {
	utils.ClearScreen()

	fmt.Print(utils.MenuUI)
	for _, item := range listMenu {
		if item.Category == category {
			fmt.Printf("| %s | %s | %d\n", item.No, item.Name, item.Price)
		}
	}

	fmt.Println("==================================")
	choice := utils.GetInput("Pilih Menu [ ID Produk ] : ")

	found := false
	for _, item := range listMenu {
		if item.Category == category && choice == item.No {
			services.AddToCart(item)
			found = true
			break
		}
	}

	if !found {
		utils.WaitForEnter("Pilihan tidak ada❌, enter untuk kembali ke home..")
	} else {
		utils.WaitForEnter("Berhasil ditambahkan ✅ , enter untuk kembali dan lihat keranjang mu..")
	}
}

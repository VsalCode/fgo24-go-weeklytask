package menu

import (
	"fmt"
	"strings"
	"sync"
	"weeklytask-8/models"
	"weeklytask-8/services"
	"weeklytask-8/utils"
)

func SearchMenu(listMenu []models.ListMenu) {
	utils.ClearScreen()

	fmt.Println("*ketik nama item atau nama menu yang anda ingin cari")
	input := utils.GetInput("\ncari item 🔎 : ")

	if input == "" {
		fmt.Println("❌ Input tidak boleh kosong!")
		utils.WaitForEnter("\nEnter untuk kembali ke home...")
		return
	}

	handleSearch(input, listMenu)
}

func handleSearch(query string, listMenu []models.ListMenu) {
	fmt.Println("\nHasil Pencarian :")

	resultChan := make(chan []models.ListMenu)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		searchResults(query, listMenu, resultChan)
	}()

	results := <-resultChan
	wg.Wait()
	close(resultChan)

	if len(results) == 0 {
		fmt.Println("❌Item Tidak Ditemukkan🔎❌")
		handleSearchAction(query, listMenu)
		return
	}

	for _, item := range results {
		fmt.Printf("[ ID Menu: %s ] %s \n", item.No, item.Name)
	}

	choice := utils.GetInput("\n Ketik ID produk untuk ditambahkan ke keranjang: ")
	
	found := false
	for _, item := range results {
		if item.No == choice {
			services.AddToCart(item)
			found = true
			break
		}
	}

	if found {
		fmt.Println("Data Berhasil Ditambahkan ✅ ")
	} else {
		fmt.Print("❌ Input Tidak Valid!")
	}

	handleSearchAction(query, listMenu)
}

func searchResults(query string, listMenu []models.ListMenu, resultChan chan []models.ListMenu) {
	var results []models.ListMenu
	
	for _, item := range listMenu {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(query)) {
			results = append(results, item)
		}
	}
	
	resultChan <- results
}

func handleSearchAction(query string, listMenu []models.ListMenu) {
	fmt.Printf("\nKetik 0 untuk melakukan pencarian kembali...")
	fmt.Printf("\nEnter untuk kembali ke home...")
	
	action := utils.GetInput("")
	if action == "0" {
		SearchMenu(listMenu)
	}
}
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weeklytask-8/data"
)

func Search(dataParams *[]data.ListMenu) {
	fmt.Print("\033[H\033[2J")

	data := *dataParams

	fmt.Println("*ketik nama item atau nama menu yang anda ingin cari")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\ncari item 🔎 : ")
	scanner.Scan()
	input := scanner.Text()

	if input != "" {
		handleSearch(input, data)
	} else {
		input = "false"
		handleSearch(input, data)
	}
}

func handleSearch(inputParams string, dataParams []data.ListMenu) {

	listMenu := dataParams

	fmt.Println("\nHasil Pencarian :")

	var resultSearch []data.ListMenu

	status := 1
	for x := range listMenu {
		true := strings.Contains(strings.ToLower(listMenu[x].Name), strings.ToLower(inputParams))
		if true {
			status = 0
			resultSearch = append(resultSearch, listMenu[x])
			fmt.Printf("[ ID Menu: %s ] %s \n", listMenu[x].No, listMenu[x].Name)
		} else if inputParams == "false" {
			status = 1
		}
	}

	if status == 1 {
		fmt.Println("❌Item Tidak Ditemukkan🔎❌")
		fmt.Printf("\nKetik 0 untuk melakukan pencarian kembali...")
		fmt.Printf("\nEnter untuk kembali ke home...")
	} else {
		fmt.Printf("ini result serach: %v", resultSearch)
		fmt.Print("\n Ketik ID produk untuk ditambahkan ke keranjang: ")

		var chooseItem string
		fmt.Scanln(&chooseItem)
		var isOnTheList string
		
		for x := range resultSearch {
			isMatch := resultSearch[x].No == chooseItem
			if isMatch {
				Cart = append(Cart, data.ListMenu{
					No:       resultSearch[x].No,
					Name:     resultSearch[x].Name,
					Price:    resultSearch[x].Price,
					Category: resultSearch[x].Category,
				})
				isOnTheList = "true"
				break
			}
			isOnTheList = "false"
			// break
		}

		if isOnTheList == "true" {
			fmt.Println("Data Berhasil Ditambahkan ✅ ")
		} else {
			fmt.Print("❌ Input Tidak Valid!")
		}

		fmt.Printf("\nKetik 0 untuk melakukan pencarian kembali...")
		fmt.Printf("\nEnter untuk kembali ke home...")
	}

	var action string
	fmt.Scanln(&action)
	switch action {
	case "0":
		Search(&dataParams)
	default:
		return
	}

}

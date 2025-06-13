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

func handleSearch(input string, dataParams []data.ListMenu) {

	data := dataParams

	fmt.Println("\nHasil Pencarian :")

	status := 1
	for x := range data {
		true := strings.Contains(strings.ToLower(data[x].Name), strings.ToLower(input))
		if true {
			status = 0
			fmt.Printf("> %s\n", data[x].Name)
		} else if input == "false" {
			status = 1
		}
	}

	if status == 1 {
		fmt.Println("❌Item Tidak Ditemukkan🔎❌")
		fmt.Printf("\nKetik 0 untuk melakukan pencarian kembali...")
		fmt.Printf("\nEnter untuk kembali ke home...")
		} else {
		fmt.Printf("\nKetik 0 untuk melakukan pencarian kembali...")
		fmt.Printf("\nEnter untuk kembali ke home dan pesan di list menu...")
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

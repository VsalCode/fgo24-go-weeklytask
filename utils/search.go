package utils

import "fmt"

var searchInteractive = `
=========================================
|  	          CARI MENU 🔎             	|
=========================================
| 1. Cari Berdasarkan Nama              |
| 2. Cari Berdasarkan Kategori          |
| 3. Lihat Berdasarkan Rating Tertinggi |
| 4. Lihat Berdasarkan Harga Termurah   |
=========================================
`

func Search() {
	fmt.Print("\033[H\033[2J")

	fmt.Print(searchInteractive)
	fmt.Print("Pilih Menu : ")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
		case "1":
		case "2":
		case "3":
		default :
			return
	}
}

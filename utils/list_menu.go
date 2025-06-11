package utils

import "fmt"

type ListMenu struct {
	no string
	name     string
	price    int
	category string
}

var listMenu []ListMenu

func manageListMenu() []ListMenu {
	listMenu = []ListMenu{
		{
			no: "1",
			name:     "Nasi Padang",
			price:    12000,
			category: "food",
		},
		{
			no: "2",
			name:     "Fried Chicken",
			price:    7000,
			category: "food",
		},
		{
			no: "3",
			name:     "Matcha",
			price:    7000,
			category: "drink",
		},
		{
			no: "4",
			name:     "Fried Rice",
			price:    15000,
			category: "food",
		},
		{
			no: "5",
			name:     "Ice Tea",
			price:    5000,
			category: "drink",
		},
		{
			no: "6",
			name:     "Coffee",
			price:    10000,
			category: "drink",
		},
	}

	fmt.Println(listMenu)
	return listMenu
}

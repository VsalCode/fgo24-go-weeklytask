package data

type ListMenu struct {
	No string
	Name     string
	Price    int
	Category string
}

var editData []ListMenu

func ManageListMenu(data *[]ListMenu) {

	editData = []ListMenu{
		{
			No: "1",
			Name:     "Nasi Padang",
			Price:    12000,
			Category: "food",
		},
		{
			No: "2",
			Name:     "Fried Chicken",
			Price:    7000,
			Category: "food",
		},
		{
			No: "3",
			Name:     "Matcha",
			Price:    7000,
			Category: "drink",
		},
		{
			No: "4",
			Name:     "Fried Rice",
			Price:    15000,
			Category: "food",
		},
		{
			No: "5",
			Name:     "Ice Tea",
			Price:    5000,
			Category: "drink",
		},
		{
			No: "6",
			Name:     "Coffee",
			Price:    10000,
			Category: "drink",
		},
	}

	*data = editData
	// fmt.Println(data)
}

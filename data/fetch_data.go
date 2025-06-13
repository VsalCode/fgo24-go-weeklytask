package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ListMenu struct {
	No       string 
	Name     string
	Price    int
	Category string
}

func fetchData(c chan []ListMenu) {

	res, err := http.Get("https://raw.githubusercontent.com/VsalCode/go-weeklytask-data/refs/heads/main/data.json")
	if err != nil {
		fmt.Println("failed to fetch!")
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("failed to fetch!")
	}


	var data []ListMenu
	json.Unmarshal([]byte(body), &data)

	c <- data
}

func ManageListMenu(data *[]ListMenu) {

	
	channel := make(chan []ListMenu)
	
	go fetchData(channel)
	
	*data = <-channel
	// fmt.Println(data)
	defer close(channel)
}

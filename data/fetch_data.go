package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type ListMenu struct {
	No       string 
	Name     string
	Price    int
	Category string
}

func fetchData(c chan []ListMenu, wg *sync.WaitGroup) {
	defer wg.Done()

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

	wg := sync.WaitGroup{}
	channel := make(chan []ListMenu)
	
	wg.Add(1)
	go fetchData(channel, &wg)
	
	*data = <-channel
	wg.Wait()

	defer close(channel)
}

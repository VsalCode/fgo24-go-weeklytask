package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"weeklytask-8/models"
)

func fetchData(c chan []models.ListMenu, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("https://raw.githubusercontent.com/VsalCode/go-weeklytask-data/refs/heads/main/data.json")
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		c <- []models.ListMenu{}
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		c <- []models.ListMenu{}
		return
	}

	var data []models.ListMenu
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		c <- []models.ListMenu{}
		return
	}

	c <- data
}

func ManageListMenu() []models.ListMenu {
	wg := sync.WaitGroup{}
	channel := make(chan []models.ListMenu)
	
	wg.Add(1)
	go fetchData(channel, &wg)
	
	data := <-channel
	wg.Wait()
	close(channel)
	
	return data
}

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var (
	data     Data
	mutex    sync.Mutex
	stopChan chan struct{}
)

func main() {

	data = Data{Water: 0, Wind: 0}

	if len(os.Args) == 1 {
		stopChan = make(chan struct{})
		go updateData()
		<-stopChan
	} else {
		if err := readInputAndDisplayStatus(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func updateData() {
	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ticker.C:
			updateRandomData()
			displayData()
			checkStatusAndLog()
		case <-stopChan:
			ticker.Stop()
			return
		}
	}
}

func updateRandomData() {
	mutex.Lock()
	defer mutex.Unlock()

	data.Water = rand.Intn(100) + 1
	data.Wind = rand.Intn(100) + 1
}

func displayData() {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Printf("Water: %d meters, Wind: %d meters/second\n", data.Water, data.Wind)
}

func checkStatusAndLog() {
	mutex.Lock()
	defer mutex.Unlock()

	var waterStatus string
	if data.Water < 5 {
		waterStatus = "Aman"
	} else if data.Water >= 5 && data.Water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	var windStatus string
	if data.Wind < 6 {
		windStatus = "Aman"
	} else if data.Wind >= 6 && data.Wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	fmt.Printf("Status: Water %s, Wind %s\n", waterStatus, windStatus)
}

func readInputAndDisplayStatus() error {
	var input Data

	jsonInput := os.Args[1]
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		return err
	}

	fmt.Printf("Input: Water %d meters, Wind %d meters/second\n", input.Water, input.Wind)

	fmt.Println("Status:")
	checkAndDisplayStatus(input.Water, input.Wind)

	return nil
}

func checkAndDisplayStatus(water, wind int) {
	var waterStatus string
	if water < 5 {
		waterStatus = "Aman"
	} else if water >= 5 && water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	var windStatus string
	if wind < 6 {
		windStatus = "Aman"
	} else if wind >= 6 && wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	fmt.Printf("Water %s, Wind %s\n", waterStatus, windStatus)
}

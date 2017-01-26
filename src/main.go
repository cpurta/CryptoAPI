package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type CoinPriceMap struct {
	PriceMap    map[string]CoinInfo
	Mutex       *sync.Mutex
	LastUpdated time.Time
}

func NewCoinPriceMap() *CoinPriceMap {
	return &CoinPriceMap{PriceMap: make(map[string]CoinInfo), Mutex: &sync.Mutex{}, LastUpdated: time.Now()}
}

func (cm *CoinPriceMap) Marshal() ([]byte, error) {
	return json.Marshal(cm.PriceMap)
}

func (cm *CoinPriceMap) MonitorPrices(server *Server) {
	ticker := time.NewTicker(time.Minute * 1)

	for range ticker.C {
		start := time.Now()
		log.Println("Pulling crytpocurrency prices...")

		url := "https://coinmarketcap-nexuist.rhcloud.com/api/all"

		req, _ := http.NewRequest("GET", url, nil)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error making request to %s: %s", url, err.Error())
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Println("Error reading response body:", err.Error())
			return
		}

		cm.Mutex.Lock()
		err = json.Unmarshal(body, &cm.PriceMap)
		if err != nil {
			log.Println("Error unmarshalling response body into pricing map:", err.Error())
		} else {
			cm.LastUpdated = time.Now()
		}
		cm.Mutex.Unlock()

		log.Printf("Finished pulling all cryptocurrency prices in %.3f", time.Since(start).Seconds())
	}
}

func main() {
	coinMap := NewCoinPriceMap()

	server := NewServer(coinMap)

	go coinMap.MonitorPrices(server)

	log.Println("Starting Crypto API...")
	server.Start()

	log.Println("Exiting Crypto API")
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

type CoinPriceMap struct {
	PriceMap    map[string]*CoinInfo
	LastUpdated time.Time
}

func (cp CoinPrices) MonitorPrices(server *Server) {
	ticker := time.NewTicker(time.Minute * 1)

	cm := &CoinPriceMap{PriceMap: make(map[string]*CoinInfo), LastUpdated: time.Now()}
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

		err = json.Unmarshal(body, &cm.PriceMap)
		if err != nil {
			log.Println("Error unmarshalling response body into pricing map:", err.Error())
			continue
		} else {
			cm.LastUpdated = time.Now()
		}

		log.Printf("Finished pulling all cryptocurrency prices in %.3f", time.Since(start).Seconds())

		updated := make(CoinPrices, 0)
		for _, info := range cm.PriceMap {
			updated = append(updated, info)
		}

		cp = updated
		sort.Sort(cp)
	}
}

func main() {
	coinPrices := make(CoinPrices, 0)

	server := NewServer(coinPrices)

	go coinPrices.MonitorPrices(server)

	log.Println("Starting Crypto API...")
	server.Start()

	log.Println("Exiting Crypto API")
}

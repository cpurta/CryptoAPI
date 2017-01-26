package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Prices CoinPrices
}

func NewServer(prices CoinPrices) *Server {
	return &Server{Prices: prices}
}

func (server *Server) Start() {
	http.HandleFunc("/api/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			response, err := json.Marshal(server.Prices)
			if err != nil {
				log.Println("Error marshalling coin map for API response:", err.Error())
				return
			}

			w.Write(response)
		} else {
			w.WriteHeader(400)
			w.Write([]byte(`{"error":true,"message":"only GET method allowed for this endpoint"}`))
		}
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err.Error())
			w.WriteHeader(500)
			w.Write([]byte(`{"error":true,"message":"Unable to register this device at this time"}`))
			return
		}

		log.Println("Registering device:", string(body))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"error":false,"message":"Succesfully registered device"}`))
	})

	certFile := os.Getenv("CRYPTO_API_CERT_FILE")
	privateKey := os.Getenv("CRYPTO_API_PRIVATE_FILE")

	log.Fatalln(http.ListenAndServeTLS(":443", certFile, privateKey, nil))
}

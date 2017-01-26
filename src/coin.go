package main

type CoinInfo struct {
	Symbol    string                 `json:"symbol"`
	Position  string                 `json:"position"`
	Name      string                 `json:"name"`
	MarketCap map[string]interface{} `json:"market_cap"`
	Price     map[string]interface{} `json:"price"`
	Supply    string                 `json:"supply"`
	Volume    map[string]interface{} `json:"volume"`
	Change    string                 `json:"change"`
	Timestamp string                 `json:"timestamp"`
}

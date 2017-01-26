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

var defaultCurrency = "usd"

type CoinPrices []*CoinInfo

func (cp CoinPrices) Len() int { return len(cp) }
func (cp CoinPrices) Less(i, j int) bool {
	if cp[i].Price != nil && cp[j].Price != nil {
		_, iOK := cp[i].Price[defaultCurrency]
		_, jOK := cp[j].Price[defaultCurrency]

		if iOK && jOK {
			return cp[i].Price[defaultCurrency].(float64) < cp[j].Price[defaultCurrency].(float64)
		}
	}

	return false
}
func (cp CoinPrices) Swap(i, j int) { cp[i], cp[j] = cp[j], cp[i] }

func (cp CoinPrices) Contains(info *CoinInfo) bool {
	for _, price := range cp {
		if info.Symbol == price.Symbol && info.Name == price.Name {
			return true
		}
	}

	return false
}

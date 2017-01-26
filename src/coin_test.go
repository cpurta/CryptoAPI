package main

import (
	"encoding/json"
	"testing"
	"time"
)

var testJSONMap = []byte(`
{
    "btc": {
        "symbol": "btc",
        "position": "1",
        "name": "Bitcoin",
        "market_cap": {
            "aud": 19609569400.612415,
            "btc": "16126562.0",
            "cad": 19350849294.928265,
            "cny": 101583191315.76682,
            "eur": 13840041642.309317,
            "gbp": 11739877012.942284,
            "hkd": 114556865396.82388,
            "jpy": 1693624309180.9878,
            "rub": 893427519569.2037,
            "usd": 14766286495.3
        },
        "price": {
            "aud": 1215.9795373999998,
            "btc": "1.0",
            "cad": 1199.9364337499999,
            "cny": 6299.122609999999,
            "eur": 858.21402245,
            "gbp": 727.9838699000001,
            "hkd": 7103.61361565,
            "jpy": 105020.7917336,
            "rub": 55400.9912075,
            "usd": 915.65
        },
        "supply": "16126562",
        "volume": {
            "aud": 158698177.99199998,
            "btc": "130522.0",
            "cad": 156604383.45,
            "cny": 822102058.8,
            "eur": 112005998.046,
            "gbp": 95009587.09200001,
            "hkd": 927096635.502,
            "jpy": 13706322998.688,
            "rub": 7230414734.100001,
            "usd": 119502000
        },
        "change": "1.69",
        "timestamp": "1485455805.514"
    },
    "btcd": {
        "symbol": "btcd",
        "position": "45",
        "name": "BitcoinDark",
        "market_cap": {
            "aud": 6789521.588879851,
            "btc": 5584.09852149,
            "cad": 6699943.602380822,
            "cny": 35171668.30211841,
            "eur": 4791908.460698816,
            "gbp": 4064757.710981487,
            "hkd": 39663609.88741794,
            "jpy": 586392213.7054638,
            "rub": 309335983.2789108,
            "usd": 5112606.95731
        },
        "price": {
            "aud": 5.267841412959999,
            "btc": 0.00433258,
            "cad": 5.198339810999999,
            "cny": 27.288928744,
            "eur": 3.71793704548,
            "gbp": 3.15375667096,
            "hkd": 30.774128046759998,
            "jpy": 454.96890276543996,
            "rub": 240.00702875800002,
            "usd": 3.96676
        },
        "supply": "1288862",
        "volume": {
            "aud": 9309.902678039998,
            "btc": 7.657,
            "cad": 9187.07188275,
            "cny": 48227.964905999994,
            "eur": 6570.74299377,
            "gbp": 5573.66203254,
            "hkd": 54387.38843049,
            "jpy": 804070.56215856,
            "rub": 424166.54272950004,
            "usd": 7010.49
        },
        "change": "-7.55",
        "timestamp": "1485455805.585"
    }
}
`)

func getCoinPrices(t *testing.T) CoinPrices {
	cm := &CoinPriceMap{PriceMap: make(map[string]*CoinInfo), LastUpdated: time.Now()}

	err := json.Unmarshal(testJSONMap, &cm.PriceMap)
	if err != nil {
		t.Error("Error unmarshalling JSON")
	}

	prices := make(CoinPrices, len(cm.PriceMap))
	i := 0
	for _, info := range cm.PriceMap {
		prices[i] = info
		i++
	}

	return prices
}

func TestLen(t *testing.T) {
	cp := make(CoinPrices, 0)

	if cp.Len() != 0 {
		t.Error("Expected a len of 0 but got:", cp.Len())
	}

	cp = getCoinPrices(t)

	if cp.Len() != 2 {
		t.Error("Expected a len of 2 but got:", cp.Len())
	}
}

func TestLess(t *testing.T) {
	prices := getCoinPrices(t)

	if prices.Less(0, 1) {
		t.Error("Expected false but got", prices.Less(0, 1))
	}

	prices = append(prices, &CoinInfo{})

	if prices.Less(1, 2) {
		t.Error("Expected false but got", prices.Less(1, 2))
	}
}

func TestSwap(t *testing.T) {
	prices := getCoinPrices(t)

	prices.Swap(0, 1)

	if prices[0].Symbol != "btcd" {
		t.Error("Expected btcd but got:", prices[0].Name)
	}

	if prices[1].Symbol != "btc" {
		t.Error("Expected btc but got:", prices[0].Name)
	}
}

func TestContains(t *testing.T) {
	c1 := &CoinInfo{Symbol: "mdc", Position: "12", Name: "MadeUpCoin", MarketCap: make(map[string]interface{}), Supply: "12904875", Change: "-120.2", Timestamp: "1485455805.585"}

	prices := getCoinPrices(t)

	if prices.Contains(c1) {
		t.Error("prices should not contain our fake coin info")
	}

	if !prices.Contains(prices[0]) {
		t.Error("We should expect that prices[0] is contained in prices")
	}
}

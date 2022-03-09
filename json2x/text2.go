package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Fruit struct {
		Name     string `json:"name"`
		PriceTag string `json:"priceTag"`
	}

	type FruitBasket struct {
		Name    string           `json:"name"`
		Fruit   map[string]Fruit `json:"fruit"`
		Id      int64            `json:"id"`
		Created time.Time        `json:"created"`
	}
	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {
              "1": {
                    "name": "Apple",
                    "priceTag": "$1"
              },
              "2": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              }
        },
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range basket.Fruit {
		fmt.Println(item.Name, item.PriceTag)
	}
}

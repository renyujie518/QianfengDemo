package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bingo := map[string]interface{}{
		"name": "Bingo Huang",
		"age":  399,
	}
	json_bytes, _ := json.Marshal(bingo)
	fmt.Printf("%s", json_bytes)
}

package main

import (
	"encoding/json"
	"fmt"
)

/*type Person struct {
	Name   string
	Age    int
	Emails []string
}*/
type Person struct {
	Name string `json:"names,omitempty"`
	Age  int    `json:"age,omitempty"`
	//Age    int   `json: "-"`
	Emails []string `json:"Emails,omitempty"`
}

func main() {
	bingo := Person{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com", "me@bingohuang.com"},
	}
	json_bytes, err := json.Marshal(bingo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", json_bytes)
}

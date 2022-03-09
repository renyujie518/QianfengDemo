package main

import (
	"encoding/json"
	"os"
)

type Person1 struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	bingo := Person1{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com", "me@bingohuang.com"},
	}
	fileWriter, _ := os.Create("/Users/renyujie/go/src/QianfengDemo/x2json/output.json")
	json.NewEncoder(fileWriter).Encode(bingo)
}

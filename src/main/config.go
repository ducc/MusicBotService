package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Keys apiKeys `json:"api_keys"`
}

type apiKeys struct {
	Youtube string `json:"youtube"`
}

func load(filename string) *config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config,", err)
		return nil
	}
	var conf config
	json.Unmarshal(body, &conf)
	return &conf
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Database struct {
		Host     string
		Username string
		Password string
		Database string
		Ssl      bool
	}
	Keys struct {
		Youtube string
	} `json:"api_keys"`
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

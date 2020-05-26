package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config ...
type Config struct {
	Address string `json:"address"`
}

// LoadConfiguration :
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

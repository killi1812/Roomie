package Helpers

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Port int `json:"port"`
	//TODO add more configs
}

var instance *Config = nil
var once sync.Once
var fileName = "config.json"

func loadConfig() *Config {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		createMockConfig()
	} else if err != nil {
		os.Exit(4)
		//TODO add logging
	}
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		os.Exit(5)
	}
	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		os.Exit(6)
	}

	return &config
}

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

func createMockConfig() {
	config := Config{
		Port: 8832,
	}
	data, err := json.Marshal(config)

	if err != nil {
		//TODO log
		return
	}

	err = SaveToFile(fileName, data)
	if err != nil {
		os.Exit(7)
	}
}

func ReloadConfig() {
	instance = loadConfig()
}

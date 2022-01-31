package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func NewConfig() Config {
	var config Config

	configFilePath := fmt.Sprintf("configs/appconfig.json")

	configFile, err := os.Open(configFilePath)

	if err != nil {
		log.Fatal("Config File could not open")
	}

	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)

	defer configFile.Close()

	log.Printf("Config loaded")

	return config
}

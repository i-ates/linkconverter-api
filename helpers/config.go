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

	var LinkConverterEnv = Environment

	configFilePath := fmt.Sprintf("configs/appconfig.%s.json", LinkConverterEnv)

	configFile, err := os.Open(configFilePath)

	if err != nil {
		log.Fatal("Config File could not open!")
	}

	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)

	defer configFile.Close()

	log.Printf("Config loaded environment: %s\n", LinkConverterEnv)

	return config
}

func GetAppName() string {
	return fmt.Sprintf("%s-%s", AppName, Environment)
}

package helpers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"linkconverter-api/libs/logging"
	"linkconverter-api/libs/runtimeenvironment"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func NewConfig() Config {
	var config Config

	var LinkConverterEnv = runtimeenvironment.Environment

	configFilePath := fmt.Sprintf("configs/appconfig.%s.json", LinkConverterEnv)

	configFile, err := os.Open(configFilePath)

	if err != nil {
		logging.Fatal("Config File could not open", zap.Error(err))
	}

	jsonParser := json.NewDecoder(configFile)
	_ = jsonParser.Decode(&config)

	defer configFile.Close()

	logging.Info("Config loaded environment")

	return config
}

func GetAppName() string {
	return fmt.Sprintf("%s-%s", AppName, runtimeenvironment.Environment)
}

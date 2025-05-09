package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error reading config file:", err)
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string `json:"token"`
	Prefix string `json:"prefix"`
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := &Config{}

		config.Token = ""
		config.Prefix = "!"

		bytes, _ := json.Marshal(config)

		if err = ioutil.WriteFile("config.json", bytes, 0777); err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	config := &Config{}

	_ = json.Unmarshal(data, config)

	return config, nil
}
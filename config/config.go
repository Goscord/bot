package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Token            string `json:"token"`
	WelcomeChannelId string `json:"welcome_channel_id"`
	MemberRoleId     string `json:"member_role_id"`
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := new(Config)

		bytes, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return nil, err
		}

		// write json to file
		err = ioutil.WriteFile("config.json", bytes, 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	config := new(Config)

	_ = json.Unmarshal(data, config)

	return config, nil
}

package steamclient

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// SteamConfig struct holds the values as read from the configuration file
// config.json on startup
type SteamConfig struct {
	SteamAPIKey     string   `json:"SteamAPIKey"`
	SteamIDs        []string `json:"SteamIDs"`
	HistoryInterval int      `json:"HistoryInterval"`
	UpdateInterval  int      `json:"UpdateInterval"`
}

func NewSteamConfig(configPath string) (SteamConfig, error) {

	conf := SteamConfig{}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		return conf, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &conf)
	return conf, err
}

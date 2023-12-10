package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Configuration struct {
	MyID        int
	MyDomain    string
	AppID       string
	AccessToken string
	RecentIP    string
	ProxyAddr   string
	ProxyUse    bool
}

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &App.config); err != nil {
		fmt.Println("Error reading configuration file:", err)
		// set default values
		// ...
	} else {
		fmt.Println("Configuration values loaded from config.toml")
	}
}

func SaveConfig() {
	file, err := os.Create("config.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if err := toml.NewEncoder(file).Encode(App.config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Configuration values saved to config.toml")
}

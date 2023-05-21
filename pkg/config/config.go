package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Color ColorConfig `json:"color"`
}

type ColorConfig struct {
	Reset       string `json:"reset"`
	Black       string `json:"black"`
	DarkGray    string `json:"darkGray"`
	Red         string `json:"red"`
	LightRed    string `json:"lightRed"`
	Green       string `json:"green"`
	LightGreen  string `json:"lightGreen"`
	Brown       string `json:"brown"`
	Yellow      string `json:"yellow"`
	Blue        string `json:"blue"`
	LightBlue   string `json:"lightBlue"`
	Purple      string `json:"purple"`
	LightPurple string `json:"lightPurple"`
	Cyan        string `json:"cyan"`
	LightCyan   string `json:"lightCyan"`
	LightGray   string `json:"lightGray"`
	White       string `json:"white"`
}

var DefaultConfig = Config{
	Color: ColorConfig{
		Reset:       "\033[0m",
		Black:       "\033[0;30m",
		DarkGray:    "\033[1;30m",
		Red:         "\033[0;31m",
		LightRed:    "\033[1;31m",
		Green:       "\033[0;32m",
		LightGreen:  "\033[1;32m",
		Brown:       "\033[0;33m",
		Yellow:      "\033[1;33m",
		Blue:        "\033[0;34m",
		LightBlue:   "\033[1;34m",
		Purple:      "\033[0;35m",
		LightPurple: "\033[1;35m",
		Cyan:        "\033[0;36m",
		LightCyan:   "\033[1;36m",
		LightGray:   "\033[0;37m",
		White:       "\033[1;37m",
	},
}

var AppConfig = DefaultConfig

const (
	configPath = "/.config/dockercolorize/config.json"
)

func LoadConfig() {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Printf("failed to get home dir: %s. Using default.", err)
		return
	}

	configFile, err := ioutil.ReadFile(home + configPath)

	if err != nil {
		log.Printf("failed to load color config: %s. Using default.", err)
		return
	}

	err = json.Unmarshal(configFile, &AppConfig)

	if err != nil {
		log.Printf("failed to parse color config: %s. Using default.", err)
		AppConfig = DefaultConfig
	}
}

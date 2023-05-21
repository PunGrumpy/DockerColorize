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
	ResetColor       string `json:"reset"`
	BlackColor       string `json:"black"`
	DarkGrayColor    string `json:"darkGray"`
	RedColor         string `json:"red"`
	LightRedColor    string `json:"lightRed"`
	GreenColor       string `json:"green"`
	LightGreenColor  string `json:"lightGreen"`
	BrownColor       string `json:"brown"`
	YellowColor      string `json:"yellow"`
	BlueColor        string `json:"blue"`
	LightBlueColor   string `json:"lightBlue"`
	PurpleColor      string `json:"purple"`
	LightPurpleColor string `json:"lightPurple"`
	CyanColor        string `json:"cyan"`
	LightCyanColor   string `json:"lightCyan"`
	LightGrayColor   string `json:"lightGray"`
	WhiteColor       string `json:"white"`
}

var DefaultConfig = Config{
	Color: ColorConfig{
		ResetColor:       "\033[0m",
		BlackColor:       "\033[0;30m",
		DarkGrayColor:    "\033[1;30m",
		RedColor:         "\033[0;31m",
		LightRedColor:    "\033[1;31m",
		GreenColor:       "\033[0;32m",
		LightGreenColor:  "\033[1;32m",
		BrownColor:       "\033[0;33m",
		YellowColor:      "\033[1;33m",
		BlueColor:        "\033[0;34m",
		LightBlueColor:   "\033[1;34m",
		PurpleColor:      "\033[0;35m",
		LightPurpleColor: "\033[1;35m",
		CyanColor:        "\033[0;36m",
		LightCyanColor:   "\033[1;36m",
		LightGrayColor:   "\033[0;37m",
		WhiteColor:       "\033[1;37m",
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

package config

import (
	"encoding/json"
	"fmt"
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

type AppConfigProvider struct {
	AppConfig Config
}

const (
	DefaultFilePermissions = 0o644
	DefaultDirPermissions  = 0o755
)

func NewAppConfigProvider() *AppConfigProvider {
	provider := &AppConfigProvider{
		AppConfig: Config{
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
		},
	}
	provider.createDefaultConfigIfNotExists()
	provider.LoadConfig()

	return provider
}

func SetConfigProvider(provider *AppConfigProvider) {
	AppConfig = provider.AppConfig
}

var AppConfig Config //nolint:gochecknoglobals

func (a *AppConfigProvider) createDefaultConfigIfNotExists() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home dir: %s", err)
	}

	defaultConfigPath := fmt.Sprintf("%s/.config/dockercolorize/config.json", home)
	_, err = os.Stat(defaultConfigPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll(home+"/.config/dockercolorize", DefaultDirPermissions)

		if err != nil {
			log.Fatalf("failed to create directory: %s", err)
		}

		a.AppConfig = Config{
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

		defaultConfigFile, err := json.MarshalIndent(a.AppConfig, "", " ")
		if err != nil {
			log.Fatalf("failed to marshal default config: %s", err)
		}

		err = os.WriteFile(defaultConfigPath, defaultConfigFile, DefaultFilePermissions)

		if err != nil {
			log.Fatalf("failed to write default config file: %s", err)
		}
	} else if err != nil {
		log.Fatalf("failed to check config file existence: %s", err)
	}
}

func (a *AppConfigProvider) LoadConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get home dir: %s", err)
	}

	configFilePath := home + "/.config/dockercolorize/config.json"

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	err = json.Unmarshal(configFile, &a.AppConfig)
	if err != nil {
		log.Fatalf("failed to unmarshal config file: %s", err)
	}
}

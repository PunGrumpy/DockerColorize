package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	DefaultConfigFile      = "config.json"
	ConfigEnvVar           = "DOCKERCOLORIZE_CONFIG"
)

var defaultConfig = Config{
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

func NewAppConfigProvider() *AppConfigProvider {
	provider := &AppConfigProvider{
		AppConfig: defaultConfig,
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
	configPath := a.getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(configPath), DefaultDirPermissions); err != nil {
			log.Printf("failed to create config directory: %v", err)
			return
		}

		if err := a.writeConfig(configPath, defaultConfig); err != nil {
			log.Printf("failed to write default config: %v", err)
		}
	} else if err != nil {
		log.Printf("failed to check config file existence: %v", err)
	}
}

func (a *AppConfigProvider) LoadConfig() {
	configPath := a.getConfigPath()
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("failed to read config file: %v", err)
		return
	}

	if err := json.Unmarshal(configFile, &a.AppConfig); err != nil {
		log.Printf("failed to unmarshal config file: %v", err)
	}
}

func (a *AppConfigProvider) getConfigPath() string {
	if configPath := os.Getenv(ConfigEnvVar); configPath != "" {
		return configPath
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("failed to get home dir: %v", err)
		return DefaultConfigFile
	}

	return filepath.Join(home, ".config", "dockercolorize", DefaultConfigFile)
}

func (a *AppConfigProvider) writeConfig(path string, config Config) error {
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, configData, DefaultFilePermissions); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}

package color

import (
	"sync"

	"github.com/PunGrumpy/dockercolorize/pkg/config"
)

type Provider struct {
	reset       string
	black       string
	darkGray    string
	red         string
	lightRed    string
	green       string
	lightGreen  string
	brown       string
	yellow      string
	blue        string
	lightBlue   string
	purple      string
	lightPurple string
	cyan        string
	lightCyan   string
	lightGray   string
	white       string
}

var (
	providerInstance *Provider //nolint:gochecknoglobals
	once             sync.Once //nolint:gochecknoglobals
)

func getColorProvider() Provider {
	once.Do(func() {
		config.SetConfigProvider(config.NewAppConfigProvider())
		providerInstance = &Provider{
			reset:       config.AppConfig.Color.Reset,
			black:       config.AppConfig.Color.Black,
			darkGray:    config.AppConfig.Color.DarkGray,
			red:         config.AppConfig.Color.Red,
			lightRed:    config.AppConfig.Color.LightRed,
			green:       config.AppConfig.Color.Green,
			lightGreen:  config.AppConfig.Color.LightGreen,
			brown:       config.AppConfig.Color.Brown,
			yellow:      config.AppConfig.Color.Yellow,
			blue:        config.AppConfig.Color.Blue,
			lightBlue:   config.AppConfig.Color.LightBlue,
			purple:      config.AppConfig.Color.Purple,
			lightPurple: config.AppConfig.Color.LightPurple,
			cyan:        config.AppConfig.Color.Cyan,
			lightCyan:   config.AppConfig.Color.LightCyan,
			lightGray:   config.AppConfig.Color.LightGray,
			white:       config.AppConfig.Color.White,
		}
	})

	return *providerInstance
}

func Black(value string) string {
	return getColorProvider().black + value + getColorProvider().reset
}

func DarkGray(value string) string {
	return getColorProvider().darkGray + value + getColorProvider().reset
}

func Red(value string) string {
	return getColorProvider().red + value + getColorProvider().reset
}

func LightRed(value string) string {
	return getColorProvider().lightRed + value + getColorProvider().reset
}

func Green(value string) string {
	return getColorProvider().green + value + getColorProvider().reset
}

func LightGreen(value string) string {
	return getColorProvider().lightGreen + value + getColorProvider().reset
}

func Brown(value string) string {
	return getColorProvider().brown + value + getColorProvider().reset
}

func Yellow(value string) string {
	return getColorProvider().yellow + value + getColorProvider().reset
}

func Blue(value string) string {
	return getColorProvider().blue + value + getColorProvider().reset
}

func LightBlue(value string) string {
	return getColorProvider().lightBlue + value + getColorProvider().reset
}

func Purple(value string) string {
	return getColorProvider().purple + value + getColorProvider().reset
}

func LightPurple(value string) string {
	return getColorProvider().lightPurple + value + getColorProvider().reset
}

func Cyan(value string) string {
	return getColorProvider().cyan + value + getColorProvider().reset
}

func LightCyan(value string) string {
	return getColorProvider().lightCyan + value + getColorProvider().reset
}

func LightGray(value string) string {
	return getColorProvider().lightGray + value + getColorProvider().reset
}

func White(value string) string {
	return getColorProvider().white + value + getColorProvider().reset
}

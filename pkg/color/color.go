package color

import (
	"github.com/PunGrumpy/dokcercolorize/pkg/config"
)

var (
	reset       = config.AppConfig.Color.Reset
	black       = config.AppConfig.Color.Black
	darkGray    = config.AppConfig.Color.DarkGray
	red         = config.AppConfig.Color.Red
	lightRed    = config.AppConfig.Color.LightRed
	green       = config.AppConfig.Color.Green
	lightGreen  = config.AppConfig.Color.LightGreen
	brown       = config.AppConfig.Color.Brown
	yellow      = config.AppConfig.Color.Yellow
	blue        = config.AppConfig.Color.Blue
	lightBlue   = config.AppConfig.Color.LightBlue
	purple      = config.AppConfig.Color.Purple
	lightPurple = config.AppConfig.Color.LightPurple
	cyan        = config.AppConfig.Color.Cyan
	lightCyan   = config.AppConfig.Color.LightCyan
	lightGray   = config.AppConfig.Color.LightGray
	white       = config.AppConfig.Color.White
)

func Black(value string) string {
	return black + value + reset
}

func DarkGray(value string) string {
	return darkGray + value + reset
}

func Red(value string) string {
	return red + value + reset
}

func LightRed(value string) string {
	return lightRed + value + reset
}

func Green(value string) string {
	return green + value + reset
}

func LightGreen(value string) string {
	return lightGreen + value + reset
}

func Brown(value string) string {
	return brown + value + reset
}

func Yellow(value string) string {
	return yellow + value + reset
}

func Blue(value string) string {
	return blue + value + reset
}

func LightBlue(value string) string {
	return lightBlue + value + reset
}

func Purple(value string) string {
	return purple + value + reset
}

func LightPurple(value string) string {
	return lightPurple + value + reset
}

func Cyan(value string) string {
	return cyan + value + reset
}

func LightCyan(value string) string {
	return lightCyan + value + reset
}

func LightGray(value string) string {
	return lightGray + value + reset
}

func White(value string) string {
	return white + value + reset
}

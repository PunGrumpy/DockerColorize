package color

import (
	"github.com/PunGrumpy/dokcercolorize/pkg/config"
)

var (
	reset       = config.AppConfig.Color.ResetColor
	black       = config.AppConfig.Color.BlackColor
	darkGray    = config.AppConfig.Color.DarkGrayColor
	red         = config.AppConfig.Color.RedColor
	lightRed    = config.AppConfig.Color.LightRedColor
	green       = config.AppConfig.Color.GreenColor
	lightGreen  = config.AppConfig.Color.LightGreenColor
	brown       = config.AppConfig.Color.BrownColor
	yellow      = config.AppConfig.Color.YellowColor
	blue        = config.AppConfig.Color.BlueColor
	lightBlue   = config.AppConfig.Color.LightBlueColor
	purple      = config.AppConfig.Color.PurpleColor
	lightPurple = config.AppConfig.Color.LightPurpleColor
	cyan        = config.AppConfig.Color.CyanColor
	lightCyan   = config.AppConfig.Color.LightCyanColor
	lightGray   = config.AppConfig.Color.LightGrayColor
	white       = config.AppConfig.Color.WhiteColor
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

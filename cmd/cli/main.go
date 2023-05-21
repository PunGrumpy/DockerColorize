//go:build !test

package main

import (
	"os"

	"dockercolorize/internal/app"
	"dockercolorize/internal/stdin"
	"dockercolorize/internal/stdout"
	"dockercolorize/pkg/config"
)

func main() {
	provider := config.NewAppConfigProvider()
	config.SetConfigProvider(provider)

	if err := run(); err != nil {
		app.Usage(err)
		os.Exit(1)
	}
}

func run() error {
	in, err := stdin.Get()
	if err != nil {
		return err
	}

	rows, err := app.Run(in)
	if err != nil {
		return err
	}

	for _, row := range rows {
		stdout.Println(row)
	}

	return nil
}

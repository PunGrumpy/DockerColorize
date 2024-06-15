//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/PunGrumpy/dockercolorize/internal/app"
	"github.com/PunGrumpy/dockercolorize/internal/stdin"
	"github.com/PunGrumpy/dockercolorize/internal/stdout"
	"github.com/PunGrumpy/dockercolorize/pkg/config"
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
	err := stdin.Stream(func(line string) {
		rows, err := app.Run([]string{line})
		if err != nil {
			stdout.Println(err.Error())
			return
		}

		for _, row := range rows {
			stdout.Println(row)
		}
	})
	if err != nil {
		return fmt.Errorf("failed to stream stdin: %w", err)
	}

	return nil
}

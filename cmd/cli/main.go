//go:build !test

package main

import (
	"os"

	"github.com/PunGrumpy/dokcercolorize/internal/app"
	"github.com/PunGrumpy/dokcercolorize/internal/stdin"
	"github.com/PunGrumpy/dokcercolorize/internal/stdout"
)

func main() {
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

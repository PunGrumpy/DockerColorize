package cmd

import (
	"errors"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/internal/util"
)

var ErrInvalidFirstLine = errors.New("invalid first line")

func Parse(header layout.Header) (Command, error) {
	columns := make([]string, len(header))
	for i, col := range header {
		columns[i] = string(col.Name)
	}

	commands := []struct {
		instance Command
		columns  []string
	}{
		{&DockerPs{}, (&DockerPs{}).Columns()},
		{&DockerImages{}, (&DockerImages{}).Columns()},
		{&DockerComposePs{}, (&DockerComposePs{}).Columns()},
		{&DockerStats{}, (&DockerStats{}).Columns()},
		{&DockerHistory{}, (&DockerHistory{}).Columns()},
	}

	for _, cmd := range commands {
		if util.Intersect(columns, cmd.columns) {
			return cmd.instance, nil
		}
	}

	return nil, ErrInvalidFirstLine
}

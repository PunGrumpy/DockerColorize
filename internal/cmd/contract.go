package cmd

import (
	"github.com/PunGrumpy/dockercolorize/internal/layout"
)

const ValidTotalParts = 2

type Command interface {
	Format(row layout.Row, column layout.Column) string
}

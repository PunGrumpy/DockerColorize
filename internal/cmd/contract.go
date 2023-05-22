package cmd

import (
	"github.com/PunGrumpy/dockercolorize/internal/layout"
)

const ValidTotalParts = 2

type Command interface {
	Format(layout.Row, layout.Column) string
}

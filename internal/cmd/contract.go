package cmd

import (
	"dockercolorize/internal/layout"
)

const ValidTotalParts = 2

type Command interface {
	Format(layout.Row, layout.Column) string
}

package cmd

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
	"github.com/PunGrumpy/dockercolorize/pkg/util/number"
)

const (
	DockerHistoryImage     = "IMAGE"
	DockerHistoryCreated   = "CREATED"
	DockerHistoryCreatedBy = "CREATED BY"
	DockerHistorySize      = "SIZE"
	DockerHistoryComment   = "COMMENT"
)

type DockerHistory struct{}

func (c *DockerHistory) Columns() []string {
	return []string{
		DockerHistoryImage,     //
		DockerHistoryCreated,   //
		DockerHistoryCreatedBy, //
		DockerHistorySize,      //
		DockerHistoryComment,   //
	}
}

func (c *DockerHistory) Format(rows layout.Row, col layout.Column) string {
	v := string(rows[col])

	switch col {
	case DockerHistoryImage:
		return c.Image(v)
	case DockerHistoryCreated:
		return c.Created(v)
	case DockerHistoryCreatedBy:
		return c.CreatedBy(v)
	case DockerHistorySize:
		return c.Size(v)
	case DockerHistoryComment:
		return c.Comment(v)
	}

	return v
}

func (c *DockerHistory) Image(v string) string {
	parts := strings.Split(v, ":") //nolint:ifshort
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	if strings.Contains(v, "<missing>") {
		return color.Red(v)
	}

	return color.Yellow(v)
}

func (c *DockerHistory) Created(v string) string {
	if strings.Contains(v, "months") {
		return color.Brown(v)
	}

	if strings.Contains(v, "years") {
		return color.Red(v)
	}

	return color.Green(v)
}

func (c *DockerHistory) CreatedBy(v string) string {
	return color.LightGreen(v)
}

func (c *DockerHistory) Size(v string) string {
	if strings.Contains(v, "GB") {
		return color.Red(v)
	}

	if strings.Contains(v, "MB") && number.ParseFloat(v) >= 500 {
		return color.Brown(v)
	}

	return v
}

func (c *DockerHistory) Comment(v string) string {
	return color.DarkGray(v)
}

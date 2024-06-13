package cmd

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
)

const (
	DockerComposePsName    = "NAME"
	DockerComposePsImage   = "IMAGE"
	DockerComposePsCommand = "COMMAND"
	DockerComposePsService = "SERVICE"
	DockerComposePsCreated = "CREATED"
	DockerComposePsStatus  = "STATUS"
	DockerComposePsPorts   = "PORTS"
)

type DockerComposePs struct{}

func (c *DockerComposePs) Columns() []string {
	return []string{
		DockerComposePsName,
		DockerComposePsImage,
		DockerComposePsCommand,
		DockerComposePsService,
		DockerComposePsCreated,
		DockerComposePsStatus,
		DockerComposePsPorts,
	}
}

func (c *DockerComposePs) Format(row layout.Row, col layout.Column) string {
	v := string(row[col])

	formatters := map[string]func(string, layout.Row) string{
		DockerComposePsName:    c.Name,
		DockerComposePsImage:   func(v string, _ layout.Row) string { return c.Image(v) },
		DockerComposePsCommand: func(v string, _ layout.Row) string { return c.Command(v) },
		DockerComposePsService: c.Service,
		DockerComposePsCreated: func(v string, _ layout.Row) string { return c.Created(v) },
		DockerComposePsStatus:  func(v string, _ layout.Row) string { return c.Status(v) },
		DockerComposePsPorts:   func(v string, _ layout.Row) string { return c.Ports(v) },
	}

	colString := string(col)
	if formatter, exists := formatters[colString]; exists {
		return formatter(v, row)
	}

	return v
}

func (*DockerComposePs) Name(v string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(v)
	}

	return color.White(v)
}

func (*DockerComposePs) Image(v string) string {
	parts := strings.Split(v, ":")
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(v)
}

func (*DockerComposePs) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerComposePs) Service(v string, row layout.Row) string {
	if strings.Contains(string(row[DockerComposePsStatus]), "exited") {
		return color.DarkGray(v)
	}

	return color.Yellow(v)
}

func (*DockerComposePs) Created(v string) string {
	switch {
	case strings.Contains(v, "months"):
		return color.Brown(v)
	case strings.Contains(v, "years"):
		return color.Red(v)
	default:
		return color.Green(v)
	}
}

func (*DockerComposePs) Status(v string) string {
	if strings.Contains(v, "exited") {
		return color.Red(v)
	}

	return color.LightGreen(v)
}

func (*DockerComposePs) Ports(v string) string {
	var ports []string

	for _, port := range strings.Split(v, ", ") {
		parts := strings.Split(port, "->")
		if len(parts) == ValidTotalParts {
			port = color.LightCyan(parts[0]) + "->" + parts[1]
		}

		ports = append(ports, port)
	}

	return strings.Join(ports, ", ")
}

package cmd

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
)

const (
	DockerPsContainerID = "CONTAINER ID"
	DockerPsImage       = "IMAGE"
	DockerPsCommand     = "COMMAND"
	DockerPsCreatedAt   = "CREATED AT"
	DockerPsCreated     = "CREATED"
	DockerPsPorts       = "PORTS"
	DockerPsState       = "STATE"
	DockerPsStatus      = "STATUS"
	DockerPsSize        = "SIZE"
	DockerPsNames       = "NAMES"
	DockerPsLabels      = "LABELS"
	DockerPsMounts      = "MOUNTS"
	DockerPsNetworks    = "NETWORKS"
)

type DockerPs struct{}

func (c *DockerPs) Columns() []string {
	return []string{
		DockerPsContainerID,
		DockerPsImage,
		DockerPsCommand,
		DockerPsCreatedAt,
		DockerPsCreated,
		DockerPsPorts,
		DockerPsState,
		DockerPsStatus,
		DockerPsSize,
		DockerPsNames,
		DockerPsLabels,
		DockerPsMounts,
		DockerPsNetworks,
	}
}

func (c *DockerPs) Format(rows layout.Row, col layout.Column) string {
	v := string(rows[col])

	formatters := map[string]func(string) string{
		DockerPsContainerID: c.ContainerID,
		DockerPsImage:       c.Image,
		DockerPsCommand:     c.Command,
		DockerPsCreated:     c.Created,
		DockerPsStatus:      c.Status,
		DockerPsPorts:       c.Ports,
		DockerPsNames:       c.Names,
	}

	colString := string(col)
	if formatter, exists := formatters[colString]; exists {
		return formatter(v)
	}

	return v
}

func (*DockerPs) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (*DockerPs) Image(v string) string {
	parts := strings.Split(v, ":")
	if len(parts) == ValidTotalParts {
		return color.Yellow(parts[0]) + color.LightGreen(":"+parts[1])
	}

	return color.Yellow(v)
}

func (*DockerPs) Command(v string) string {
	return color.DarkGray(v)
}

func (*DockerPs) Created(v string) string {
	switch {
	case strings.Contains(v, "months"):
		return color.Brown(v)
	case strings.Contains(v, "years"):
		return color.Red(v)
	default:
		return color.Green(v)
	}
}

func (*DockerPs) Status(v string) string {
	if strings.Contains(v, "Exited") {
		return color.Red(v)
	}

	return color.LightGreen(v)
}

func (*DockerPs) Ports(v string) string {
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

func (*DockerPs) Names(v string) string {
	return color.White(v)
}

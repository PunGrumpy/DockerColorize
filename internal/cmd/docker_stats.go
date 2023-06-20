package cmd

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
	"github.com/PunGrumpy/dockercolorize/pkg/util/number"
)

const (
	DockerStatsContainerID = "CONTAINER ID"
	DockerStatsName        = "NAME"
	DockerStatsCPUPerc     = "CPU %"
	DockerStatsMemUsage    = "MEM USAGE / LIMIT"
	DockerStatsMemPerc     = "MEM %"
	DockerStatsNetIO       = "NET I/O"
	DockerStatsBlockIO     = "BLOCK I/O"
	DockerStatsPIDs        = "PIDS"
)

const (
	cpuThresholdLow        = 25
	cpuThresholdMedium     = 50
	cpuThresholdHigh       = 75
	memThresholdLow        = 25 * 1024 * 1024
	memThresholdMedium     = 50 * 1024 * 1024
	memThresholdHigh       = 75 * 1024 * 1024
	netIOThresholdLow      = 25 * 1024 * 1024
	netIOThresholdMedium   = 50 * 1024 * 1024
	netIOThresholdHigh     = 75 * 1024 * 1024
	blockIOThresholdLow    = 1 * 1024 * 1024
	blockIOThresholdMedium = 250 * 1024 * 1024
	blockIOThresholdHigh   = 500 * 1024 * 1024
)

type DockerStats struct{}

func (c *DockerStats) Columns() []string {
	return []string{
		DockerStatsContainerID, //
		DockerStatsName,        //
		DockerStatsCPUPerc,     //
		DockerStatsMemUsage,    //
		DockerStatsMemPerc,     //
		DockerStatsNetIO,       //
		DockerStatsBlockIO,     //
		DockerStatsPIDs,        //
	}
}

func (c *DockerStats) Format(rows layout.Row, col layout.Column) string {
	v := string(rows[col])

	switch col {
	case DockerStatsContainerID:
		return c.ContainerID(v)
	case DockerStatsName:
		return c.Name(v)
	case DockerStatsCPUPerc:
		return c.CPUPerc(v)
	case DockerStatsMemUsage:
		return c.MemUsage(v)
	case DockerStatsMemPerc:
		return c.MemPerc(v)
	case DockerStatsNetIO:
		return c.NetIO(v)
	case DockerStatsBlockIO:
		return c.BlockIO(v)
	case DockerStatsPIDs:
		return c.PIDs(v)
	default:
		return v
	}
}

func (c *DockerStats) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (c *DockerStats) Name(v string) string {
	return color.White(v)
}

func (c *DockerStats) CPUPerc(v string) string {
	switch {
	case number.ParseFloat(v) >= cpuThresholdHigh:
		return color.Brown(v)
	case number.ParseFloat(v) >= cpuThresholdMedium:
		return color.Yellow(v)
	case number.ParseFloat(v) >= cpuThresholdLow:
		return color.Red(v)
	default:
		return color.Green(v)
	}
}

func (c *DockerStats) MemUsage(v string) string {
	memUsage := strings.Split(v, "/")[0]
	limit := color.DarkGray("/" + strings.Split(v, "/")[1])

	switch {
	case number.ParseBytes(memUsage) >= memThresholdHigh:
		return color.Red(v) + limit
	case number.ParseBytes(memUsage) >= memThresholdMedium:
		return color.Yellow(v) + limit
	case number.ParseBytes(memUsage) >= memThresholdLow:
		return color.Brown(memUsage) + limit
	default:
		return color.Green(memUsage) + limit
	}
}

func (c *DockerStats) MemPerc(v string) string {
	switch {
	case number.ParseFloat(v) >= cpuThresholdHigh:
		return color.Red(v)
	case number.ParseFloat(v) >= cpuThresholdMedium:
		return color.Yellow(v)
	case number.ParseFloat(v) >= cpuThresholdLow:
		return color.Brown(v)
	default:
		return color.Green(v)
	}
}

func (*DockerStats) NetIO(v string) string {
	netIO := number.ParseBytes(v)

	switch {
	case netIO >= netIOThresholdHigh || strings.Contains(v, "GB"):
		return color.Red(v)
	case netIO >= netIOThresholdMedium || strings.Contains(v, "MB"):
		return color.Yellow(v)
	case netIO >= netIOThresholdLow || strings.Contains(v, "kB"):
		return color.Brown(v)
	default:
		return v
	}
}

func (*DockerStats) BlockIO(v string) string {
	blockIO := number.ParseBytes(v)

	switch {
	case blockIO >= blockIOThresholdHigh || strings.Contains(v, "GB"):
		return color.Red(v)
	case blockIO >= blockIOThresholdMedium || strings.Contains(v, "MB"):
		return color.Yellow(v)
	case blockIO >= blockIOThresholdLow || strings.Contains(v, "kB"):
		return color.Brown(v)
	default:
		return v
	}
}

func (*DockerStats) PIDs(v string) string {
	return v
}

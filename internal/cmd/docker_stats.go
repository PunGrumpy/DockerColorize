package cmd

import (
	"strconv"
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
	cpuThresholdLow         = 50
	cpuThresholdMedium      = 80
	cpuThresholdHigh        = 90
	memThresholdLow         = 50
	memThresholdMedium      = 80
	memThresholdHigh        = 90
	memUsageThresholdHigh   = 0.8
	memUsageThresholdMedium = 0.5
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
	case number.ParseFloat(v) >= cpuThresholdLow:
		return color.Yellow(v)
	case number.ParseFloat(v) >= cpuThresholdMedium:
		return color.Brown(v)
	case number.ParseFloat(v) >= cpuThresholdHigh:
		return color.Red(v)
	default:
		return color.Green(v)
	}
}

func (c *DockerStats) MemUsage(v string) string {
	memUsage := strings.Split(v, "/")[0]
	limit := strings.Split(v, "/")[1]

	memUsageInt, _ := strconv.Atoi(strings.TrimSpace(memUsage))
	limitInt, _ := strconv.Atoi(strings.TrimSpace(limit))

	memPerLimit := float64(memUsageInt) / float64(limitInt)

	if memPerLimit >= memUsageThresholdHigh {
		return color.Red(v)
	}

	if memPerLimit >= memUsageThresholdMedium {
		return color.Brown(v)
	}

	return color.Green(v)
}

func (c *DockerStats) MemPerc(v string) string {
	switch {
	case number.ParseFloat(v) >= memThresholdLow:
		return color.Brown(v)
	case number.ParseFloat(v) >= memThresholdMedium:
		return color.Yellow(v)
	case number.ParseFloat(v) >= memThresholdHigh:
		return color.Red(v)
	default:
		return color.Green(v)
	}
}

func (*DockerStats) NetIO(v string) string {
	netIO := number.ParseBytes(v)

	if strings.Contains(v, "GB") {
		return color.Red(v)
	}

	if strings.Contains(v, "MB") && netIO >= 500 {
		return color.Brown(v)
	}

	return color.LightGreen(v)
}

func (*DockerStats) BlockIO(v string) string {
	blockIO := number.ParseBytes(v)

	if strings.Contains(v, "GB") {
		return color.Red(v)
	}

	if strings.Contains(v, "MB") && blockIO >= 500 {
		return color.Brown(v)
	}

	return color.LightGreen(v)
}

func (*DockerStats) PIDs(v string) string {
	return color.LightGreen(v)
}

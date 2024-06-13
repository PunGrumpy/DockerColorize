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
	memUsageParse           = 2
)

type DockerStats struct{}

func (c *DockerStats) Columns() []string {
	return []string{
		DockerStatsContainerID,
		DockerStatsName,
		DockerStatsCPUPerc,
		DockerStatsMemUsage,
		DockerStatsMemPerc,
		DockerStatsNetIO,
		DockerStatsBlockIO,
		DockerStatsPIDs,
	}
}

func (c *DockerStats) Format(rows layout.Row, col layout.Column) string {
	v := string(rows[col])

	formatters := map[string]func(string) string{
		DockerStatsContainerID: c.ContainerID,
		DockerStatsName:        c.Name,
		DockerStatsCPUPerc:     c.CPUPerc,
		DockerStatsMemUsage:    c.MemUsage,
		DockerStatsMemPerc:     c.MemPerc,
		DockerStatsNetIO:       c.NetIO,
		DockerStatsBlockIO:     c.BlockIO,
		DockerStatsPIDs:        c.PIDs,
	}

	colString := string(col)
	if formatter, exists := formatters[colString]; exists {
		return formatter(v)
	}

	return v
}

func (c *DockerStats) ContainerID(v string) string {
	return color.DarkGray(v)
}

func (c *DockerStats) Name(v string) string {
	return color.White(v)
}

func (c *DockerStats) CPUPerc(v string) string {
	cpu := number.ParseFloat(v)

	switch {
	case cpu >= cpuThresholdHigh:
		return color.Red(v)
	case cpu >= cpuThresholdMedium:
		return color.Brown(v)
	case cpu >= cpuThresholdLow:
		return color.Yellow(v)
	default:
		return color.Green(v)
	}
}

func (c *DockerStats) MemUsage(v string) string {
	parts := strings.Split(v, "/")
	if len(parts) != memUsageParse {
		return v
	}

	memUsage := strings.TrimSpace(parts[0])
	limit := strings.TrimSpace(parts[1])

	memUsageInt, err1 := strconv.Atoi(memUsage)
	limitInt, err2 := strconv.Atoi(limit)

	if err1 != nil || err2 != nil {
		return v
	}

	memPerLimit := float64(memUsageInt) / float64(limitInt)

	switch {
	case memPerLimit >= memUsageThresholdHigh:
		return color.Red(v)
	case memPerLimit >= memUsageThresholdMedium:
		return color.Brown(v)
	default:
		return color.Green(v)
	}
}

func (c *DockerStats) MemPerc(v string) string {
	mem := number.ParseFloat(v)

	switch {
	case mem >= memThresholdHigh:
		return color.Red(v)
	case mem >= memThresholdMedium:
		return color.Brown(v)
	case mem >= memThresholdLow:
		return color.Yellow(v)
	default:
		return color.Green(v)
	}
}

func (c *DockerStats) NetIO(v string) string {
	return c.ioFormatter(v)
}

func (c *DockerStats) BlockIO(v string) string {
	return c.ioFormatter(v)
}

func (c *DockerStats) ioFormatter(v string) string {
	io := number.ParseBytes(v)

	switch {
	case strings.Contains(v, "GB"):
		return color.Red(v)
	case strings.Contains(v, "MB") && io >= 500:
		return color.Brown(v)
	default:
		return color.LightGreen(v)
	}
}

func (c *DockerStats) PIDs(v string) string {
	return color.LightGreen(v)
}

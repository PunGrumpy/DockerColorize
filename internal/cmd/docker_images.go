package cmd

import (
	"strings"

	"github.com/PunGrumpy/dockercolorize/internal/layout"
	"github.com/PunGrumpy/dockercolorize/pkg/color"
	"github.com/PunGrumpy/dockercolorize/pkg/util/number"
)

const (
	DockerImagesRepository = "REPOSITORY"
	DockerImagesTag        = "TAG"
	DockerImagesDigest     = "DIGEST"
	DockerImagesImageID    = "IMAGE ID"
	DockerImagesCreated    = "CREATED"
	DockerImagesCreatedAt  = "CREATED AT"
	DockerImagesSize       = "SIZE"
)

type DockerImages struct{}

func (c *DockerImages) Columns() []string {
	return []string{
		DockerImagesImageID,
		DockerImagesRepository,
		DockerImagesTag,
		DockerImagesDigest,
		DockerImagesCreated,
		DockerImagesCreatedAt,
		DockerImagesSize,
	}
}

func (c *DockerImages) Format(row layout.Row, col layout.Column) string {
	v := string(row[col])

	formatters := map[string]func(string) string{
		DockerImagesRepository: c.Repository,
		DockerImagesTag:        c.Tag,
		DockerImagesImageID:    c.ImageID,
		DockerImagesCreated:    c.Created,
		DockerImagesSize:       c.Size,
	}

	colString := string(col)
	if formatter, exists := formatters[colString]; exists {
		return formatter(v)
	}

	return v
}

func (*DockerImages) Repository(v string) string {
	if strings.Contains(v, "/") {
		return color.DarkGray(v)
	}

	return color.White(v)
}

func (*DockerImages) Tag(v string) string {
	if v == "latest" {
		return color.LightGreen(v)
	}

	return v
}

func (*DockerImages) ImageID(v string) string {
	return color.DarkGray(v)
}

func (*DockerImages) Created(v string) string {
	switch {
	case strings.Contains(v, "hour"):
		return color.Green(v)
	case strings.Contains(v, "days"), strings.Contains(v, "weeks"):
		return color.Green(v)
	case strings.Contains(v, "months"):
		return color.Brown(v)
	case strings.Contains(v, "years"):
		return color.Red(v)
	default:
		return v
	}
}

func (*DockerImages) Size(v string) string {
	switch {
	case strings.Contains(v, "GB"):
		return color.Red(v)
	case strings.Contains(v, "MB") && number.ParseFloat(v) >= 500:
		return color.Brown(v)
	default:
		return v
	}
}

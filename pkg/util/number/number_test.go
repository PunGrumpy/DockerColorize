package number_test

import (
	"testing"

	"github.com/PunGrumpy/dockercolorize/pkg/util/assert"
	"github.com/PunGrumpy/dockercolorize/pkg/util/number"
)

func TestParseFloat(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in   string
		want float64
	}{
		"plain":       {in: "100", want: 100},
		"float":       {in: "100.10", want: 100.10},
		"measurement": {in: "100MB", want: 100},
		"trim":        {in: " 100 ", want: 100},
		"empty":       {in: "", want: 0},
		"unparsed":    {in: "-", want: 0},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, number.ParseFloat(tt.in))
		})
	}
}

func TestParseBytes(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in   string
		want float64
	}{
		"plain":    {in: "100", want: 100},
		"float":    {in: "100.10", want: 100.10},
		"KB":       {in: "50KB", want: 50 * 1024},
		"MB":       {in: "2MB", want: 2 * 1024 * 1024},
		"GB":       {in: "5GB", want: 5 * 1024 * 1024 * 1024},
		"trim":     {in: " 100 ", want: 100},
		"empty":    {in: "", want: 0},
		"unparsed": {in: "-", want: 0},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := number.ParseBytes(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

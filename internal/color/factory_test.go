package color

import (
	"testing"

	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	vmath "github.com/chrispotter/trace/internal/math"
)

func TestColorConfigFromYaml(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *ColorConfig
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no Config",
			Expected:    &ColorConfig{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: &ColorConfig{
				Color: vmath.Vector3d{
					X: 100.0,
					Y: 100.0,
					Z: 100.0,
				},
			},
			Bytes: []byte(`
    color: 
      - 100.0
      - 100.0
      - 100.0
`),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			colorConfig := &ColorConfig{}
			yaml, err := simpleyaml.NewYaml(test.Bytes)
			require.NoError(t, err)
			err = colorConfig.FromYaml(yaml)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, colorConfig)
		})
	}
}

func TestColorFactory(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    map[string]Color
		Configs     []*ColorConfig
	}{
		{
			Description: "Empty makes empty",
			Expected:    map[string]Color{},
			Configs:     []*ColorConfig{},
		},
		{
			Description: "one makes one",
			Expected: map[string]Color{
				"color1": &ColorValue{
					Name: "color1",
					Color: vmath.Vector3d{
						X: 100.0,
						Y: 100.0,
						Z: 100.0,
					},
				},
			},
			Configs: []*ColorConfig{
				&ColorConfig{
					Name: "color1",
					Color: vmath.Vector3d{
						X: 100.0,
						Y: 100.0,
						Z: 100.0,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			colors, err := ColorFactory(test.Configs)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, colors)
		})
	}
}

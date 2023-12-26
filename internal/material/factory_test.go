package material

import (
	"testing"

	"github.com/chrispotter/trace/internal/color"
	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaterialConfigFactory(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    []MaterialConfig
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no Config",
			Expected:    []MaterialConfig{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: []MaterialConfig{
				&LambertConfig{
					Colors: []color.Color{
						&color.ColorValue{},
						&color.ColorValue{},
					},
				},
			},
			Bytes: []byte(`
    lambert1:
      type: lambert
      color:
          - color1
          - color2
`),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			yaml, err := simpleyaml.NewYaml(test.Bytes)
			require.NoError(t, err)
			configs, err := ConfigFactory(yaml,
				map[string]color.Color{
					"color1": &color.ColorValue{},
					"color2": &color.ColorValue{},
				},
			)
			require.NoError(t, err)
			assert.Equal(t, len(test.Expected), len(configs))
		})
	}
}

func TestMaterialFactory(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    map[string]Material
		Configs     []MaterialConfig
	}{
		{
			Description: "Empty makes empty",
			Expected:    map[string]Material{},
			Configs:     []MaterialConfig{},
		},
		{
			Description: "one makes one",
			Expected: map[string]Material{
				"lambert1": &Lambert{
					Name:    "lambert1",
					Diffuse: &color.ColorValue{},
					Ambient: &color.ColorValue{},
				},
			},
			Configs: []MaterialConfig{
				&LambertConfig{
					Name: "Lambert",
					Colors: []color.Color{
						&color.ColorValue{},
						&color.ColorValue{},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			materials, err := Factory(test.Configs)
			require.NoError(t, err)
			// i don't care to test that the configs are correct, just that they
			// don't fail and have the correct amount
			assert.Equal(t, len(test.Expected), len(materials))
		})
	}
}

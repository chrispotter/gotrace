package material

import (
	"errors"
	"testing"

	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/lights"
	vmath "github.com/chrispotter/trace/internal/math"
	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLambertConfigFromYaml(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    MaterialConfig
		Config      []byte
		Colors      map[string]color.Color
		ExpectedErr error
	}{
		{
			Description: "Test Empty color map returns error",
			Expected:    &LambertConfig{},
			Config: []byte(`
    type: lambert
    color: 
      - lakersPurple 
      - lakersYellow
`),
			Colors:      map[string]color.Color{},
			ExpectedErr: errors.New("not enough colors"),
		},
		{
			Description: "Test Empty config color map returns error",
			Expected:    &LambertConfig{},
			Config:      []byte(``),
			Colors:      map[string]color.Color{"one": &color.ColorValue{}, "two": &color.ColorValue{}},
			ExpectedErr: errors.New("not enough colors in lambert config"),
		},
		{
			Description: "Test improper config color map returns error",
			Expected:    &LambertConfig{},
			Config: []byte(`
    type: lambert
    color: 
      - lakersYellow
`),
			Colors:      map[string]color.Color{"one": &color.ColorValue{}, "two": &color.ColorValue{}},
			ExpectedErr: errors.New("not enough colors in lambert config"),
		},
		{
			Description: "Test improper config color map returns error",
			Expected:    &LambertConfig{},
			Config: []byte(`
    type: lambert
    color: 
      - lakersPurple
      - one
`),
			Colors:      map[string]color.Color{"one": &color.ColorValue{}, "two": &color.ColorValue{}},
			ExpectedErr: errors.New("color lakersPurple does not exist in scene."),
		},
		{
			Description: "Happy Path",
			Expected: &LambertConfig{
				Colors: []color.Color{
					&color.ColorValue{},
					&color.ColorValue{},
				},
			},
			Config: []byte(`
    type: lambert
    color: 
      - two
      - one
`),
			Colors: map[string]color.Color{"one": &color.ColorValue{}, "two": &color.ColorValue{}},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			yaml, err := simpleyaml.NewYaml(test.Config)
			require.NoError(t, err)
			config := &LambertConfig{}
			err = config.FromYaml(yaml, test.Colors)
			if test.ExpectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, test.ExpectedErr, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.Expected, config)
			}
		})
	}
}

func TestLambertConfigNewMaterial(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *Lambert
		Config      *LambertConfig
	}{
		{
			Description: "Test Empty config color map returns error",
			Expected: &Lambert{
				Name: "Lambert",
				Ambient: &color.ColorValue{
					Name: "color1",
					Color: vmath.Vector3d{
						X: 1.0,
						Y: 1.0,
						Z: 1.0,
					},
				},
				Diffuse: &color.ColorValue{
					Name: "color2",
					Color: vmath.Vector3d{
						X: 2.0,
						Y: 2.0,
						Z: 2.0,
					},
				},
				SH: 1.0,
			},
			Config: &LambertConfig{
				Name: "Lambert",
				Colors: []color.Color{
					&color.ColorValue{
						Name: "color1",
						Color: vmath.Vector3d{
							X: 1.0,
							Y: 1.0,
							Z: 1.0,
						},
					},
					&color.ColorValue{
						Name: "color2",
						Color: vmath.Vector3d{
							X: 2.0,
							Y: 2.0,
							Z: 2.0,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			material, err := test.Config.NewMaterial()
			require.NoError(t, err)
			assert.Equal(t, test.Expected, material)
		})
	}
}

func TestLambertReturnColor(t *testing.T) {
	var tests = []struct {
		Description string
		Material    *Lambert
		Expected    vmath.Vector3d
		Angle       float64
	}{
		{
			Description: "Test 1 degree angle",
			Material: &Lambert{
				Ambient: &color.ColorValue{
					Name: "color1",
					Color: vmath.Vector3d{
						X: 1.0,
						Y: 1.0,
						Z: 1.0,
					},
				},
				Diffuse: &color.ColorValue{
					Name: "color2",
					Color: vmath.Vector3d{
						X: 2.0,
						Y: 2.0,
						Z: 2.0,
					},
				},
				SH: 1.0,
			},
			Angle: 1.0,
			Expected: vmath.Vector3d{
				X: 1.0,
				Y: 1.0,
				Z: 1.0,
			},
		},
		{
			Description: "test 45 degree angle",
			Material: &Lambert{
				Ambient: &color.ColorValue{
					Name: "color1",
					Color: vmath.Vector3d{
						X: 1.0,
						Y: 1.0,
						Z: 1.0,
					},
				},
				Diffuse: &color.ColorValue{
					Name: "color2",
					Color: vmath.Vector3d{
						X: 2.0,
						Y: 2.0,
						Z: 2.0,
					},
				},
				SH: 1.0,
			},
			Angle: .45,
			Expected: vmath.Vector3d{
				X: 1.0,
				Y: 1.0,
				Z: 1.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			color := test.Material.ReturnColor(test.Angle, 0, 0, &lights.PointLight{})
			assert.Equal(t, test.Expected, color)
		})
	}
}

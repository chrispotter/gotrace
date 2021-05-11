package color

import (
	"github.com/smallfish/simpleyaml"

	vmath "github.com/chrispotter/trace/internal/math"
)

// ColorConfig is a yaml definition of the Constructor to be read from
// util/scene.go
type ColorConfig struct {
	Name      string
	Color     vmath.Vector3d `yaml:"color"`
	ImagePath string         `yaml:"path"`
}

//FromYaml updates the CameraConfig with it's definition from the input yaml
func (cc *ColorConfig) FromYaml(config *simpleyaml.Yaml) error {
	if color, err := config.Get("color").Array(); err == nil {
		cc.Color = vmath.Vector3d{
			X: (color[0]).(float64),
			Y: (color[1]).(float64),
			Z: (color[2]).(float64),
		}
	}
	return nil
}

// ColorFactory returns a map of Colors from an array of ColorConfigs
func ColorFactory(configs []*ColorConfig) (map[string]Color, error) {
	colorMap := make(map[string]Color)
	for _, config := range configs {
		color := NewColorValue(config.Color)
		color.Name = config.Name
		colorMap[color.Name] = color
	}

	return colorMap, nil
}

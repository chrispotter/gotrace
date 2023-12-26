package camera

import (
	vmath "github.com/chrispotter/trace/internal/math"
	"github.com/smallfish/simpleyaml"
)

// Config is a yaml definition of the Constructor to be read from
// util/scene.go
type Config struct {
	Name     string
	Position vmath.Vector3d `yaml:"position"`
	Ratio    vmath.Vector2d `yaml:"ratio"`
}

// FromYaml updates the CameraConfig with it's definition from the input yaml
func (cc *Config) FromYaml(config *simpleyaml.Yaml) error {
	if position, err := config.Get("position").Array(); err == nil {
		cc.Position = vmath.Vector3d{
			X: (position[0]).(float64),
			Y: (position[1]).(float64),
			Z: (position[2]).(float64),
		}
	}
	if ratio, err := config.Get("ratio").Array(); err == nil {
		cc.Ratio = vmath.Vector2d{
			X: (ratio[0]).(float64),
			Y: (ratio[1]).(float64),
		}
	}
	return nil
}

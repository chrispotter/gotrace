package shapes

import (
	"github.com/chrispotter/trace/internal/common"
	"github.com/chrispotter/trace/internal/material"
	"github.com/smallfish/simpleyaml"
)

// ShapesConfig is an interface to define all configs able to provide to
// ShapeFactory
type ShapesConfig interface {
	FromYaml(*simpleyaml.Yaml, map[string]material.Material) error
	NewShape() (common.Traceable, error)
}

// ShapeConfigFactory generates configs for any shape
func ShapesConfigFactory(yaml *simpleyaml.Yaml, materials map[string]material.Material) ([]ShapesConfig, error) {
	configs := []ShapesConfig{}
	keys, err := yaml.GetMapKeys()
	if err != nil {
		return nil, err
	}
	for _, name := range keys {
		conf := yaml.Get(name)
		if t, err := conf.Get("type").String(); err == nil {
			switch t {
			case "sphere":
				sphereConfig := &SphereConfig{}
				err := sphereConfig.FromYaml(conf, materials)
				if err != nil {
					return nil, err
				}
				configs = append(configs, sphereConfig)
			case "plane":
				planeConfig := &PlaneConfig{}
				err := planeConfig.FromYaml(conf, materials)
				if err != nil {
					return nil, err
				}
				configs = append(configs, planeConfig)
			}
		}
	}

	return configs, nil
}

// ShapesFactory will make shapes according to type
func ShapesFactory(configs []ShapesConfig) ([]common.Traceable, error) {
	traceables := []common.Traceable{}
	for _, config := range configs {
		var traceable common.Traceable
		traceable, err := config.NewShape()
		if err != nil {
			return nil, err
		}
		traceables = append(traceables, traceable)
	}

	return traceables, nil
}

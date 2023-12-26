package material

import (
	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/lights"
	"github.com/smallfish/simpleyaml"
)

type Material interface {
	ReturnColor(angle float64,
		cam float64,
		ref float64,
		light lights.Light) color.Color
}

// MaterialConfig is a yaml definition of the Constructor to be read from
// util/scene.go
type MaterialConfig interface {
	FromYaml(*simpleyaml.Yaml, map[string]color.Color) error
	NewMaterial() (Material, error)
	GetName() string
}

// MaterialConfigFactory generates configs for any shape
func ConfigFactory(yaml *simpleyaml.Yaml, colors map[string]color.Color) ([]MaterialConfig, error) {
	configs := []MaterialConfig{}
	keys, err := yaml.GetMapKeys()
	// the only error that can happen is if yaml nil.
	if err != nil {
		return configs, nil
	}
	for _, name := range keys {
		conf := yaml.Get(name)
		if t, err := conf.Get("type").String(); err == nil {
			switch t {
			case "lambert":
				lambertConfig := &LambertConfig{}
				err := lambertConfig.FromYaml(conf, colors)
				if err != nil {
					return nil, err
				}
				lambertConfig.Name = name
				configs = append(configs, lambertConfig)
			case "cartoon":
				cartoonConfig := &CartoonConfig{}
				err := cartoonConfig.FromYaml(conf, colors)
				if err != nil {
					return nil, err
				}
				cartoonConfig.Name = name
				configs = append(configs, cartoonConfig)
			}
		}
	}

	return configs, nil
}

// Factory will make shapes according to type
func Factory(configs []MaterialConfig) (map[string]Material, error) {
	materials := map[string]Material{}
	for _, config := range configs {
		var material Material
		material, err := config.NewMaterial()
		if err != nil {
			return nil, err
		}
		materials[config.GetName()] = material
	}

	return materials, nil
}

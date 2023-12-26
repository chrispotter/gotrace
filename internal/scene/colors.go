package scene

import (
	"github.com/chrispotter/trace/internal/color"
	"github.com/smallfish/simpleyaml"
)

func (s *Scene) colors(y *simpleyaml.Yaml) error {
	if !y.Get("colors").IsFound() {
		return nil
	}

	yml := y.Get("colors")
	var colorConfigs []*color.ColorConfig
	keys, err := yml.GetMapKeys()
	if err != nil {
		return err
	}
	for _, name := range keys {
		conf := yml.Get(name)
		colorConfig := &color.ColorConfig{}
		err := colorConfig.FromYaml(conf)
		if err != nil {
			return err
		}
		colorConfig.Name = name
		colorConfigs = append(colorConfigs, colorConfig)
	}

	colors, err := color.ColorFactory(colorConfigs)
	if err != nil {
		return err
	}

	s.Colors = colors

	return nil
}

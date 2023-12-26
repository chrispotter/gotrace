package scene

import (
	"github.com/chrispotter/trace/internal/lights"
	"github.com/smallfish/simpleyaml"
)

func (s *Scene) lights(y *simpleyaml.Yaml) error {
	if !y.Get("lights").IsFound() {
		return nil
	}

	yml := y.Get("lights")
	lightConfigs, err := lights.LightsConfigFactory(yml, s.Colors)
	if err != nil {
		return err
	}

	l, err := lights.Factory(lightConfigs)
	if err != nil {
		return err
	}

	s.Renderables.Lights = l

	return nil
}

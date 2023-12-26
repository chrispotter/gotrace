package scene

import (
	"github.com/chrispotter/trace/internal/material"
	"github.com/smallfish/simpleyaml"
)

func (s *Scene) materials(y *simpleyaml.Yaml) error {
	if !y.Get("materials").IsFound() {
		return nil
	}

	yml := y.Get("materials")
	materialConfigs, err := material.ConfigFactory(yml, s.Colors)
	if err != nil {
		return err
	}

	materials, err := material.Factory(materialConfigs)
	if err != nil {
		return err
	}

	s.Materials = materials

	return nil
}

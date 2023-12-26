package scene

import (
	"github.com/chrispotter/trace/internal/common"
	"github.com/chrispotter/trace/internal/shapes"
	"github.com/smallfish/simpleyaml"
)

func (s *Scene) shapes(y *simpleyaml.Yaml) error {
	if !y.Get("shapes").IsFound() {
		return nil
	}

	yml := y.Get("shapes")
	shapeConfigs, err := shapes.ShapesConfigFactory(yml, s.Materials)
	if err != nil {
		return err
	}

	sh, err := shapes.ShapesFactory(shapeConfigs)
	if err != nil {
		return err
	}

	s.Renderables = &common.RenderableObjects{
		Shapes: sh,
	}

	return nil
}

package scene

import (
	"github.com/chrispotter/trace/internal/camera"
	"github.com/smallfish/simpleyaml"
)

func (s *Scene) cameras(y *simpleyaml.Yaml) error {
	if !y.Get("cameras").IsFound() {
		return nil
	}

	yml := y.Get("cameras")
	var cameraConfigs []*camera.Config
	keys, err := yml.GetMapKeys()
	if err != nil {
		return err
	}
	for _, name := range keys {
		conf := yml.Get(name)
		cameraConfig := &camera.Config{}
		err := cameraConfig.FromYaml(conf)
		if err != nil {
			return err
		}
		cameraConfig.Name = name
		cameraConfigs = append(cameraConfigs, cameraConfig)
	}
	s.Cameras, err = camera.Factory(cameraConfigs)
	if err != nil {
		return err
	}

	return nil
}

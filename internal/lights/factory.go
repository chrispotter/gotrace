package lights

import (
	"fmt"

	"github.com/chrispotter/trace/internal/color"
	"github.com/smallfish/simpleyaml"
)

// ShapeConfigFactory generates configs for any shape
func LightsConfigFactory(yaml *simpleyaml.Yaml, colors map[string]color.Color) ([]LightsConfig, error) {
	configs := []LightsConfig{}
	keys, err := yaml.GetMapKeys()
	if err != nil {
		return nil, err
	}
	for _, name := range keys {
		conf := yaml.Get(name)
		if t, err := conf.Get("type").String(); err == nil {
			switch t {
			case "directional":
				directionalLightConfig := &DirectionalLightConfig{}
				err := directionalLightConfig.FromYaml(conf, colors)
				if err != nil {
					return nil, err
				}
				fmt.Printf("%+v\n", directionalLightConfig)
				directionalLightConfig.Name = name
				configs = append(configs, directionalLightConfig)
			}
		}
	}

	return configs, nil
}

// Factory returns an array of Lights from an array of Light Configs
func Factory(configs []LightsConfig) ([]Light, error) {
	lightsMap := []Light{}
	for _, config := range configs {
		light, err := config.NewLight()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v\n", light)
		light.SetName(config.GetName())
		lightsMap = append(lightsMap, light)
	}

	return lightsMap, nil
}

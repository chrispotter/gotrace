package lights

import (
	"errors"
	"fmt"

	"github.com/chrispotter/trace/internal/color"
	vmath "github.com/chrispotter/trace/internal/math"
	"github.com/smallfish/simpleyaml"
)

type DirectionalLightConfig struct {
	Name  string
	View  vmath.Vector3d
	Color color.Color
}

func (dlc *DirectionalLightConfig) FromYaml(config *simpleyaml.Yaml, colors map[string]color.Color) error {
	if view, err := config.Get("view").Array(); err == nil {
		dlc.View = vmath.Vector3d{
			X: (view[0]).(float64),
			Y: (view[1]).(float64),
			Z: (view[2]).(float64),
		}
	}

	if configColor, err := config.Get("color").String(); err == nil {
		c, ok := colors[configColor]
		if ok {
			dlc.Color = c
		} else {
			return errors.New(fmt.Sprintf("color %s does not exist in scene.", configColor))
		}
	} else {
		return errors.New("Color missing from light configuration")
	}

	return nil
}

func (dlc *DirectionalLightConfig) NewLight() (Light, error) {
	position := vmath.Vector3d{0.0, 0.0, 0.0}.Cross(dlc.View)
	return &DirectionalLight{
		V:         dlc.View,
		P:         position,
		Color:     dlc.Color,
		Intensity: 1.0,
	}, nil
}

func (dlc *DirectionalLightConfig) GetName() string {
	return dlc.Name
}

type DirectionalLight struct {
	Name      string
	P         vmath.Vector3d
	V         vmath.Vector3d
	Intensity float64
	Color     color.Color
}

func (dl *DirectionalLight) Intersect(ray vmath.Vector3d) bool {
	return true
}

func (dl *DirectionalLight) ReturnLightVector(hit vmath.Vector3d) vmath.Vector3d {
	return dl.V.SMultiply(-1)
}

func (dl *DirectionalLight) GetColor() color.Color {
	return dl.Color.SMultiply(dl.Intensity)
}

func (dl *DirectionalLight) SetName(name string) {
	dl.Name = name
}

func (dl *DirectionalLight) SetPH(hit vmath.Vector3d) {}

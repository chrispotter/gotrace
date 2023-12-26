package material

import (
	"errors"
	"fmt"
	"math"

	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/lights"
)

// LambertConfig defines a sphere for the ShapeFactory
type LambertConfig struct {
	Name   string
	Colors []color.Color
}

func (lc *LambertConfig) GetName() string {
	return lc.Name
}

// NewShape generates a Shape from the config object
// satisfies the MaterialConfig interface  (1/2)
func (lc *LambertConfig) NewMaterial() (Material, error) {
	return &Lambert{
		Name:    lc.Name,
		Ambient: lc.Colors[0],
		Diffuse: lc.Colors[1],
		SH:      1.0,
	}, nil
}

// FromYaml generates Config from input yaml
// satisfies the interface ShapesConfig (2/2)
func (lc *LambertConfig) FromYaml(config *simpleyaml.Yaml, colors map[string]color.Color) error {
	if len(colors) < 2 {
		return errors.New("not enough colors")
	}
	if configColors, err := config.Get("color").Array(); err == nil {
		if len(configColors) != 2 {
			return errors.New("not enough colors in lambert config")
		}

		// color[0] is ambient
		// color[1] is diffuse
		for _, colorName := range configColors {
			c, ok := colors[colorName.(string)]
			if ok {
				lc.Colors = append(lc.Colors, c)
			} else {
				return errors.New(fmt.Sprintf("color %s does not exist in scene.", colorName.(string)))
			}
		}
	} else if err != nil {
		return errors.New("not enough colors in lambert config")
	}

	return nil
}

// Lambert is a default flat shader with a Ambient and Diffuse color
type Lambert struct {
	Name                                             string
	Ambient, Diffuse                                 color.Color
	LightIndex, DistanceLightHit, U, V, N, SH        float64
	Reflect, Iridesent, Refract, Glossy, Transparent bool
}

func (l *Lambert) ReturnColor(angle float64, cam float64, ref float64, light lights.Light) color.Color {
	beta := 1.0

	c := ((angle + 1.0) / 2.0)
	if c > .55 {
		c = .55
	}
	if c < .45 {
		c = .45
	}
	d := (c - .45) / (.55 - .45)

	e := (3 * math.Pow(d, 2)) - (2 * math.Pow(d, 3))
	if e > 1 {
		e = 0
	}
	if e < 0 {
		e = 1
	}

	ca := math.Pow(e, beta)

	//(l.diffuse.ReturnColor(u, v)*(1.0-ca) + l.ambient.ReturnColor(u, v)*ca) * sh
	matColor := l.Diffuse.GetColor(l.U, l.V).SMultiply(1.0 - ca).Add(l.Ambient.GetColor(l.U, l.V).SMultiply(ca)).SMultiply(l.SH)

	matColor = matColor.Add(light.GetColor().GetColor(0, 0))

	return color.NewColorValue(matColor)
}

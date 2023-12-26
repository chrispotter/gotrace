package material

import (
	"errors"
	"fmt"

	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/lights"
)

// LambertConfig defines a sphere for the ShapeFactory
type CartoonConfig struct {
	Name     string
	Segments int
	Colors   []color.Color
}

func (cc *CartoonConfig) GetName() string {
	return cc.Name
}

// NewShape generates a Shape from the config object
// satisfies the MaterialConfig interface  (1/2)
func (cc *CartoonConfig) NewMaterial() (Material, error) {
	return &Cartoon{
		Name:     cc.Name,
		Ambient:  cc.Colors[0],
		Diffuse:  cc.Colors[1],
		Specular: cc.Colors[2],
		Outline:  cc.Colors[3],
		Segments: cc.Segments,
		SH:       1.0,
	}, nil
}

// FromYaml generates Config from input yaml
// satisfies the interface ShapesConfig (2/2)
func (cc *CartoonConfig) FromYaml(config *simpleyaml.Yaml, colors map[string]color.Color) error {
	if segments, err := config.Get("segments").Int(); err == nil {
		cc.Segments = segments
	} else {
		cc.Segments = 3
	}
	if len(colors) < 2 {
		return errors.New("not enough colors for cartoon shader")
	}
	if configColors, err := config.Get("color").Array(); err == nil {
		if len(configColors) != 4 {
			return errors.New("not enough colors in cartoon config")
		}

		// color[0] is ambient
		// color[1] is diffuse
		// color[2] is specular
		// color[3] is outline
		for _, colorName := range configColors {
			c, ok := colors[colorName.(string)]
			if ok {
				cc.Colors = append(cc.Colors, c)
			} else {
				return errors.New(fmt.Sprintf("color %s does not exist in scene.", colorName.(string)))
			}
		}
	} else if err != nil {
		return errors.New("not enough colors in cartoon config")
	}

	return nil
}

// Cartoon is a default flat shader with a Ambient, Diffuse, Specular, and
// Outline color
type Cartoon struct {
	Name                                             string
	Ambient, Diffuse, Specular, Outline              color.Color
	LightIndex, DistanceLightHit, U, V, N, SH        float64
	Reflect, Iridesent, Refract, Glossy, Transparent bool
	Segments                                         int
}

func (c *Cartoon) ReturnColor(angle float64, cam float64, ref float64, light lights.Light) color.Color {
	normalizedAngle := ((angle + 1.0) / 2.0)

	if cam < .5 {
		fmt.Println(cam)
	}
	// split Ambient and Diffuse colors in X number of Segments
	fract := normalizedAngle * float64(c.Segments)
	rest := int(fract) % c.Segments

	// set fraction of the color blend between ambient and diffuse
	matColor := c.Ambient.GetColor(c.U, c.V).SMultiply(float64(rest) / float64(c.Segments)).
		Add(c.Diffuse.GetColor(c.U, c.V).SMultiply(float64(c.Segments-rest) / float64(c.Segments)))
		/*
			      int seg2 = 4;
			      float s = (ref+1.0)/2.0;
			      float sa = pow(s,30);
			      int fract2 = sa * seg2;
			      int rest2 = fract2 % seg2;

			      matColor = specular->rtnColor() * (float)rest2/(float)seg2 +
							 matColor * (float)(seg2 - abs(rest2))/(float)seg2;

			      int seg3 = 4;
			      float sha = pow(s,30);
			      int fract3 = sha * seg3;
			      int rest3 = fract3 % seg3;

			      matColor = matColor ^light->getColor(light->getPH())* light ->getRatio();
		*/
	if cam <= 0.6 {
		matColor = c.Outline.GetColor(c.U, c.V)
	}

	//matColor = matColor.Add(light.GetColor().GetColor(0, 0))

	return color.NewColorValue(matColor)
}

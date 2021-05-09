package shapes

import (
	"fmt"
	"math"

	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/common"
	vmath "github.com/chrispotter/trace/internal/math"
)

// ShapesConfig is an interface to define all configs able to provide to
// ShapeFactory
type ShapesConfig interface {
	FromYaml(config *simpleyaml.Yaml) error
	NewShape() (common.Traceable, error)
}

//ShapeConfigFactory generates configs for any shape
func ShapesConfigFactory(yaml *simpleyaml.Yaml) ([]ShapesConfig, error) {
	configs := []ShapesConfig{}
	keys, err := yaml.GetMapKeys()
	if err != nil {
		return nil, err
	}
	for _, name := range keys {
		conf := yaml.Get(name)
		if t, err := conf.Get("type").String(); err == nil {
			switch t {
			case "sphere":
				sphereConfig := &SphereConfig{}
				err := sphereConfig.FromYaml(conf)
				if err != nil {
					return nil, err
				}
				configs = append(configs, sphereConfig)
			}
		}
	}

	return configs, nil
}

// ShapesFactory will make shapes according to type
func ShapesFactory(configs []ShapesConfig) ([]common.Traceable, error) {
	traceables := []common.Traceable{}
	for _, config := range configs {
		var traceable common.Traceable
		traceable, err := config.NewShape()
		if err != nil {
			return nil, err
		}
		traceables = append(traceables, traceable)
	}

	return traceables, nil
}

// SphereConfig defines a sphere for the ShapeFactory
type SphereConfig struct {
	Name     string
	Radius   float64
	Position vmath.Vector3d
}

// NewShape generates a Shape from the config object
// satisfies the interface ShapesConfig (1/2)
func (sc *SphereConfig) NewShape() (common.Traceable, error) {
	sphere := NewSphere(sc.Position, sc.Radius)
	sphere.Name = sc.Name
	return sphere, nil
}

// FromYaml generates Config from input yaml
// satisfies the interface ShapesConfig (2/2)
func (sc *SphereConfig) FromYaml(config *simpleyaml.Yaml) error {
	if position, err := config.Get("position").Array(); err == nil {
		sc.Position = vmath.Vector3d{
			X: (position[0]).(float64),
			Y: (position[1]).(float64),
			Z: (position[2]).(float64),
		}
	}
	if radius, err := config.Get("radius").Float(); err == nil {
		sc.Radius = radius
	}
	return nil
}

// Vector3d position,lposition;
//Vector3d axis[3];
//float s[
//float a[5];
//Material *mat;
//Vector3d ph;
//int sample;
// float e;
type Sphere struct {
	P        vmath.Vector3d
	axis     []vmath.Vector3d
	radius   float64
	PlaceHit vmath.Vector3d
	s        []float64
	a        []float64
	Name     string
}

func NewSphere(pos vmath.Vector3d, rad float64) *Sphere {
	return &Sphere{
		P:      pos,
		radius: rad,
		axis: []vmath.Vector3d{
			vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
			vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
			vmath.Vector3d{X: 0.0, Y: 0.0, Z: 1.0},
		},
		s: []float64{rad, rad, rad},
		a: []float64{1.0, 1.0, 1.0, 0.0, -1.0},
	}
}

// Intersect satisfies the qualifications for
// Render object interface for a scene
func (s *Sphere) Intersect(ray *vmath.Ray) bool {
	var a, b, c = 0.0, 0.0, 0.0
	for index, axis := range s.axis {
		a += s.a[index] * math.Pow(axis.Dot(ray.Direction)/s.s[index], 2)
		b += s.a[index] * axis.Dot(ray.Direction) * axis.Dot(ray.Origin.Subtract(s.P)) * 2.0 / math.Pow(s.s[index], 2)
		fmt.Printf("%+v\n", ray.Origin.Subtract(s.P))
		c += s.a[index]*math.Pow(axis.Dot(ray.Origin.Subtract(s.P))/s.s[index], 2) + s.a[4]
	}

	d := math.Pow(b, 2) - 4.0*a*c

	return (-b - math.Sqrt(d)/(2.0*a)) > 0
}

// GetPosition satisfies requirements for Object
// interface for a scene
func (s *Sphere) GetPosition() vmath.Vector3d {
	return s.P
}

func (s *Sphere) GetName() string {
	return "shape"
}

// GetType satisfies requirements for Object
// interface for a scene
func (s *Sphere) GetType() string {
	return "sphere"
}

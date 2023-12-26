package shapes

import (
	"errors"
	"fmt"
	"math"

	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/common"
	"github.com/chrispotter/trace/internal/material"
	vmath "github.com/chrispotter/trace/internal/math"
)

// SphereConfig defines a sphere for the ShapeFactory
type SphereConfig struct {
	Name     string
	Radius   float64
	Position vmath.Vector3d
	Material material.Material
}

// NewShape generates a Shape from the config object
// satisfies the interface ShapesConfig (1/2)
func (sc *SphereConfig) NewShape() (common.Traceable, error) {
	sphere := NewSphere(sc.Position, sc.Radius)
	sphere.Name = sc.Name
	sphere.Material = sc.Material
	return sphere, nil
}

// FromYaml generates Config from input yaml
// satisfies the interface ShapesConfig (2/2)
func (sc *SphereConfig) FromYaml(config *simpleyaml.Yaml, materials map[string]material.Material) error {
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
	if material, err := config.Get("material").String(); err == nil {

		m, ok := materials[material]
		if ok {
			sc.Material = m
		} else {
			return errors.New(fmt.Sprintf("material %s does not exist in scene.", material))
		}
	}
	return nil
}

// Vector3d position,lposition;
// Vector3d axis[3];
// float s[
// float a[5];
// Material *mat;
// Vector3d ph;
// int sample;
// float e;
type Sphere struct {
	Name              string
	axis              []vmath.Vector3d
	s                 []float64
	a                 []float64
	PlaceHit          vmath.Vector3d
	P                 vmath.Vector3d
	intersectionRatio float64
	radius            float64
	delta             float64

	Material material.Material
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
	a, b, c := 0.0, 0.0, 0.0
	for index, axis := range s.axis {
		a += s.a[index] * math.Pow(axis.Dot(ray.Direction)/s.s[index], 2)
		b += s.a[index] * axis.Dot(ray.Direction) * axis.Dot(ray.Origin.Subtract(s.P)) * 2.0 / math.Pow(s.s[index], 2)
		c += s.a[index]*math.Pow(axis.Dot(ray.Origin.Subtract(s.P))/s.s[index], 2) + s.a[4]
	}

	d := math.Pow(b, 2) - 4.0*a*c

	quadPlus := (-b + math.Sqrt(d)/(2.0*a))
	quadMinus := (-b - math.Sqrt(d)/(2.0*a))

	s.intersectionRatio = quadMinus
	if quadPlus < quadMinus || quadMinus < 0 {
		s.intersectionRatio = quadPlus
	}
	if s.intersectionRatio > 0 && d > 0 {
		s.PlaceHit = ray.Origin.Add(ray.Direction.SMultiply(s.intersectionRatio))
		s.delta = d
		return true
	}
	return false
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

// GetIntersectionRatio
func (s *Sphere) GetIntersectionRatio() float64 {
	return s.intersectionRatio
}

func (s *Sphere) CalculateNorm(hit vmath.Vector3d) vmath.Vector3d {
	grad := vmath.Vector3d{0.0, 0.0, 0.0}
	for index, axis := range s.axis {
		//(2.0*a[i]*(axis[i] * (hit - getPosition()))/pow(s[i],2))*axis[i]
		grad = grad.Add(axis.SMultiply(2.0 * s.a[index] * axis.Dot(hit.Subtract(s.P)) / math.Pow(s.s[index], 2)))
	}
	grad.Normalize()

	return grad
}

func (s *Sphere) ReturnColor(ray *vmath.Ray, objs *common.RenderableObjects) color.Color {
	color := &color.ColorValue{
		Color: vmath.Vector3d{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	nh := s.CalculateNorm(s.PlaceHit)     //normalized surface normal
	nc := ray.Origin.Subtract(s.PlaceHit) //normalized camera and ph vector
	nc.Normalize()
	cameraAngle := nh.Dot(nc) //angle between camera and surface normal
	for _, light := range objs.Lights {
		nlh := light.ReturnLightVector(s.PlaceHit) // normalized light direction
		lightAngle := nh.Dot(nlh)
		color.Add(s.Material.ReturnColor(lightAngle, cameraAngle, 0, light))
	}

	return color
}

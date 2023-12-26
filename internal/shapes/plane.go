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
type PlaneConfig struct {
	Name     string
	Position vmath.Vector3d
	Normal   vmath.Vector3d
	Material material.Material
}

// NewShape generates a Shape from the config object
// satisfies the interface ShapesConfig (1/2)
func (pc *PlaneConfig) NewShape() (common.Traceable, error) {
	plane := NewPlane(pc.Position, pc.Normal)
	plane.Name = pc.Name
	plane.Material = pc.Material
	return plane, nil
}

// FromYaml generates Config from input yaml
// satisfies the interface ShapesConfig (2/2)
func (pc *PlaneConfig) FromYaml(config *simpleyaml.Yaml, materials map[string]material.Material) error {
	if position, err := config.Get("position").Array(); err == nil {
		pc.Position = vmath.Vector3d{
			X: (position[0]).(float64),
			Y: (position[1]).(float64),
			Z: (position[2]).(float64),
		}
	}
	if normal, err := config.Get("normal").Array(); err == nil {
		pc.Normal = vmath.Vector3d{
			X: (normal[0]).(float64),
			Y: (normal[1]).(float64),
			Z: (normal[2]).(float64),
		}
	}
	if material, err := config.Get("material").String(); err == nil {

		m, ok := materials[material]
		if ok {
			pc.Material = m
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
type Plane struct {
	Name              string
	axis              []vmath.Vector3d
	s                 []float64
	a                 []float64
	PlaceHit          vmath.Vector3d
	P                 vmath.Vector3d
	intersectionRatio float64
	delta             float64

	Material material.Material
}

func NewPlane(pos vmath.Vector3d, normal vmath.Vector3d) *Plane {
	plane := &Plane{
		P:    pos,
		axis: []vmath.Vector3d{},
		a:    []float64{0.0, 0.0, 0.0, 1.0, 0.0},
	}
	// build 3 axis from supplied normal
	normal.Normalize()
	xaxis := normal.Cross(vmath.Vector3d{1.0, 0.0, 0.0})
	if normal.Equals(vmath.Vector3d{1.0, 0.0, 0.0}) {
		xaxis = normal.Cross(vmath.Vector3d{0.0, 1.0, 0.0})
	}
	yaxis := xaxis.Cross(normal)
	yaxis.Normalize()
	xaxis = normal.Cross(yaxis)
	xaxis.Normalize()

	plane.axis = []vmath.Vector3d{xaxis, yaxis, normal}

	return plane

}

// Intersect satisfies the qualifications for
// Render object interface for a scene
func (p *Plane) Intersect(ray *vmath.Ray) bool {
	// the ray is parallel if the dot of the z axis and the direction are 0
	if p.axis[2].Dot(ray.Direction) == 0 {
		return false
	}

	t := p.axis[2].Dot(p.P.Subtract(ray.Origin)) / (p.axis[2].Dot(ray.Direction))
	if t < 0 || math.IsNaN(t) {
		return false
	}

	p.PlaceHit = ray.Origin.Add(ray.Direction.SMultiply(t))
	return true
}

// GetPosition satisfies requirements for Object
// interface for a scene
func (p *Plane) GetPosition() vmath.Vector3d {
	return p.P
}

func (p *Plane) GetName() string {
	return p.Name
}

// GetType satisfies requirements for Object
// interface for a scene
func (p *Plane) GetType() string {
	return "plane"
}

func (p *Plane) CalculateNorm(hit vmath.Vector3d) vmath.Vector3d {
	ds := vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0}
	norm := p.axis[2].SMultiply(ds.Z).Add(p.axis[0].SMultiply(ds.X)).Add(p.axis[1].SMultiply(ds.Y))
	norm.Normalize()
	return norm
}

func (p *Plane) ReturnColor(ray *vmath.Ray, objs *common.RenderableObjects) color.Color {
	color := &color.ColorValue{
		Color: vmath.Vector3d{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	nh := p.CalculateNorm(p.PlaceHit)     //normalized surface normal
	nc := ray.Origin.Subtract(p.PlaceHit) //normalized camera and ph vector
	nc.Normalize()
	cameraAngle := nh.Dot(nc) //angle between camera and surface normal
	for _, light := range objs.Lights {
		nlh := light.ReturnLightVector(p.PlaceHit) // normalized light direction
		lightAngle := nh.Dot(nlh)
		color.Add(p.Material.ReturnColor(lightAngle, cameraAngle, 0, light))
	}

	return color
}

func (p *Plane) GetIntersectionRatio() float64 {
	return 0.0
}

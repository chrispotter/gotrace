package internal

import (
	vmath "github.com/chrispotter/trace/internal/math"
)

type ColorValue interface {
	GetColor(u float64, v float64) vmath.Vector3d
}

type VectorValue struct {
	Color vmath.Vector3d
}

func NewVectorValue(c vmath.Vector3d) *VectorValue {
	return &VectorValue{
		Color: c,
	}
}

func (vv *VectorValue) GetColor(u float64, v float64) vmath.Vector3d {
	return vv.Color
}

type Color struct {
	Ambient, Diffuse, Specular, Outline ColorValue
}

// NewColor expects up to 4 colors expecting in order, Ambient, Diffuse,
// Specular, Outline
func NewColor(colors ...ColorValue) *Color {
	var c *Color
	if len(colors) > 0 {
		c = &Color{
			Ambient: colors[0],
			Diffuse: colors[1],
		}
	}

	if len(colors) > 2 {
		c.Specular = colors[2]
	}

	if len(colors) > 3 {
		c.Outline = colors[0]
	}

	return c
}

// ReturnAmbient returns ambient vector3d
func (c *Color) ReturnAmbient(u float64, v float64) vmath.Vector3d {
	return c.Ambient.GetColor(u, v)
}

// ReturnDiffuse returns diffuse vector3d
func (c *Color) ReturnDiffuse(u float64, v float64) vmath.Vector3d {
	return c.Diffuse.GetColor(u, v)
}

// ReturnSpecular returns specular vector3d
func (c *Color) ReturnSpecular(u float64, v float64) vmath.Vector3d {
	return c.Specular.GetColor(u, v)
}

// ReturnOutline returns outline vector3d
func (c *Color) ReturnOutline(u float64, v float64) vmath.Vector3d {
	return c.Outline.GetColor(u, v)
}

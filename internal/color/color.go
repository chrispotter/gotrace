package color

import (
	vmath "github.com/chrispotter/trace/internal/math"
)

// Color is an interface that defines RGBA values for an object.
type Color interface {
	GetColor(u float64, v float64) vmath.Vector3d
}

/*
type ImageValue struct {
	Image vmath.Vector3d
}

func NewImageValue(c vmath.Vector3d) *ImageValue {
	return &ImageValue{
		Image: c,
	}
}

func (iv *ImageValue) GetColor(u float64, v float64) vmath.Vector3d {
	return iv.Image
} */

type ColorValue struct {
	Name  string
	Color vmath.Vector3d
}

// NewColorValue constructor for Color value
func NewColorValue(color vmath.Vector3d) *ColorValue {
	return &ColorValue{
		Color: color,
	}
}

// GetColor returns the color value for a ColorValue
// Satisfies Color interface
func (cv *ColorValue) GetColor(u float64, v float64) vmath.Vector3d {
	return cv.Color
}

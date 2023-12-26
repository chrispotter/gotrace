package common

import (
	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/lights"
	vmath "github.com/chrispotter/trace/internal/math"
)

type Traceable interface {
	Intersect(cast *vmath.Ray) bool
	GetIntersectionRatio() float64
	ReturnColor(*vmath.Ray, *RenderableObjects) color.Color
}

type Object interface {
	GetPosition() vmath.Vector3d
	GetName() string
	GetType() string
}

type RenderableObjects struct {
	Shapes []Traceable
	Lights []lights.Light
}

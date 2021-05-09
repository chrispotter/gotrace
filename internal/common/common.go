package common

import (
	vmath "github.com/chrispotter/trace/internal/math"
)

type Traceable interface {
	Intersect(cast *vmath.Ray) bool
}

type Object interface {
	GetPosition() vmath.Vector3d
	GetName() string
	GetType() string
}

type Light interface {
}

//type Renderable interface {
//	GetMaterial() Material
//}
//
type Material interface {
	ReturnColor(angle float64,
		cam float64,
		ref float64,
		light Light) vmath.Vector3d
}

type RenderableObjects struct {
	Shapes []Traceable
	Lights []Light
}

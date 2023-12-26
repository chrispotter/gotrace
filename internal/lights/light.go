package lights

import (
	"github.com/chrispotter/trace/internal/color"
	vmath "github.com/chrispotter/trace/internal/math"
	"github.com/smallfish/simpleyaml"
)

type Light interface {
	Intersect(vmath.Vector3d) bool
	ReturnLightVector(vmath.Vector3d) vmath.Vector3d
	GetColor() color.Color
	SetName(string)
	SetPH(vmath.Vector3d)
}

// LightsConfig is an interface to define all configs able to provide to
// LightsFactory
type LightsConfig interface {
	FromYaml(*simpleyaml.Yaml, map[string]color.Color) error
	NewLight() (Light, error)
	GetName() string
}

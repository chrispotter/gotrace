package shapes

import (
	"testing"

	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	vmath "github.com/chrispotter/trace/internal/math"
)

func TestSphereConfigFromYaml(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    interface{}
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no Config",
			Expected:    &SphereConfig{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: &SphereConfig{
				Position: vmath.Vector3d{
					X: 0.0,
					Y: 1.0,
					Z: -25.0,
				},
				Radius: 0.5,
			},
			Bytes: []byte(`
    position: 
      - 0.0
      - 1.0
      - -25.0
    radius: 0.5
`),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			sphereConfig := &SphereConfig{}
			yaml, err := simpleyaml.NewYaml(test.Bytes)
			require.NoError(t, err)
			err = sphereConfig.FromYaml(yaml)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, sphereConfig)
		})
	}
}

func TestSphereNew(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *Sphere
		Position    vmath.Vector3d
		Radius      float64
	}{
		{
			Description: "Test Simple",
			Expected: &Sphere{
				P:      vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
				radius: 1.0,
				axis: []vmath.Vector3d{
					vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
					vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
					vmath.Vector3d{X: 0.0, Y: 0.0, Z: 1.0},
				},
				s: []float64{1.0, 1.0, 1.0},
				a: []float64{1.0, 1.0, 1.0, 0.0, -1.0},
			},
			Position: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			Radius:   1.0,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			sphere := NewSphere(test.Position, test.Radius)
			assert.Equal(t, test.Expected, sphere)
		})
	}
}

func TestSphereIntersect(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    bool
		Ray         *vmath.Ray
	}{
		{
			Description: "Test Simple Hit",
			Expected:    true,
			Ray: &vmath.Ray{
				Origin: vmath.Vector3d{
					X: 0.0,
					Y: 0.0,
					Z: 10.0,
				},
				Direction: vmath.Vector3d{
					X: 0.0,
					Y: 0.0,
					Z: -1.0,
				},
			},
		},
		{
			Description: "Test Simple Miss",
			Expected:    false,
			Ray: &vmath.Ray{
				Origin: vmath.Vector3d{
					X: 0.0,
					Y: 0.0,
					Z: 10.0,
				},
				Direction: vmath.Vector3d{
					X: 0.0,
					Y: 0.0,
					Z: 1.0,
				},
			},
		},
		{
			Description: "Test Complex Hit",
			Expected:    true,
			Ray: &vmath.Ray{
				Origin: vmath.Vector3d{
					X: 0.0,
					Y: 0.0,
					Z: 10.0,
				},
				Direction: vmath.Vector3d{
					X: -0.10,
					Y: -0.2,
					Z: -10.0,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			sphere := NewSphere(vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0}, 1.0)
			err := test.Ray.Direction.Normalize()
			require.NoError(t, err)
			hit := sphere.Intersect(test.Ray)
			assert.Equal(t, test.Expected, hit)
		})
	}
}

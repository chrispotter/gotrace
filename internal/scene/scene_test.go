package scene

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chrispotter/trace/internal/camera"
	vmath "github.com/chrispotter/trace/internal/math"
)

func TestSceneFromYaml(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *Scene
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no scene",
			Expected:    &Scene{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: &Scene{
				Cameras: []*camera.Camera{
					&camera.Camera{
						Name:  "camera1",
						P:     vmath.Vector3d{X: 0.0, Y: 1.0, Z: -25.0},
						N:     vmath.Vector3d{X: 0.0, Y: 0.0, Z: -1.0},
						U:     vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
						V:     vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
						XMax:  1280.0,
						YMax:  720.0,
						Depth: 3.0,
						Sx:    1.0,
						Sy:    (720 * 1.0) / 1280,
					},
				},
			},
			Bytes: []byte(`
cameras:
  camera1:
    position: 
      - 0.0
      - 1.0
      - -25.0
    ratio: 
      - 1280.0
      - 720.0
`),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			scene := &Scene{}
			err := scene.FromYaml(test.Bytes)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, scene)
		})
	}
}

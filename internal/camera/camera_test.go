package camera

import (
	"testing"

	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	vmath "github.com/chrispotter/trace/internal/math"
)

func TestCameraConfigFromYaml(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    interface{}
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no Config",
			Expected:    &CameraConfig{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: &CameraConfig{
				Position: vmath.Vector3d{
					X: 0.0,
					Y: 1.0,
					Z: -25.0,
				},
				Ratio: vmath.Vector2d{
					X: 1280.0,
					Y: 720.0,
				},
			},
			Bytes: []byte(`
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
			cameraConfig := &CameraConfig{}
			yaml, err := simpleyaml.NewYaml(test.Bytes)
			require.NoError(t, err)
			err = cameraConfig.FromYaml(yaml)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, cameraConfig)
		})
	}
}

func TestCameraGetImagePosition(t *testing.T) {
	var tests = []struct {
		Description    string
		Expected       vmath.Vector3d
		CameraPosition vmath.Vector3d
		depth          float64
	}{
		{
			Description:    "Simple Image Position",
			CameraPosition: vmath.Vector3d{},
			Expected:       vmath.Vector3d{},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			camera, err := NewCamera(test.CameraPosition, vmath.Vector2d{X: 1, Y: 1})
			require.NoError(t, err)
			actualPosition := camera.getImagePosition()
			assert.Equal(t, test.Expected, actualPosition)
		})
	}
}

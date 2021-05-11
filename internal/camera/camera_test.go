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

func TestCameraNew(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *Camera
		Position    vmath.Vector3d
		Ratio       vmath.Vector2d
	}{
		{
			Description: "Empty constructor",
			Expected: &Camera{
				P:     vmath.Vector3d{X: 0, Y: 0, Z: 0},
				U:     vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
				V:     vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
				N:     vmath.Vector3d{X: 0.0, Y: 0.0, Z: 1.0},
				XMax:  1,
				YMax:  1,
				Depth: 1.0,
				Sx:    1.0,
				Sy:    1.0,
			},
		},
		{
			Description: "Camera with positon and ratio",
			Expected: &Camera{
				P:     vmath.Vector3d{X: 1, Y: 3, Z: 5},
				U:     vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
				V:     vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
				N:     vmath.Vector3d{X: 0, Y: 0, Z: 1},
				XMax:  100,
				YMax:  100,
				Depth: 1.0,
				Sx:    1.0,
				Sy:    1.0,
			},
			Position: vmath.Vector3d{X: 1, Y: 3, Z: 5},
			Ratio:    vmath.Vector2d{X: 100, Y: 100},
		},
	}
	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			camera, err := NewCamera(test.Position, test.Ratio)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, camera)
		})
	}
}

func TestCameraLookAt(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    *Camera
		Camera      *Camera
		Look        vmath.Vector3d
		Up          vmath.Vector3d
	}{
		{
			Description: "Zero LookAt",
			Expected: &Camera{
				P: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
				N: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 1.0},
				U: vmath.Vector3d{X: 1.0, Y: 0.0, Z: 0.0},
				V: vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
			},
			Camera: &Camera{
				P: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			},
			Look: vmath.Vector3d{X: 0, Y: 0, Z: -1},
			Up:   vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			err := test.Camera.LookAt(test.Look, test.Up)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, test.Camera)
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
			Expected:       vmath.Vector3d{X: 0, Y: 0, Z: -1},
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

func TestCameraGetOriginPixel(t *testing.T) {
	var tests = []struct {
		Description    string
		Expected       vmath.Vector3d
		CameraPosition vmath.Vector3d
	}{
		{
			Description:    "Simple GetOriginPixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			Expected:       vmath.Vector3d{X: -0.5, Y: 0.5, Z: -1},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			camera, err := NewCamera(test.CameraPosition, vmath.Vector2d{X: 1, Y: 1})
			require.NoError(t, err)
			actualPosition := camera.getOriginPixel()
			assert.Equal(t, test.Expected, actualPosition)
		})
	}
}

func TestCameraGetPixel(t *testing.T) {
	var tests = []struct {
		Description    string
		Expected       vmath.Vector3d
		CameraPosition vmath.Vector3d
		XY             vmath.Vector2d
	}{
		{
			Description:    "Simple getPixel same as OriginPixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			Expected:       vmath.Vector3d{X: -0.5, Y: 0.5, Z: -1},
			XY:             vmath.Vector2d{X: 0, Y: 0},
		},
		{
			Description:    "Simple getPixel max pixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			Expected:       vmath.Vector3d{X: 0.5, Y: -0.5, Z: -1},
			XY:             vmath.Vector2d{X: 1, Y: 1},
		},
		{
			Description:    "Simple getPixel middle pixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 0.0},
			Expected:       vmath.Vector3d{X: 49.5, Y: -49.5, Z: -1},
			XY:             vmath.Vector2d{X: 50, Y: 50},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			camera, err := NewCamera(test.CameraPosition, vmath.Vector2d{X: 100, Y: 100})
			require.NoError(t, err)
			actualPosition := camera.getPixel(float64(test.XY.X), float64(test.XY.Y))
			assert.Equal(t, test.Expected, actualPosition)
		})
	}
}

func TestCameraGetPickRay(t *testing.T) {
	var tests = []struct {
		Description    string
		CameraPosition vmath.Vector3d
		XY             vmath.Vector2d
	}{
		{
			Description:    "Simple getPixel same as OriginPixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 5.0},
			XY:             vmath.Vector2d{X: 0, Y: 0},
		},
		{
			Description:    "Simple getPixel max pixel",
			CameraPosition: vmath.Vector3d{X: 0.0, Y: 0.0, Z: 5.0},
			XY:             vmath.Vector2d{X: 1, Y: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			camera, err := NewCamera(test.CameraPosition, vmath.Vector2d{X: 1, Y: 1})
			require.NoError(t, err)
			actualPosition := camera.getPixel(float64(test.XY.X)/camera.XMax, float64(test.XY.Y)/camera.YMax)
			actualDirection := actualPosition.Subtract(camera.P)
			err = actualDirection.Normalize()
			require.NoError(t, err)
			expectedRay := &vmath.Ray{
				Origin:    camera.P,
				Direction: actualDirection,
			}
			actualRay, err := camera.GetPickRay(int(test.XY.X), int(test.XY.Y))
			require.NoError(t, err)
			assert.Equal(t, expectedRay, actualRay)
		})
	}
}

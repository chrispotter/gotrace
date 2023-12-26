package camera

import (
	"testing"

	vmath "github.com/chrispotter/trace/internal/math"
	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCameraConfigFromYaml(t *testing.T) {
	tests := []struct {
		Description string
		Expected    interface{}
		Bytes       []byte
	}{
		{
			Description: "Test Empty Map creates no Config",
			Expected:    &Config{},
			Bytes:       []byte(``),
		},
		{
			Description: "Simple YAML makes simple config",
			Expected: &Config{
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
			cameraConfig := &Config{}
			yaml, err := simpleyaml.NewYaml(test.Bytes)
			require.NoError(t, err)
			err = cameraConfig.FromYaml(yaml)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, cameraConfig)
		})
	}
}

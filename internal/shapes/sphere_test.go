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

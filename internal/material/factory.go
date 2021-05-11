package material

/*
import ()

// MaterialConfig is a yaml definition of the Constructor to be read from
// util/scene.go
type MaterialConfig struct {
	Name     string
	Position vmath.Vector3d `yaml:"position"`
	Ratio    vmath.Vector2d `yaml:"ratio"`
}

//FromYaml updates the CameraConfig with it's definition from the input yaml
func (cc *CameraConfig) FromYaml(config *simpleyaml.Yaml) error {
	if position, err := config.Get("position").Array(); err == nil {
		cc.Position = vmath.Vector3d{
			X: (position[0]).(float64),
			Y: (position[1]).(float64),
			Z: (position[2]).(float64),
		}
	}
	if ratio, err := config.Get("ratio").Array(); err == nil {
		cc.Ratio = vmath.Vector2d{
			X: (ratio[0]).(float64),
			Y: (ratio[1]).(float64),
		}
	}
	return nil
}

// NewCamera makes new camera with position and image ratio
func NewCamera(pos vmath.Vector3d, ratio vmath.Vector2d) (*Camera, error) {
	// image pixels shouldn't be 0 but if they are, just make them 1
	if ratio.X == 0 {
		ratio.X = 1
	}
	if ratio.Y == 0 {
		ratio.Y = 1
	}

	cam := &Camera{
		P:     pos,
		U:     vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0},
		XMax:  ratio.X,
		YMax:  ratio.Y,
		Depth: 1.0,
		Sx:    1.0,
	}

	cam.Sy = (cam.YMax * cam.Sx) / cam.XMax

	err := cam.LookAt(cam.P.Add(vmath.Vector3d{X: 0.0, Y: 0.0, Z: -1.0}), vmath.Vector3d{X: 0.0, Y: 1.0, Z: 0.0})
	if err != nil {
		return nil, err
	}

	return cam, nil
}
*/

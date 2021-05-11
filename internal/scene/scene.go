package scene

import (
	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/camera"
	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/common"
	"github.com/chrispotter/trace/internal/shapes"
)

/*
vector <Shape*> shapes;
vector <Light*> lights;
int depth = 5;
int e,f,m,n,k;
float globalN, currentN;
int M = 2;
Vector3d rotation(0.0,0.0,0.0);

Sphere envSph = Sphere(Vector3d (0.0,0.0,0.0), 1.0);
obj geodome = obj("ls/Geodesic/162.obj", Vector3d(0.0,0.0,0.0));
bool environment = true;
bool aOcclusion = false;
*/
type Scene struct {
	Cameras     []*camera.Camera
	Renderables *common.RenderableObjects
	Colors      map[string]color.Color
}

// Render each camera in the scene to the output
// `output/<cameraname>-filename.png
func (s *Scene) Render(filename string) error {
	for _, camera := range s.Cameras {
		err := camera.RenderImage(filename, s.Renderables)
		if err != nil {
			return err
		}
	}

	return nil
}

// FromYaml creates a scene calling the FromYaml of each Config that is passed
// to an objects Factory method
func (s *Scene) FromYaml(data []byte) error {
	y, err := simpleyaml.NewYaml(data)
	if err != nil {
		return err
	}

	// read cameras
	err = s.cameras(y)
	if err != nil {
		return err
	}
	// read colors before materials, light so they can reference colors
	err = s.colors(y)
	if err != nil {
		return err
	}
	//read materials before shapes so they can reference materials
	// read shapes
	err = s.shapes(y)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scene) cameras(y *simpleyaml.Yaml) error {
	if y.Get("cameras").IsFound() {
		yml := y.Get("cameras")
		var cameraConfigs []*camera.CameraConfig
		keys, err := yml.GetMapKeys()
		if err != nil {
			return err
		}
		for _, name := range keys {
			conf := yml.Get(name)
			cameraConfig := &camera.CameraConfig{}
			err := cameraConfig.FromYaml(conf)
			if err != nil {
				return err
			}
			cameraConfig.Name = name
			cameraConfigs = append(cameraConfigs, cameraConfig)
		}
		s.Cameras, err = camera.CameraFactory(cameraConfigs)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scene) colors(y *simpleyaml.Yaml) error {
	if y.Get("colors").IsFound() {
		yml := y.Get("colors")
		var colorConfigs []*color.ColorConfig
		keys, err := yml.GetMapKeys()
		if err != nil {
			return err
		}
		for _, name := range keys {
			conf := yml.Get(name)
			colorConfig := &color.ColorConfig{}
			err := colorConfig.FromYaml(conf)
			if err != nil {
				return err
			}
			colorConfig.Name = name
			colorConfigs = append(colorConfigs, colorConfig)
		}

		colors, err := color.ColorFactory(colorConfigs)
		if err != nil {
			return err
		}

		s.Colors = colors
	}
	return nil
}

func (s *Scene) shapes(y *simpleyaml.Yaml) error {
	if y.Get("shapes").IsFound() {
		yml := y.Get("shapes")
		shapeConfigs, err := shapes.ShapesConfigFactory(yml)
		if err != nil {
			return err
		}

		sh, err := shapes.ShapesFactory(shapeConfigs)
		if err != nil {
			return err
		}

		s.Renderables = &common.RenderableObjects{
			Shapes: sh,
		}
	}
	return nil
}

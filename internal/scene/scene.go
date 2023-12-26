package scene

import (
	"fmt"

	"github.com/smallfish/simpleyaml"

	"github.com/chrispotter/trace/internal/camera"
	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/common"
	"github.com/chrispotter/trace/internal/material"
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
	Materials   map[string]material.Material
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
	// read colors before materials, lights so they can reference colors
	err = s.colors(y)
	if err != nil {
		return err
	}
	// read materials before shapes so they can reference materials
	err = s.materials(y)
	if err != nil {
		return err
	}
	// read shapes
	err = s.shapes(y)
	if err != nil {
		return err
	}
	// read lights
	err = s.lights(y)
	if err != nil {
		return err
	}
	fmt.Printf("Scene: %+v", s)
	return nil
}

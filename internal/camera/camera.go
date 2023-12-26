package camera

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"

	"github.com/chrispotter/trace/internal/color"
	"github.com/chrispotter/trace/internal/common"
	vmath "github.com/chrispotter/trace/internal/math"
)

// Camera is a struct definition of a camera that will render an image or
// sequence of images
type Camera struct {
	// identifier for camera
	Name string
	// sidevector, upvector, viewvector, position
	U, V, N, P vmath.Vector3d
	// sx, sy, image width, image heigh, depth of field, focus length, focal radius
	Sx, Sy, XMax, YMax, Depth, Focus, Radius float64

	// depth of field enabled
	dof bool

	hit *vmath.Vector2d
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

// LookAt point camera at selected vector with specified up
// vector
func (c *Camera) LookAt(l vmath.Vector3d, u vmath.Vector3d) error {
	c.N = c.P.Subtract(l)
	err := c.N.Normalize()
	if err != nil {
		return err
	}

	c.U = u.Cross(c.N)
	err = c.U.Normalize()
	if err != nil {
		return err
	}

	c.V = c.N.Cross(c.U)
	err = c.V.Normalize()
	if err != nil {
		return err
	}

	return nil
}

// Slide along selected vactor in camera coordinates
func (c *Camera) Slide(s vmath.Vector3d) {
	input := vmath.Vector3d{
		X: s.X*c.U.X + s.Y*c.V.X + s.Z*c.N.X,
		Y: s.X*c.U.Y + s.Y*c.V.Y + s.Z*c.N.Y,
		Z: s.X*c.U.Z + s.Y*c.V.Z + s.Z*c.N.Z,
	}

	c.P = c.P.Add(input)
}

// Pitch rotates around U axis
func (cam *Camera) Pitch(angle float64) {
	cosVal := math.Cos(angle * math.Pi / 180)
	sinVal := math.Sin(angle * math.Pi / 180)

	N := vmath.Vector3d{
		X: cosVal*cam.N.X + sinVal*cam.V.X,
		Y: cosVal*cam.N.Y + sinVal*cam.V.Y,
		Z: cosVal*cam.N.Z + sinVal*cam.V.Z,
	}
	V := vmath.Vector3d{
		X: cosVal*cam.V.X - sinVal*cam.N.X,
		Y: cosVal*cam.V.Y - sinVal*cam.N.Y,
		Z: cosVal*cam.V.Y - sinVal*cam.N.Z,
	}

	cam.N = N
	cam.V = V
}

// Yaw rotates camera around V axis
func (cam *Camera) Yaw(angle float64) {
	cosVal := math.Cos(angle * math.Pi / 180)
	sinVal := math.Sin(angle * math.Pi / 180)

	U := vmath.Vector3d{
		X: cosVal*cam.U.X + sinVal*cam.N.X,
		Y: cosVal*cam.U.Y + sinVal*cam.N.Y,
		Z: cosVal*cam.U.Z + sinVal*cam.N.Z,
	}
	N := vmath.Vector3d{
		X: cosVal*cam.N.X - sinVal*cam.U.X,
		Y: cosVal*cam.N.Y - sinVal*cam.U.Y,
		Z: cosVal*cam.N.Y - sinVal*cam.U.Z,
	}

	cam.U = U
	cam.N = N
}

// Roll rotates camera around N axis
func (cam *Camera) Roll(angle float64) {
	cosVal := math.Cos(angle * math.Pi / 180)
	sinVal := math.Sin(angle * math.Pi / 180)

	V := vmath.Vector3d{
		X: cosVal*cam.V.X + sinVal*cam.U.X,
		Y: cosVal*cam.V.Y + sinVal*cam.U.Y,
		Z: cosVal*cam.V.Z + sinVal*cam.U.Z,
	}
	U := vmath.Vector3d{
		X: cosVal*cam.U.X - sinVal*cam.V.X,
		Y: cosVal*cam.U.Y - sinVal*cam.V.Y,
		Z: cosVal*cam.U.Y - sinVal*cam.V.Z,
	}

	cam.V = V
	cam.U = U
}

// Rotate camera at angle around axis
func (cam *Camera) Rotate(axis vmath.Vector3d, angle float64) {
	cosVal := math.Cos(angle * math.Pi / 180)
	sinVal := math.Sin(angle * math.Pi / 180)

	a := 1 + (1-cosVal)*(axis.X*axis.X-1)
	b := (1-cosVal)*axis.X*axis.Y - axis.Z*sinVal
	c := (1-cosVal)*axis.X*axis.Z + axis.Y*sinVal
	e := (1-cosVal)*axis.X*axis.Y + axis.Z*sinVal
	f := 1 + (1-cosVal)*(axis.Y*axis.Y-1)
	g := (1-cosVal)*axis.Y*axis.Z - axis.X*sinVal
	i := (1-cosVal)*axis.X*axis.Z - axis.Y*sinVal
	j := (1-cosVal)*axis.Y*axis.Z + axis.X*sinVal
	k := 1 + (1-cosVal)*(axis.Z*axis.Z-1)

	U := vmath.Vector3d{
		X: cam.U.X*a + cam.V.X*b + cam.N.X*c,
		Y: cam.U.Y*a + cam.V.Y*b + cam.N.Y*c,
		Z: cam.U.Z*a + cam.V.Z*b + cam.N.Z*c,
	}
	V := vmath.Vector3d{
		X: cam.U.X*e + cam.V.X*f + cam.N.X*g,
		Y: cam.U.Y*e + cam.V.Y*f + cam.N.Y*g,
		Z: cam.U.Z*e + cam.V.Z*f + cam.N.Z*g,
	}
	N := vmath.Vector3d{
		X: cam.U.X*i + cam.V.X*j + cam.N.X*k,
		Y: cam.U.Y*i + cam.V.Y*j + cam.N.Y*k,
		Z: cam.U.Z*i + cam.V.Z*j + cam.N.Z*k,
	}
	cam.U = U
	cam.V = V
	cam.N = N
}

// getImagePosition in world space
func (cam *Camera) getImagePosition() vmath.Vector3d {
	// Image should be along negative N axis, so we'll negate depth
	return cam.P.Add(cam.N.SMultiply(-cam.Depth))
}

// upper right hand pixel
func (cam *Camera) getOriginPixel() vmath.Vector3d {
	return cam.getImagePosition().Subtract(cam.U.SMultiply(cam.Sx / 2.0)).Add(cam.V.SMultiply(cam.Sy / 2.0))
}

func (cam *Camera) getPixel(x float64, y float64) vmath.Vector3d {
	return cam.getOriginPixel().Add(cam.U.SMultiply(float64(x) * cam.Sx)).Subtract(cam.V.SMultiply(float64(y) * cam.Sy))
}

// GetImage returns Vector representing Image Ratio
func (cam *Camera) GetImage() *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(cam.XMax), int(cam.YMax)}

	// set up the image to be rendered for this camera
	return image.NewRGBA(image.Rectangle{upLeft, lowRight})
}

// GetPickRay cast a ray from x,y image plane of Camera
func (cam *Camera) GetPickRay(x int, y int) (*vmath.Ray, error) {
	pp := cam.getPixel(float64(x)/cam.XMax, float64(y)/cam.YMax)

	pDirection := pp.Subtract(cam.P)
	err := pDirection.Normalize()
	if err != nil {
		return nil, err
	}

	if cam.dof {
		r := cam.U.SMultiply(vmath.RandNum(-cam.Radius, cam.Radius)).Add(cam.V.SMultiply(vmath.RandNum(-cam.Radius, cam.Radius)))
		pRand := cam.P.Add(r)

		C := cam.P.Add(pp.SMultiply(cam.Focus)).Subtract(pRand)
		err := C.Normalize()
		if err != nil {
			return nil, err
		} else {
			return &vmath.Ray{
				Origin:    pRand,
				Direction: C,
			}, nil
		}
	}

	return &vmath.Ray{
		Origin:    cam.P,
		Direction: pDirection,
	}, nil
}

func (c *Camera) trace(ray *vmath.Ray, objs *common.RenderableObjects) color.Color {
	// determine if ray intersects any shapes in scene
	// if no hit, return background color
	// if hit then return material color
	for index, shape := range objs.Shapes {
		if shape.Intersect(ray) {
			if c.hit == nil {
				c.hit = &vmath.Vector2d{
					X: float64(index),
					Y: shape.GetIntersectionRatio(),
				}
			} else if shape.GetIntersectionRatio() < c.hit.Y {
				c.hit = &vmath.Vector2d{
					X: float64(index),
					Y: shape.GetIntersectionRatio(),
				}
			}
		}
	}

	if c.hit == nil {
		backgroundColor := &color.ColorValue{
			Color: vmath.Vector3d{0.0, 0.0, 0.0},
		}
		return backgroundColor
	}

	hitShape := objs.Shapes[int(c.hit.X)]
	materialColor := hitShape.ReturnColor(ray, objs)
	c.hit = nil
	return materialColor
}

// RenderImage will render the objects provided to the file name
func (c *Camera) RenderImage(filename string, objs *common.RenderableObjects) error {
	// we'll generate & render image when RenderImage called rather than at
	// NewCamera to limit info passed with Cameras
	image := c.GetImage()
	for x := 0; x < image.Rect.Max.X; x++ {
		for y := 0; y < image.Rect.Max.Y; y++ {
			ray, err := c.GetPickRay(x, y)
			if err != nil {
				return err
			}

			color := c.trace(ray, objs)

			image.Set(x, y, color.GetRGBA())
		}
	}

	output := fmt.Sprintf("output/%s-%s.png", c.Name, filename)

	f, err := os.Create(output)
	if err != nil {
		return err
	}

	err = png.Encode(f, image)
	if err != nil {
		return err
	}

	return nil
}

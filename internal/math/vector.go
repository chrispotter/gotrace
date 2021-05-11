package math

import (
	"errors"
	"fmt"
	"math"
)

type Vector interface {
	Normalize() error
}

// Vector2d is a 2d Vector
type Vector2d struct {
	X, Y float64
}

// UNegate negates a Vector2d
func (v Vector2d) UNegate() Vector2d {
	return Vector2d{X: -v.X, Y: -v.Y}
}

// Subtract two Vector2d
func (v Vector2d) Subtract(vec Vector2d) Vector2d {
	return Vector2d{
		X: v.X - vec.X,
		Y: v.Y - vec.Y,
	}
}

// Add two Vector2d
func (v Vector2d) Add(vec Vector2d) Vector2d {
	return Vector2d{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

// SMultiply a scalar to a Vector2d
func (v Vector2d) SMultiply(s float64) Vector2d {
	return Vector2d{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Dot product of 2 Vector2d
func (v Vector2d) Dot(vec Vector2d) float64 {
	return v.X*vec.X + v.Y*vec.Y
}

// Compt component-wise multiplication of Vector2d
func (v Vector2d) Compt(vec Vector2d) Vector2d {
	return Vector2d{
		X: v.X * vec.X,
		Y: v.Y * vec.Y,
	}
}

// Divide Vector2d by a scalar
func (v Vector2d) Divide(s float64) Vector2d {
	return Vector2d{
		X: v.X / s,
		Y: v.Y / s,
	}
}

// Equals between two Vector2d
func (v Vector2d) Equals(vec Vector2d) bool {
	return (v.X == vec.X) && (v.Y == vec.Y)
}

// String representation of Vector2d
func (v Vector2d) String() string {
	return fmt.Sprintf("{X: %f, Y: %f}", v.X, v.Y)
}

// Norm computes a vectors magnitude
func (v Vector2d) Norm() float64 {
	return math.Sqrt(v.Normsqr())
}

// Normsqr computer the squared magnitude of a vector
func (v Vector2d) Normsqr() float64 {
	return Sqr(v.X) + Sqr(v.Y)
}

func (v *Vector2d) Normalize() error {
	magnitude := v.Norm()
	if math.Abs(v.X) > magnitude*math.MaxFloat64 ||
		math.Abs(v.Y) > magnitude*math.MaxFloat64 {
		return errors.New("attempting to take the norm of a zero 2d vector")
	}
	v.X = v.X / magnitude
	v.Y = v.Y / magnitude

	return nil
}

// Vector3d is a 3d Vector
type Vector3d struct {
	X, Y, Z float64
}

// UNegate negates a Vector3d
func (v Vector3d) UNegate() Vector3d {
	return Vector3d{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Subtract two Vector3d
func (v Vector3d) Subtract(vec Vector3d) Vector3d {
	return Vector3d{
		X: v.X - vec.X,
		Y: v.Y - vec.Y,
		Z: v.Z - vec.Z,
	}
}

// Add two Vector3d
func (v Vector3d) Add(vec Vector3d) Vector3d {
	return Vector3d{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
		Z: v.Z + vec.Z,
	}
}

// SMultiply a scalar to a Vector3d
func (v Vector3d) SMultiply(s float64) Vector3d {
	return Vector3d{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

// Dot product of 2 Vector3d
func (v Vector3d) Dot(vec Vector3d) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Compt component-wise multiplication of Vector3d
func (v Vector3d) Compt(vec Vector3d) Vector3d {
	return Vector3d{
		X: v.X * vec.X,
		Y: v.Y * vec.Y,
		Z: v.Z * vec.Z,
	}
}

// Cross computes cross-product of Vector3d
//
func (v Vector3d) Cross(vec Vector3d) Vector3d {
	return Vector3d{
		X: v.Y*vec.Z - v.Z*vec.Y,
		Y: v.Z*vec.X - v.X*vec.Z,
		Z: v.X*vec.Y - v.Y*vec.X,
	}
}

// Divide Vector3d by a scalar
func (v Vector3d) Divide(s float64) Vector3d {
	return Vector3d{
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
	}
}

// Equals between two Vector3d
func (v Vector3d) Equals(vec Vector3d) bool {
	return (v.X == vec.X) && (v.Y == vec.Y) && (v.Z == vec.Z)
}

// String representation of Vector2d
func (v Vector3d) String() string {
	return fmt.Sprintf("{X: %f, Y: %f, Z: %f}", v.X, v.Y, v.Z)
}

// Norm computes a vectors magnitude
func (v Vector3d) Norm() float64 {
	return math.Sqrt(v.Normsqr())
}

// Normsqr computer the squared magnitude of a vector
func (v Vector3d) Normsqr() float64 {
	return Sqr(v.X) + Sqr(v.Y) + Sqr(v.Z)
}

// Normalize Vector3d to make each axis 0<1
func (v *Vector3d) Normalize() error {
	magnitude := v.Norm()
	if v.IsZero() {
		return errors.New("attempting to take the norm of a zero 3d vector")
	}
	v.X = v.X / magnitude
	v.Y = v.Y / magnitude
	v.Z = v.Z / magnitude

	return nil
}

func (v *Vector3d) IsZero() bool {
	return math.Abs(v.X) == 0.0 &&
		math.Abs(v.Y) == 0.0 &&
		math.Abs(v.Z) == 0.0
}

// Vector4d is a 4d Vector
type Vector4d struct {
	X, Y, Z, W float64
}

// UNegate negates a Vector4d
func (v Vector4d) UNegate() Vector4d {
	return Vector4d{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
}

// Subtract two Vector4d
func (v Vector4d) Subtract(vec Vector4d) Vector4d {
	return Vector4d{
		X: v.X - vec.X,
		Y: v.Y - vec.Y,
		Z: v.Z - vec.Z,
		W: v.W - vec.W,
	}
}

// Add two Vector4d
func (v Vector4d) Add(vec Vector4d) Vector4d {
	return Vector4d{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
		Z: v.Y + vec.Z,
		W: v.W + vec.W,
	}
}

// SMultiply a scalar to a Vector4d
func (v Vector4d) SMultiply(s float64) Vector4d {
	return Vector4d{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
		W: v.W * s,
	}
}

// Dot product of 2 Vector4d
func (v Vector4d) Dot(vec Vector4d) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z + v.W*vec.W
}

// Compt component-wise multiplication of Vector4d
func (v Vector4d) Compt(vec Vector4d) Vector4d {
	return Vector4d{
		X: v.X * vec.X,
		Y: v.Y * vec.Y,
		Z: v.Z * vec.Z,
		W: v.W * vec.W,
	}
}

// Divide Vector4d by a scalar
func (v Vector4d) Divide(s float64) Vector4d {
	return Vector4d{
		X: v.X / s,
		Y: v.Y / s,
		Z: v.Z / s,
		W: 1.0,
	}
}

// Equals between two Vector4d
func (v Vector4d) Equals(vec Vector4d) bool {
	return (v.X == vec.X) && (v.Y == vec.Y) && (v.Z == vec.Z) && (v.W == vec.W)
}

// String representation of Vector4d
func (v *Vector4d) String() string {
	return fmt.Sprintf("{X: %f, Y: %f, Z: %f, W: %f}", v.X, v.Y, v.Z, v.W)
}

// Norm computes a vectors magnitude
func (v Vector4d) Norm() float64 {
	return math.Sqrt(v.Normsqr())
}

// Normsqr computer the squared magnitude of a vector
func (v Vector4d) Normsqr() float64 {
	return Sqr(v.X) + Sqr(v.Y) + Sqr(v.Z) + Sqr(v.W)
}

// Normalize Vector4d to make each axis 0<1
func (v *Vector4d) Normalize() error {
	magnitude := v.Norm()
	if math.Abs(v.X) > magnitude*math.MaxFloat64 ||
		math.Abs(v.Y) > magnitude*math.MaxFloat64 ||
		math.Abs(v.Z) > magnitude*math.MaxFloat64 {
		return errors.New("attempting to take the norm of a zero 4d vector")
	}

	v.X = v.X / magnitude
	v.Y = v.Y / magnitude
	v.Z = v.Z / magnitude
	v.W = 1.0

	return nil
}

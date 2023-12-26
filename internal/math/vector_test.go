package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVector2d(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    interface{}
		Result      interface{}
	}{
		{
			Description: "UNegate-posneg",
			Expected:    Vector2d{X: -1, Y: -1},
			Result:      (&Vector2d{X: 1, Y: 1}).UNegate(),
		},
		{
			Description: "UNegate-negpos",
			Expected:    Vector2d{X: 1, Y: 1},
			Result:      (&Vector2d{X: -1, Y: -1}).UNegate(),
		},
		{
			Description: "Subtract-post-result",
			Expected:    Vector2d{X: 1, Y: 1},
			Result:      (&Vector2d{X: 2, Y: 2}).Subtract(Vector2d{X: 1, Y: 1}),
		},
		{
			Description: "Subtract-double-neg",
			Expected:    Vector2d{X: -3, Y: -3},
			Result:      (&Vector2d{X: -2, Y: -2}).Subtract(Vector2d{X: 1, Y: 1}),
		},
		{
			Description: "Subtract-neg-result",
			Expected:    Vector2d{X: -1, Y: -1},
			Result:      (&Vector2d{X: -2, Y: -2}).Subtract(Vector2d{X: -1, Y: -1}),
		},
		{
			Description: "Add-pos-result",
			Expected:    Vector2d{X: 3, Y: 3},
			Result:      (&Vector2d{X: 2, Y: 2}).Add(Vector2d{X: 1, Y: 1}),
		},
		{
			Description: "Add-double-neg",
			Expected:    Vector2d{X: -3, Y: -3},
			Result:      (&Vector2d{X: -2, Y: -2}).Add(Vector2d{X: -1, Y: -1}),
		},
		{
			Description: "Add-neg-result",
			Expected:    Vector2d{X: 1, Y: 1},
			Result:      (&Vector2d{X: 2, Y: 2}).Add(Vector2d{X: -1, Y: -1}),
		},
		{
			Description: "SMultiply",
			Expected:    Vector2d{X: 15, Y: 15},
			Result:      (&Vector2d{X: 3, Y: 3}).SMultiply(5),
		},
		{
			Description: "Dot",
			Expected:    float64(20),
			Result:      (&Vector2d{X: 5, Y: 5}).Dot(Vector2d{X: 2, Y: 2}),
		},
		{
			Description: "Compt",
			Expected:    Vector2d{X: 10, Y: 10},
			Result:      (&Vector2d{X: 2, Y: 5}).Compt(Vector2d{X: 5, Y: 2}),
		},
		{
			Description: "Divide",
			Expected:    Vector2d{X: 5, Y: 2},
			Result:      (&Vector2d{X: 10, Y: 4}).Divide(2),
		},
		{
			Description: "Equals",
			Expected:    false,
			Result:      (&Vector2d{X: 10, Y: 4}).Equals(Vector2d{X: 4, Y: 10}),
		},
		{
			Description: "Equals",
			Expected:    true,
			Result:      (&Vector2d{X: 10, Y: 4}).Equals(Vector2d{X: 10, Y: 4}),
		},
		{
			Description: "String",
			Expected:    "{X: 10.000000, Y: 4.000000}",
			Result:      (&Vector2d{X: 10, Y: 4}).String(),
		},
		{
			Description: "Normsqr",
			Expected:    float64(116),
			Result:      (&Vector2d{X: 10, Y: 4}).Normsqr(),
		},
		{
			Description: "Norm",
			Expected:    math.Sqrt(Sqr(10) + Sqr(4)),
			Result:      (&Vector2d{X: 10, Y: 4}).Norm(),
		}}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Result)
		})
	}
}

func TestVector3d(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    interface{}
		Result      interface{}
	}{
		{
			Description: "UNegate-posneg",
			Expected:    Vector3d{X: -1, Y: -2, Z: -3},
			Result:      (&Vector3d{X: 1, Y: 2, Z: 3}).UNegate(),
		},
		{
			Description: "UNegate-negpos",
			Expected:    Vector3d{X: 1, Y: 2, Z: 3},
			Result:      (&Vector3d{X: -1, Y: -2, Z: -3}).UNegate(),
		},
		{
			Description: "Subtract-post-result",
			Expected:    Vector3d{X: 1, Y: 1, Z: 1},
			Result: (&Vector3d{X: 2, Y: 2, Z: 2}).Subtract(Vector3d{X: 1, Y: 1,
				Z: 1}),
		},
		{
			Description: "Subtract-double-neg",
			Expected:    Vector3d{X: -3, Y: -3, Z: -3},
			Result:      (&Vector3d{X: -2, Y: -2, Z: -2}).Subtract(Vector3d{X: 1, Y: 1, Z: 1}),
		},
		{
			Description: "Subtract-neg-result",
			Expected:    Vector3d{X: -1, Y: -1, Z: -1},
			Result:      (&Vector3d{X: -2, Y: -2, Z: -2}).Subtract(Vector3d{X: -1, Y: -1, Z: -1}),
		},
		{
			Description: "Add-pos-result",
			Expected:    Vector3d{X: 3, Y: 3, Z: 3},
			Result:      (&Vector3d{X: 2, Y: 2, Z: 2}).Add(Vector3d{X: 1, Y: 1, Z: 1}),
		},
		{
			Description: "Add-double-neg",
			Expected:    Vector3d{X: -3, Y: -3, Z: -3},
			Result:      (&Vector3d{X: -2, Y: -2, Z: -2}).Add(Vector3d{X: -1, Y: -1, Z: -1}),
		},
		{
			Description: "Add-neg-result",
			Expected:    Vector3d{X: 1, Y: 1, Z: 1},
			Result:      (&Vector3d{X: 2, Y: 2, Z: 2}).Add(Vector3d{X: -1, Y: -1, Z: -1}),
		},
		{
			Description: "SMultiply",
			Expected:    Vector3d{X: 15, Y: 15, Z: 15},
			Result:      (&Vector3d{X: 3, Y: 3, Z: 3}).SMultiply(5),
		},
		{
			Description: "Dot",
			Expected:    float64(30),
			Result:      (&Vector3d{X: 5, Y: 5, Z: 5}).Dot(Vector3d{X: 2, Y: 2, Z: 2}),
		},
		{
			Description: "Compt",
			Expected:    Vector3d{X: 10, Y: 10, Z: 10},
			Result:      (&Vector3d{X: 2, Y: 5, Z: 10}).Compt(Vector3d{X: 5, Y: 2, Z: 1}),
		},
		{
			Description: "Cross",
			Expected:    Vector3d{X: -8, Y: 5, Z: -3},
			Result:      (&Vector3d{X: 2, Y: 5, Z: 3}).Cross(Vector3d{X: 3, Y: 6, Z: 2}),
		},
		{
			Description: "Divide",
			Expected:    Vector3d{X: 5, Y: 2, Z: 3},
			Result:      (&Vector3d{X: 10, Y: 4, Z: 6}).Divide(2),
		},
		{
			Description: "Equals",
			Expected:    false,
			Result:      (&Vector3d{X: 10, Y: 4, Z: 6}).Equals(Vector3d{X: 4, Y: 10, Z: 6}),
		},
		{
			Description: "Equals",
			Expected:    true,
			Result:      (&Vector3d{X: 10, Y: 4, Z: 6}).Equals(Vector3d{X: 10, Y: 4, Z: 6}),
		},
		{
			Description: "String",
			Expected:    "{X: 10.000000, Y: 4.000000, Z: 6.000000}",
			Result:      (&Vector3d{X: 10, Y: 4, Z: 6}).String(),
		},
		{
			Description: "Normsqr",
			Expected:    float64(120),
			Result:      (&Vector3d{X: 10, Y: 4, Z: 2}).Normsqr(),
		},
		{
			Description: "Norm",
			Expected:    math.Sqrt(Sqr(10) + Sqr(4) + Sqr(2)),
			Result:      (&Vector3d{X: 10, Y: 4, Z: 2}).Norm(),
		}}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Result)
		})
	}
}

func TestVector4d(t *testing.T) {
	var tests = []struct {
		Description string
		Expected    interface{}
		Result      interface{}
	}{
		{
			Description: "UNegate-posneg",
			Expected:    Vector4d{X: -1, Y: -2, Z: -3, W: -4},
			Result:      (&Vector4d{X: 1, Y: 2, Z: 3, W: 4}).UNegate(),
		},
		{
			Description: "UNegate-negpos",
			Expected:    Vector4d{X: 1, Y: 2, Z: 3, W: 4},
			Result:      (&Vector4d{X: -1, Y: -2, Z: -3, W: -4}).UNegate(),
		},
		{
			Description: "Subtract-post-result",
			Expected:    Vector4d{X: 1, Y: 1, Z: 1, W: 1},
			Result: (&Vector4d{X: 2, Y: 2, Z: 2, W: 2}).Subtract(Vector4d{X: 1, Y: 1,
				Z: 1, W: 1}),
		},
		{
			Description: "Subtract-double-neg",
			Expected:    Vector4d{X: -3, Y: -3, Z: -3, W: -3},
			Result:      (&Vector4d{X: -2, Y: -2, Z: -2, W: -2}).Subtract(Vector4d{X: 1, Y: 1, Z: 1, W: 1}),
		},
		{
			Description: "Subtract-neg-result",
			Expected:    Vector4d{X: -1, Y: -1, Z: -1, W: -1},
			Result:      (&Vector4d{X: -2, Y: -2, Z: -2, W: -2}).Subtract(Vector4d{X: -1, Y: -1, Z: -1, W: -1}),
		},
		{
			Description: "Add-pos-result",
			Expected:    Vector4d{X: 3, Y: 3, Z: 3, W: 3},
			Result:      (&Vector4d{X: 2, Y: 2, Z: 2, W: 2}).Add(Vector4d{X: 1, Y: 1, Z: 1, W: 1}),
		},
		{
			Description: "Add-double-neg",
			Expected:    Vector4d{X: -3, Y: -3, Z: -3, W: -3},
			Result:      (&Vector4d{X: -2, Y: -2, Z: -2, W: -2}).Add(Vector4d{X: -1, Y: -1, Z: -1, W: -1}),
		},
		{
			Description: "Add-neg-result",
			Expected:    Vector4d{X: 1, Y: 1, Z: 1, W: 1},
			Result:      (&Vector4d{X: 2, Y: 2, Z: 2, W: 2}).Add(Vector4d{X: -1, Y: -1, Z: -1, W: -1}),
		},
		{
			Description: "SMultiply",
			Expected:    Vector4d{X: 15, Y: 15, Z: 15, W: 15},
			Result:      (&Vector4d{X: 3, Y: 3, Z: 3, W: 3}).SMultiply(5),
		},
		{
			Description: "Dot",
			Expected:    float64(40),
			Result:      (&Vector4d{X: 5, Y: 5, Z: 5, W: 2}).Dot(Vector4d{X: 2, Y: 2, Z: 2, W: 5}),
		},
		{
			Description: "Compt",
			Expected:    Vector4d{X: 10, Y: 10, Z: 10, W: 10},
			Result:      (&Vector4d{X: 2, Y: 5, Z: 10, W: 1}).Compt(Vector4d{X: 5, Y: 2, Z: 1, W: 10}),
		},
		{
			Description: "Divide",
			Expected:    Vector4d{X: 5, Y: 2, Z: 3, W: 1.0},
			Result:      (&Vector4d{X: 10, Y: 4, Z: 6, W: 14}).Divide(2),
		},
		{
			Description: "Equals",
			Expected:    false,
			Result:      (&Vector4d{X: 10, Y: 4, Z: 6, W: 12}).Equals(Vector4d{X: 4, Y: 10, Z: 6, W: 12}),
		},
		{
			Description: "Equals",
			Expected:    true,
			Result:      (&Vector4d{X: 10, Y: 4, Z: 6, W: 1}).Equals(Vector4d{X: 10, Y: 4, Z: 6, W: 1}),
		},
		{
			Description: "String",
			Expected:    "{X: 10.000000, Y: 4.000000, Z: 6.000000, W: 3.000000}",
			Result:      (&Vector4d{X: 10, Y: 4, Z: 6, W: 3}).String(),
		},
		{
			Description: "Normsqr",
			Expected:    float64(129),
			Result:      (&Vector4d{X: 10, Y: 4, Z: 2, W: 3}).Normsqr(),
		},
		{
			Description: "Norm",
			Expected:    math.Sqrt(Sqr(10) + Sqr(4) + Sqr(2) + Sqr(3)),
			Result:      (&Vector4d{X: 10, Y: 4, Z: 2, W: 3}).Norm(),
		}}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Result)
		})
	}

}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		Description string
		Assertion   assert.ErrorAssertionFunc
		Expected    interface{}
		Result      Vector
	}{
		{
			Description: "Vector2d normalize",
			Assertion:   assert.NoError,
			Expected:    &Vector2d{X: 10 / math.Sqrt(Sqr(10)+Sqr(6)), Y: 6 / math.Sqrt(Sqr(10)+Sqr(6))},
			Result:      &Vector2d{X: 10, Y: 6},
		},
		{
			Description: "Vector3d normalize",
			Assertion:   assert.NoError,
			Expected: &Vector3d{X: 10 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)),
				Y: 6 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)),
				Z: 4 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)),
			},
			Result: &Vector3d{X: 10, Y: 6, Z: 4},
		},
		{
			Description: "Vector4d normalize",
			Assertion:   assert.NoError,
			Expected: &Vector4d{X: 10 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)+Sqr(2)),
				Y: 6 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)+Sqr(2)),
				Z: 4 / math.Sqrt(Sqr(10)+Sqr(6)+Sqr(4)+Sqr(2)),
				W: 1,
			},
			Result: &Vector4d{X: 10, Y: 6, Z: 4, W: 2},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			err := test.Result.Normalize()
			require.NoError(t, err)
			assert.Equal(t, test.Expected, test.Result)
		})
	}
}

package math

import (
	"math"
	"math/rand"
)

// RandNum returns a number between h and l
func RandNum(h float64, l float64) float64 {
	return l + rand.Float64()*(h-l)
}

// Sqr wrapper for x*x
func Sqr(x float64) float64 {
	return math.Pow(x, 2)
}

// Pythag computers sqrt(a^2 + b^2) without destructive
// underflow or overflow
func Pythag(a float64, b float64) float64 {
	var absa, absb float64

	absa = math.Abs(a)
	absb = math.Abs(b)

	if absa > absb {
		return absa * math.Sqrt(1.0+Sqr(absb/absa))
	} else if absa > 0 {
		return absb * math.Sqrt(1.0+Sqr(absa/absb))
	} else {
		return 0
	}
}

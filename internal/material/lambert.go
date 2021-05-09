package material

// Lambert is a default flat shader with a Ambient and Diffuse color
type Lambert struct {
	LightIndex, DistanceLightHit, U, V, N            float64
	Reflect, Iridesent, Refract, Glossy, Transparent bool
}

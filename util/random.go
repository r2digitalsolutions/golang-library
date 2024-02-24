package util

import "math/rand"

// Generate a random integer
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// Generate a random integer64
func RandomInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// Generate a random float

// Generate a random float32
func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// Generate a random float64
func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Generate a random string
func RandomString(length int) string {
	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}

	return string(bytes)
}

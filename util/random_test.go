package util

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	min := 0
	max := 100
	result := RandomInt(min, max)
	if result < min || result >= max {
		t.Fatalf("RandomInt() generated an out of range value: %d", result)
	}
}

func TestRandomInt64(t *testing.T) {
	min := int64(0)
	max := int64(100)
	result := RandomInt64(min, max)
	if result < min || result >= max {
		t.Fatalf("RandomInt64() generated an out of range value: %d", result)
	}
}

func TestRandomFloat32(t *testing.T) {
	min := float32(0)
	max := float32(100)
	result := RandomFloat32(min, max)
	if result < min || result >= max {
		t.Fatalf("RandomFloat32() generated an out of range value: %f", result)
	}
}

func TestRandomFloat64(t *testing.T) {
	min := float64(0)
	max := float64(100)
	result := RandomFloat64(min, max)
	if result < min || result >= max {
		t.Fatalf("RandomFloat64() generated an out of range value: %f", result)
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	result := RandomString(length)
	if len(result) != length {
		t.Fatalf("RandomString() generated an out of range value: %s", result)
	}
}

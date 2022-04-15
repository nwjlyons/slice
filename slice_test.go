package slice_test

import (
	"slice"
	"testing"
)

func TestReduce(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	got := slice.Reduce(planets, func(planet string, acc string) string {
		return acc + planet
	}, "")
	expected := "MercuryVenusEarthMarsJupiterSaturnUranusNeptune"
	assert(t, got, expected)
}

func assert[T comparable](t *testing.T, got T, expected T) {
	if got != expected {
		t.Errorf("got:\n%v\nexpected:\n%v\n", got, expected)
	}
}

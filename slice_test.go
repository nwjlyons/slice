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
	if got != expected {
		t.Errorf("got:\n%v\nexpected:\n%v\n", got, expected)
	}
}

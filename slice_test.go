package slice_test

import (
	"reflect"
	"slice"
	"testing"
)

func TestReduceWhile(t *testing.T) {
	numbers := []int{40, 2, 8}
	got := slice.ReduceWhile(numbers, func(number int, total int) (slice.Reduction, int) {
		if total >= 42 {
			return slice.Halt, total
		}
		return slice.Cont, number + total
	}, 0)
	assert(t, got, 42)
}

func TestReduce(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	got := slice.Reduce(planets, func(planet string, acc string) string {
		return acc + planet
	}, "")
	expected := "MercuryVenusEarthMarsJupiterSaturnUranusNeptune"
	assert(t, got, expected)
}

func TestMap(t *testing.T) {
	trafficLights := []string{"red", "amber", "green"}
	got := slice.Map(trafficLights, func(light string) string {
		return light + "!"
	})
	expected := []string{"red!", "amber!", "green!"}
	assert(t, reflect.DeepEqual(got, expected), true)
}

func TestIsMember(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assert(t, slice.IsMember(planets, "Earth"), true)
	assert(t, slice.IsMember(planets, "Pluto"), false)
}

func TestMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assert(t, slice.Max(numbers), 9)
}

func TestMin(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assert(t, slice.Min(numbers), 1)
}

func assert[T comparable](t *testing.T, got T, expected T) {
	if got != expected {
		t.Errorf("got:\n%v\nexpected:\n%v\n", got, expected)
	}
}

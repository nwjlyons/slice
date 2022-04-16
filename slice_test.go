package slice_test

import (
	"reflect"
	"slice"
	"strings"
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

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := slice.Filter(numbers, func(number int) bool {
		return number%2 == 0
	})
	assert(t, reflect.DeepEqual(got, []int{2, 4, 6, 8}), true)
}

func TestReject(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := slice.Reject(numbers, func(number int) bool {
		return number%2 == 0
	})
	assert(t, reflect.DeepEqual(got, []int{1, 3, 5, 7, 9}), true)
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

func TestMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	maxPlanet := slice.MaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assert(t, maxPlanet.Name, jupiter.Name)
}

func TestMin(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assert(t, slice.Min(numbers), 1)
}

type planet struct {
	Name   string
	Radius int
}

func TestMinBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet := slice.MinBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assert(t, minPlanet.Name, mars.Name)
}

func TestMinMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	min, max := slice.MinMax(numbers)
	assert(t, min, 1)
	assert(t, max, 9)
}

func TestMinMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet, maxPlanet := slice.MinMaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assert(t, minPlanet.Name, mars.Name)
	assert(t, maxPlanet.Name, jupiter.Name)
}

func TestSum(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assert(t, slice.Sum(numbers), 46)
}

func TestSumBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	sum := slice.SumBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assert(t, sum, mars.Radius+jupiter.Radius+neptune.Radius)
}

func TestAny(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assert(t, slice.Any([]int{1, 3, 5, 7, 9}, isEven), false)
	assert(t, slice.Any([]int{1, 3, 2, 7, 9}, isEven), true)
}

func TestAll(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assert(t, slice.All([]int{1, 2, 4, 6, 8}, isEven), false)
	assert(t, slice.All([]int{2, 4, 6, 8}, isEven), true)
}

func TestCount(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert(t, slice.Count(numbers), 9)
}

func TestCountBy(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assert(t, slice.CountBy(numbers, isEven), 4)
}

func TestFrequencies(t *testing.T) {
	frequencies := slice.Frequencies([]string{"aa", "aa", "bb", "cc"})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assert(t, reflect.DeepEqual(frequencies, expected), true)
}

func TestFrequenciesBy(t *testing.T) {
	frequencies := slice.FrequenciesBy([]string{"aa", "aA", "bb", "cc"}, func(element string) string {
		return strings.ToLower(element)
	})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assert(t, reflect.DeepEqual(frequencies, expected), true)
}

func assert[T comparable](t *testing.T, got T, expected T) {
	if got != expected {
		t.Errorf("got:\n%v\nexpected:\n%v\n", got, expected)
	}
}

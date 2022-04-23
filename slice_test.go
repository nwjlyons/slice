package slice_test

import (
	"fmt"
	"github.com/nwjlyons/slice"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type planet struct {
	Name   string
	Radius int
}

func TestAll(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, slice.All([]int{1, 2, 4, 6, 8}, isEven), false)
	assertEqual(t, slice.All([]int{2, 4, 6, 8}, isEven), true)
}

func TestAny(t *testing.T) {
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, slice.Any([]int{1, 3, 5, 7, 9}, isEven), false)
	assertEqual(t, slice.Any([]int{1, 3, 2, 7, 9}, isEven), true)
}

func TestCount(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEqual(t, slice.Count(numbers), 9)
}

func TestCountBy(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	isEven := func(number int) bool {
		return number%2 == 0
	}
	assertEqual(t, slice.CountBy(numbers, isEven), 4)
}

func TestEach(t *testing.T) {
	countdown := []string{"3", "2", "1", "Go!"}
	slice.Each(countdown, func(tick string) { fmt.Println(tick) })
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := slice.Filter(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, got, []int{2, 4, 6, 8})
}

func TestFlatMap(t *testing.T) {
	numbers := []int{1, 2, 3}
	assertEqual(t, slice.FlatMap(numbers, func(number int) []int {
		return []int{number, number}
	}), []int{1, 1, 2, 2, 3, 3})
}

func TestFrequencies(t *testing.T) {
	frequencies := slice.Frequencies([]string{"aa", "aa", "bb", "cc"})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assertEqual(t, frequencies, expected)
}

func TestFrequenciesBy(t *testing.T) {
	frequencies := slice.FrequenciesBy([]string{"aa", "aA", "bb", "cc"}, func(element string) string {
		return strings.ToLower(element)
	})
	expected := map[string]int{"aa": 2, "bb": 1, "cc": 1}
	assertEqual(t, frequencies, expected)
}

func TestGroupBy(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	expected := map[int][]string{7: {"Mercury", "Jupiter", "Neptune"}, 6: {"Saturn", "Uranus"}, 5: {"Venus", "Earth"}, 4: {"Mars"}}
	got := slice.GroupBy(planets, func(planet string) int {
		return len(planet)
	})
	assertEqual(t, got, expected)
}

func TestIsMember(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, slice.IsMember(planets, "Earth"), true)
	assertEqual(t, slice.IsMember(planets, "Pluto"), false)
}

func TestIsMemberBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	neptuneWithDifferentRadius := planet{Name: "Neptune", Radius: 42_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}
	planets := []planet{neptune, mars, jupiter}

	assertEqual(t, slice.IsMemberBy(planets, neptuneWithDifferentRadius, func(planet planet) string {
		return planet.Name
	}), true)
}

func TestMap(t *testing.T) {
	trafficLights := []string{"red", "amber", "green"}
	got := slice.Map(trafficLights, func(light string) string {
		return light + "!"
	})
	expected := []string{"red!", "amber!", "green!"}
	assertEqual(t, got, expected)

	numbers := []int{1, 2, 3}
	numbersAsStrings := slice.Map(numbers, func(number int) string {
		return strconv.Itoa(number)
	})
	assertEqual(t, numbersAsStrings, []string{"1", "2", "3"})
}

func TestMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, slice.Max(numbers), 9)
}

func TestMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	maxPlanet := slice.MaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, maxPlanet.Name, jupiter.Name)
}

func TestMin(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, slice.Min(numbers), 1)
}

func TestMinBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet := slice.MinBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, minPlanet.Name, mars.Name)
}

func TestMinMax(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	min, max := slice.MinMax(numbers)
	assertEqual(t, min, 1)
	assertEqual(t, max, 9)
}

func TestMinMaxBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	minPlanet, maxPlanet := slice.MinMaxBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, minPlanet.Name, mars.Name)
	assertEqual(t, maxPlanet.Name, jupiter.Name)
}

func TestProduct(t *testing.T) {
	assertEqual(t, slice.Product([]int{2, 3, 4}), 24)
	assertEqual(t, slice.Product([]int{2.0, 3.0, 4.0}), 24.0)
	assertEqual(t, slice.Product([]int{42}), 42)
}

func TestProductBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 2}
	mars := planet{Name: "Mars", Radius: 3}
	jupiter := planet{Name: "Jupiter", Radius: 4}
	assertEqual(t, slice.ProductBy([]planet{neptune, mars, jupiter}, func(planet planet) int {
		return planet.Radius
	}), 24)
}

func TestRandom(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, slice.Random(planets, 42), "Venus")
}

func TestReduce(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	got := slice.Reduce(planets, func(planet string, acc string) string {
		return acc + planet
	}, "")
	expected := "MercuryVenusEarthMarsJupiterSaturnUranusNeptune"
	assertEqual(t, got, expected)
}

func TestReduceWhile(t *testing.T) {
	numbers := []int{40, 2, 8}
	got := slice.ReduceWhile(numbers, func(number int, total int) (slice.Reduction, int) {
		if total >= 42 {
			return slice.Halt, total
		}
		return slice.Cont, number + total
	}, 0)
	assertEqual(t, got, 42)
}

func TestReject(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := slice.Reject(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, got, []int{1, 3, 5, 7, 9})
}

func TestReverse(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	expected := []string{"Neptune", "Uranus", "Saturn", "Jupiter", "Mars", "Earth", "Venus", "Mercury"}
	assertEqual(t, slice.Reverse(planets), expected)
}

func TestShuffle(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, slice.Shuffle(planets, 42), []string{"Saturn", "Neptune", "Jupiter", "Uranus", "Venus", "Mars", "Mercury", "Earth"})
}

func TestSort(t *testing.T) {
	numbers := []int{5, 6, 1, 3, 7, 8, 2, 4, 9}
	assertEqual(t, slice.Sort(numbers, slice.Asc), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assertEqual(t, slice.Sort(numbers, slice.Desc), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}

	assertEqual(t, slice.SortBy(planets, func(planet planet) int {
		return planet.Radius
	}, slice.Asc), []planet{mars, neptune, jupiter})

	assertEqual(t, slice.SortBy(planets, func(planet planet) int {
		return planet.Radius
	}, slice.Desc), []planet{jupiter, neptune, mars})
}

func TestSplitWhile(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	left, right := slice.SplitWhile(numbers, func(number int) bool {
		return number <= 5
	})
	assertEqual(t, left, []int{1, 2, 3, 4, 5})
	assertEqual(t, right, []int{6, 7, 8, 9})
}

func TestSplitWith(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers, oddNumbers := slice.SplitWith(numbers, func(number int) bool {
		return number%2 == 0
	})
	assertEqual(t, evenNumbers, []int{2, 4, 6, 8})
	assertEqual(t, oddNumbers, []int{1, 3, 5, 7, 9})
}

func TestSum(t *testing.T) {
	numbers := []int{6, 4, 8, 2, 1, 9, 4, 7, 5}
	assertEqual(t, slice.Sum(numbers), 46)
}

func TestSumBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	jupiter := planet{Name: "Jupiter", Radius: 69_911_000}

	planets := []planet{neptune, mars, jupiter}
	sum := slice.SumBy(planets, func(planet planet) int {
		return planet.Radius
	})
	assertEqual(t, sum, mars.Radius+jupiter.Radius+neptune.Radius)
}

func TestTake(t *testing.T) {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	assertEqual(t, slice.Take(planets, 2), []string{"Mercury", "Venus"})
	assertEqual(t, slice.Take(planets, 10), planets)
	assertEqual(t, slice.Take(planets, 0), []string{})
}

func TestTakeWhile(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assertEqual(t, slice.TakeWhile(numbers, func(number int) bool {
		return number <= 5
	}), []int{1, 2, 3, 4, 5})
}

func TestUniq(t *testing.T) {
	moves := []string{"Up", "Down", "Up", "Up", "Down", "Left", "Right", "Right", "Right", "Left"}
	assertEqual(t, slice.Uniq(moves), []string{"Up", "Down", "Left", "Right"})
}

func TestUniqBy(t *testing.T) {
	neptune := planet{Name: "Neptune", Radius: 24_622_000}
	mars := planet{Name: "Mars", Radius: 3_389_500}
	sameRadiusAsMars := planet{Name: "Same Radius as Mars", Radius: 3_389_500}

	planets := []planet{mars, neptune, sameRadiusAsMars}
	assertEqual(t, slice.UniqBy(planets, func(planet planet) int {
		return planet.Radius
	}), []planet{mars, neptune})
}

func assertEqual[T any](t *testing.T, got T, expected T) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}

func assertNotEqual[T any](t *testing.T, got T, expected T) {
	if reflect.DeepEqual(got, expected) {
		t.Errorf("\n     got: %v\nexpected: %v\n", got, expected)
	}
}

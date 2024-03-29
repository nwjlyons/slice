package slice

import (
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

type Reduction int

const (
	Cont Reduction = iota
	Halt
)

type Order int

const (
	Asc Order = iota
	Desc
)

type pair[Element any] struct {
	left  Element
	right Element
}

type Number interface {
	constraints.Integer | constraints.Float
}

// All returns true if fun returns true for all elements in the slice.
func All[Element any](elements []Element, fun func(Element) bool) bool {
	return ReduceWhile(elements, func(element Element, accumulator bool) (Reduction, bool) {
		if fun(element) {
			return Cont, true
		}
		return Halt, false
	}, false)
}

// Any returns true if fun returns true for at least one element in the slice.
func Any[Element any](elements []Element, fun func(Element) bool) bool {
	return ReduceWhile(elements, func(element Element, accumulator bool) (Reduction, bool) {
		if fun(element) {
			return Halt, true
		}
		return Cont, false
	}, false)
}

// At finds the element at the given index (zero-based).
func At[Element any](elements []Element, index int, defaultValue Element) Element {
	if index >= 0 && index < len(elements) {
		return elements[index]
	} else {
		return defaultValue
	}
}

// Concat concatenates the enumerable on the right with the enumerable on the left.
func Concat[Element any](left []Element, right []Element) []Element {
	return append(left, right...)
}

// Count counts the number of elements in the slice.
func Count[Element any](elements []Element) int {
	return len(elements)
}

// CountBy counts the number of elements in slice where fun returns true.
func CountBy[Element any](elements []Element, fun func(Element) bool) int {
	return Reduce(elements, func(element Element, accumulator int) int {
		if fun(element) {
			return accumulator + 1
		}
		return accumulator
	}, 0)
}

// Each invokes fun on each element in the slice.
func Each[Element any](elements []Element, fun func(Element)) {
	Reduce(elements, func(element Element, accumulator interface{}) interface{} {
		fun(element)
		return accumulator
	}, nil)
}

// Filter returns elements where fun returns true.
func Filter[Element any](elements []Element, fun func(Element) bool) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		if fun(element) {
			return append(accumulator, element)
		}
		return accumulator
	}, make([]Element, 0))
}

// FlatMap maps the given fun over slice and flattens the result.
func FlatMap[Element any](elements []Element, fun func(Element) []Element) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		return append(accumulator, fun(element)...)
	}, make([]Element, 0))
}

// Frequencies returns a map with keys as unique elements and values as the count of every element.
func Frequencies[Element comparable](elements []Element) map[Element]int {
	return FrequenciesBy(elements, func(element Element) Element {
		return element
	})
}

// FrequenciesBy returns a map with keys as unique elements given by key_fun and values as the count of every element.
func FrequenciesBy[Element any, Key comparable](elements []Element, fun func(Element) Key) map[Key]int {
	return Reduce(elements, func(element Element, accumulator map[Key]int) map[Key]int {
		accumulator[fun(element)]++
		return accumulator
	}, make(map[Key]int))
}

// GroupBy splits the slice into groups based on key_fun.
func GroupBy[Element any, GroupBy comparable](elements []Element, fun func(Element) GroupBy) map[GroupBy][]Element {
	return Reduce(elements, func(element Element, accumulator map[GroupBy][]Element) map[GroupBy][]Element {
		accumulator[fun(element)] = append(accumulator[fun(element)], element)
		return accumulator
	}, make(map[GroupBy][]Element))
}

// IsMember checks if element exists in the slice.
func IsMember[Element comparable](elements []Element, member Element) bool {
	return IsMemberBy(elements, member, func(element Element) Element {
		return element
	})
}

// IsMemberBy checks if element exists in the slice according to fun.
func IsMemberBy[Element any, IsMemberBy comparable](elements []Element, member Element, fun func(Element) IsMemberBy) bool {
	return ReduceWhile(elements, func(element Element, accumulator bool) (Reduction, bool) {
		if fun(element) == fun(member) {
			return Halt, true
		}
		return Cont, false
	}, false)
}

// Map invokes fun on each element in the slice.
func Map[Element any, ReturnElement any](elements []Element, fun func(Element) ReturnElement) []ReturnElement {
	return Reduce(elements, func(element Element, accumulator []ReturnElement) []ReturnElement {
		return append(accumulator, fun(element))
	}, make([]ReturnElement, 0))
}

// Max returns the maximum element in the slice.
func Max[Element constraints.Ordered](elements []Element) Element {
	return MaxBy(elements, func(element Element) Element {
		return element
	})
}

// MaxBy returns the maximum element in the slice according to fun.
func MaxBy[Element any, CompareBy constraints.Ordered](elements []Element, fun func(Element) CompareBy) Element {
	return Reduce(elements, func(element Element, max Element) Element {
		if fun(element) > fun(max) {
			max = element
		}
		return max
	}, elements[0])
}

// Min returns the minimum element in the slice.
func Min[Element constraints.Ordered](elements []Element) Element {
	return MinBy(elements, func(element Element) Element {
		return element
	})
}

// MinBy returns the minimum element in the slice according to fun.
func MinBy[Element any, CompareBy constraints.Ordered](elements []Element, fun func(Element) CompareBy) Element {
	return Reduce(elements, func(element Element, min Element) Element {
		if fun(element) < fun(min) {
			min = element
		}
		return min
	}, elements[0])
}

// MinMax returns the minimum and maximum element in the slice.
func MinMax[Element constraints.Ordered](elements []Element) (Element, Element) {
	return MinMaxBy(elements, func(element Element) Element {
		return element
	})
}

// MinMaxBy returns the minimum and maximum element in the slice according to fun.
func MinMaxBy[Element any, CompareBy constraints.Ordered](elements []Element, fun func(Element) CompareBy) (Element, Element) {
	result := Reduce(elements, func(element Element, accumulator pair[Element]) pair[Element] {
		if fun(element) < fun(accumulator.left) {
			accumulator.left = element
		}
		if fun(element) > fun(accumulator.right) {
			accumulator.right = element
		}
		return accumulator
	}, pair[Element]{left: elements[0], right: elements[0]})

	return result.left, result.right
}

// Product returns the product of all elements.
func Product[Element Number](elements []Element) Element {
	return ProductBy(elements, func(element Element) Element {
		return element
	})
}

// ProductBy returns the product of all elements according to fun.
func ProductBy[Element any, Product Number](elements []Element, fun func(Element) Product) Product {
	return Reduce(elements[1:], func(element Element, accumulator Product) Product {
		return fun(element) * accumulator
	}, fun(elements[0]))
}

// Random returns a random element from the slice.
func Random[Element any](elements []Element, seed ...int64) Element {
	if len(seed) == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else if len(seed) == 1 {
		rand.Seed(seed[0])
	} else {
		panic("unexpected value for seed parameter")
	}
	return elements[rand.Intn(len(elements))]
}

// Reduce invokes fun on each element in the slice with the accumulator.
func Reduce[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) Accumulator, accumulator Accumulator) Accumulator {
	return ReduceWhile(elements, func(element Element, accumulator Accumulator) (Reduction, Accumulator) {
		return Cont, fun(element, accumulator)
	}, accumulator)
}

// ReduceWhile invokes fun on each element in the slice with the accumulator until Halt is returned.
func ReduceWhile[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) (Reduction, Accumulator), accumulator Accumulator) Accumulator {
	reduction := Cont
	for _, element := range elements {
		reduction, accumulator = fun(element, accumulator)
		if reduction == Halt {
			return accumulator
		}
	}
	return accumulator
}

// Reject returns elements excluding those where fun returns true.
func Reject[Element any](elements []Element, fun func(Element) bool) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		if !fun(element) {
			return append(accumulator, element)
		}
		return accumulator
	}, make([]Element, 0))
}

// Reverse returns a slice of elements in reverse order.
func Reverse[Element comparable](elements []Element) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		return append([]Element{element}, accumulator...)
	}, make([]Element, 0))
}

// Shuffle returns a list with the elements of the slice shuffled.
func Shuffle[Element any](elements []Element, seed ...int64) []Element {
	if len(seed) == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else if len(seed) == 1 {
		rand.Seed(seed[0])
	} else {
		panic("unexpected value for seed parameter")
	}
	shuffledElements := make([]Element, len(elements))
	copy(shuffledElements, elements)
	rand.Shuffle(len(elements), func(i, j int) {
		shuffledElements[i], shuffledElements[j] = shuffledElements[j], shuffledElements[i]
	})
	return shuffledElements
}

// Sort returns a slice sorted according to fun.
func Sort[Element constraints.Ordered](elements []Element, order Order) []Element {
	return SortBy(elements, func(element Element) Element {
		return element
	}, order)
}

// SortBy returns a slice sorted according to fun.
func SortBy[Element any, SortBy constraints.Ordered](elements []Element, fun func(Element) SortBy, order Order) []Element {
	sortedElements := make([]Element, len(elements))
	copy(sortedElements, elements)
	sort.Slice(sortedElements, func(i, j int) bool {
		if order == Asc {
			return fun(sortedElements[i]) < fun(sortedElements[j])
		}
		return fun(sortedElements[i]) > fun(sortedElements[j])
	})
	return sortedElements
}

// SplitWhile splits the slice in two at the position of the element for which fun returns a false for the first time.
func SplitWhile[Element any](elements []Element, fun func(Element) bool) ([]Element, []Element) {
	addToLeft := true
	result := Reduce(elements, func(element Element, accumulator pair[[]Element]) pair[[]Element] {

		if addToLeft == true && fun(element) == false {
			addToLeft = false
		}

		if addToLeft {
			accumulator.left = append(accumulator.left, element)
		} else {
			accumulator.right = append(accumulator.right, element)
		}
		return accumulator
	}, pair[[]Element]{})
	return result.left, result.right
}

// SplitWith splits the slice in two lists according to the given function fun.
func SplitWith[Element any](elements []Element, fun func(Element) bool) ([]Element, []Element) {
	result := Reduce(elements, func(element Element, accumulator pair[[]Element]) pair[[]Element] {
		if fun(element) {
			accumulator.left = append(accumulator.left, element)
		} else {
			accumulator.right = append(accumulator.right, element)
		}
		return accumulator
	}, pair[[]Element]{})
	return result.left, result.right
}

// Sum returns the sum of all elements.
func Sum[Element Number](elements []Element) Element {
	return SumBy(elements, func(element Element) Element {
		return element
	})
}

// SumBy returns the sum of all elements according to fun.
func SumBy[Element any, SumBy Number](elements []Element, fun func(Element) SumBy) SumBy {
	return Reduce(elements[1:], func(element Element, accumulator SumBy) SumBy {
		return fun(element) + accumulator
	}, fun(elements[0]))
}

// Take takes an amount of elements from the beginning of the slice.
func Take[Element any](elements []Element, amount uint) []Element {
	return ReduceWhile(elements, func(element Element, accumulator []Element) (Reduction, []Element) {
		if len(accumulator) < int(amount) {
			accumulator = append(accumulator, element)
			return Cont, accumulator
		}
		return Halt, accumulator
	}, make([]Element, 0))
}

// TakeWhile takes the elements from the beginning of the slice while fun returns a truthy value.
func TakeWhile[Element any](elements []Element, fun func(Element) bool) []Element {
	return ReduceWhile(elements, func(element Element, accumulator []Element) (Reduction, []Element) {
		if fun(element) {
			return Cont, append(accumulator, element)
		}
		return Halt, accumulator
	}, make([]Element, 0))
}

// Uniq iterates over the slice, removing all duplicated elements.
func Uniq[Element comparable](elements []Element) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		if !IsMember(accumulator, element) {
			accumulator = append(accumulator, element)
		}
		return accumulator
	}, make([]Element, 0))
}

// UniqBy iterates over the slice, removing all duplicated elements according to fun.
func UniqBy[Element any, UniqBy comparable](elements []Element, fun func(Element) UniqBy) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		if !IsMemberBy(accumulator, element, fun) {
			accumulator = append(accumulator, element)
		}
		return accumulator
	}, make([]Element, 0))
}

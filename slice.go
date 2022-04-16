package slice

import (
	"golang.org/x/exp/constraints"
)

type Reduction int

const (
	Cont Reduction = iota
	Halt
)

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

// Reduce invokes fun on each element in the slice with the accumulator.
func Reduce[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) Accumulator, accumulator Accumulator) Accumulator {
	return ReduceWhile(elements, func(element Element, accumulator Accumulator) (Reduction, Accumulator) {
		return Cont, fun(element, accumulator)
	}, accumulator)
}

// Map invokes fun on each element in the slice.
func Map[Element any](elements []Element, fun func(Element) Element) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		return append(accumulator, fun(element))
	}, make([]Element, 0))
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

// Reject returns elements excluding those where fun returns true.
func Reject[Element any](elements []Element, fun func(Element) bool) []Element {
	return Reduce(elements, func(element Element, accumulator []Element) []Element {
		if !fun(element) {
			return append(accumulator, element)
		}
		return accumulator
	}, make([]Element, 0))
}

// IsMember checks if element exists in the slice.
func IsMember[Element comparable](elements []Element, member Element) bool {
	return ReduceWhile(elements, func(element Element, accumulator bool) (Reduction, bool) {
		if element == member {
			return Halt, true
		}
		return Cont, false
	}, false)
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

type minMax[Element constraints.Ordered] struct {
	min Element
	max Element
}

// MinMax returns the minimum and maximum element in the slice.
func MinMax[Element constraints.Ordered](elements []Element) (Element, Element) {
	result := Reduce(elements, func(element Element, accumulator minMax[Element]) minMax[Element] {
		if element < accumulator.min {
			accumulator.min = element
		}
		if element > accumulator.max {
			accumulator.max = element
		}
		return accumulator
	}, minMax[Element]{min: elements[0], max: elements[0]})

	return result.min, result.max
}

// Sum returns the sum of all elements.
func Sum[Element constraints.Ordered](elements []Element) Element {
	return Reduce(elements[1:], func(element Element, accumulator Element) Element {
		return element + accumulator
	}, elements[0])
}

// Any returns true if fun returns true for at least one element in the slice.
func Any[Element any](elements []Element, fun func(Element) bool) bool {
	return Reduce(elements, func(element Element, accumulator bool) bool {
		if fun(element) {
			accumulator = true
		}
		return accumulator
	}, false)
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

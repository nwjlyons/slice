package slice

import "golang.org/x/exp/constraints"

// Reduce invokes fun for each element in the slice with the accumulator.
func Reduce[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) Accumulator, accumulator Accumulator) Accumulator {
	for _, element := range elements {
		accumulator = fun(element, accumulator)
	}
	return accumulator
}

// Max returns the maximal element in the slice.
func Max[Element constraints.Ordered](elements []Element) Element {
	return Reduce(elements, func(element Element, max Element) Element {
		if element > max {
			max = element
		}
		return max
	}, elements[0])
}

// Min returns the minimal element in the slice.
func Min[Element constraints.Ordered](elements []Element) Element {
	return Reduce(elements, func(element Element, min Element) Element {
		if element < min {
			min = element
		}
		return min
	}, elements[0])
}

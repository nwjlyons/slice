package slice

import (
	"golang.org/x/exp/constraints"
)

type Reduction int

const (
	Cont Reduction = iota
	Halt
)

// ReduceWhile reduces slice until fun returns Halt.
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

// Reduce invokes fun for each element in the slice with the accumulator.
func Reduce[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) Accumulator, accumulator Accumulator) Accumulator {
	return ReduceWhile(elements, func(element Element, accumulator Accumulator) (Reduction, Accumulator) {
		return Cont, fun(element, accumulator)
	}, accumulator)
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

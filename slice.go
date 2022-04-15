package slice

// Reduce invokes fun for each element in the slice with the accumulator.
func Reduce[Element any, Accumulator any](elements []Element, fun func(Element, Accumulator) Accumulator, accumulator Accumulator) Accumulator {
	for _, element := range elements {
		accumulator = fun(element, accumulator)
	}
	return accumulator
}

package main

import "fmt"

/*
	Generic allow call function with more types
*/

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, val := range m {
		sum += val
	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, val := range m {
		sum += val
	}

	return sum
}

// Generic function accept both int64 and float64 => 2 funcs -> 1 generic func
// Note: Key of map only accept Comparable type, generic type must support all operations
func SumGeneric[K comparable, V int64 | float64](m map[K]V) V {
	var sum V

	for _, val := range m {
		sum += val
	}

	return sum
}

type Number interface {
	int64 | float64
}

// Generic with type constraint
func SumGenericTypeConstraint[K comparable, V Number](m map[K]V) V {
	var sum V

	for _, val := range m {
		sum += val
	}

	return sum
}

func main() {
	ints := map[string]int64{
		"first":  3,
		"second": 4,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sum : %v and %v\n", SumGeneric(ints), SumGeneric(floats))

	fmt.Printf("Type constraint Generic Sum : %v and %v\n", SumGenericTypeConstraint(ints), SumGenericTypeConstraint(floats))

}

package main

import "fmt"

func main() {
	ints := map[string]int{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("非泛型计算结果，SumInts: %v, SumFloats: %v\n", SumNumbers(ints), SumNumbers(floats))
}

type Number interface {
	int | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

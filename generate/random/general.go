package random

import "fmt"

type conditioner func(float64) bool
type combiner func(float64, float64) float64

func assert(test bool, msg string) {
	if !test {
		panic(fmt.Sprintf("AssertionError: %s.", msg))
	}
}

func reduce(items []float64, f combiner) float64 {
	assert(len(items) > 0, "Cannot reduce empty slice")
	out := items[0]
	for i := 1; i < len(items); i++ {
		out = f(out, items[i])
	}
	return out
}

func sum(items []float64) float64 {
	return reduce(items, func(x, y float64) float64 { return x + y })
}

func product(items []float64) float64 {
	return reduce(items, func(x, y float64) float64 { return x * y })
}

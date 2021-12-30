package random

import (
	"math"
	"math/rand"
)

type any interface{}
type DiscreteRandomVariates []any
type DiscreteDistribution interface {
	PMF(x float64) float64
	CMF(x float64) float64
	PPF(p float64) float64
	Sample() float64
}

// A RandomVariable X on N elements is

// A RandomVariable X on N elements in the range [Lower, Upper] can be distributed with
// the DiscreteUniformDistribution if all samples are equally likely
type DiscreteUniform struct {
	N, Lower, Upper float64
}

func (dist *DiscreteUniform) PMF(x float64) float64 {
	if !(dist.Lower <= x && x <= dist.Upper) {
		return 0.0
	}
	return 1.0 / dist.N
}

func (dist *DiscreteUniform) CDF(x float64) float64 {
	if x < dist.Lower {
		return 0.0
	} else if x >= dist.N {
		return 1.0
	} else {
		return (math.Floor(x) + 1.0) / dist.N
	}
}

func (dist *DiscreteUniform) PPF(p float64) float64 {
	return math.Floor(dist.N * p)
}

func (dist *DiscreteUniform) Sample() float64 {
	return math.Floor(dist.PPF(rand.Float64()))
}

// We can make a DiscreteUniform-ly distributed RandomVariable X for a given set of
// DiscreteRandomVariates
func DiscreteUniformFromSlice(s DiscreteRandomVariates) *DiscreteUniform {
	N := float64(len(s))
	return &DiscreteUniform{N: N, Lower: 0, Upper: N - 1}
}

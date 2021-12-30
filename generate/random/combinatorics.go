package random

func UniformChoice(options []any) any {
	U := DiscreteUniformFromSlice(options)
	return options[uint(U.Sample())]
}

func UniformChoices(k uint, options []any, dist DiscreteDistribution) []any {
	U := DiscreteUniformFromSlice(options)
	out := make([]any, k)
	for i := uint(0); i < k; i++ {
		out[i] = options[uint(U.Sample())]
	}
	return out
}

package generate

type StateMachine interface {
	Next(i uint) State
	NextN(i, n uint) State
	Current(i uint) State
}

type BooleanNetwork struct {
	states      []State
	actions     [][]Coupling
	transitions [][]Probability
	policy      func(int) bool
}

/*
Make a random array of booleans based on the given distribution and a threshold
it must beat.
*/
func ThresholdBooleanNetwork(n uint, threshold int) BooleanNetwork {
	policy := func(x int) bool { return x >= threshold }

	// Make random arrays for the nodes and their transitions
	nodes := randomStates(n, choices)
	actions := make([][]Coupling, n)
	for row := range actions {
		actions[row] = random.Choice()
	}
	return BooleanNetwork{nodes, transition, policy}
}

func assert(test bool, msg string) {
	if !test {
		panic(msg)
	}
}

func dot(u, v []int) int {
	assert(len(u) == len(v), "dot product of vectors with different lengths.")

	out := 0
	for i := 0; i < len(u); i++ {
		out += u[i] * v[i]
	}
	return out
}

func coinFlip(G *BooleanNetwork, i uint) State {
	T, A := G.transitions[i], G.actions[i]

}

func (G *BooleanNetwork) Next(i uint) {
	return
}

func (G *BooleanNetwork) NextN(i, n uint) []State {

}

func (G *BooleanNetwork) Current(i uint) {
	return G.states[i]
}

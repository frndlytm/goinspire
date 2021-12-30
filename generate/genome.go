package generate

/*
	A Genome stores the lineage of a genetic algorithm sample. As we progress through
	the program execution, we gradually accumulate more and more trials resulting from
	splicing together two previous executions
*/
type Genome struct {
	xchrom *Genome
	ychrom *Genome
	child  interface{}
	score  int
}

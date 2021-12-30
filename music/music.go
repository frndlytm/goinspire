package music

type Any interface{}

type Transposable interface {
	Transpose(t int8) *Transposable
}

type Repeatable interface {
	Repeat(n uint8) *Repeatable
}

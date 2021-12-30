package generate

type Any interface{}

type Coupling int8

const (
	Inhibit Coupling = iota - 1
	Ignore
	Excite
)

type State int8

const (
	Off State = iota
	On
)

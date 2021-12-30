package music

import "github.com/go-audio/midi"

/**/
type TimeSignature struct {
	Top, Bottom int
}

func CommonTime() TimeSignature {
	return TimeSignature{4, 4}
}

/**/
type Rhythm struct {
	Events []*midi.Event
}

type PolyRhythm struct {
	Rhythms []Rhythm
}

func NoverK(T TimeSignature, N, K int, quant int) PolyRhythm {
	var q float64 = 1 / float64(quant)

}

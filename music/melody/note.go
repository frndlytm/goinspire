package melody

import (
	"fmt"
	"sort"
)

type Tone uint8

const (
	// Using iota, establish the base 12-Tones
	C Tone = iota
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
	A
	ASharp
	B

	// Set common names for sharps/flats to the same tones
	DFlat  = CSharp
	EFlat  = DSharp
	ESharp = F
	FFlat  = E
	GFlat  = FSharp
	AFlat  = GSharp
	BFlat  = ASharp
	BSharp = C
)

func (T *Tone) String() string {
	var s string
	switch *T {
	case C:
		s = "C"
	case CSharp:
		s = "C#/Db"
	case D:
		s = "D"
	case DSharp:
		s = "D#/Eb"
	case E:
		s = "E"
	case F:
		s = "F"
	case FSharp:
		s = "F#/Gb"
	case G:
		s = "G"
	case GSharp:
		s = "G#/Ab"
	case A:
		s = "A"
	case ASharp:
		s = "A#/Bb"
	case B:
		s = "B"
	}
	return s
}

type Interval uint8

const (
	Unison Interval = iota
	Min2
	Maj2
	Min3
	Maj3
	Maj4
	Min5
	Maj5
	Min6
	Maj6
	Min7
	Maj7
	Octave
)

type Note struct {
	Tone   Tone
	Octave int8
}

func (N *Note) String() string {
	return fmt.Sprintf("%s%d", N.Tone.String(), N.Octave)
}

func divmod(n, d int8) (q, r int8) {
	q, r = n/d, n%d
	return
}

func FromMidiValue(m uint8) Note {
	if m > 127 {
		panic(fmt.Sprintf("MidiOverflow: %d", m))
	}
	q, r := divmod(int8(m), 12)
	return Note{Tone: Tone(r), Octave: q - 1}
}

func (N Note) MidiValue() uint8 {
	// C0 has MidiValue() == 12, so Note{0, 0}.MidiValue() == 12
	if N.Octave < -1 {
		panic("MidiUnderflow: Note.Octave too low.")
	}
	return uint8(int8(N.Tone) + 12*(N.Octave+1))
}

func (N Note) Diff(O Note) int8            { return int8(O.MidiValue() - N.MidiValue()) }
func (N Note) Up(interval Interval) Note   { return N.Transpose(int8(interval)) }
func (N Note) Down(interval Interval) Note { return N.Transpose(-int8(interval)) }
func (N Note) Transpose(t int8) Note {
	return FromMidiValue(uint8(int8(N.MidiValue()) + t))
}

type ByMidiValue []Note

func (S ByMidiValue) Len() int           { return len(S) }
func (S ByMidiValue) Swap(i, j int)      { S[i], S[j] = S[j], S[i] }
func (S ByMidiValue) Less(i, j int) bool { return S[i].MidiValue() < S[j].MidiValue() }
func (S ByMidiValue) Insert(n Note) {
	// Partition S where we want to insert Note n
	k := sort.Search(S.Len(), func(i int) bool {
		return S[i].MidiValue() >= n.MidiValue()
	})
	left, right := S[:k], S[k:]

	// Update S to be an empty slice that's longer by 1 element.
	S = make(ByMidiValue, len(S)+1)
	copy(S[:k], left)
	S[k] = n
	copy(S[k+1:], right)
}

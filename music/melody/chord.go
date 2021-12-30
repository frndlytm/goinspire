package melody

type Chord struct {
	Tone  Tone
	Notes []Note
}

/*
Named Chords
*/
func MajorTriad(root Note) Chord {
	return Chord{
		Tone:  root.Tone,
		Notes: []Note{root, root.Up(Maj3), root.Up(Maj5)},
	}
}

func MinorTriad(root Note) Chord {
	return Chord{
		Tone:  root.Tone,
		Notes: []Note{root, root.Up(Min3), root.Up(Maj5)},
	}
}

func DiminishedTriad(root Note) Chord {
	return Chord{
		Tone:  root.Tone,
		Notes: []Note{root, root.Up(Min3), root.Up(Min5)},
	}
}

func AugmentedTriad(root Note) Chord {
	return Chord{
		Tone:  root.Tone,
		Notes: []Note{root, root.Up(Maj3), root.Up(Min6)},
	}
}

func (C *Chord) Root() Note {
	return C.Notes[0]
}

/*
	Add a Note to the Chord via an interval. This is commonly modeled
		in Chord builder notation, like CMaj7add9, in the way musicians
		notate chords.

	NOTE: An Inversion must retain the sort of the Notes in the Chord by
		their MidiValues, therefore, modifying an Inverted chord with an Add
		operation is an undefined operation.

	Examples:
		MajorTriad(D).Add(Maj7).Add(Maj9) == 'CMaj7add9'
*/
func (C *Chord) Add(interval Interval) *Chord {
	if isInverted(C) {
		panic("ChordBuildError: Cannot Add to Inverted chord")
	}
	ByMidiValue(C.Notes).Insert(C.Root().Up(interval))
	return C
}

func (C Chord) Transpose(t int8) Chord {
	notes := make([]Note, len(C.Notes))
	for i := range notes {
		notes[i] = C.Notes[i].Transpose(t)
	}
	return Chord{Notes: notes}
}

/*
	An Inversion modifies the Root of a Chord by slice rotation.
	An inversionUp takes the Bass/Root note of the chord and pushes
		it up an Octave
	An inversionDown takes the top note of the chord and pulls
		it down an Octave.
*/
func isInverted(C *Chord) bool {
	return C.Tone != C.Root().Tone
}

func (C Chord) Inversion(i int8) Chord {
	notes := C.Notes

	var doInversion func([]Note) []Note
	p, i := parity(i), abs(i)
	if p == -1 {
		doInversion = inversionDown
	} else {
		doInversion = inversionUp
	}

	for i > 0 {
		notes = doInversion(notes)
		i--
	}
	return Chord{Tone: C.Tone, Notes: notes}
}

func inversionUp(notes []Note) []Note {
	N, rest := notes[0], notes[1:]
	return append(rest, N.Up(Octave))
}

func inversionDown(notes []Note) []Note {
	// Pop the last Note N from the rest of the notes
	end := len(notes) - 1
	rest, N := notes[:end], notes[end]

	// Put the inverted element at the state and copy the rest
	out := make([]Note, len(notes))
	out[0] = N.Down(Octave)
	copy(out[1:], rest)

	return out
}

func parity(i int8) int8 {
	if i < 0 {
		return -1
	} else if i == 0 {
		return 0
	} else {
		return 1
	}
}

func abs(i int8) int8 {
	switch parity(i) {
	case -1:
		return -i
	default:
		return i
	}
}

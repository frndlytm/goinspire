package melody_tests

import (
	"fmt"
	"testing"

	"github.com/frndlytm/go-inspire/music/melody"
	"github.com/stretchr/testify/assert"
)

func TestNamedChords(t *testing.T) {
	cases := []struct {
		name    string
		root    melody.Note
		factory func(melody.Note) melody.Chord
		notes   []melody.Note
	}{
		{
			name:    "Test: melody.MajorTriad(%v)",
			root:    melody.Note{Tone: melody.A, Octave: 4},
			factory: melody.MajorTriad,
			notes: []melody.Note{
				{Tone: melody.A, Octave: 4},
				{Tone: melody.CSharp, Octave: 5},
				{Tone: melody.E, Octave: 5},
			},
		},
		{
			name:    "Test: melody.MinorTriad(%v)",
			root:    melody.Note{Tone: melody.A, Octave: 4},
			factory: melody.MinorTriad,
			notes: []melody.Note{
				{Tone: melody.A, Octave: 4},
				{Tone: melody.C, Octave: 5},
				{Tone: melody.E, Octave: 5},
			},
		},
		{
			name:    "Test: melody.AugmentedTriad(%v)",
			root:    melody.Note{Tone: melody.A, Octave: 4},
			factory: melody.AugmentedTriad,
			notes: []melody.Note{
				{Tone: melody.A, Octave: 4},
				{Tone: melody.CSharp, Octave: 5},
				{Tone: melody.ESharp, Octave: 5},
			},
		},
		{
			name:    "Test: melody.DiminishedTriad(%v)",
			root:    melody.Note{Tone: melody.A, Octave: 4},
			factory: melody.DiminishedTriad,
			notes: []melody.Note{
				{Tone: melody.A, Octave: 4},
				{Tone: melody.C, Octave: 5},
				{Tone: melody.EFlat, Octave: 5},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(c.name, c.root), func(t *testing.T) {
			chord := c.factory(c.root)
			assert.ElementsMatch(t, chord.Notes, c.notes)
		})
	}
}

func TestChordInversion(t *testing.T) {
	cases := []struct {
		name          string
		chord         melody.Chord
		numInversions int8
		expectedRoot  melody.Note
	}{
		{
			name:          "Test: Amin 1st Inversion",
			chord:         melody.MinorTriad(melody.Note{Tone: melody.A, Octave: 4}),
			numInversions: 1,
			expectedRoot:  melody.Note{Tone: melody.C, Octave: 5},
		},
		{
			name:          "Test: Amin 2nd Inversion",
			chord:         melody.MinorTriad(melody.Note{Tone: melody.A, Octave: 4}),
			numInversions: 2,
			expectedRoot:  melody.Note{Tone: melody.E, Octave: 5},
		},
		{
			name:          "Test: Amin -1st Inversion",
			chord:         melody.MinorTriad(melody.Note{Tone: melody.A, Octave: 4}),
			numInversions: -1,
			expectedRoot:  melody.Note{Tone: melody.E, Octave: 4},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			chord := c.chord.Inversion(c.numInversions)
			assert.Equal(t, c.chord.Tone, chord.Tone)
			assert.Equal(t, c.expectedRoot, chord.Root())
		})
	}
}

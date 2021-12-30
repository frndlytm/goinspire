package melody_tests

import (
	"testing"

	"github.com/frndlytm/go-inspire/music/melody"
	"github.com/stretchr/testify/assert"
)

func TestNote_ToMidiValue(t *testing.T) {
	cases := []struct {
		Name string
		In   melody.Note
		Out  uint8
	}{
		{
			Name: "Test: C-1.ToMidiValue() (Midi Min)",
			In:   melody.Note{Tone: melody.C, Octave: -1},
			Out:  uint8(0),
		},
		{
			Name: "Test: A0.ToMidiValue() (Piano Min)",
			In:   melody.Note{Tone: melody.A, Octave: 0},
			Out:  uint8(21),
		},
		{
			Name: "Test: C4.ToMidiValue() (Middle C)",
			In:   melody.Note{Tone: melody.C, Octave: 4},
			Out:  uint8(60),
		},
		{
			Name: "Test: C8.ToMidiValue() (Piano Max)",
			In:   melody.Note{Tone: melody.C, Octave: 8},
			Out:  uint8(108),
		},
		{
			Name: "Test: G9.ToMidiValue() (Midi Max)",
			In:   melody.Note{Tone: melody.G, Octave: 9},
			Out:  uint8(127),
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, c.In.MidiValue(), c.Out)
		})
	}
}

func TestNote_FromMidiValue(t *testing.T) {
	cases := []struct {
		Name string
		Out  melody.Note
		In   uint8
	}{
		{
			Name: "Test: C-1.FromMidiValue() (Midi Min)",
			Out:  melody.Note{Tone: melody.C, Octave: -1},
			In:   uint8(0),
		},
		{
			Name: "Test: A0.FromMidiValue() (Piano Min)",
			Out:  melody.Note{Tone: melody.A, Octave: 0},
			In:   uint8(21),
		},
		{
			Name: "Test: C4.FromMidiValue() (Middle C)",
			Out:  melody.Note{Tone: melody.C, Octave: 4},
			In:   uint8(60),
		},
		{
			Name: "Test: C8.FromMidiValue() (Piano Max)",
			Out:  melody.Note{Tone: melody.C, Octave: 8},
			In:   uint8(108),
		},
		{
			Name: "Test: G9.FromMidiValue() (Midi Max)",
			Out:  melody.Note{Tone: melody.G, Octave: 9},
			In:   uint8(127),
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, melody.FromMidiValue(c.In), c.Out)
		})
	}
}

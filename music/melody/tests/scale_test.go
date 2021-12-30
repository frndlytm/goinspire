package melody_tests

import (
	"testing"

	"github.com/frndlytm/go-inspire/music/melody"
	"github.com/stretchr/testify/assert"
)

func TestBooleanScale(t *testing.T) {
	cases := []struct {
		Name string
		In   melody.Scale
		Out  [12]bool
	}{
		{
			Name: "TestBooleanScale: Chromatic",
			In:   melody.Chromatic,
			Out:  [12]bool{true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			Name: "TestBooleanScale: WholeTone",
			In:   melody.WholeTone,
			Out:  [12]bool{true, false, true, false, true, false, true, false, true, false, true, false},
		},
		{
			Name: "TestBooleanScale: Major",
			In:   melody.Major,
			Out:  [12]bool{true, false, true, false, true, true, false, true, false, true, false, true},
		},
		{
			Name: "TestBooleanScale: MajorPentatonic",
			In:   melody.MajorPentatonic,
			Out:  [12]bool{true, false, true, false, true, false, false, true, false, true, false, false},
		},
		{
			Name: "TestBooleanScale: Minor",
			In:   melody.Minor,
			Out:  [12]bool{true, false, true, true, false, true, false, true, true, false, true, false},
		},
		{
			Name: "TestBooleanScale: HarmonicMinor",
			In:   melody.HarmonicMinor,
			Out:  [12]bool{true, false, true, true, false, true, false, true, true, false, false, true},
		},
		{
			Name: "TestBooleanScale: MinorPentatonic",
			In:   melody.MinorPentatonic,
			Out:  [12]bool{true, false, false, true, false, true, false, true, false, false, true, false},
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			assert.Equal(t, melody.BooleanScale(c.In), c.Out)
		})
	}
}

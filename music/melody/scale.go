package melody

type Scale []Interval

var (
	Chromatic       = Scale{Unison, Min2, Maj2, Min3, Maj3, Maj4, Min5, Maj5, Min6, Maj6, Min7, Maj7}
	WholeTone       = Scale{Unison, Maj2, Maj3, Min5, Min6, Min7}
	Major           = Scale{Unison, Maj2, Maj3, Maj4, Maj5, Maj6, Maj7}
	MajorPentatonic = Scale{Unison, Maj2, Maj3, Maj5, Maj6}
	Minor           = Scale{Unison, Maj2, Min3, Maj4, Maj5, Min6, Min7}
	HarmonicMinor   = Scale{Unison, Maj2, Min3, Maj4, Maj5, Min6, Maj7}
	MinorPentatonic = Scale{Unison, Min3, Maj4, Maj5, Min7}
)

func BooleanScale(intervals Scale) (scale [12]bool) {
	for _, i := range intervals {
		scale[i%12] = true
	}
	return
}

func (S Scale) Match(notes []Note) bool {
	if len(S) != len(notes) {
		return false
	} else {
		// TODO: This doesn't work exactly if our notes are shifted
		//   to a different mode
		root := notes[0]
		for i := 1; i < len(S); i++ {
			if root.Diff(notes[i]) != int8(S[i]) {
				return false
			}
		}
	}
	return true
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/frndlytm/goinspire/music/melody"
	"github.com/go-audio/midi"
)

func main() {
	// Command-Line Flags
	var key melody.Note
	flag.Var(
		&NoteVar{&key},
		"Key",
		"Choose the starting note of the scale.",
	)
	var scale *melody.Scale
	flag.Var(
		&ScaleVar{scale},
		"Scale",
		"Choose from a predefined set of scaels (see exported constants)",
	)

	var destination = flag.String(
		"Destination",
		"output.mid",
		"Path to output the file (example: /path/to/output.mid)",
	)
	var velocity = flag.Int(
		"Velocity",
		90,
		"A constant velocity for every note in the scale.",
	)

	flag.Parse()

	encoder := midi.NewEncoder(mustCreate(*destination), midi.SingleTrack, 96)
	track := encoder.NewTrack()

	var note int
	for _, interval := range *scale {
		note = int(key.Up(interval).MidiValue())
		track.Add(1, midi.NoteOn(0, note, *velocity))
		track.Add(1, midi.NoteOff(0, note))
	}
	track.Add(1, midi.EndOfTrack())

	if err := encoder.Write(); err != nil {
		log.Fatal(err)
	}
}

func mustCreate(fp string) *os.File {
	f, err := os.Create(fp)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

type ScaleVar struct {
	scale *melody.Scale
}

func (v ScaleVar) String() string {
	switch scale := v.scale; {
	case reflect.DeepEqual(scale, melody.Chromatic):
		return "Chromatic"
	case reflect.DeepEqual(scale, melody.WholeTone):
		return "WholeTone"
	case reflect.DeepEqual(scale, melody.Major):
		return "Major"
	case reflect.DeepEqual(scale, melody.MajorPentatonic):
		return "MajorPentatonic"
	case reflect.DeepEqual(scale, melody.Minor):
		return "Minor"
	case reflect.DeepEqual(scale, melody.HarmonicMinor):
		return "HarmonicMinor"
	case reflect.DeepEqual(scale, melody.MinorPentatonic):
		return "MinorPentatonic"
	default:
		return "None"
	}
}

func (v ScaleVar) Set(scale string) error {
	switch s := strings.ToLower(scale); {
	case s == "chromatic":
		*v.scale = melody.Chromatic
	case s == "wholetone":
		*v.scale = melody.WholeTone
	case s == "major":
		*v.scale = melody.Major
	case s == "majorpentatonic":
		*v.scale = melody.MajorPentatonic
	case s == "minor":
		*v.scale = melody.Minor
	case s == "harmonicminor":
		*v.scale = melody.HarmonicMinor
	case s == "minorpentatonic":
		*v.scale = melody.MinorPentatonic
	default:
		return errors.New(fmt.Sprintf("ScaleNotFound: %s", scale))
	}
	return nil
}

type NoteVar struct {
	note *melody.Note
}

func (v NoteVar) String() string {
	return v.note.String()
}

func (v NoteVar) Set(scale string) error {
	switch s := strings.ToLower(scale); {
	case s == "chromatic":
		*v.scale = melody.Chromatic
	case s == "wholetone":
		*v.scale = melody.WholeTone
	case s == "major":
		*v.scale = melody.Major
	case s == "majorpentatonic":
		*v.scale = melody.MajorPentatonic
	case s == "minor":
		*v.scale = melody.Minor
	case s == "harmonicminor":
		*v.scale = melody.HarmonicMinor
	case s == "minorpentatonic":
		*v.scale = melody.MinorPentatonic
	default:
		return errors.New(fmt.Sprintf("ScaleNotFound: %s", scale))
	}
	return nil
}

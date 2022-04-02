package wordle

import (
	"fmt"
	"github.com/mgutz/ansi"
	"io"
)

type LetterMatch int

const (
	Unknown LetterMatch = iota
	NoMatch
	WrongPosition
	ExactMatch
)

type GameLetter struct {
	char  int32
	match LetterMatch
}

func (letter *GameLetter) Display(output io.Writer) error {
	_, err := fmt.Fprint(output, letter.color(), string(letter.char), " ")
	return err
}

func (letter *GameLetter) color() string {
	switch letter.match {
	default:
		return ansi.LightWhite
	case NoMatch:
		return ansi.Red
	case WrongPosition:
		return ansi.Yellow
	case ExactMatch:
		return ansi.Green
	}
}

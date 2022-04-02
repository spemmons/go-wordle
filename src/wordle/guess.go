package wordle

import (
	"fmt"
	"github.com/mgutz/ansi"
	"io"
)

type GameGuess []GameLetter

func NewGuess(str string) (GameGuess, error) {
	word, err := NewWord(str)
	if err != nil {
		return nil, err
	}
	guess := make([]GameLetter, len(word.string))
	for index, char := range word.string {
		guess[index].char = char
	}
	return guess, nil
}

func (guess *GameGuess) Display(output io.Writer) error {
	for _, letter := range *guess {
		if err := letter.Display(output); err != nil {
			return err
		}
	}
	_, err := fmt.Fprint(output, ansi.Reset)
	return err
}

func (guess *GameGuess) includes(targetLetter GameLetter) bool {
	for _, guessLetter := range *guess {
		if guessLetter.char == targetLetter.char {
			return true
		}
	}
	return false
}

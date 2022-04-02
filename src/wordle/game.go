package wordle

import (
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"strings"
)

type GameStatus int

const (
	InProgress GameStatus = iota
	Failure
	Success
)

type GameState struct {
	target  *GameWord
	guesses []GameGuess
	letters []GameLetter
	status  GameStatus
}

func NewGame(str string) (*GameState, error) {
	target, err := NewWord(str)
	if err != nil {
		return nil, err
	} else {
		game := new(GameState)
		game.target = target
		game.guesses = make([]GameGuess, 0, 6)
		game.letters = makeLetters()
		return game, nil
	}
}

func (game *GameState) Status() GameStatus {
	return game.status
}

func (game *GameState) Display(output io.Writer) (err error) {
	if len(game.guesses) > 0 {
		err = firstError(
			onlyError(fmt.Fprintln(output, "guesses:")),
			displayGuesses(output, game.guesses),
			onlyError(fmt.Fprint(output, "\nletters: ")),
			displayLetters(output, game.letters),
			onlyError(fmt.Fprintln(output, ansi.Reset)))
	}
	return
}

func (game *GameState) AddGuess(str string) error {
	if game.status != InProgress {
		return nil
	}
	guess, err := NewGuess(str)
	if err == nil {
		evaluateGuess(game.target, guess)
		if !checkUniqueGuess(game.guesses, guess) {
			err = errors.New("guess already used")
		} else if !allHintsUsed(game.letters, guess) {
			err = errors.New("guess missing known hints")
		} else {
			updateLetters(game.letters, guess)
			game.guesses = append(game.guesses, guess)
			if exactMatch(guess) {
				game.status = Success
			} else if len(game.guesses) >= cap(game.guesses) {
				game.status = Failure
			}
		}
	}
	return err
}

func evaluateGuess(target *GameWord, guess GameGuess) {
	for index, letter := range guess {
		match := strings.Index(target.string, string(letter.char))
		if match == index {
			letter.match = ExactMatch
		} else if match >= 0 {
			letter.match = WrongPosition
		} else {
			letter.match = NoMatch
		}
		guess[index] = letter
	}
}

func checkUniqueGuess(guesses []GameGuess, guess GameGuess) bool {
	if len(guesses) == 0 {
		return true
	}
	for _, previous := range guesses {
		for index := range previous {
			if previous[index].char != guess[index].char {
				return true
			}
		}
	}
	return false
}

func allHintsUsed(hintLetters []GameLetter, guess GameGuess) bool {
	for _, hintLetter := range hintLetters {
		if (hintLetter.match == ExactMatch || hintLetter.match == WrongPosition) && !guess.includes(hintLetter) {
			return false
		}
	}
	return true
}

func updateLetters(hintLetters []GameLetter, guess GameGuess) {
	for _, guessLetter := range guess {
		updateLetter(hintLetters, guessLetter)
	}
}

func updateLetter(hintLetters []GameLetter, adjustedLetter GameLetter) {
	index := adjustedLetter.char - 'A'
	hintLetters[index].match = bestMatch(hintLetters[index].match, adjustedLetter.match)
}

func bestMatch(oldMatch LetterMatch, newMatch LetterMatch) LetterMatch {
	if oldMatch == ExactMatch || newMatch == ExactMatch {
		return ExactMatch
	} else {
		return newMatch
	}
}

func exactMatch(guess GameGuess) bool {
	for _, letter := range guess {
		if letter.match != ExactMatch {
			return false
		}
	}
	return true
}

func makeLetters() []GameLetter {
	first, last := 'A', 'Z'
	letters := make([]GameLetter, last-first+1)
	for index := range letters {
		letters[index].char = first + int32(index)
	}
	return letters
}

func displayGuesses(output io.Writer, guesses []GameGuess) (err error) {
	for index, guess := range guesses {
		err = firstError(err,
			onlyError(fmt.Fprintf(output, "%v - ", index+1)),
			guess.Display(output),
			onlyError(fmt.Fprintln(output)))
	}
	return
}

func displayLetters(output io.Writer, letters []GameLetter) (err error) {
	for _, letter := range letters {
		err = firstError(err, letter.Display(output))
	}
	return
}

func firstError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func onlyError(_ any, err error) error {
	return err
}

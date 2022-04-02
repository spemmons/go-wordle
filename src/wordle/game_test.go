package wordle

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestNewGame(t *testing.T) {
	game, err := NewGame("today")
	assert.Nil(t, err)
	assert.Equal(t, "TODAY", game.target.string)
	assert.Equal(t, InProgress, game.Status())
	assert.Equal(t, 0, len(game.guesses))
	assert.Equal(t, 26, len(game.letters))

	_, err = NewGame("tiny")
	assert.Equal(t, err, errors.New("invalid word"))
}

func TestGameSuccess(t *testing.T) {
	game, _ := NewGame("today")

	err := game.AddGuess("today")
	assert.Nil(t, err)
	assert.Equal(t, game.Status(), Success)
	assert.Equal(t, len(game.guesses), 1)
}

func TestGameFailure(t *testing.T) {
	game, _ := NewGame("today")

	err := game.AddGuess("todaa")
	assert.Nil(t, err)
	assert.Equal(t, game.Status(), InProgress)

	err = game.AddGuess("todaa")
	assert.Equal(t, errors.New("guess already used"), game.AddGuess("todaa"))

	assert.Equal(t, errors.New("guess missing known hints"), game.AddGuess("xxxxx"))

	assert.Nil(t, game.AddGuess("todab"))
	assert.Equal(t, InProgress, game.Status())

	assert.Nil(t, game.AddGuess("todac"))
	assert.Equal(t, InProgress, game.Status())

	assert.Nil(t, game.AddGuess("todad"))
	assert.Equal(t, InProgress, game.Status())

	assert.Nil(t, game.AddGuess("todae"))
	assert.Equal(t, InProgress, game.Status())

	assert.Nil(t, game.AddGuess("todaf"))
	assert.Equal(t, Failure, game.Status())

	assert.Nil(t, game.AddGuess("todaf"))
	assert.Equal(t, Failure, game.Status())
	assert.Equal(t, 6, len(game.guesses))
}

func TestGameDisplay(t *testing.T) {
	var buffer bytes.Buffer
	game, _ := NewGame("today")
	assert.Nil(t, game.Display(&buffer))
	assert.Equal(t, "", buffer.String())
	assert.Nil(t, game.AddGuess("today"))
	assert.Nil(t, game.Display(&buffer))

	re, _ := regexp.Compile("^guesses:\\n1 - .*T .*O .*D .*A .*Y .*\\n\\nletters: .*\n$")
	assert.True(t, re.MatchString(buffer.String()))
}

func TestFirstError(t *testing.T) {
	first, second := errors.New("first"), errors.New("second")
	assert.Equal(t, first, firstError(nil, first, second))
}

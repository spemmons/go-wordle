package wordle

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewGuess(t *testing.T) {
	guess, err := NewGuess("today")
	assert.Nil(t, err)
	assert.Equal(t, "[{84 0} {79 0} {68 0} {65 0} {89 0}]", fmt.Sprint(guess)) // TODO - make better test

	_, err = NewGuess("invalid")
	assert.Equal(t, errors.New("invalid word"), err)
}

type MockWriter struct {
	mock.Mock
}

func (*MockWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failure")
}

func TestGuessDisplay(t *testing.T) {
	var buffer bytes.Buffer
	guess, _ := NewGuess("today")
	assert.Nil(t, guess.Display(&buffer))
	assert.Equal(t, ansi.LightWhite+"T "+ansi.LightWhite+"O "+ansi.LightWhite+"D "+ansi.LightWhite+"A "+ansi.LightWhite+"Y "+ansi.Reset, buffer.String())

	failedWriter := new(MockWriter)
	assert.Equal(t, errors.New("write failure"), guess.Display(failedWriter))
}

func TestGuessIncludes(t *testing.T) {
	guess, _ := NewGuess("today")
	assert.True(t, guess.includes(GameLetter{char: 'A', match: Unknown}))
	assert.False(t, guess.includes(GameLetter{char: 'Z', match: Unknown}))
}

package wordle

import (
	"bytes"
	"github.com/mgutz/ansi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterColor(t *testing.T) {
	assert.Equal(t, ansi.LightWhite, (&GameLetter{char: 'A', match: Unknown}).color())
	assert.Equal(t, ansi.Red, (&GameLetter{char: 'A', match: NoMatch}).color())
	assert.Equal(t, ansi.Yellow, (&GameLetter{char: 'A', match: WrongPosition}).color())
	assert.Equal(t, ansi.Green, (&GameLetter{char: 'A', match: ExactMatch}).color())
}

func TestLetterDisplay(t *testing.T) {
	var buffer bytes.Buffer
	assert.Nil(t, (&GameLetter{char: 'A', match: ExactMatch}).Display(&buffer))
	assert.Equal(t, ansi.Green+"A ", buffer.String())
}

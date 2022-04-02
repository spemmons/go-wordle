package wordle

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWord(t *testing.T) {
	word, err := NewWord("today")
	assert.Nil(t, err)
	assert.Equal(t, "TODAY", word.String())

	_, err = NewWord("tiny")
	assert.Equal(t, errors.New("invalid word"), err)

	_, err = NewWord("toolarge")
	assert.Equal(t, errors.New("invalid word"), err)

	_, err = NewWord("non-word1")
	assert.Equal(t, errors.New("invalid word"), err)
}

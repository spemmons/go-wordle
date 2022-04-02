package cli

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go-wordle/src/wordle"
	"testing"
)

func TestPlayGame(t *testing.T) {
	var input bytes.Buffer
	var output bytes.Buffer
	game, err := wordle.NewGame("today")
	assert.Nil(t, err)

	assert.Nil(t, PlayGame(game, &input, &output))
	assert.Equal(t, "> ", output.String())
	assert.Equal(t, wordle.InProgress, game.Status())

	output.Reset()
	input.WriteString("test\n")
	assert.Nil(t, PlayGame(game, &input, &output))
	assert.Equal(t, wordle.InProgress, game.Status())
	assert.Equal(t, "> invalid word\n", output.String())

	output.Reset()
	input.Reset()
	input.WriteString("today\n")
	assert.Nil(t, PlayGame(game, &input, &output))
	assert.Equal(t, wordle.Success, game.Status())
}

package cli

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupAndPlayGame(t *testing.T) {
	var input bytes.Buffer
	var output bytes.Buffer

	assert.Nil(t, SetupAndPlayGame([]string{"test"}, &input, &output))
	assert.Equal(t, "usage: test <target>\n", output.String())

	output.Reset()
	assert.Nil(t, SetupAndPlayGame([]string{"test", "test"}, &input, &output))
	assert.Equal(t, "invalid word\n", output.String())

	output.Reset()
	assert.Nil(t, SetupAndPlayGame([]string{"test", "today"}, &input, &output))
	assert.Equal(t, "> ", output.String())
}

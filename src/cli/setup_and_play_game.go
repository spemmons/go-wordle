package cli

import (
	"fmt"
	"go-wordle/src/wordle"
	"io"
)

func SetupAndPlayGame(args []string, input io.Reader, output io.Writer) (err error) {
	if len(args) != 2 {
		_, err = fmt.Fprintf(output, "usage: %v <target>\n", args[0])
	} else {
		var game *wordle.GameState
		game, err = wordle.NewGame(args[1])
		if err != nil {
			_, err = fmt.Fprintln(output, err)
		} else {
			err = PlayGame(game, input, output)
		}
	}
	return
}

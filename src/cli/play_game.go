package cli

import (
	"bufio"
	"fmt"
	"go-wordle/src/wordle"
	"io"
	"strings"
)

func PlayGame(game *wordle.GameState, input io.Reader, output io.Writer) (err error) {
	err = game.Display(output)
	reader := bufio.NewReader(input)
	for err == nil && game.Status() == wordle.InProgress {
		_, err = fmt.Fprint(output, "> ")
		for err == nil {
			str, _ := reader.ReadString('\n')
			str = strings.TrimSuffix(str, "\n")
			if len(str) == 0 {
				return
			}
			guessErr := game.AddGuess(str)
			if guessErr == nil {
				break
			}
			_, err = fmt.Fprintln(output, guessErr)
		}

		err = game.Display(output)
	}
	return
}

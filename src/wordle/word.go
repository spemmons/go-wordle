package wordle

import (
	"errors"
	"regexp"
	"strings"
)

var validWordRegexp *regexp.Regexp

func init() {
	validWordRegexp, _ = regexp.Compile("^[[:alpha:]]{5}$")
}

type GameWord struct {
	string
}

func (word GameWord) String() string {
	return word.string
}

func NewWord(str string) (*GameWord, error) {
	if validWordRegexp.MatchString(str) {
		return &GameWord{strings.ToUpper(str)}, nil
	} else {
		return nil, errors.New("invalid word")
	}
}

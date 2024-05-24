package tokenizer

import (
	"fmt"
	"professor/pkg/language"
	"strings"
	"text/scanner"
)

type LanguageToken struct {
	token language.Tokenized
	depth int
}

func MakeLanguageTokensFromInput(input string) []*LanguageToken {
	//depth := 0

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Filename = "input"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	return []*LanguageToken{}
}

package main

import (
	_ "embed"
	"log"

	"github.com/antlr4-go/antlr/v4"

	parser "github.com/slonegd/go-st/g4"
	external_parser "github.com/slonegd/go-st/g4/external"
)

//go:embed 001_simpliest.st
var simpliest string

func main() {
	simpliest_main()
	simpliest_external_main()
}

func simpliest_main() {
	log.Print("\n\n\t\t simpliest_main")
	// Setup the input
	is := antlr.NewInputStream(simpliest)

	// Create the Lexer
	lexer := parser.NewstLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		log.Printf("%s (%q): %#v", lexer.SymbolicNames[t.GetTokenType()], t.GetText(), t)
	}
}

func simpliest_external_main() {
	log.Print("\n\n\t\t simpliest_external_main")
	// Setup the input
	is := antlr.NewInputStream(simpliest)

	// Create the Lexer
	lexer := external_parser.NewstLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		log.Printf("%s (%q): %#v", lexer.SymbolicNames[t.GetTokenType()], t.GetText(), t)
	}
}

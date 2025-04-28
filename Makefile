.PHONY: all vendor

gen_external:
	antlr4 -Dlanguage=Go ./antlr/external/ST_lexer.g4 ./antlr/external/ST_parser.g4

gen_antlr:
	antlr4 -Dlanguage=Go ./antlr/ST.g4 

vendor:
	go mod tidy
	go mod vendor

tree:
	antlr4-parse ./antlr/ST.g4 programm -gui < ./tests/001_simpliest.st 

example:
	go run ./tests/examples.go
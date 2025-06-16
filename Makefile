.PHONY: all vendor

gen-external:
	antlr4 -Dlanguage=Go ./antlr/external/ST_lexer.g4 ./antlr/external/ST_parser.g4

gen-antlr:
	antlr4 -visitor -no-listener -Dlanguage=Go ./antlr/ST.g4 

vendor:
	go mod tidy
	go mod vendor

tree:
	antlr4-parse ./antlr/ST.g4 program -gui < ./tests/006_floats.st

example:
	go run ./tests/examples.go

test:
	go test ./...

bench:
	go test -bench=. ./tests/
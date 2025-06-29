.PHONY: all vendor bench

gen-external:
	antlr4 -Dlanguage=Go ./antlr/external/ST_lexer.g4 ./antlr/external/ST_parser.g4

gen-antlr:
	antlr4 -visitor -no-listener -Dlanguage=Go ./antlr/ST.g4 

vendor:
	go mod tidy
	go mod vendor

tree:
	antlr4-parse ./antlr/ST.g4 program -gui < ./tests/003_iftest.st

example:
	go run ./tests/examples.go

test:
	go test ./...

bench:
	cd bench && go test -bench=. ./arithmetic_if && cd ..
	cd bench && go test -bench=. ./arithmetic_if_while && cd ..
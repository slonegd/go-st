.PHONY: all vendor bench

gen-antlr:
	antlr4 -visitor -no-listener -Dlanguage=Go ./antlr/ST.g4 

vendor:
	go mod tidy
	go mod vendor

tree:
	antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/001_simpliest.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/002_arithmetic.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/003_iftest.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/005_cast.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/006_floats.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/007_while.st
	# antlr4-parse ./antlr/ST.g4 pous -gui < ./tests/008_func.st

example:
	go run ./tests/examples.go

test:
	go test ./...

bench:
	cd bench && go test -bench=. ./arithmetic_if && cd ..
	cd bench && go test -bench=. ./arithmetic_if_while && cd ..
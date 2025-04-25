.PHONY: all vendor

gen_external:
	antlr4 -Dlanguage=Go ./g4/external/stParser.g4 ./g4/external/stLexer.g4

gen_g4:
	antlr4 -Dlanguage=Go ./g4/st.g4 

vendor:
	go mod tidy
	go mod vendor
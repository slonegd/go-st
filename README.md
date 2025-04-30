# go-st
IEC 61131-3 Structured Text транслятор скрипта, исполняемый на go

## Implemented

### Basic ST Syntax 
The antlr/ST.g4 file defines a grammar for a subset of ST, including program structure, variable declarations (limited to INT), assignment statements, and IF statements.
### Integer Variables
The code supports integer variables and basic arithmetic operations (+, -, *, /, MOD).
### Conditional Execution
The code implements IF, THEN, and ELSE statements for conditional execution.
### Interpreter
The st.go and visitor.go files provide the structure for an interpreter that executes the parsed ST code.
### antlr-based Parsing
The project uses antlr to create a parser and lexer.

## Not Implemented

### Data Types
The current implementation only supports INT. Adding support for BOOL, REAL, STRING, and TIME would be a high priority. Modify the type_name rule in antlr/ST.g4 and the variable declaration and assignment logic in visitor.go.
### More statement
Implement CASE, FOR, WHILE, and REPEAT statements will be useful. Add appropriate rules in antlr/ST.g4 and corresponding logic in visitor.go.
### Boolean expressions
Condition contains only expression, bool expression require.
### Functions and Function Blocks
Add support for functions and function blocks. This would involve extending the grammar in antlr/ST.g4 and implementing the necessary execution logic in st.go and visitor.go.
### Array
Need to implement Array type expression.
### And more...
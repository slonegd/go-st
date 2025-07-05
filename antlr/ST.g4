// файл сгенерирован в deepseek за несколько итераций
grammar ST;

options {
    contextSuperClass = CustomContext;
}

// Лексер
INT_TYPE_NAME   : 'SINT' | 'INT' | 'DINT' | 'LINT' | 'USINT' | 'UINT' | 'UDINT' | 'ULINT';
REAL_TYPE_NAME  : 'REAL' | 'LREAL';
TIME_TYPE_NAME  : 'TIME' | 'LTIME';
DATE_TYPE_NAME  : 'DATE' | 'LDATE';
TOD_TYPE_NAME   : 'TIME_OF_DAY' | 'TOD' | 'LTOD';
DT_TYPE_NAME    : 'DATE_AND_TIME' | 'DT' | 'LDT';
BOOL_TYPE_NAME  : 'BOOL';
BIT_TYPE_NAME   : 'BYTE' | 'WORD' | 'DWORD' | 'LWORD';
STRING_TYPE_NAME: 'STRING' ('[' [0-9]+ ']')? | 'WSTRING' ('[' [0-9]+ ']')?;

VAR         : 'VAR';
VAR_INPUT   : 'VAR_INPUT';
VAR_OUTPUT  : 'VAR_OUTPUT';
VAR_IN_OUT  : 'VAR_IN_OUT';
VAR_EXTERNAL: 'VAR_EXTERNAL';
VAR_GLOBAL  : 'VAR_GLOBAL';
VAR_TEMP    : 'VAR_TEMP';
END_VAR     : 'END_VAR';
FUNCTION    : 'FUNCTION';
END_FUNCTION: 'END_FUNCTION';
FUNCTION_BLOCK : 'FUNCTION_BLOCK';
END_FUNCTION_BLOCK : 'END_FUNCTION_BLOCK';
PROGRAM     : 'PROGRAM';
END_PROGRAM : 'END_PROGRAM';
IF          : 'IF';
THEN        : 'THEN';
ELSE        : 'ELSE';
ELSIF       : 'ELSIF';
END_IF      : 'END_IF';
FOR         : 'FOR';
TO          : 'TO';
BY          : 'BY';
DO          : 'DO';
END_FOR     : 'END_FOR';
WHILE       : 'WHILE';
END_WHILE   : 'END_WHILE';
REPEAT      : 'REPEAT';
UNTIL       : 'UNTIL';
END_REPEAT  : 'END_REPEAT';
CASE        : 'CASE';
OF          : 'OF';
END_CASE    : 'END_CASE';
RETURN      : 'RETURN';
NOT         : 'NOT';
AND         : 'AND';
OR          : 'OR';
XOR         : 'XOR';
ARRAY       : 'ARRAY';
STRUCT      : 'STRUCT';
END_STRUCT  : 'END_STRUCT';
TYPE        : 'TYPE';
END_TYPE    : 'END_TYPE';
CONSTANT    : 'CONSTANT';
RETAIN      : 'RETAIN';
NON_RETAIN  : 'NON_RETAIN';
CONTINUE    : 'CONTINUE';
EXIT        : 'EXIT';

ASSIGN      : ':=';
EQUAL       : '=';
NOT_EQUAL   : '<>';
LESS        : '<';
LESS_EQ     : '<=';
GREATER     : '>';
GREATER_EQ  : '>=';
PLUS        : '+';
MINUS       : '-';
MULT        : '*';
DIV         : '/';
MOD         : 'MOD';
POWER       : '**';
DOT         : '.';
COMMA       : ',';
COLON       : ':';
SEMICOLON   : ';';
LPAREN      : '(';
RPAREN      : ')';
LBRACK      : '[';
RBRACK      : ']';
DOTDOT      : '..';

IDENTIFIER  : [A-Za-z_][A-Za-z0-9_]*;

// Числовые литералы должны обрабатываться до операторов
INT_LITERAL
    : DECIMAL_LITERAL
    | BINARY_LITERAL
    | OCTAL_LITERAL
    | HEX_LITERAL
    ;
fragment DECIMAL_LITERAL
    : [0-9]+ ('_' [0-9]+)*
    ;

fragment BINARY_LITERAL
    : '2#' [01]+ ('_' [01]+)*
    ;

fragment OCTAL_LITERAL
    : '8#' [0-7]+ ('_' [0-7]+)*
    ;

fragment HEX_LITERAL
    : '16#' [0-9A-Fa-f]+ ('_' [0-9A-Fa-f]+)*
    ;

REAL_LITERAL
    : (REAL_TYPE_NAME '#')? 
      ( 
        // Полная форма: 123.456e+78
        DECIMAL_LITERAL '.' DECIMAL_LITERAL EXPONENT?
        
        // Сокращенные формы:
        | DECIMAL_LITERAL '.' EXPONENT?  // 123.
        | '.' DECIMAL_LITERAL EXPONENT?  // .456
        | DECIMAL_LITERAL EXPONENT       // 123e+78
      )
    ;

fragment EXPONENT
    : [Ee] [+-]? DECIMAL_LITERAL
    ;


BOOL_LITERAL    : 'TRUE' | 'FALSE' | '0' | '1';
TIME_LITERAL    : DURATION | TIME_OF_DAY | DATE | DATE_AND_TIME;
STRING_LITERAL  : '\'' (~['\r\n] | '\'\'')* '\'' | '"' (~["\r\n] | '""')* '"';
DIRECT_VARIABLE : '%' [IQM] [XBWD]? [0-9]+ ('.' [0-9]+)*;

COMMENT     : '(*' .*? '*)' -> skip;
COMMENT2    : '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
WS          : [ \t\r\n]+ -> skip;

fragment DURATION : (TIME_TYPE_NAME '#')? ('+' | '-')? INTERVAL;
fragment INTERVAL : DAYS | HOURS | MINUTES | SECONDS;
fragment DAYS     : [0-9]+ 'd' ('_'? HOURS)?;
fragment HOURS    : [0-9]+ 'h' ('_'? MINUTES)?;
fragment MINUTES  : [0-9]+ 'm' ('_'? SECONDS)?;
fragment SECONDS  : [0-9]+ ('_'? [0-9]*)? 's';
fragment TIME_OF_DAY : TOD_TYPE_NAME '#' DAY_HOUR ':' DAY_MINUTE ':' DAY_SECOND;
fragment DATE     : DATE_TYPE_NAME '#' YEAR '-' MONTH '-' DAY;
fragment DATE_AND_TIME : DT_TYPE_NAME '#' YEAR '-' MONTH '-' DAY '-' DAY_HOUR ':' DAY_MINUTE ':' DAY_SECOND;
fragment DAY_HOUR : [0-9]+;
fragment DAY_MINUTE : [0-9]+;
fragment DAY_SECOND : [0-9]+ ('.' [0-9]+)?;
fragment YEAR      : [0-9]+;
fragment MONTH     : [0-9]+;
fragment DAY       : [0-9]+;

// Парсер
pous
    : (program | function_decl | function_block_decl | type_declaration | global_var_declaration)* EOF 
    ;

program
    : PROGRAM id=IDENTIFIER 
      (vars+=var_declaration_block)* 
      stmts=statement_list
      END_PROGRAM
    ;

function_decl
    : FUNCTION IDENTIFIER 
      ((':' data_type)? (VAR_INPUT input_decl* END_VAR)?)
      (VAR_OUTPUT output_decl* END_VAR)?
      (VAR var_decl* END_VAR)?
      statement_list
      END_FUNCTION
    ;

function_block_decl
    : FUNCTION_BLOCK IDENTIFIER
      (VAR_INPUT input_decl* END_VAR)?
      (VAR_OUTPUT output_decl* END_VAR)?
      (VAR var_decl* END_VAR)?
      statement_list
      END_FUNCTION_BLOCK
    ;

type_declaration
    : TYPE type_definition+ END_TYPE
    ;

type_definition
    : IDENTIFIER ':' data_type SEMICOLON
    ;

global_var_declaration
    : VAR_GLOBAL var_decl* END_VAR
    ;

var_declaration_block
    : (VAR | VAR_INPUT | VAR_OUTPUT | VAR_IN_OUT | VAR_EXTERNAL | VAR_GLOBAL | VAR_TEMP)
      decls+=var_decl+
      END_VAR
    ;

var_decl
    : id=IDENTIFIER (COMMA IDENTIFIER)* COLON type=data_type (ASSIGN expr=expression)? SEMICOLON
    ;

data_type
    : elementary_type_name 
    | array_type 
    | structured_type
    | IDENTIFIER
    ;

elementary_type_name
    : INT_TYPE_NAME
    | REAL_TYPE_NAME
    | BOOL_TYPE_NAME
    | TIME_TYPE_NAME
    | DATE_TYPE_NAME
    | STRING_TYPE_NAME
    | BIT_TYPE_NAME
    ;

array_type
    : ARRAY '[' subrange (COMMA subrange)* ']' OF data_type
    ;

subrange
    : expression DOTDOT expression
    ;

structured_type
    : STRUCT var_decl+ END_STRUCT
    ;

statement_list
    : (statement)+
    ;

statement
    : assignment_statement SEMICOLON
    | function_invocation SEMICOLON
    | continue_statement SEMICOLON?
    | exit_statement SEMICOLON?
    | return_statement SEMICOLON?
    | if_statement SEMICOLON?
    | case_statement SEMICOLON?
    | for_statement SEMICOLON?
    | while_statement SEMICOLON?
    | repeat_statement SEMICOLON?
    ;

assignment_statement
    : variable ASSIGN expression
    ;

if_statement
    : IF conds+=expression THEN thens+=statement_list
      (ELSIF conds+=expression THEN thens+=statement_list)*
      (ELSE else=statement_list)?
      END_IF
    ;

case_statement
    : CASE expression OF
      case_element+
      (ELSE statement_list)?
      END_CASE
    ;

case_element
    : case_label (COMMA case_label)* COLON statement_list
    ;

case_label
    : expression
    | subrange
    ;

for_statement
    : FOR IDENTIFIER ASSIGN expression TO expression (BY expression)? 
      DO statement_list 
      END_FOR
    ;

while_statement
    : WHILE expression DO statement_list END_WHILE
    ;

repeat_statement
    : REPEAT statement_list UNTIL expression END_REPEAT
    ;

function_invocation
    : IDENTIFIER LPAREN (args+=expression (COMMA args+=expression)*)? RPAREN
    ;

return_statement
    : RETURN expression?
    ;

continue_statement
    : CONTINUE
    ;

exit_statement
    : EXIT
    ;


// Иерархия выражений с правильными приоритетами
expression
    : LPAREN expression RPAREN                                       # parenExpression
    | operator=(PLUS|MINUS|NOT) expression                           # unaryExpression
    | left=expression operator=POWER right=expression                # binaryExpression
    | left=expression operator=(MULT|DIV|MOD) right=expression       # binaryExpression
    | left=expression operator=(PLUS|MINUS) right=expression         # binaryExpression
    | left=expression operator=(LESS|LESS_EQ|GREATER|GREATER_EQ) 
      right=expression                                               # binaryExpression
    | left=expression operator=(EQUAL|NOT_EQUAL) right=expression    # binaryExpression
    | left=expression operator=AND right=expression                  # binaryExpression
    | left=expression operator=XOR right=expression                  # binaryExpression
    | left=expression operator=OR right=expression                   # binaryExpression
    | literal                                                        # literalExpression
    | variable                                                       # varExpression
    | function_invocation                                            # funcCallExpression
    ;

literal
    : INT_LITERAL
    | REAL_LITERAL
    | BOOL_LITERAL
    | TIME_LITERAL
    | STRING_LITERAL
    | typed_literal
    ;

typed_literal
    : type_name '#' (INT_LITERAL | REAL_LITERAL | BIT_STRING_LITERAL)
    ;

type_name
    : INT_TYPE_NAME
    | REAL_TYPE_NAME
    | BIT_TYPE_NAME
    | BOOL_TYPE_NAME
    ;

BIT_STRING_LITERAL
    : BIT_TYPE_NAME '#' (BINARY_LITERAL | OCTAL_LITERAL | HEX_LITERAL | DECIMAL_LITERAL)
    ;

variable
    : name=IDENTIFIER (DOT IDENTIFIER | LBRACK expression RBRACK)*
    | DIRECT_VARIABLE
    ;

// Дополнение к существующей грамматике

input_decl
    : var_decl
    | edge_decl
    | array_conform_decl
    ;

output_decl
    : var_decl
    | array_conform_decl
    ;

edge_decl
    : variable_list COLON BOOL_TYPE_NAME (R_EDGE | F_EDGE)
    ;

array_conform_decl
    : variable_list COLON array_conformand
    ;

array_conformand
    : ARRAY LBRACK MULT (COMMA MULT)* RBRACK OF data_type_access
    ;

variable_list
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

data_type_access
    : elementary_type_name
    | derived_type_access
    ;

derived_type_access
    : single_elem_type_access
    | array_type_access
    | struct_type_access
    | string_type_access
    ;

single_elem_type_access
    : simple_type_access
    | subrange_type_access
    | enum_type_access
    ;

simple_type_access
    : (namespace_name DOT)* simple_type_name
    ;

subrange_type_access
    : (namespace_name DOT)* subrange_type_name
    ;

enum_type_access
    : (namespace_name DOT)* enum_type_name
    ;

array_type_access
    : (namespace_name DOT)* array_type_name
    ;

struct_type_access
    : (namespace_name DOT)* struct_type_name
    ;

string_type_access
    : (namespace_name DOT)* STRING_TYPE_NAME
    ;

namespace_name
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

simple_type_name
    : IDENTIFIER
    ;

subrange_type_name
    : IDENTIFIER
    ;

enum_type_name
    : IDENTIFIER
    ;

array_type_name
    : IDENTIFIER
    ;

struct_type_name
    : IDENTIFIER
    ;

R_EDGE: 'R_EDGE';
F_EDGE: 'F_EDGE';
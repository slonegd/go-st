grammar ST;

prorgamm :
    'PROGRAM' identifier=ID 
        var_declaration_blocks
        statement_list 
    'END_PROGRAM'
;

var_declaration_blocks : (var_declaration_block)*;

var_declaration_block :
	'VAR'
	var_declaration
	'END_VAR'
;

var_declaration : (identifier=ID ':' type=type_name ';')*; // TODO значения по умолчанию

type_name: 'INT'; // TODO другие типы

statement_list : (statement)*;

statement : assignment_statement ';'; // TODO другие виды выражений

assignment_statement : left=ID ':=' right=expression; // TODO другие виды присваиваний?

expression : constant; // TODO не только константы, но и выражения

constant : number; // TODO не только число

number : Integer; // TODO не только целое

Integer : ('0'..'9')('0'..'9')*; // TODO другие записи целого числа: 16#

ID : ([a-zA-Z_]) ([$a-zA-Z0-9_])*;

WHITESPACE: [ \r\n\t]+ -> skip;
grammar ST;

program :
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

var_declaration : (identifier+=ID ':' type+=type_name ';')*; // TODO значения по умолчанию

type_name: 'INT'; // TODO другие типы

statement_list : (statement)*;

statement : assignment_statement ';'
          | if_statement ';'?
; // TODO другие виды выражений

assignment_statement : left=ID ':=' right=expression; // TODO другие виды присваиваний?

if_statement
:
	'IF' cond+=condition 'THEN' thenlist+=then_list
	('ELSEIF' cond+=condition 'THEN' thenlist+=then_list)* 
	('ELSE' elselist = else_list)?
	'END_IF' ';'?
;

// семантически нужно отличать (для ветвления)
condition : expression; // TODO bool_expression
then_list : statement_list;
else_list: statement_list;

expression : number                                                             #constant
           | ID                                                                 #variable
           | '(' sub=expression ')'                                             #parenExpr
           | left=expression op =('*'|'/'|'MOD') right=expression               #binaryPowerExpr
           | left=expression op =('+'|'-') right=expression                     #binaryPlusExpr
           | left=expression op =('>'|'>='|'<'|'<='|'='|'<>') right=expression  #binaryCompareExpr
; // TODO остальные типы выражений

number : Integer; // TODO не только целое

Integer : ('0'..'9')('0'..'9')*; // TODO другие записи целого числа: 16#

ID : ([a-zA-Z_]) ([$a-zA-Z0-9_])*;

WHITESPACE: [ \r\n\t]+ -> skip;
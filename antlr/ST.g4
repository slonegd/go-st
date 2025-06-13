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
	var_declaration*
	'END_VAR'
;

var_declaration : identifier=ID ':' type=type_name ( ':=' default=number )? ';'; 

type_name: 'SINT'
         | 'INT'
         | 'DINT'
         | 'LINT'
         | 'USINT'
         | 'UINT'
         | 'UDINT'
         | 'ULINT'
         | 'REAL'
         | 'LREAL'
         | 'BOOL'
         | 'STRING'
; // TODO другие типы

statement_list : (statement)*;

statement : assignment_statement ';'
          | if_statement ';'?
; // TODO другие виды выражений

assignment_statement : left=ID ':=' right=expression; // TODO другие виды присваиваний?

if_statement
:
	'IF' cond+=condition 'THEN' thenlist+=then_list
	('ELSIF' cond+=condition 'THEN' thenlist+=then_list)* 
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

number : signed_integer 
       | integer
; // TODO не только целое

signed_integer: PLUS integer
              | MINUS integer
;

integer: Integer;

Integer : ('0'..'9')('0'..'9')*; // TODO другие записи целого числа: 16#

ID : ([a-zA-Z_]) ([$a-zA-Z0-9_])*;

PLUS : '+';
MINUS: '-';

Comment : (
               '//' ~( '\n' | '\r' )* '\r' ? '\n' 
             | '(*' .*? '*)' 
             | '/*' .*? '*/' 
          ) -> channel(HIDDEN);
WHITESPACE: [ \r\n\t]+ -> skip;
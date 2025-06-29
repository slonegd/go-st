grammar ST;

options {
    // для пробрасывания по дереву доп объектов (логгер, исходный код и тп)
    contextSuperClass = CustomContext;
}

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
          | while_statement ';'?
          | continue_statement ';'?
          | exit_statement ';'?
; // TODO другие виды выражений

assignment_statement : left=ID ':=' right=expression; // TODO другие виды присваиваний?

if_statement :
	'IF' cond+=condition 'THEN' thenlist+=then_list
	('ELSIF' cond+=condition 'THEN' thenlist+=then_list)* 
	('ELSE' elselist = else_list)?
	'END_IF' ';'?
;

while_statement : 'WHILE' cond=expression 'DO' body=statement_list 'END_WHILE';

continue_statement : 'CONTINUE';
exit_statement     : 'EXIT';

// семантически нужно отличать (для ветвления)
condition : expression;
then_list : statement_list;
else_list: statement_list;

expression : number                                                             #constant
           | id=ID '(' sub=expression ')'                                       #callExpr
           | ID                                                                 #variable
           | '(' sub=expression ')'                                             #parenExpr
           | left=expression op =('*'|'/'|'MOD') right=expression               #binaryPowerExpr
           | left=expression op =('+'|'-') right=expression                     #binaryPlusExpr
           | left=expression op =('>'|'>='|'<'|'<='|'='|'<>') right=expression  #binaryCompareExpr
; // TODO остальные типы выражений

number : FLOAT 
       | integer  
;

integer: (signed_integer|unsign_integer);

signed_integer: sign=(PLUS | MINUS) unsign_integer;

unsign_integer: Integer | Hex_Int;

Integer : Digit+;
Hex_Int : '16#' ( '_' ? Hex_Digit )+;
FLOAT   :   MINUS? ( Digit+  '.' Digit*  // -0.5 или 3.14
                     | '.' Digit+)       // .5 
                   ('E' (PLUS|MINUS)? Digit+)?     // экспонента (1E2 или 1e-3)
;

ID : ([a-zA-Z_]) ([$a-zA-Z0-9_])*;

Digit : '0'..'9';
Hex_Digit : '0'..'9' | 'A'..'F';

PLUS : '+';
MINUS: '-';

Comment : (
               '//' ~( '\n' | '\r' )* '\r' ? '\n' 
             | '(*' .*? '*)' 
             | '/*' .*? '*/' 
          ) -> channel(HIDDEN);
WHITESPACE: [ \r\n\t]+ -> skip;
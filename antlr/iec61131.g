// копипаст из стандарта, для сравнения, не генерации
// cудя по всему в стандарте для ANTLRv3
grammar iec61131;

// Table 1 - Character sets
// Table 2 - Identifiers
Letter      : 'A'..'Z' | '_';
Digit       : '0'..'9';
Bit         : '0'..'1';
Octal_Digit : '0'..'7';
Hex_Digit   : '0'..'9' | 'A'..'F';
Identifier  : Letter ( Letter | Digit )*;

// Table 3 - Comments
Comment : (
               '//' ~( '\n' | '\r' )* '\r' ? '\n' 
             | '(*' .*? '*)' 
             | '/*' .*? '*/' 
          ) -> channel(HIDDEN);
WS : ( ' ' | '\t' | '\r' | '\n' ) -> channel(HIDDEN); // white space
EOL : '\n'; 

// Table 4 - Pragma
Pragma : '{' *? '}' -> channel(HIDDEN);

// Table 5 - Numeric literal
Constant        : Numeric_Literal | Char_Literal | Time_Literal | Bit_Str_Literal | Bool_Literal;
Numeric_Literal : Int_Literal | Real_Literal;
Int_Literal     : ( Int_Type_Name '#' )? ( Signed_Int | Binary_Int | Octal_Int | Hex_Int );
Unsigned_Int    : Digit ( '_' ? Digit )*;
Signed_Int      : ( '+' | '-' )? Unsigned_Int;
Binary_Int      : '2#' ( '_' ? Bit )+;
Octal_Int       : '8#' ( '_' ? Octal_Digit )+;
Hex_Int         : '16#' ( '_' ? Hex_Digit )+;
Real_Literal    : ( Real_Type_Name '#' )? Signed_Int '.' Unsigned_Int ( 'E' Signed_Int )?;
Bit_Str_Literal : ( Multibits_Type_Name '#' )? ( Unsigned_Int | Binary_Int | Octal_Int | Hex_Int );
Bool_Literal    : ( Bool_Type_Name '#' )? ( '0' | '1' | 'FALSE' | 'TRUE' ); 

// Table 6 - Character String literals
// Table 7 - Two-character combinations in character strings
Char_Literal      : ( 'STRING#' )? Char_Str;
Char_Str          : S_Byte_Char_Str | D_Byte_Char_Str;
S_Byte_Char_Str   : '\'' S_Byte_Char_Value + '\'';
D_Byte_Char_Str   : '"' D_Byte_Char_Value + '"';
S_Byte_Char_Value : Common_Char_Value | '$\'' | '"' | '$' Hex_Digit Hex_Digit;
D_Byte_Char_Value : Common_Char_Value | '\'' | '$"' | '$' Hex_Digit Hex_Digit Hex_Digit Hex_Digit;
Common_Char_Value : ' ' | '!' | '#' | '%' | '&' | '('..'/' | '0'..'9' | ':'..'@' | 'A'..'Z' | '['..'`' | 'a'..'z' 
                  | '{'..'~' | '$$' | '$L' | '$N' | '$P' | '$R' | '$T'; // any printable characters except $, " and '

// Table 8 - Duration literals
// Table 9 – Date and time of day literals
Time_Literal  : Duration | Time_Of_Day | Date | Date_And_Time;
Duration      : ( Time_Type_Name | 'T' | 'LT' ) '#' ( '+' | '-' )? Interval;
Fix_Point     : Unsigned_Int ( '.' Unsigned_Int )?;
Interval      : Days | Hours | Minutes | Seconds | Milliseconds | Microseconds | Nanoseconds;
Days          : ( Fix_Point 'd' ) | ( Unsigned_Int 'd' '_' ? )? Hours ?;
Hours         : ( Fix_Point 'h' ) | ( Unsigned_Int 'h' '_' ? )? Minutes ?;
Minutes       : ( Fix_Point 'm' ) | ( Unsigned_Int 'm' '_' ? )? Seconds ?;
Seconds       : ( Fix_Point 's' ) | ( Unsigned_Int 's' '_' ? )? Milliseconds ?;
Milliseconds  : ( Fix_Point 'ms' ) | ( Unsigned_Int 'ms' '_' ? )? Microseconds ?;
Microseconds  : ( Fix_Point 'us' ) | ( Unsigned_Int 'us' '_' ? )? Nanoseconds ?;
Nanoseconds   : Fix_Point 'ns';
Time_Of_Day   : ( Tod_Type_Name | 'LTIME_OF_DAY' ) '#' Daytime;
Daytime       : Day_Hour ':' Day_Minute ':' Day_Second;
Day_Hour      : Unsigned_Int;
Day_Minute    : Unsigned_Int;
Day_Second    : Fix_Point;
Date          : ( Date_Type_Name | 'D' | 'LD' ) '#' Date_Literal;
Date_Literal  : Year '-' Month '-' Day;
Year          : Unsigned_Int;
Month         : Unsigned_Int;
Day           : Unsigned_Int;
Date_And_Time : ( DT_Type_Name | 'LDATE_AND_TIME' ) '#' Date_Literal '-' Daytime; 

// Table 10 - Elementary data types
Data_Type_Access     : Elem_Type_Name | Derived_Type_Access;
Elem_Type_Name       : Numeric_Type_Name | Bit_Str_Type_Name
                     | String_Type_Name | Date_Type_Name | Time_Type_Name;
Numeric_Type_Name    : Int_Type_Name | Real_Type_Name;
Int_Type_Name        : Sign_Int_Type_Name | Unsign_Int_Type_Name;
Sign_Int_Type_Name   : 'SINT' | 'INT' | 'DINT' | 'LINT';
Unsign_Int_Type_Name : 'USINT' | 'UINT' | 'UDINT' | 'ULINT'; 
Real_Type_Name       : 'REAL' | 'LREAL';
String_Type_Name     : 'STRING' ( '[' Unsigned_Int ']' )? | 'WSTRING' ( '[' Unsigned_Int ']' )? | 'CHAR' | 'WCHAR';
Time_Type_Name       : 'TIME' | 'LTIME';
Date_Type_Name       : 'DATE' | 'LDATE';
Tod_Type_Name        : 'TIME_OF_DAY' | 'TOD' | 'LTOD';
DT_Type_Name         : 'DATE_AND_TIME' | 'DT' | 'LDT';
Bit_Str_Type_Name    : Bool_Type_Name | Multibits_Type_Name;
Bool_Type_Name       : 'BOOL';
Multibits_Type_Name  : 'BYTE' | 'WORD' | 'DWORD' | 'LWORD'; 

// Table 11 - Declaration of user-defined data types and initialization
Derived_Type_Access     : Single_Elem_Type_Access | Array_Type_Access | Struct_Type_Access
                        | String_Type_Access | Class_Type_Access | Ref_Type_Access | Interface_Type_Access;
String_Type_Access      : ( Namespace_Name '.' )* String_Type_Name;
Single_Elem_Type_Access : Simple_Type_Access | Subrange_Type_Access | Enum_Type_Access;
Simple_Type_Access      : ( Namespace_Name '.' )* Simple_Type_Name;
Subrange_Type_Access    : ( Namespace_Name '.' )* Subrange_Type_Name;
Enum_Type_Access        : ( Namespace_Name '.' )* Enum_Type_Name;
Array_Type_Access       : ( Namespace_Name '.' )* Array_Type_Name;
Struct_Type_Access      : ( Namespace_Name '.' )* Struct_Type_Name;
Simple_Type_Name        : Identifier;
Subrange_Type_Name      : Identifier;
Enum_Type_Name          : Identifier;
Array_Type_Name         : Identifier;
Struct_Type_Name        : Identifier;
Data_Type_Decl          : 'TYPE' ( Type_Decl ';' )+ 'END_TYPE';
Type_Decl               : Simple_Type_Decl | Subrange_Type_Decl | Enum_Type_Decl
                        | Array_Type_Decl | Struct_Type_Decl
                        | Str_Type_Decl | Ref_Type_Decl;
Simple_Type_Decl        : Simple_Type_Name ':' Simple_Spec_Init;
Simple_Spec_Init        : Simple_Spec ( ':=' Constant_Expr )?;
Simple_Spec             : Elem_Type_Name | Simple_Type_Access;
Subrange_Type_Decl      : Subrange_Type_Name ':' Subrange_Spec_Init;
Subrange_Spec_Init      : Subrange_Spec ( ':=' Signed_Int )?;
Subrange_Spec           : Int_Type_Name '(' Subrange ')' | Subrange_Type_Access;
Subrange                : Constant_Expr '..' Constant_Expr;
Enum_Type_Decl          : Enum_Type_Name ':' ( ( Elem_Type_Name ? Named_Spec_Init ) | Enum_Spec_Init );
Named_Spec_Init         : '(' Enum_Value_Spec ( ',' Enum_Value_Spec )* ')' ( ':=' Enum_Value )?;
Enum_Spec_Init          : ( ( '(' Identifier ( ',' Identifier )* ')' ) | Enum_Type_Access ) ( ':=' Enum_Value )?;
Enum_Value_Spec         : Identifier ( ':=' ( Int_Literal | Constant_Expr ) )?;
Enum_Value              : ( Enum_Type_Name '#' )? Identifier;
Array_Type_Decl         : Array_Type_Name ':' Array_Spec_Init;
Array_Spec_Init         : Array_Spec ( ':=' Array_Init )?;
Array_Spec              : Array_Type_Access | 'ARRAY' '[' Subrange ( ',' Subrange )* ']' 'OF' Data_Type_Access;
Array_Init              : '[' Array_Elem_Init ( ',' Array_Elem_Init )* ']';
Array_Elem_Init         : Array_Elem_Init_Value | Unsigned_Int '(' Array_Elem_Init_Value ? ')';
Array_Elem_Init_Value   : Constant_Expr | Enum_Value | Struct_Init | Array_Init;
Struct_Type_Decl        : Struct_Type_Name ':' Struct_Spec;
Struct_Spec             : Struct_Decl | Struct_Spec_Init;
Struct_Spec_Init        : Struct_Type_Access ( ':=' Struct_Init )?;
Struct_Decl             :'STRUCT' 'OVERLAP' ? ( Struct_Elem_Decl ';' )+ 'END_STRUCT';
Struct_Elem_Decl        : Struct_Elem_Name ( Located_At Multibit_Part_Access ? )? ':'
                          (Simple_Spec_Init | Subrange_Spec_Init | Enum_Spec_Init | Array_Spec_Init | Struct_Spec_Init);
Struct_Elem_Name        : Identifier;
Struct_Init             : '(' Struct_Elem_Init ( ',' Struct_Elem_Init )* ')';
Struct_Elem_Init        : Struct_Elem_Name ':=' ( Constant_Expr | Enum_Value | Array_Init | Struct_Init | Ref_Value );
Str_Type_Decl           : String_Type_Name ':' String_Type_Name ( ':=' Char_Str )?;

// Table 16 - Directly represented variables
Direct_Variable : '%' ( 'I' | 'Q' | 'M' ) ( 'X' | 'B' | 'W' | 'D' | 'L' )? Unsigned_Int ( '.' Unsigned_Int )*; 

// Table 12 - Reference operations
Ref_Type_Decl   : Ref_Type_Name ':' Ref_Spec_Init;
Ref_Spec_Init   : Ref_Spec ( ':=' Ref_Value )?;
Ref_Spec        : 'REF_TO' + Data_Type_Access;
Ref_Type_Name   : Identifier;
Ref_Type_Access : ( Namespace_Name '.' )* Ref_Type_Name;
Ref_Name        : Identifier;
Ref_Value       : Ref_Addr | 'NULL';
Ref_Addr        : 'REF' '(' ( Symbolic_Variable | FB_Instance_Name | Class_Instance_Name ) ')';
Ref_Assign      : Ref_Name ':=' ( Ref_Name | Ref_Deref | Ref_Value );
Ref_Deref       : Ref_Name '^' +; 

// Table 13 - Declaration of variables/Table 14 – Initialization of variables
Variable             : Direct_Variable | Symbolic_Variable; 
Symbolic_Variable    : ( ( 'THIS' '.' ) | ( Namespace_Name '.' )+ )? ( Var_Access | Multi_Elem_Var );
Var_Access           : Variable_Name | Ref_Deref;
Variable_Name        : Identifier;
Multi_Elem_Var       : Var_Access ( Subscript_List | Struct_Variable )+;
Subscript_List       : '[' Subscript ( ',' Subscript )* ']';
Subscript            : Expression;
Struct_Variable      : '.' Struct_Elem_Select;
Struct_Elem_Select   : Var_Access;
Input_Decls          : 'VAR_INPUT' ( 'RETAIN' | 'NON_RETAIN' )? ( Input_Decl ';' )* 'END_VAR';
Input_Decl           : Var_Decl_Init | Edge_Decl | Array_Conform_Decl;
Edge_Decl            : Variable_List ':' 'BOOL' ( 'R_EDGE' | 'F_EDGE' );
Var_Decl_Init        : Variable_List ':' ( Simple_Spec_Init | Str_Var_Decl | Ref_Spec_Init )
                     | Array_Var_Decl_Init | Struct_Var_Decl_Init | FB_Decl_Init | Interface_Spec_Init;
Ref_Var_Decl         : Variable_List ':' Ref_Spec;
Interface_Var_Decl   : Variable_List ':' Interface_Type_Access;
Variable_List        : Variable_Name ( ',' Variable_Name )*;
Array_Var_Decl_Init  : Variable_List ':' Array_Spec_Init;
Array_Conformand     : 'ARRAY' '[' '*' ( ',' '*' )* ']' 'OF' Data_Type_Access;
Array_Conform_Decl   : Variable_List ':' Array_Conformand;
Struct_Var_Decl_Init : Variable_List ':' Struct_Spec_Init;
FB_Decl_No_Init      : FB_Name ( ',' FB_Name )* ':' FB_Type_Access;
FB_Decl_Init         : FB_Decl_No_Init ( ':=' Struct_Init )?;
FB_Name              : Identifier;
FB_Instance_Name     : ( Namespace_Name '.' )* FB_Name '^' *;
Output_Decls         : 'VAR_OUTPUT' ( 'RETAIN' | 'NON_RETAIN' )? ( Output_Decl ';' )* 'END_VAR';
Output_Decl          : Var_Decl_Init | Array_Conform_Decl;
In_Out_Decls         : 'VAR_IN_OUT' ( In_Out_Var_Decl ';' )* 'END_VAR';
In_Out_Var_Decl      : Var_Decl | Array_Conform_Decl | FB_Decl_No_Init;
Var_Decl             : Variable_List ':' ( Simple_Spec | Str_Var_Decl | Array_Var_Decl | Struct_Var_Decl );
Array_Var_Decl       : Variable_List ':' Array_Spec;
Struct_Var_Decl      : Variable_List ':' Struct_Type_Access;
Var_Decls            : 'VAR' 'CONSTANT' ? Access_Spec ? ( Var_Decl_Init ';' )* 'END_VAR';
Retain_Var_Decls     : 'VAR' 'RETAIN' Access_Spec ? ( Var_Decl_Init ';' )* 'END_VAR';
Loc_Var_Decls        : 'VAR' ( 'CONSTANT' | 'RETAIN' | 'NON_RETAIN' )? ( Loc_Var_Decl ';' )* 'END_VAR';
Loc_Var_Decl         : Variable_Name ? Located_At ':' Loc_Var_Spec_Init;
Temp_Var_Decls       : 'VAR_TEMP' ( ( Var_Decl | Ref_Var_Decl | Interface_Var_Decl ) ';' )* 'END_VAR';
External_Var_Decls   : 'VAR_EXTERNAL' 'CONSTANT' ? ( External_Decl ';' )* 'END_VAR';
External_Decl        : Global_Var_Name ':' ( Simple_Spec | Array_Spec | Struct_Type_Access | FB_Type_Access 
                         | Ref_Type_Access 
                     );
Global_Var_Name      : Identifier;
Global_Var_Decls     : 'VAR_GLOBAL' ( 'CONSTANT' | 'RETAIN' )? ( Global_Var_Decl ';' )* 'END_VAR';
Global_Var_Decl      : Global_Var_Spec ':' ( Loc_Var_Spec_Init | FB_Type_Access );
Global_Var_Spec      : ( Global_Var_Name ( ',' Global_Var_Name )* ) | ( Global_Var_Name Located_At );
Loc_Var_Spec_Init    : Simple_Spec_Init | Array_Spec_Init | Struct_Spec_Init | S_Byte_Str_Spec | D_Byte_Str_Spec;
Located_At           : 'AT' Direct_Variable;
Str_Var_Decl         : S_Byte_Str_Var_Decl | D_Byte_Str_Var_Decl;
S_Byte_Str_Var_Decl  : Variable_List ':' S_Byte_Str_Spec;
S_Byte_Str_Spec      : 'STRING' ( '[' Unsigned_Int ']' )? ( ':=' S_Byte_Char_Str )?;
D_Byte_Str_Var_Decl  : Variable_List ':' D_Byte_Str_Spec;
D_Byte_Str_Spec      : 'WSTRING' ( '[' Unsigned_Int ']' )? ( ':=' D_Byte_Char_Str )?;
Loc_Partly_Var_Decl  : 'VAR' ( 'RETAIN' | 'NON_RETAIN' )? Loc_Partly_Var * 'END_VAR';
Loc_Partly_Var       : Variable_Name 'AT' '%' ( 'I' | 'Q' | 'M' ) '*' ':' Var_Spec ';';
Var_Spec             : Simple_Spec | Array_Spec | Struct_Type_Access
                     | ( 'STRING' | 'WSTRING' ) ( '[' Unsigned_Int ']' )?; 

// Table 19 - Function declaration
Func_Name         : Std_Func_Name | Derived_Func_Name;
Func_Access       : ( Namespace_Name '.' )* Func_Name;
Std_Func_Name     : 'TRUNC' | 'ABS' | 'SQRT' | 'LN' | 'LOG' | 'EXP'
                  | 'SIN' | 'COS' | 'TAN' | 'ASIN' | 'ACOS' | 'ATAN' | 'ATAN2 '
                  | 'ADD' | 'SUB' | 'MUL' | 'DIV' | 'MOD' | 'EXPT' | 'MOVE '
                  | 'SHL' | 'SHR' | 'ROL' | 'ROR'
                  | 'AND' | 'OR' | 'XOR' | 'NOT'
                  | 'SEL' | 'MAX' | 'MIN' | 'LIMIT' | 'MUX '
                  | 'GT' | 'GE' | 'EQ' | 'LE' | 'LT' | 'NE'
                  | 'LEN' | 'LEFT' | 'RIGHT' | 'MID' | 'CONCAT' | 'INSERT' | 'DELETE' | 'REPLACE' 
                  | 'FIND'; // incomplete list
Derived_Func_Name : Identifier;
Func_Decl         : 'FUNCTION' Derived_Func_Name ( ':' Data_Type_Access )? Using_Directive *
                    ( IO_Var_Decls | Func_Var_Decls | Temp_Var_Decls )* Func_Body 'END_FUNCTION';
IO_Var_Decls      : Input_Decls | Output_Decls | In_Out_Decls;
Func_Var_Decls    : External_Var_Decls | Var_Decls;
Func_Body         : Ladder_Diagram | FB_Diagram | Instruction_List | Stmt_List | Other_Languages; 

// Table 40 – Function block type declaration
// Table 41 - Function block instance declaration
FB_Type_Name        : Std_FB_Name | Derived_FB_Name;
FB_Type_Access      : ( Namespace_Name '.' )* FB_Type_Name;
Std_FB_Name         : 'SR' | 'RS' | 'R_TRIG' | 'F_TRIG' | 'CTU'| 'CTD' | 'CTUD' | 'TP' | 'TON' 
                    | 'TOF'; // incomplete list
Derived_FB_Name     : Identifier;
FB_Decl             : 'FUNCTION_BLOCK' ( 'FINAL' | 'ABSTRACT' )? Derived_FB_Name Using_Directive *
                      ( 'EXTENDS' ( FB_Type_Access | Class_Type_Access ) )?
                      ( 'IMPLEMENTS' Interface_Name_List )?
                      ( FB_IO_Var_Decls | Func_Var_Decls | Temp_Var_Decls | Other_Var_Decls )*
                      ( Method_Decl )* FB_Body ? 'END_FUNCTION_BLOCK';
FB_IO_Var_Decls     : FB_Input_Decls | FB_Output_Decls | In_Out_Decls;
FB_Input_Decls      : 'VAR_INPUT' ( 'RETAIN' | 'NON_RETAIN' )? ( FB_Input_Decl ';' )* 'END_VAR';
FB_Input_Decl       : Var_Decl_Init | Edge_Decl | Array_Conform_Decl;
FB_Output_Decls     : 'VAR_OUTPUT' ( 'RETAIN' | 'NON_RETAIN' )? ( FB_Output_Decl ';' )* 'END_VAR';
FB_Output_Decl      : Var_Decl_Init | Array_Conform_Decl;
Other_Var_Decls     : Retain_Var_Decls | No_Retain_Var_Decls | Loc_Partly_Var_Decl;
No_Retain_Var_Decls : 'VAR' 'NON_RETAIN' Access_Spec ? ( Var_Decl_Init ';' )* 'END_VAR';
FB_Body             : SFC | Ladder_Diagram | FB_Diagram | Instruction_List | Stmt_List | Other_Languages;
Method_Decl         : 'METHOD' Access_Spec ( 'FINAL' | 'ABSTRACT' )? 'OVERRIDE' ?
                       Method_Name ( ':' Data_Type_Access )?
                      ( IO_Var_Decls | Func_Var_Decls | Temp_Var_Decls )* Func_Body 'END_METHOD';
Method_Name         : Identifier; 

// Table 48 - Class
// Table 50 Textual call of methods – Formal and non-formal parameter list
Class_Decl            : 'CLASS' ( 'FINAL' | 'ABSTRACT' )? Class_Type_Name Using_Directive *
                        ( 'EXTENDS' Class_Type_Access )? ( 'IMPLEMENTS' Interface_Name_List )?
                        ( Func_Var_Decls | Other_Var_Decls )* ( Method_Decl )* 'END_CLASS';
Class_Type_Name       : Identifier;
Class_Type_Access     : ( Namespace_Name '.' )* Class_Type_Name;
Class_Name            : Identifier;
Class_Instance_Name   : ( Namespace_Name '.' )* Class_Name '^' *;
Interface_Decl        : 'INTERFACE' Interface_Type_Name Using_Directive *
                        ( 'EXTENDS' Interface_Name_List )? Method_Prototype * 'END_INTERFACE';
Method_Prototype      : 'METHOD' Method_Name ( ':' Data_Type_Access )? IO_Var_Decls * 'END_METHOD';
Interface_Spec_Init   : Variable_List ( ':=' Interface_Value )?;
Interface_Value       : Symbolic_Variable | FB_Instance_Name | Class_Instance_Name | 'NULL';
Interface_Name_List   : Interface_Type_Access ( ',' Interface_Type_Access )*;
Interface_Type_Name   : Identifier;
Interface_Type_Access : ( Namespace_Name '.' )* Interface_Type_Name;
Interface_Name        : Identifier;
Access_Spec           : 'PUBLIC' | 'PROTECTED' | 'PRIVATE' | 'INTERNAL'; 

// Table 47 - Program declaration
Prog_Decl         : 'PROGRAM' Prog_Type_Name
                       ( IO_Var_Decls | Func_Var_Decls | Temp_Var_Decls | Other_Var_Decls
                   |  Loc_Var_Decls | Prog_Access_Decls )* FB_Body 'END_PROGRAM';
Prog_Type_Name    : Identifier;
Prog_Type_Access  : ( Namespace_Name '.' )* Prog_Type_Name;
Prog_Access_Decls : 'VAR_ACCESS' ( Prog_Access_Decl ';' )* 'END_VAR';
Prog_Access_Decl  : Access_Name ':' Symbolic_Variable Multibit_Part_Access ?
                    ':' Data_Type_Access Access_Direction ?; 

// Table 54 - 61 - Sequential Function Chart (SFC)
SFC                : Sfc_Network +;
Sfc_Network        : Initial_Step ( Step | Transition | Action )*;
Initial_Step       : 'INITIAL_STEP' Step_Name ':' ( Action_Association ';' )* 'END_STEP';
Step               : 'STEP' Step_Name ':' ( Action_Association ';' )* 'END_STEP';
Step_Name          : Identifier;
Action_Association : Action_Name '(' Action_Qualifier ? ( ',' Indicator_Name )* ')';
Action_Name        : Identifier;
Action_Qualifier   : 'N' | 'R' | 'S' | 'P' | ( ( 'L' | 'D' | 'SD' | 'DS' | 'SL' ) ',' Action_Time );
Action_Time        : Duration | Variable_Name;
Indicator_Name     : Variable_Name;
Transition         : 'TRANSITION' Transition_Name ? ( '(' 'PRIORITY' ':=' Unsigned_Int ')' )?
                     'FROM' Steps 'TO' Steps ':' Transition_Cond 'END_TRANSITION';
Transition_Name    : Identifier;
Steps              : Step_Name | '(' Step_Name ( ',' Step_Name )+ ')';
Transition_Cond    : ':=' Expression ';' | ':' ( FBD_Network | LD_Rung ) | ':=' IL_Simple_Inst;
Action             : 'ACTION' Action_Name ':' FB_Body 'END_ACTION'; 

// Table 62 - Configuration and resource declaration
Config_Name          : Identifier;
Resource_Type_Name   : Identifier;
Config_Decl          : 'CONFIGURATION' Config_Name Global_Var_Decls ?
                       ( Single_Resource_Decl | Resource_Decl + ) Access_Decls ? Config_Init ?
                       'END_CONFIGURATION';
Resource_Decl        : 'RESOURCE' Resource_Name 'ON' Resource_Type_Name
                       Global_Var_Decls ? Single_Resource_Decl
                       'END_RESOURCE';
Single_Resource_Decl : ( Task_Config ';' )* ( Prog_Config ';' )+;
Resource_Name        : Identifier;
Access_Decls         : 'VAR_ACCESS' ( Access_Decl ';' )* 'END_VAR';
Access_Decl          : Access_Name ':' Access_Path ':' Data_Type_Access Access_Direction ?;
Access_Path          : ( Resource_Name '.' )? Direct_Variable
                     | ( Resource_Name '.' )? ( Prog_Name '.' )?
                       ( ( FB_Instance_Name | Class_Instance_Name ) '.' )* Symbolic_Variable;
Global_Var_Access    : ( Resource_Name '.' )? Global_Var_Name ( '.' Struct_Elem_Name )?;
Access_Name          : Identifier;
Prog_Output_Access   : Prog_Name '.' Symbolic_Variable;
Prog_Name            : Identifier;
Access_Direction     : 'READ_WRITE' | 'READ_ONLY';
Task_Config          : 'TASK' Task_Name Task_Init;
Task_Name            : Identifier;
Task_Init            : '(' ( 'SINGLE' ':=' Data_Source ',' )?
                       ( 'INTERVAL' ':=' Data_Source ',' )?
                       'PRIORITY' ':=' Unsigned_Int ')';
Data_Source          : Constant | Global_Var_Access | Prog_Output_Access | Direct_Variable;
Prog_Config          : 'PROGRAM' ( 'RETAIN' | 'NON_RETAIN' )? Prog_Name ( 'WITH' Task_Name )? ':'
                        Prog_Type_Access ( '(' Prog_Conf_Elems ')' )?;
Prog_Conf_Elems      : Prog_Conf_Elem ( ',' Prog_Conf_Elem )*;
Prog_Conf_Elem       : FB_Task | Prog_Cnxn;
FB_Task              : FB_Instance_Name 'WITH' Task_Name;
Prog_Cnxn            : Symbolic_Variable ':=' Prog_Data_Source | Symbolic_Variable '=>' Data_Sink;
Prog_Data_Source     : Constant | Enum_Value | Global_Var_Access | Direct_Variable;
Data_Sink            : Global_Var_Access | Direct_Variable;
Config_Init          : 'VAR_CONFIG' ( Config_Inst_Init ';' )* 'END_VAR';
Config_Inst_Init     : Resource_Name '.' Prog_Name '.' ( ( FB_Instance_Name | Class_Instance_Name ) '.' )*
                       ( Variable_Name Located_At ? ':' Loc_Var_Spec_Init
                         | ( ( FB_Instance_Name ':' FB_Type_Access )
                         | ( Class_Instance_Name ':' Class_Type_Access ) ) ':=' Struct_Init 
                       ); 

// Table 64 - Namespace
Namespace_Decl     : 'NAMESPACE' 'INTERNAL' ? Namespace_H_Name Using_Directive * Namespace_Elements 'END_NAMESPACE';
Namespace_Elements : ( Data_Type_Decl | Func_Decl | FB_Decl
                       | Class_Decl | Interface_Decl | Namespace_Decl 
                     )+;
Namespace_H_Name   : Namespace_Name ( '.' Namespace_Name )*;
Namespace_Name     : Identifier;
Using_Directive    : 'USING' Namespace_H_Name ( ',' Namespace_H_Name )* ';';
POU_Decl           : Using_Directive *
                     ( Global_Var_Decls | Data_Type_Decl | Access_Decls
                       | Func_Decl | FB_Decl | Class_Decl | Interface_Decl | Namespace_Decl 
                     )+; 

// Table 67 - 70 - Instruction List (IL)
Instruction_List       : IL_Instruction +;
IL_Instruction         : ( IL_Label ':' )? ( IL_Simple_Operation | IL_Expr | IL_Jump_Operation
                       | IL_Invocation | IL_Formal_Func_Call
                       | IL_Return_Operator )? EOL +;
IL_Simple_Inst         : IL_Simple_Operation | IL_Expr | IL_Formal_Func_Call;
IL_Label               : Identifier;
IL_Simple_Operation    : IL_Simple_Operator IL_Operand ? | Func_Access IL_Operand_List ?;
IL_Expr                : IL_Expr_Operator '(' IL_Operand ? EOL + IL_Simple_Inst_List ? ')';
IL_Jump_Operation      : IL_Jump_Operator IL_Label;
IL_Invocation          : IL_Call_Operator ((( FB_Instance_Name | Func_Name | Method_Name | 'THIS '
                       | ( ( 'THIS' '.' ( ( FB_Instance_Name | Class_Instance_Name ) '.' )* ) Method_Name ) )
                         ( '(' ( ( EOL + IL_Param_List ? ) | IL_Operand_List ? ) ')' )? ) | 'SUPER' '(' ')' );
IL_Formal_Func_Call    : Func_Access '(' EOL + IL_Param_List ? ')';
IL_Operand             : Constant | Enum_Value | Variable_Access;
IL_Operand_List        : IL_Operand ( ',' IL_Operand )*;
IL_Simple_Inst_List    : IL_Simple_Instruction +;
IL_Simple_Instruction  : ( IL_Simple_Operation | IL_Expr | IL_Formal_Func_Call ) EOL +;
IL_Param_List          : IL_Param_Inst * IL_Param_Last_Inst;
IL_Param_Inst          : ( IL_Param_Assign | IL_Param_Out_Assign ) ',' EOL +;
IL_Param_Last_Inst     : ( IL_Param_Assign | IL_Param_Out_Assign ) EOL +;
IL_Param_Assign        : IL_Assignment ( IL_Operand | ( '(' EOL + IL_Simple_Inst_List ')' ) );
IL_Param_Out_Assign    : IL_Assign_Out_Operator Variable_Access; 
IL_Simple_Operator     : 'LD' | 'LDN' | 'ST' | 'STN' | 'ST?' | 'NOT' | 'S' | 'R'
                       | 'S1' | 'R1' | 'CLK' | 'CU' | 'CD' | 'PV'
                       | 'IN' | 'PT' | IL_Expr_Operator;
IL_Expr_Operator       : 'AND' | '&' | 'OR' | 'XOR' | 'ANDN' | '&N' | 'ORN'
                       | 'XORN' | 'ADD' | 'SUB' | 'MUL' | 'DIV'
                       | 'MOD' | 'GT' | 'GE' | 'EQ' | 'LT' | 'LE' | 'NE';
IL_Assignment          : Variable_Name ':=';
IL_Assign_Out_Operator : 'NOT' ? Variable_Name '=>';
IL_Call_Operator       : 'CAL' | 'CALC' | 'CALCN';
IL_Return_Operator     : 'RT' | 'RETC' | 'RETCN';
IL_Jump_Operator       : 'JMP' | 'JMPC' | 'JMPCN'; 

// Table 71 - 72 - Language Structured Text (ST)
Expression           : Xor_Expr ( 'OR' Xor_Expr )*;
Constant_Expr        : Expression; // a constant expression must evaluate to a constant value at compile time
Xor_Expr             : And_Expr ( 'XOR' And_Expr )*;
And_Expr             : Compare_Expr ( ( '&' | 'AND' ) Compare_Expr )*;
Compare_Expr         : ( Equ_Expr ( ( '=' | '<>' ) Equ_Expr )* );
Equ_Expr             : Add_Expr ( ( '<' | '>' | '<=' | '>=' ) Add_Expr )*;
Add_Expr             : Term ( ( '+' | '-' ) Term )*;
Term                 : Power_Expr ( '*' | '/' | 'MOD' Power_Expr )*;
Power_Expr           : Unary_Expr ( '**' Unary_Expr )*;
Unary_Expr           : '-' | '+' | 'NOT' ? Primary_Expr;
Primary_Expr         : Constant | Enum_Value | Variable_Access | Func_Call | Ref_Value| '(' Expression ')';
Variable_Access      : Variable Multibit_Part_Access ?;
Multibit_Part_Access : '.' ( Unsigned_Int | '%' ( 'X' | 'B' | 'W' | 'D' | 'L' ) ? Unsigned_Int );
Func_Call            : Func_Access '(' ( Param_Assign ( ',' Param_Assign )* )? ')';
Stmt_List            : ( Stmt ? ';' )*;
Stmt                 : Assign_Stmt | Subprog_Ctrl_Stmt | Selection_Stmt | Iteration_Stmt;
Assign_Stmt          : ( Variable ':=' Expression ) | Ref_Assign | Assignment_Attempt;
Assignment_Attempt   : ( Ref_Name | Ref_Deref ) '?=' ( Ref_Name | Ref_Deref | Ref_Value );
Invocation           : ( FB_Instance_Name | Method_Name | 'THIS'
                     | ( ( 'THIS' '.' )? ( ( ( FB_Instance_Name | Class_Instance_Name ) '.' )+ ) Method_Name ) )
                       '(' ( Param_Assign ( ',' Param_Assign )* )? ')';
Subprog_Ctrl_Stmt    : Func_Call | Invocation | 'SUPER' '(' ')' | 'RETURN';
Param_Assign         : ( ( Variable_Name ':=' )? Expression ) | Ref_Assign | ( 'NOT' ? Variable_Name '=>' Variable );
Selection_Stmt       : IF_Stmt | Case_Stmt;
IF_Stmt              : 'IF' Expression 'THEN' Stmt_List ( 'ELSIF' Expression 'THEN' Stmt_List )* ( 'ELSE' Stmt_List )?
                       'END_IF';
Case_Stmt            : 'CASE' Expression 'OF' Case_Selection + ( 'ELSE' Stmt_List )? 'END_CASE';
Case_Selection       : Case_List ':' Stmt_List;
Case_List            : Case_List_Elem ( ',' Case_List_Elem )*;
Case_List_Elem       : Subrange | Constant_Expr;
Iteration_Stmt       : For_Stmt | While_Stmt | Repeat_Stmt | 'EXIT' | 'CONTINUE';
For_Stmt             : 'FOR' Control_Variable ':=' For_List 'DO' Stmt_List 'END_FOR';
Control_Variable     : Identifier;
For_List             : Expression 'TO' Expression ( 'BY' Expression )?;
While_Stmt           : 'WHILE' Expression 'DO' Stmt_List 'END_WHILE';
Repeat_Stmt          : 'REPEAT' Stmt_List 'UNTIL' Expression 'END_REPEAT';

// Table 73 - 76 - Graphic languages elements
Ladder_Diagram  : LD_Rung *;
LD_Rung         : 'LD_Rung syntax for graphical languages not shown here';
FB_Diagram      : FBD_Network *;
FBD_Network     : 'FBD_Network syntax for graphical languages not shown here';
Other_Languages : 'Other_Languages syntax for other languages not shown here'; 
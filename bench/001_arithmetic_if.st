PROGRAM arithmetic
  VAR
    a : INT;
    j : INT;
    k : INT := 42;

    i       : INT;
    divisor : INT;
    mult    : INT;
    test    : INT;
  END_VAR

  a := 15+j;
  j := 5 + 2*a*(30-a)/5; 
  k := k + -2;

  i := i+1;
  IF i MOD 5 = 0 THEN
    divisor := 5;
    IF i/5 = 1 THEN    // тест условия внутри условия
      mult := 1;
    ELSIF i/5 = 2 THEN // тест блока ELSIF
      mult := 2;
    ELSE               // тест блока ELSE
      mult := 3;
    END_IF
  ELSIF i MOD 4 = 0 THEN
    divisor := 4;
    mult := i/4;
  ELSIF i MOD 3 = 0 THEN
    divisor := 3;
    mult := i/3;
  ELSIF i MOD 2 = 0 THEN
    divisor := 2;
    mult := i/2;
  ELSIF i MOD 1 = 0 THEN
    divisor := 1;
    mult := i;
  END_IF

  IF mult <> 1 THEN
    test := mult*divisor;
  END_IF
END_PROGRAM
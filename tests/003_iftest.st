PROGRAM iftest
  VAR
    i       : INT;
    divisor : INT;
  END_VAR

  i := i+1;
  IF i MOD 5 = 0 THEN
    divisor := 5;
  ELSEIF i MOD 4 = 0 THEN
    divisor := 4;
  ELSEIF i MOD 3 = 0 THEN
    divisor := 3;
  ELSEIF i MOD 2 = 0 THEN
    divisor := 2;
  ELSE
    divisor := 1;
  END_IF
END_PROGRAM
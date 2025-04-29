PROGRAM iftest
  VAR
    i : INT;
  END_VAR

  i := i+10;
  IF i MOD 2 = 0 THEN
    i := i / 2;
  END_IF
END_PROGRAM
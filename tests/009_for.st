PROGRAM ForTest
  VAR
    v : REAL;
    i : REAL;
  END_VAR

  FOR i := 10 TO 8 BY -0.5 DO
    v := v + i; // 10+9.5+9+8.5+8=45
  END_FOR
END_PROGRAM
PROGRAM cast
  VAR
    f1 : REAL := 1; 
    f2 : LREAL := -0.5;
    f3 : REAL := 5.E1;
    test: INT;
  END_VAR
  test := test + 1;
  IF test = 1 THEN
    // f1 := f2; // ошибка каста
    f1 := LREAL_TO_REAL(f2);
    f2 := f3; // скрытый каст
  ELSIF test = 2 THEN
    f3 := f1 * 2; // -0,5*2=-1
    f2 := f1 + f3;  // -0,5+-1 = -1,5
  ELSIF test = 3 THEN
    f2 := f3/f1-f2; // -1/-0,5 - -1,5 = 3,5
  END_IF
END_PROGRAM
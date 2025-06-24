PROGRAM while_test
  VAR
    i : INT; 
    r : INT;
  END_VAR
  WHILE i<5 DO
    r := r+i;
    i := i+1;
  END_WHILE
  i := i-3;
END_PROGRAM
PROGRAM while_test
  VAR
    i : INT; 
    counter : INT; 
    r : INT;
    r2 : INT;
  END_VAR
  WHILE i<5 DO
    r := r+i;
    i := i+1;
  END_WHILE
  i := i-3;

  // ===== ТЕСТ: WHILE с CONTINUE (пропуск итерации) =====
  counter := 0;
  r2 := 0;
  WHILE counter < 5 DO
      counter := counter + 1;
      IF counter MOD 2 = 0 THEN
          CONTINUE; // Пропускаем четные числа
      END_IF
      r2 := r2 + counter; // Суммируем только нечетные
  END_WHILE
END_PROGRAM
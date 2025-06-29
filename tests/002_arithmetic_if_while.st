PROGRAM arithmetic
  VAR
    // арифметика
    a : INT;
    j : INT;
    k : INT := 42;

    // условия
    i       : INT;
    divisor : INT;
    mult    : INT;
    test    : INT;

    // циклы
    counter1 : INT;
    r        : INT;
    counter2 : INT;
    r2       : INT;
    counter3 : INT;
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

  WHILE counter1<100 DO
    r := r+counter1;
    counter1 := counter1+1;
  END_WHILE
  counter1 := counter1-55;

  WHILE counter2 < 100 DO
      counter2 := counter2 + 1;
      IF counter2 MOD 2 = 0 THEN
          CONTINUE; // Пропускаем четные числа
      END_IF
      r2 := r2 + counter2; // Суммируем только нечетные
  END_WHILE
  counter2 := 0;

  WHILE counter3 < 100 DO
      counter3 := counter3 + 1;
      IF counter3 >= 70 THEN
          EXIT; // Выходим после 70 итераций
      END_IF
  END_WHILE
  counter3 := 0;
END_PROGRAM
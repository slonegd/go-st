PROGRAM arithmetic
  VAR
    i : INT;
    j : INT;
    k : INT := 42;
  END_VAR

  i := 15+j;
  j := 5 + 2*i*(30-i)/5; // проверка приоритета арифмитических операций
  k := k + -2;           // проверка, что унарный минус отработает
END_PROGRAM
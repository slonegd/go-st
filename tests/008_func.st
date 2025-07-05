PROGRAM FunctionTest
  VAR
    in        : INT := 0;
    square    : INT := 0;
    factorial : INT := 0;
  END_VAR

  in := 5;
  square := Square(in);
  // factorial := Factorial(in);
END_PROGRAM

// Функция вычисления квадрата числа
FUNCTION Square : INT
  VAR_INPUT
    num : INT;
  END_VAR
    Square := num * num;
END_FUNCTION

// // Функция вычисления факториала (рекурсивная)
// FUNCTION FactorialRecurce : INT
//   VAR_INPUT
//     n : INT;
//   END_VAR
//   VAR
//     temp : INT;
//   END_VAR
//     IF n <= 1 THEN
//         Factorial := 1;
//     ELSE
//         Factorial := n * FactorialRecurce(n - 1);
//     END_IF
// END_FUNCTION

// FUNCTION Factorial : INT
// VAR_INPUT
//     n : INT;
// END_VAR
// VAR
//     result : INT := 1;
//     i : INT;
// END_VAR
//     FOR i := 1 TO n BY 1 DO
//         result := result * i;
//     END_FOR
//     Factorial := result;
// END_FUNCTION
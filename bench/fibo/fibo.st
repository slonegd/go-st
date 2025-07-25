PROGRAM fibo
    VAR
        fibo5  : INT;
        fibo10 : INT;
    END_VAR

    fibo5  := Fibonacci(5);
    fibo10 := Fibonacci(10);
END_PROGRAM

FUNCTION Fibonacci : INT
VAR_INPUT
    n : INT;
END_VAR
VAR
    a : INT := 0;
    b : INT := 1;
    i : INT;
    temp : INT;
END_VAR

    IF n <= 1 THEN
        Fibonacci := n;
    ELSE
        FOR i := 2 TO n DO
            temp := a + b;
            a := b;
            b := temp;
        END_FOR;
        Fibonacci := b;
    END_IF;
END_FUNCTION

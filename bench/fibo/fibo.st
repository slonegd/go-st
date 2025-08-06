PROGRAM fibo
    VAR
        r  : INT;
        n  : INT;
    END_VAR

    n := n + 1;
    IF n > 100 THEN
        n := 100;
    END_IF;
    r := Fibonacci(n);
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

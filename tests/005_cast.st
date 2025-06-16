PROGRAM cast
  VAR
    i8   : SINT;
    i16  : INT;
    i32  : DINT;
    i64  : LINT;
    ui8  : USINT;
    ui16 : UINT;
    ui32 : UDINT;
    ui64 : ULINT;
    test : INT;
  END_VAR

  test := test + 1;
  IF test = 1 THEN
    // max unsigned
    i8 := i8-1;
    i16 := i16-1;
    i32 := i32-1;
    i64 := i64-1;
    ui8 := ui8-1;
    // ui8 := ui16-1; // если раскомментировать, можно увидеть ошибку
    ui16 := ui16-1;
    ui32 := ui32-1;
    ui64 := ui64-1;
    // i8 := i8 + ui64; // если раскомментировать, можно увидеть ошибку
  ELSIF test = 2 THEN
    // max signed
    i8 := 16#7F;
    i16 := 16#7FFF;
    i32 := 16#7FFF_FFFF;
    i64 := 16#7FFF_FFFF_FFFF_FFFF;
  ELSIF test = 3 THEN
    // max signed
    i8 := i8 + 1;
    i16 := i16 + 1;
    i32 := i32 + 1;
    i64 := i64 + 1;
    ui16 := INT_TO_UINT(i16);
    // ui16 := DINT_TO_UINT(i16); // если раскомментировать, можно увидеть ошибку
  END_IF
END_PROGRAM
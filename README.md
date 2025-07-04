# go-st

go-st - это транслятор и интерпретатор скриптов на языке ST (Structured Text) стандарта IEC 61131-3, написанный на Go. Он реализует подмножество языка ST, фокусируясь на целочисленных переменных, арифметических операциях и условном выполнении. Бенчмарки показывают, что эта реализация обеспечивает один из самых быстрых интерпретаторов по сравнению с другими скриптовыми движками.

```
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
Benchmark_go_001-4        0.0000165 ns/op (x0.08)  // pure go
Benchmark_st_001-4        0.0002076 ns/op (x1)     // this
Benchmark_yaegi_001-4     0.001653 ns/op  (x8)     // github.com/traefik/yaegi v0.16.1
Benchmark_goscript_001-4  0.002492 ns/op  (x12)    // github.com/japm/goScript
Benchmark_lua_001-4       0.002826 ns/op  (x13)    // github.com/Shopify/go-lua
Benchmark_expr_001-4      0.005552 ns/op  (x26)    // github.com/expr-lang/expr v1.17.5
Benchmark_tengo_001-4     0.01662 ns/op   (x80)    // github.com/d5/tengo/v2
```
Подробнее см: [bench/arithmetic_if/0_bench_test.go](https://github.com/slonegd/go-st/blob/main/bench/arithmetic_if/0_bench_test.go)

## Реализовано

*   **Базовый синтаксис ST:** Включает структуру программы, объявления переменных (ограничено типом INT), операторы присваивания и операторы IF. (Определено в `antlr/ST.g4`)
*   **Целочисленные переменные:** Поддерживает целочисленные переменные и основные арифметические операции (+, -, \*, /, MOD).
*   **Условное выполнение:** Реализованы операторы IF, THEN и ELSE.
*   **Интерпретатор:** Предоставляет структуру интерпретатора для выполнения разобранного кода ST. (`st.go`, `ast/visitor.go`)
*   **Парсинг на основе antlr:** Использует antlr для генерации лексера и парсера.

## Не реализовано

*   **Типы данных:** Поддерживается только INT. Требуется добавление поддержки BOOL, REAL, STRING и TIME.
*   **Больше операторов:** Реализовать операторы CASE, FOR, WHILE и REPEAT.
*   **Логические выражения:** Условия содержат только выражения, требуются логические выражения.
*   **Функции и функциональные блоки:** Добавить поддержку функций и функциональных блоков.
*   **Массивы:** Реализовать выражения типа Array.
*   **И многое другое...**

## TODO

*   Реализовать другие типы данных: BOOL, REAL, STRING, TIME и т.д.
    *   Изменить правило `type_name` в [antlr/ST.g4](https://github.com/slonegd/go-st/blob/main/antlr/ST.g4).
    *   Обновить логику объявления переменных и присваивания в [ast/visitor.go](https://github.com/slonegd/go-st/blob/main/ast/visitor.go).
*   Реализовать операторы потока управления: CASE, FOR, WHILE, REPEAT.
    *   Добавить правила в [antlr/ST.g4](https://github.com/slonegd/go-st/blob/main/antlr/ST.g4).
    *   Реализовать соответствующую логику в [ast/visitor.go](https://github.com/slonegd/go-st/blob/main/ast/visitor.go).
*   Реализовать логические выражения
*   Реализовать функции и функциональные блоки.
    *   Расширить грамматику в [antlr/ST.g4](https://github.com/slonegd/go-st/blob/main/antlr/ST.g4).
    *   Реализовать логику выполнения в [st.go](https://github.com/slonegd/go-st/blob/main/st.go) и [ast/visitor.go](https://github.com/slonegd/go-st/blob/main/ast/visitor.go).
*   Реализовать выражения типа Array.

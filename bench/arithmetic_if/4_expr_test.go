package bench

import (
	_ "embed"
	"testing"

	"github.com/expr-lang/expr"
	_ "github.com/expr-lang/expr"
	_ "github.com/expr-lang/expr/vm"
	"github.com/stretchr/testify/assert"
)

const arithmetic_expr = `
	set("a", int(15 + j));
	set("j", int(5 + 2*a*(30-a)/5));
	set("k", k + -2);

	set("i", i + 1);

	// так как в expr поддерживаются только выражения и нет присваиваний
	// приходится выкручиваться через функцию set и неиспользуемые переменные, как шаги скрипта
	let _ = i%5 == 0 ?
		let _ = set("divisor", 5);
		i/5 == 1 ?
			set("mult", 1)
		: i/5 == 1 ?
			set("mult", 2)
		:
			set("mult", 3)
	:
		i%4 == 0 ?
			let _ = set("divisor", 4);
			set("mult", int(i/4))
		: i%3 == 0 ?
			let _ = set("divisor", 3);
			set("mult", int(i/3))
		: i%2 == 0 ?
			let _ = set("divisor", 2);
			set("mult", int(i/2))
		: i%1 == 0 ?
			let _ = set("divisor", 1);
			set("mult", 1)
		: 0;
	
	mult != 1 ? 
		set("test", mult * divisor)
	: 0
`

func Benchmark_expr_001(b *testing.B) {
	env := map[string]any{
		"a": 0,
		"j": 0,
		"k": 42,

		"i":       0,
		"divisor": 0,
		"mult":    0,
		"test":    0,
	}
	env["set"] = func(k string, v any) struct{} { env[k] = v; return struct{}{} }
	script, _ := expr.Compile(arithmetic_expr, expr.Env(env))

	for b.Loop() {
		expr.Run(script, env)
	}
}

func Test_expr_001(t *testing.T) {
	env := map[string]any{
		"a": 0,
		"j": 0,
		"k": 42,

		"i":       0,
		"divisor": 0,
		"mult":    0,
		"test":    0,
	}
	env["set"] = func(k string, v any) struct{} { env[k] = v; return struct{}{} }
	script, err := expr.Compile(arithmetic_expr, expr.Env(env))
	assert.NoError(t, err)

	_, err = expr.Run(script, env)
	assert.NoError(t, err)
	assert.Equal(t, 1, env["i"])

	_, err = expr.Run(script, env)
	_, err = expr.Run(script, env)
	_, err = expr.Run(script, env)
	_, err = expr.Run(script, env)
	_, err = expr.Run(script, env)
	_, err = expr.Run(script, env)
	assert.Equal(t, 7, env["i"])
	assert.Equal(t, 1, env["divisor"])
	assert.Equal(t, 1, env["mult"])
	assert.Equal(t, 6, env["test"])
}

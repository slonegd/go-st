package bench

var (
	fibo5  int
	fibo10 int
)

func fibo_go(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func exec() {
	fibo5 = fibo_go(5)
	fibo10 = fibo_go(10)
}

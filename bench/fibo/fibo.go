package bench

var (
	r int
	n int
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
	n++
	if n > 100 {
		n = 100
	}
	r = fibo_go(n)
}

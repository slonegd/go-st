package bench

var (
	a int
	j int
	k int = 42

	i       int
	divisor int
	mult    int
	test    int
)

func exec() {
	a = 15 + j
	j = 5 + 2*a*(30-a)/5
	k = k + -2

	i++
	if i%5 == 0 {
		divisor = 5
		if i/5 == 1 {
			mult = 1
		} else if i/5 == 2 {
			mult = 2
		} else {
			mult = 3
		}
	} else if i%4 == 0 {
		divisor = 4
		mult = i / 4
	} else if i%3 == 0 {
		divisor = 3
		mult = i / 3
	} else if i%2 == 0 {
		divisor = 2
		mult = i / 2
	} else if i%1 == 0 {
		divisor = 1
		mult = 1
	}

	if mult != 1 {
		test = mult * divisor
	}
}

package main

func fib(n int) int {
	a := 0
	b := 0
	c := 1

	for i := 1; i < n-1; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}

func main() {
	return fib(15)
}

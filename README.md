# goi

An interpreter for Go.

### Use

```
$ go build main.go
$ cat examples/fib.go
package main

func fib(a int) int {
	if a == 1 {
		return 0
	}

	if a == 2 {
		return 1
	}

	return fib(a-1) + fib(a-2)
}

func main() {
	return fib(15)
}
$ ./main examples/fib.go
377
```

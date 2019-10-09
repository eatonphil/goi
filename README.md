# goi

An interpreter for Go.

### Use

```
$ go build main.go
$ cat examples/if.go
package main

func main() {
	if 1 == 1 {
		return 3
	}

	return 4
}
$ ./main examples/if.go
3
```

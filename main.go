package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/eatonphil/goi/ast"
	"github.com/eatonphil/goi/llvm"
)

func main() {
	implementation := ast.Interpret
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--trace":
			TRACE = true
		case "--llvm":
			implementation = llvm.Compile
		}
	}

	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	implementation(f)
}

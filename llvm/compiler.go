package llvm

type valueType uint

type context map[string]valueType

func (c context) register(l string) {

}

type program struct {
	code []string
}

func (p *program) emit(code string, args ...interface{}) {
	p.code = append(p.code, fmt.Sprintf(code, args...))
}

func compileFuncDecl(ctx context, p program, fd *ast.FuncDecl) {
	safe := ctx.register(fd.Name.String())
	params := []string{}
	for _, arg := range fd.Type.Params.List {
		a := arg.Names[0].Name
		params = append(params, "i64 "+a)
	}
	p.emit("define %s @%s(%s) {", "i64", safe, strings.Join(params, ", "))
}

func Compile(f *ast.File) {
	ctx := context{}
	var p program

	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			compileFuncDecl(ctx, p, fd)
		}
	}
}

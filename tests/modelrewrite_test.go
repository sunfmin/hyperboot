package tests

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"testing"
)

func TestRewrite(t *testing.T) {
	fset := token.NewFileSet()
	fl, err := os.Open("dummy/models.go")
	if err != nil {
		t.Error(err)
	}
	defer fl.Close()
	f, err := parser.ParseFile(fset, "dummy/models.go", fl, parser.ParseComments)
	ast.Print(fset, f)
	t.Errorf("%+v, %+v", f, err)
}

package formatter_test

import (
	"bytes"
	"github.com/VKCOM/php-parser/pkg/token"
	"github.com/VKCOM/php-parser/pkg/visitor/formatter"
	"github.com/VKCOM/php-parser/pkg/visitor/printer"
	"testing"

	"github.com/VKCOM/php-parser/pkg/ast"
)

func TestFormatter_Root(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter()
	n.Accept(f)

	p := printer.NewPrinter(o)
	n.Accept(p)

	expected := `<?php 

;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Nullable(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Nullable{
		Expr: &ast.Identifier{
			Value: []byte("array"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `?array`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Parameter{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Parameter{
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `&$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Variadic(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Parameter{
		VariadicTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `...$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Type(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Parameter{
		Type: &ast.Identifier{
			Value: []byte("array"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `array $var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Parameter_Default(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Parameter{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		DefaultValue: &ast.ScalarString{
			Value: []byte("'default'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$var = 'default'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Identifier(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Identifier{
		Value: []byte("foo"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Argument{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Argument{
		AmpersandTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `&$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Argument_Variadic(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Argument{
		VariadicTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `...$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtBreak{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `break;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtBreak_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtBreak{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `break $var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Case(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `case $var:
        ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Catch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtCatch{
		Types: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$baz"),
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `catch (foo | bar $baz) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `class foo {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Modifier(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClass{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("final"),
			},
		},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `final class foo {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Anonymous(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `class foo($a, $b) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Extends(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Extends: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `class foo extends bar {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_Class_Implements(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClass{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Implements: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `class foo implements bar {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `const foo = 'foo', bar = 'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtClassConstList_Modifier(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassConstList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
		},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `public const foo = 'foo', bar = 'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtNop{},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo() ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Modifier(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
		},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `public function foo() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassMethod{
		AmpersandTkn: &token.Token{},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function &foo() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_Params(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo($a, $b) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ClassMethod_ReturnType(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtClassMethod{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo(): bar {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `const foo = 'foo', bar = 'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtConstant{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Expr: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo = 'bar'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtContinue{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `continue;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtContinue_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtContinue{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `continue $var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDeclare(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtDeclare{
		Stmt: &ast.StmtNop{},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `declare() ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDeclare_Consts(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("foo"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'foo'"),
				},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{
					Value: []byte("bar"),
				},
				Expr: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `declare(foo = 'foo', bar = 'bar') {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDefault(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `default:
        ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtDo(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtDo{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `do {
        ;
    } while($var);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtEcho(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ScalarString{
				Value: []byte("'foo'"),
			},
			&ast.ScalarString{
				Value: []byte("'bar'"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `echo 'foo', 'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtElse(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `else {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtElseIf(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `elseif($var) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtExpression{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$var"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFinally{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `finally {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFor(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `for($foo, $bar; $foo, $bar; $foo, $bar) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foreach($foo as $val) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foreach($foo as &$val) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtForeach_Key(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$key"),
			},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$val"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foreach($foo as $key => $val) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFunction{
		AmpersandTkn: &token.Token{},
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function &foo() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_Params(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo($a, $b) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtFunction_ReturnType(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtFunction{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo(): bar {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `global $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtGoto{
		Label: &ast.Identifier{
			Value: []byte("FOO"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtHaltCompiler{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `if ($foo) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf_ElseIf(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$baz"),
					},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtNop{},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `if ($foo) {
        ;
    } elseif($bar) {
        ;
    } elseif($baz) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtIf_Else(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
		Else: &ast.StmtElse{
			Stmt: &ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `if ($foo) {
        ;
    } else {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
					&ast.StmtInlineHtml{
						Value: []byte("<div></div>"),
					},
					&ast.StmtEcho{
						Exprs: []ast.Vertex{
							&ast.ExprVariable{
								Name: &ast.Identifier{
									Value: []byte("$foo"),
								},
							},
						},
					},
					&ast.StmtInlineHtml{
						Value: []byte("<div></div>"),
					},
					&ast.StmtNop{},
				},
			},
		},
	}

	f := formatter.NewFormatter()
	n.Accept(f)

	p := printer.NewPrinter(o)
	n.Accept(p)

	expected := `<?php 

{
    ;?><div></div><?php 
    echo $foo;?><div></div><?php 
    ;
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInterface(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtInterface{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `interface foo {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtInterface_Extends(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtInterface{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Extends: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `interface foo extends bar {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtLabel(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtLabel{
		Name: &ast.Identifier{
			Value: []byte("FOO"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `FOO:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNamespace_Name(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtNamespace{
		Name: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `namespace foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNamespace_Stmts(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtNamespace{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `namespace {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtNop(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtNop{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtProperty{
		Var: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtProperty_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtProperty{
		Var: &ast.Identifier{
			Value: []byte("$foo"),
		},
		Expr: &ast.Identifier{
			Value: []byte("$bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo = $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtPropertyList{
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo, $bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList_Modifiers(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{
				Value: []byte("public"),
			},
			&ast.Identifier{
				Value: []byte("static"),
			},
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `public static $foo, $bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtPropertyList_Type(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtPropertyList{
		Type: &ast.Identifier{
			Value: []byte("array"),
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$foo"),
				},
			},
			&ast.StmtProperty{
				Var: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `array $foo, $bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtReturn(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtReturn{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `return;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtReturn_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtReturn{
		Expr: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `return $foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStatic(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `static $a, $b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStaticVar(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStaticVar_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.Identifier{
			Value: []byte("$bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo = $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtNop{},
						},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter()
	n.Accept(f)

	p := printer.NewPrinter(o)
	n.Accept(p)

	expected := `<?php 

{
    {
        ;
    }
}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{
					Value: []byte("'bar'"),
				},
				Stmts: []ast.Vertex{
					&ast.StmtBreak{},
				},
			},
			&ast.StmtDefault{
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `switch($foo) {
        case 'bar':
            break;
        default:
            ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtThrow(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtThrow{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `throw $foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTrait(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTrait{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `trait foo {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use foo, bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUse_Adaptations(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("foo"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
		},
		Adaptations: []ast.Vertex{
			&ast.StmtTraitUseAlias{
				Method: &ast.Identifier{
					Value: []byte("foo"),
				},
				Alias: &ast.Identifier{
					Value: []byte("baz"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use foo, bar {
        foo as baz;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo as public bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Trait(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUseAlias{
		Trait: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
		Alias: &ast.Identifier{
			Value: []byte("baz"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::bar as public baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Alias(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo as bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUseAlias_Modifier(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUseAlias{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Modifier: &ast.Identifier{
			Value: []byte("public"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo as public;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUsePrecedence{
		Method: &ast.Identifier{
			Value: []byte("foo"),
		},
		Insteadof: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("bar"),
					},
				},
			},
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("baz"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo insteadof bar, baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTraitUsePrecedence_Trait(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTraitUsePrecedence{
		Trait: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Insteadof: []ast.Vertex{
			&ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("baz"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::bar insteadof baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `try {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry_Catch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		Catches: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Value: []byte("foo"),
							},
						},
					},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Value: []byte("foo"),
							},
						},
					},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtNop{},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `try {
        ;
    } catch (foo $bar) {
        ;
    } catch (foo $bar) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtTry_Finally(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
		Finally: &ast.StmtFinally{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `try {
        ;
    } finally {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `unset($a, $b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUse(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUseList{
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("foo"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use foo, bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUse_Type(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUseList{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("foo"),
						},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use function foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtGroupUseList{
		Prefix: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("baz"),
						},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use foo\{bar, baz};`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtGroupUse_Type(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtGroupUseList{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Prefix: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("bar"),
						},
					},
				},
			},
			&ast.StmtUse{
				Use: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Value: []byte("baz"),
						},
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `use function foo\{bar, baz};`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUse{
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration_Type(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUse{
		Type: &ast.Identifier{
			Value: []byte("function"),
		},
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtUseDeclaration_Alias(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtUse{
		Use: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Alias: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo as bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_StmtWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Stmt: &ast.StmtNop{},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `while($foo) ;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `array($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayDimFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Dim: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo[$bar]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrayItem{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem_Key(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrayItem{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo => $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrayItem_Variadic(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrayItem{
		EllipsisTkn: &token.Token{},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `...$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrowFunction{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `fn() => $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrowFunction{
		AmpersandTkn: &token.Token{},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `fn&() => $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_Params(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrowFunction{
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `fn($a, $b) => $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprArrowFunction_ReturnType(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprArrowFunction{
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `fn(): foo => $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBitwiseNot(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `~$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBooleanNot(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `!$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBrackets(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBrackets{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `($foo)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClassConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClassConstFetch{
		Class: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Const: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo::bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClone(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClone{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `clone $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosure{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Ref(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosure{
		AmpersandTkn: &token.Token{},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function&() {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Params(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosure{
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function($a, $b) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_ReturnType(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosure{
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function(): foo {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosure_Use(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosure{
		Uses: []ast.Vertex{
			&ast.ExprClosureUse{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$foo"),
					},
				},
			},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `function() use($foo) {
        ;
    }`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosureUse(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosureUse{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$a"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprClosureUse_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprClosureUse{
		AmpersandTkn: &token.Token{},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$a"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `&$a`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprConstFetch{
		Const: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("FOO"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `FOO`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprEmpty{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `empty($foo)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprErrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprErrorSuppress{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `@$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprEval(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprEval{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `eval($foo)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprExit(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprExit{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `exit`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprExit_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprExit{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `exit($foo)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprFunctionCall{
		Function: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo()`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprFunctionCall_Arguments(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprFunctionCall{
		Function: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$bar"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo($bar)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprInclude(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprInclude{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `include 'foo.php'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprIncludeOnce{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `include_once 'foo.php'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprInstanceOf{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("bar"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo instanceof bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprIsset(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$a"),
				},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$b"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `isset($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprList(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `list($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo->bar()`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo->{'bar'}()`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprMethodCall_Arguments(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Method: &ast.Identifier{
			Value: []byte("bar"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo->bar($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprNew(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `new foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprNew_Arguments(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `new foo($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPreDec(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPreDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `--$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPreInc(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPreInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `++$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPostDec(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPostDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo--`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPostInc(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPostInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo++`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPrint(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPrint{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `print $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Prop: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprPropertyFetch_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Prop: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo->{'bar'}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprRequire{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `require 'foo.php'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprRequireOnce{
		Expr: &ast.ScalarString{
			Value: []byte("'foo.php'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `require_once 'foo.php'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprShellExec{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := "``"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec_Part(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := "`foo`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprShellExec_Parts(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := "`foo $bar baz`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.Identifier{
			Value: []byte("bar"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::bar()`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall_Expr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::{'bar'}()`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticCall_Arguments(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprStaticCall{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Call: &ast.Identifier{
			Value: []byte("bar"),
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$a"),
					},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{
						Value: []byte("$b"),
					},
				},
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::bar($a, $b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprStaticPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprStaticPropertyFetch{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("foo"),
				},
			},
		},
		Prop: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo::$bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprTernary(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		IfTrue: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$baz"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ? $bar : $baz`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprTernary_short(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ?: $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `-$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprUnaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `+$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprVariable{
		Name: &ast.Identifier{
			Value: []byte("$foo"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable_Variable(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprVariable{
		Name: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprVariable_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprVariable{
		Name: &ast.ScalarString{
			Value: []byte("'foo'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `${'foo'}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYield(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprYield{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `yield $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYield_Key(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprYield{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `yield $foo => $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `yield from $foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssign{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo = $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignReference(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignReference{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo =& $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignBitwiseAnd{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo &= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignBitwiseOr{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo |= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignBitwiseXor{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ^= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignCoalesce{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ??= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignConcat{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo .= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignDiv{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo /= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignMinus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo -= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMod(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignMod{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo %= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignMul(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignMul{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo *= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignPlus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo += $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignPow(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignPow{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo **= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignShiftLeft{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo <<= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprAssignShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprAssignShiftRight{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo >>= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryBitwiseAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo & $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryBitwiseOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo | $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryBitwiseXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ^ $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBooleanAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryBooleanAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo && $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryBooleanOr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryBooleanOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo || $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryCoalesce{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ?? $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryConcat{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo . $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryDiv{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo / $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo == $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryGreater(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryGreater{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo > $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryGreaterOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryGreaterOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo >= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo === $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryLogicalAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo and $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalOr(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryLogicalOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo or $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryLogicalXor(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryLogicalXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo xor $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryMinus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo - $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMod(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryMod{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo % $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryMul(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryMul{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo * $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryNotEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryNotEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo != $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryNotIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryNotIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo !== $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryPlus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo + $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryPow(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryPow{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo ** $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryShiftLeft{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo << $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinaryShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinaryShiftRight{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo >> $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySmaller(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinarySmaller{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo < $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySmallerOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinarySmallerOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo <= $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprBinarySpaceship(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprBinarySpaceship{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `$foo <=> $bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastArray(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastArray{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(array)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastBool(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastBool{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(bool)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastDouble{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(float)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastInt(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastInt{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(int)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastObject(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastObject{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(object)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastString(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastString{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(string)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ExprCastUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ExprCastUnset{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `(unset)$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarDnumber(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarDnumber{
		Value: []byte("1234"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `1234`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsed{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `""`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed_Part(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `"foo"`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsed_Parts(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `"foo $bar baz"`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsedStringPart{
		Value: []byte("foo"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringVar(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsedStringVar{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `${foo}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringVar_Dim(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsedStringVar{
		Name: &ast.Identifier{
			Value: []byte("foo"),
		},
		Dim: &ast.ScalarString{
			Value: []byte("'bar'"),
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `${foo['bar']}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarEncapsedStringBrackets(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarEncapsedStringBrackets{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{
				Value: []byte("$foo"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `{$foo}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarHeredoc{}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `<<<EOT
EOT`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc_Part(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo\n"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `<<<EOT
foo
EOT`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarHeredoc_Parts(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{
				Value: []byte("foo "),
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{
					Value: []byte("$bar"),
				},
			},
			&ast.ScalarEncapsedStringPart{
				Value: []byte(" baz\n"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `<<<EOT
foo $bar baz
EOT`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarLnumber(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarLnumber{
		Value: []byte("1234"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `1234`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarMagicConstant{
		Value: []byte("__DIR__"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `__DIR__`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_ScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.ScalarString{
		Value: []byte("'foo'"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `'foo'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameName(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo\bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.NameFullyQualified{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `\foo\bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.NameRelative{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("foo"),
			},
			&ast.NamePart{
				Value: []byte("bar"),
			},
		},
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `namespace\foo\bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestFormatter_NameNamePart(t *testing.T) {
	o := bytes.NewBufferString("")

	n := &ast.NamePart{
		Value: []byte("foo"),
	}

	f := formatter.NewFormatter().WithState(formatter.FormatterStatePHP).WithIndent(1)
	n.Accept(f)

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n.Accept(p)

	expected := `foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

package printer_test

import (
	"bytes"
	"github.com/VKCOM/php-parser/pkg/token"
	"github.com/VKCOM/php-parser/pkg/visitor/printer"
	"testing"

	"github.com/VKCOM/php-parser/pkg/ast"
)

func TestPrinterPrintFile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Foo")},
					},
				},
			},
			&ast.StmtClass{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
				Name: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Bar")},
					},
				},
				Extends: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{Value: []byte("Baz")},
					},
				},
				Stmts: []ast.Vertex{
					&ast.StmtClassMethod{
						Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
						Name:      &ast.Identifier{Value: []byte("greet")},
						Stmt: &ast.StmtStmtList{
							Stmts: []ast.Vertex{
								&ast.StmtEcho{
									Exprs: []ast.Vertex{
										&ast.ScalarString{Value: []byte("'Hello world'")},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `<?php namespace Foo;abstract class Bar extends Baz{public function greet(){echo'Hello world';}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFileInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ScalarString{
						Value: []byte(`"a"`),
					},
				},
			},
			&ast.StmtInlineHtml{Value: []byte("<div>HTML</div>")},
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ScalarString{
						Value: []byte(`"b"`),
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `<div>HTML</div><?php echo"a";?><div>HTML</div><?php echo"b";`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// node

func TestPrinterPrintIdentifier(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Identifier{
		Value: []byte("test"),
	}
	n.Accept(p)

	expected := `test`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintParameter(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Parameter{
		Type: &ast.NameFullyQualified{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		VariadicTkn: &token.Token{
			Value: []byte("..."),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		DefaultValue: &ast.ScalarString{
			Value: []byte("'default'"),
		},
	}
	n.Accept(p)

	expected := "\\Foo...$var='default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNullable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Nullable{
		Expr: &ast.Parameter{
			Type: &ast.NameFullyQualified{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Value: []byte("Foo"),
					},
				},
			},
			AmpersandTkn: &token.Token{
				Value: []byte("&"),
			},
			Var: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			DefaultValue: &ast.ScalarString{
				Value: []byte("'default'"),
			},
		},
	}
	n.Accept(p)

	expected := "?\\Foo&$var='default'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintArgument(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Argument{
		VariadicTkn: &token.Token{
			Value: []byte("..."),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := "...$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}
func TestPrinterPrintArgumentByRef(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Argument{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := "&$var"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// name

func TestPrinterPrintNameNamePart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.NamePart{
		Value: []byte("foo"),
	}
	n.Accept(p)

	expected := "foo"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameName(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Name{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	n.Accept(p)

	expected := "Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameFullyQualified(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.NameFullyQualified{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	n.Accept(p)

	expected := "\\Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNameRelative(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.NameRelative{
		Parts: []ast.Vertex{
			&ast.NamePart{
				Value: []byte("Foo"),
			},
			&ast.NamePart{
				Value: []byte("Bar"),
			},
		},
	}
	n.Accept(p)

	expected := "namespace\\Foo\\Bar"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// scalar

func TestPrinterPrintScalarLNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarLnumber{
		Value: []byte("1"),
	}
	n.Accept(p)

	expected := "1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarDNumber(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarDnumber{
		Value: []byte(".1"),
	}
	n.Accept(p)

	expected := ".1"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarString{
		Value: []byte("'hello world'"),
	}
	n.Accept(p)

	expected := `'hello world'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsedStringPart(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarEncapsedStringPart{
		Value: []byte("hello world"),
	}
	n.Accept(p)

	expected := `hello world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarEncapsed(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarEncapsed{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world")},
		},
	}
	n.Accept(p)

	expected := `"hello $var world"`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarHeredoc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarHeredoc{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$var")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte(" world\n")},
		},
	}
	n.Accept(p)

	expected := `<<<EOT
hello $var world
EOT`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintScalarMagicConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ScalarMagicConstant{
		Value: []byte("__DIR__"),
	}
	n.Accept(p)

	if o.String() != `__DIR__` {
		t.Errorf("TestPrintScalarMagicConstant is failed\n")
	}
}

// assign

func TestPrinterPrintAssign(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssign{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignReference{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a=&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignBitwiseAnd{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a&=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignBitwiseOr{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a|=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignBitwiseXor{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a^=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignCoalesce{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a??=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignConcat{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a.=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignDiv{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a/=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignMinus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a-=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignMod{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a%=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignMul{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a*=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignPlus{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a+=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignPow{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a**=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignShiftLeft{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a<<=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAssignShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprAssignShiftRight{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a>>=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// binary

func TestPrinterPrintBinaryBitwiseAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryBitwiseAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryBitwiseOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a|$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBitwiseXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryBitwiseXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a^$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryBooleanAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a&&$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryBooleanOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryBooleanOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a||$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryCoalesce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryCoalesce{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a??$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryConcat(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryConcat{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a.$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryDiv(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryDiv{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a/$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreaterOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryGreaterOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a>=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryGreater(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryGreater{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a===$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalAnd(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryLogicalAnd{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a and$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalOr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryLogicalOr{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a or$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryLogicalXor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryLogicalXor{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a xor$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryMinus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a-$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryMod{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a%$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryMul(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryMul{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a*$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryNotEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a!=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryNotIdentical(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryNotIdentical{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a!==$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryPlus{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a+$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryPow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryPow{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a**$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftLeft(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryShiftLeft{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a<<$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinaryShiftRight(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinaryShiftRight{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a>>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmallerOrEqual(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinarySmallerOrEqual{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a<=$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySmaller(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinarySmaller{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a<$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBinarySpaceship(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBinarySpaceship{
		Left: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Right: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a<=>$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// cast

func TestPrinterPrintArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastArray{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(array)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintBool(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastBool{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(bool)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDouble(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastDouble{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(float)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInt(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastInt{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(int)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintObject(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastObject{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(object)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintString(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastString{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(string)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprCastUnset{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `(unset)$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// expr

func TestPrinterPrintExprArrayDimFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArrayDimFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Dim: &ast.ScalarLnumber{Value: []byte("1")},
	}
	n.Accept(p)

	expected := `$var[1]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemWithKey(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArrayItem{
		Key: &ast.ScalarString{Value: []byte("'Hello'")},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	n.Accept(p)

	expected := `'Hello'=>$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItem(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArrayItem{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	n.Accept(p)

	expected := `$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItem_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArrayItem{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	n.Accept(p)

	expected := `&$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrayItemUnpack(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArrayItem{
		EllipsisTkn: &token.Token{
			Value: []byte("..."),
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$world")},
		},
	}
	n.Accept(p)

	expected := `...$world`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArray{
		ArrayTkn: &token.Token{
			Value: []byte("array"),
		},
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `array('Hello'=>$world,2=>&$var,$var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBitwiseNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBitwiseNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `~$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBooleanNot(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `!$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprBracket(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprBooleanNot{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `!$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClassConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprClassConstFetch{
		Class: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Const: &ast.Identifier{
			Value: []byte("CONST"),
		},
	}
	n.Accept(p)

	expected := `$var::CONST`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClone(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprClone{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `clone$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprClosureUse{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
	}
	n.Accept(p)

	expected := `$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosureUse_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprClosureUse{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
	}
	n.Accept(p)

	expected := `&$foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprClosure(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprClosure{
		StaticTkn: &token.Token{
			Value: []byte("static"),
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Params: []ast.Vertex{
			&ast.Parameter{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
		Uses: []ast.Vertex{
			&ast.ExprClosureUse{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprClosureUse{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	n.Accept(p)

	expected := `static function&($var)use(&$a,$b):\Foo{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprArrowFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtExpression{
		Expr: &ast.ExprArrowFunction{
			StaticTkn: &token.Token{
				Value: []byte("static"),
			},
			AmpersandTkn: &token.Token{
				Value: []byte("&"),
			},
			Params: []ast.Vertex{
				&ast.Parameter{
					AmpersandTkn: &token.Token{
						Value: []byte("&"),
					},
					Var: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$var")},
					},
				},
			},
			ReturnType: &ast.NameFullyQualified{
				Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
			},
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
	}
	n.Accept(p)

	expected := `static fn&(&$var):\Foo=>$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprConstFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprConstFetch{
		Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}},
	}
	n.Accept(p)

	expected := "null"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprEmpty{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `empty($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrettyPrinterrorSuppress(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprErrorSuppress{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `@$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintEval(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprEval{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `eval($var)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExit(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprExit{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `exit$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintDie(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprExit{
		ExitTkn: &token.Token{
			Value: []byte("die"),
		},
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `die$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintFunctionCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprFunctionCall{
		Function: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `$var(&$a,...$b,$c)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInclude(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprInclude{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	n.Accept(p)

	expected := `include'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIncludeOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprIncludeOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	n.Accept(p)

	expected := `include_once'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInstanceOf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprInstanceOf{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Class: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
	}
	n.Accept(p)

	expected := `$var instanceof Foo`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIsset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprIsset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	n.Accept(p)

	expected := `isset($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprList{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							},
						},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `list($a,list($b,$c))`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintMethodCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprMethodCall{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
		Method: &ast.Identifier{Value: []byte("bar")},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `$foo->bar($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNew(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprNew{
		Class: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `new Foo($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPostDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `$var--`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPostInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPostInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `$var++`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreDec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPreDec{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `--$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPreInc(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPreInc{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `++$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPrint(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPrint{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `print$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprPropertyFetch{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$foo")},
		},
		Prop: &ast.Identifier{Value: []byte("bar")},
	}
	n.Accept(p)

	expected := `$foo->bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequire(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprRequire{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	n.Accept(p)

	expected := `require'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintRequireOnce(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprRequireOnce{
		Expr: &ast.ScalarString{Value: []byte("'path'")},
	}
	n.Accept(p)

	expected := `require_once'path'`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShellExec(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprShellExec{
		Parts: []ast.Vertex{
			&ast.ScalarEncapsedStringPart{Value: []byte("hello ")},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$world")},
			},
			&ast.ScalarEncapsedStringPart{Value: []byte("!")},
		},
	}
	n.Accept(p)

	expected := "`hello $world!`"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExprShortArray(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprArray{
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Key: &ast.ScalarString{Value: []byte("'Hello'")},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$world")},
				},
			},
			&ast.ExprArrayItem{
				Key: &ast.ScalarLnumber{Value: []byte("2")},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `['Hello'=>$world,2=>&$var,$var]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintShortList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprList{
		OpenBracketTkn: &token.Token{
			Value: []byte("["),
		},
		Items: []ast.Vertex{
			&ast.ExprArrayItem{
				Val: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.ExprArrayItem{
				Val: &ast.ExprList{
					Items: []ast.Vertex{
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$b")},
							},
						},
						&ast.ExprArrayItem{
							Val: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							},
						},
					},
				},
			},
		},
		CloseBracketTkn: &token.Token{
			Value: []byte("]"),
		},
	}
	n.Accept(p)

	expected := `[$a,list($b,$c)]`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticCall(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprStaticCall{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Call:  &ast.Identifier{Value: []byte("bar")},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `Foo::bar($a,$b)`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticPropertyFetch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprStaticPropertyFetch{
		Class: &ast.Identifier{Value: []byte("Foo")},
		Prop: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$bar")},
		},
	}
	n.Accept(p)

	expected := `Foo::$bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernary(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
	}
	n.Accept(p)

	expected := `$a?:$b`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTernaryFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprTernary{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		IfTrue: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$b")},
		},
		IfFalse: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$c")},
		},
	}
	n.Accept(p)

	expected := `$a?$b:$c`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryMinus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprUnaryMinus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `-$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUnaryPlus(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprUnaryPlus{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `+$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintVariable(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprVariable{
		DollarTkn: &token.Token{
			Value: []byte("$"),
		},
		OpenCurlyBracketTkn: &token.Token{
			Value: []byte("{"),
		},
		Name: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		CloseCurlyBracketTkn: &token.Token{
			Value: []byte("}"),
		},
	}
	n.Accept(p)

	expected := `${$var}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFrom(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprYieldFrom{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `yield from$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYield(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprYield{
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `yield$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintYieldFull(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.ExprYield{
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		Val: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `yield$k=>$var`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

// stmt

func TestPrinterPrintAltElseIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `elseif($a):$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseIfEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{},
	}
	n.Accept(p)

	expected := `elseif($a):`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElse{
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `else:$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltElseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElse{
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{},
	}
	n.Accept(p)

	expected := `else:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$c")},
			},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `for($a;$b;$c):$d;endfor;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$key")},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$val")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `foreach($var as$key=>$val):$d;endforeach;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltForeach_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$key")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$val")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `foreach($var as$key=>&$val):$d;endforeach;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltIf(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$d")},
				}},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
				ColonTkn: &token.Token{
					Value: []byte(":"),
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$b")},
						}},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
				ColonTkn: &token.Token{
					Value: []byte(":"),
				},
				Stmt: &ast.StmtStmtList{},
			},
		},
		Else: &ast.StmtElse{
			ColonTkn: &token.Token{
				Value: []byte(":"),
			},
			Stmt: &ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	n.Accept(p)

	expected := `if($a):$d;elseif($b):$b;elseif($c):else:$b;endif;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAltSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'a'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					}},
				},
			},
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'b'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	n.Accept(p)

	expected := `switch($var):case'a':$a;case'b':$b;endswitch;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintAltWhile(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		ColonTkn: &token.Token{
			Value: []byte(":"),
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `while($a):$b;endwhile;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtBreak(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtBreak{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
		},
	}
	n.Accept(p)

	expected := "break 1;"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCase(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	n.Accept(p)

	expected := `case$a:$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCaseEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtCase{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmts: []ast.Vertex{},
	}
	n.Accept(p)

	expected := "case$a:"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtCatch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtCatch{
		Types: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Exception")}}},
			&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("RuntimeException")}}},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$e")},
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	n.Accept(p)

	expected := `catch(Exception|\RuntimeException$e){$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("int")}}}},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `public function&foo(?int&$a=null,...$b):void{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAbstractClassMethod(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtClassMethod{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: &ast.Nullable{Expr: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("int")}}}},
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				DefaultValue: &ast.ExprConstFetch{Const: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("null")}}}},
			},
			&ast.Parameter{
				VariadicTkn: &token.Token{
					Value: []byte("..."),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		ReturnType: &ast.Name{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("void")}},
		},
		Stmt: &ast.StmtNop{},
	}
	n.Accept(p)

	expected := `public static function&foo(?int&$a=null,...$b):void;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtClass{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
		Name:      &ast.Identifier{Value: []byte("Foo")},
		Extends:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		Implements: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Quuz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassConstList{
				Modifiers: []ast.Vertex{
					&ast.Identifier{Value: []byte("public")},
					&ast.Identifier{Value: []byte("static")},
				},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `abstract class Foo extends Bar implements Baz,Quuz{public static const FOO='bar';}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtAnonymousClass(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtClass{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("abstract")}},
		Args: []ast.Vertex{
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.Argument{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
		Extends: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		Implements: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Quuz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassConstList{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Consts: []ast.Vertex{
					&ast.StmtConstant{
						Name: &ast.Identifier{Value: []byte("FOO")},
						Expr: &ast.ScalarString{Value: []byte("'bar'")},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `abstract class($a,$b)extends Bar implements Baz,Quuz{public const FOO='bar';}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtClassConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtClassConstList{
		Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("BAR")},
				Expr: &ast.ScalarString{Value: []byte("'b'")},
			},
		},
	}
	n.Accept(p)

	expected := `public const FOO='a',BAR='b';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtConstList{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("BAR")},
				Expr: &ast.ScalarString{Value: []byte("'b'")},
			},
		},
	}
	n.Accept(p)

	expected := `const FOO='a',BAR='b';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtConstant(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("FOO")},
		Expr: &ast.ScalarString{Value: []byte("'BAR'")},
	}
	n.Accept(p)

	expected := "FOO='BAR'"
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtContinue(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtContinue{
		Expr: &ast.ScalarLnumber{
			Value: []byte("1"),
		},
	}
	n.Accept(p)

	expected := `continue 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `declare(FOO='bar'){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	n.Accept(p)

	expected := `declare(FOO='bar')'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDeclareNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDeclare{
		Consts: []ast.Vertex{
			&ast.StmtConstant{
				Name: &ast.Identifier{Value: []byte("FOO")},
				Expr: &ast.ScalarString{Value: []byte("'bar'")},
			},
		},
		Stmt: &ast.StmtNop{},
	}
	n.Accept(p)

	expected := `declare(FOO='bar');`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalut(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDefault{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	n.Accept(p)

	expected := `default:$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDefalutEmpty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDefault{
		Stmts: []ast.Vertex{},
	}
	n.Accept(p)

	expected := `default:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_Expression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
		},
	}
	n.Accept(p)

	expected := `do$a;while(1);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtDo_StmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtDo{
		Cond: &ast.ScalarLnumber{Value: []byte("1")},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `do{$a;}while(1);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoHtmlState(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtEcho{
				Exprs: []ast.Vertex{
					&ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					},
					&ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `<?php echo$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtEchoPhpState(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtEcho{
		Exprs: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	n.Accept(p)

	expected := `echo$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `elseif($a){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	n.Accept(p)

	expected := `elseif($a)'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElseIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
	}
	n.Accept(p)

	expected := `elseif($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElse{
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `else{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseExpr(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElse{
		Stmt: &ast.StmtExpression{Expr: &ast.ScalarString{Value: []byte("'bar'")}},
	}
	n.Accept(p)

	expected := `else'bar';`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtElseNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtElse{
		Stmt: &ast.StmtNop{},
	}
	n.Accept(p)

	expected := `else;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtExpression{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
	}
	n.Accept(p)

	expected := `$a;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFinally(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtFinally{
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}
	n.Accept(p)

	expected := `finally{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFor(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtFor{
		Init: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		Cond: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$c")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$d")},
			},
		},
		Loop: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$e")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$f")},
			},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `for($a,$b;$c,$d;$e,$f){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$v")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `foreach($a as$k=>$v){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtForeach_Reference(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtForeach{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Key: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$k")},
		},
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$v")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `foreach($a as$k=>&$v){;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtFunction(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtFunction{
		AmpersandTkn: &token.Token{
			Value: []byte("&"),
		},
		Name: &ast.Identifier{Value: []byte("foo")},
		Params: []ast.Vertex{
			&ast.Parameter{
				AmpersandTkn: &token.Token{
					Value: []byte("&"),
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$var")},
				},
			},
		},
		ReturnType: &ast.NameFullyQualified{
			Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtNop{},
		},
	}
	n.Accept(p)

	expected := `function&foo(&$var):\Foo{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGlobal(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtGlobal{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	n.Accept(p)

	expected := `global$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGoto(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtGoto{
		Label: &ast.Identifier{Value: []byte("FOO")},
	}
	n.Accept(p)

	expected := `goto FOO;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintHaltCompiler(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtHaltCompiler{}
	n.Accept(p)

	expected := `__halt_compiler();`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfExpression(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtExpression{
			Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
		ElseIf: []ast.Vertex{
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$c")},
				},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{
							Expr: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$d")},
							},
						},
					},
				},
			},
			&ast.StmtElseIf{
				Cond: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$e")},
				},
				Stmt: &ast.StmtNop{},
			},
		},
		Else: &ast.StmtElse{
			Stmt: &ast.StmtExpression{
				Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$f")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `if($a)$b;elseif($c){$d;}elseif($e);else$f;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{
					Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `if($a){$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintIfNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtIf{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtNop{},
	}
	n.Accept(p)

	expected := `if($a);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInlineHtml(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.StmtInlineHtml{
				Value: []byte("test"),
			},
		},
	}
	n.Accept(p)

	expected := `test`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintInterface(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtInterface{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Extends: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
		},
		Stmts: []ast.Vertex{
			&ast.StmtClassMethod{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Name:      &ast.Identifier{Value: []byte("foo")},
				Params:    []ast.Vertex{},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `interface Foo extends Bar,Baz{public function foo(){$a;}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintLabel(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtLabel{
		Name: &ast.Identifier{Value: []byte("FOO")},
	}
	n.Accept(p)

	expected := `FOO:`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespace(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtNamespace{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
	}
	n.Accept(p)

	expected := `namespace Foo;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNamespaceWithStmts(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtNamespace{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
	}
	n.Accept(p)

	expected := `namespace Foo{$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintNop(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtNop{}
	n.Accept(p)

	expected := `;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintPropertyList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtPropertyList{
		Modifiers: []ast.Vertex{
			&ast.Identifier{Value: []byte("public")},
			&ast.Identifier{Value: []byte("static")},
		},
		Type: &ast.Name{
			Parts: []ast.Vertex{
				&ast.NamePart{
					Value: []byte("Foo"),
				},
			},
		},
		Props: []ast.Vertex{
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
				Expr: &ast.ScalarString{Value: []byte("'a'")},
			},
			&ast.StmtProperty{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `public static Foo$a='a',$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintProperty(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtProperty{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	n.Accept(p)

	expected := `$a=1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintReturn(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtReturn{
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	n.Accept(p)

	expected := `return 1;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStaticVar(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtStaticVar{
		Var: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	n.Accept(p)

	expected := `$a=1`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStatic(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtStatic{
		Vars: []ast.Vertex{
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				},
			},
			&ast.StmtStaticVar{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$b")},
				},
			},
		},
	}
	n.Accept(p)

	expected := `static$a,$b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			}},
		},
	}
	n.Accept(p)

	expected := `{$a;$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtListNested(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
			&ast.StmtStmtList{
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
					&ast.StmtStmtList{
						Stmts: []ast.Vertex{
							&ast.StmtExpression{Expr: &ast.ExprVariable{
								Name: &ast.Identifier{Value: []byte("$c")},
							}},
						},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `{$a;{$b;{$c;}}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtSwitch(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtSwitch{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
		Cases: []ast.Vertex{
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'a'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$a")},
					}},
				},
			},
			&ast.StmtCase{
				Cond: &ast.ScalarString{Value: []byte("'b'")},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
	}
	n.Accept(p)

	expected := `switch($var){case'a':$a;case'b':$b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtThrow(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtThrow{
		Expr: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$var")},
		},
	}
	n.Accept(p)

	expected := `throw$var;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUseAlias(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTraitUseAlias{
		Trait:    &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Method:   &ast.Identifier{Value: []byte("a")},
		Modifier: &ast.Identifier{Value: []byte("public")},
		Alias:    &ast.Identifier{Value: []byte("b")},
	}
	n.Accept(p)

	expected := `Foo::a as public b;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUsePrecedence(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTraitUsePrecedence{
		Trait:  &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Method: &ast.Identifier{Value: []byte("a")},
		Insteadof: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
		},
	}
	n.Accept(p)

	expected := `Foo::a insteadof Bar,Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		},
	}
	n.Accept(p)

	expected := `use Foo,Bar;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTraitAdaptations(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTraitUse{
		Traits: []ast.Vertex{
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
			&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Bar")}}},
		},
		Adaptations: []ast.Vertex{
			&ast.StmtTraitUseAlias{
				Trait:  &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Method: &ast.Identifier{Value: []byte("a")},
				Alias:  &ast.Identifier{Value: []byte("b")},
			},
		},
	}
	n.Accept(p)

	expected := `use Foo,Bar{Foo::a as b;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintTrait(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTrait{
		Name: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Stmts: []ast.Vertex{
			&ast.StmtClassMethod{
				Modifiers: []ast.Vertex{&ast.Identifier{Value: []byte("public")}},
				Name:      &ast.Identifier{Value: []byte("foo")},
				Params:    []ast.Vertex{},
				Stmt: &ast.StmtStmtList{
					Stmts: []ast.Vertex{
						&ast.StmtExpression{Expr: &ast.ExprVariable{
							Name: &ast.Identifier{Value: []byte("$a")},
						}},
					},
				},
			},
		},
	}
	n.Accept(p)

	expected := `trait Foo{public function foo(){$a;}}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtTry(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtTry{
		Stmts: []ast.Vertex{
			&ast.StmtExpression{Expr: &ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			}},
		},
		Catches: []ast.Vertex{
			&ast.StmtCatch{
				Types: []ast.Vertex{
					&ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Exception")}}},
					&ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("RuntimeException")}}},
				},
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$e")},
				},
				Stmts: []ast.Vertex{
					&ast.StmtExpression{Expr: &ast.ExprVariable{
						Name: &ast.Identifier{Value: []byte("$b")},
					}},
				},
			},
		},
		Finally: &ast.StmtFinally{
			Stmts: []ast.Vertex{
				&ast.StmtNop{},
			},
		},
	}
	n.Accept(p)

	expected := `try{$a;}catch(Exception|\RuntimeException$e){$b;}finally{;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtUnset(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtUnset{
		Vars: []ast.Vertex{
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$a")},
			},
			&ast.ExprVariable{
				Name: &ast.Identifier{Value: []byte("$b")},
			},
		},
	}
	n.Accept(p)

	expected := `unset($a,$b);`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtUseList{
		Type: &ast.Identifier{Value: []byte("function")},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUse{
				Use: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			},
		},
	}
	n.Accept(p)

	expected := `use function Foo as Bar,Baz;`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintStmtGroupUse(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtGroupUseList{
		Type:   &ast.Identifier{Value: []byte("function")},
		Prefix: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Uses: []ast.Vertex{
			&ast.StmtUse{
				Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
				Alias: &ast.Identifier{Value: []byte("Bar")},
			},
			&ast.StmtUse{
				Use: &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Baz")}}},
			},
		},
	}
	n.Accept(p)

	expected := `use function Foo\{Foo as Bar,Baz};`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintUseDeclaration(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtUse{
		Type:  &ast.Identifier{Value: []byte("function")},
		Use:   &ast.Name{Parts: []ast.Vertex{&ast.NamePart{Value: []byte("Foo")}}},
		Alias: &ast.Identifier{Value: []byte("Bar")},
	}
	n.Accept(p)

	expected := `function Foo as Bar`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

func TestPrinterPrintWhileStmtList(t *testing.T) {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o).WithState(printer.PrinterStatePHP)
	n := &ast.StmtWhile{
		Cond: &ast.ExprVariable{
			Name: &ast.Identifier{Value: []byte("$a")},
		},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{
				&ast.StmtExpression{Expr: &ast.ExprVariable{
					Name: &ast.Identifier{Value: []byte("$a")},
				}},
			},
		},
	}
	n.Accept(p)

	expected := `while($a){$a;}`
	actual := o.String()

	if expected != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", expected, actual)
	}
}

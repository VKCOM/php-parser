package php8_test

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/tester"
)

func TestClassReadonlyModifier(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
readonly class Foo {
	public string $a;
}

final readonly class Boo {
	public string $a;
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtClass{
			Modifiers: []ast.Vertex{
				&ast.Identifier{
					Val: []byte("readonly"),
				},
			},
			Name: &ast.Identifier{
				Val: []byte("Foo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("string"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
					},
				},
			},
		},
		&ast.StmtClass{
			Modifiers: []ast.Vertex{
				&ast.Identifier{
					Val: []byte("final"),
				},
				&ast.Identifier{
					Val: []byte("readonly"),
				},
			},
			Name: &ast.Identifier{
				Val: []byte("Boo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Type: &ast.Name{
						Parts: []ast.Vertex{
							&ast.NamePart{
								Val: []byte("string"),
							},
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

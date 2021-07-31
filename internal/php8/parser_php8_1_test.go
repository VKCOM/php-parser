package php8_test

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/tester"
)

func TestReadonlyModifier(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
class Foo {
	readonly string $a;
	private readonly string $a;
	private string $a;
	private readonly $a = 100;

	public function __construct(
		readonly string $a,
		private readonly string $a,
		private string $a,
		private readonly $a = 100,
	) {}
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtClass{
			Name: &ast.Identifier{
				Val: []byte("Foo"),
			},
			Stmts: []ast.Vertex{
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("readonly"),
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
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("private"),
						},
						&ast.Identifier{
							Val: []byte("readonly"),
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
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("private"),
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
				&ast.StmtPropertyList{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("private"),
						},
						&ast.Identifier{
							Val: []byte("readonly"),
						},
					},
					Props: []ast.Vertex{
						&ast.StmtProperty{
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
							Expr: &ast.ScalarLnumber{
								Val: []byte("100"),
							},
						},
					},
				},
				&ast.StmtClassMethod{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							Val: []byte("public"),
						},
					},
					Name: &ast.Identifier{
						Val: []byte("__construct"),
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Modifiers: []ast.Vertex{
								&ast.Identifier{
									Val: []byte("readonly"),
								},
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("string"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Modifiers: []ast.Vertex{
								&ast.Identifier{
									Val: []byte("private"),
								},
								&ast.Identifier{
									Val: []byte("readonly"),
								},
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("string"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Modifiers: []ast.Vertex{
								&ast.Identifier{
									Val: []byte("private"),
								},
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("string"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
						},
						&ast.Parameter{
							Modifiers: []ast.Vertex{
								&ast.Identifier{
									Val: []byte("private"),
								},
								&ast.Identifier{
									Val: []byte("readonly"),
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$a"),
								},
							},
							DefaultValue: &ast.ScalarLnumber{
								Val: []byte("100"),
							},
						},
					},
					Stmt: &ast.StmtStmtList{
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestNeverType(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
function f(): never {}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtFunction{
			Name: &ast.Identifier{
				Val: []byte("f"),
			},
			ReturnType: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("never"),
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
	},
},`

	suite.Run()
}

func TestEnum(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
enum A {}
enum B implements Bar, Baz {
}
enum C: int implements Bar {}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtEnum{
			Name: &ast.Identifier{
				Val: []byte("A"),
			},
			Stmts: []ast.Vertex{},
		},
		&ast.StmtEnum{
			Name: &ast.Identifier{
				Val: []byte("B"),
			},
			Implements: []ast.Vertex{
				&ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Val: []byte("Bar"),
						},
					},
				},
				&ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Val: []byte("Baz"),
						},
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
		&ast.StmtEnum{
			Name: &ast.Identifier{
				Val: []byte("C"),
			},
			Type: &ast.Name{
				Parts: []ast.Vertex{
					&ast.NamePart{
						Val: []byte("int"),
					},
				},
			},
			Implements: []ast.Vertex{
				&ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Val: []byte("Bar"),
						},
					},
				},
			},
			Stmts: []ast.Vertex{},
		},
	},
},`

	suite.Run()
}

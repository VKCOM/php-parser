package php8_test

import (
	"testing"

	"github.com/z7zmey/php-parser/internal/tester"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/conf"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/parser"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
	"github.com/z7zmey/php-parser/pkg/version"
	"gotest.tools/assert"
)

func TestNullsafeExprMethodCall(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<? $a?->foo();`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    14,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    14,
				},
				Expr: &ast.ExprNullsafeMethodCall{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    13,
					},
					Var: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  3,
									EndPos:    5,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_OPEN_TAG,
										Value: []byte("<?"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  0,
											EndPos:    2,
										},
									},
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  2,
											EndPos:    3,
										},
									},
								},
							},
							Value: []byte("$a"),
						},
					},
					ObjectOperatorTkn: &token.Token{
						ID:    token.T_NULLSAFE_OBJECT_OPERATOR,
						Value: []byte("?->"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  5,
							EndPos:    8,
						},
					},
					Method: &ast.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    11,
						},
						IdentifierTkn: &token.Token{
							ID:    token.T_STRING,
							Value: []byte("foo"),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    11,
							},
						},
						Value: []byte("foo"),
					},
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    12,
						},
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  12,
							EndPos:    13,
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  13,
						EndPos:    14,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestNullsafePropertyFetch(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php $a?->foo;`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    15,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    15,
				},
				Expr: &ast.ExprNullsafePropertyFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    14,
					},
					Var: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    8,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    8,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  6,
									EndPos:    8,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_OPEN_TAG,
										Value: []byte("<?php"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  0,
											EndPos:    5,
										},
									},
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  5,
											EndPos:    6,
										},
									},
								},
							},
							Value: []byte("$a"),
						},
					},
					ObjectOperatorTkn: &token.Token{
						ID:    token.T_NULLSAFE_OBJECT_OPERATOR,
						Value: []byte("?->"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    11,
						},
					},
					Prop: &ast.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    14,
						},
						IdentifierTkn: &token.Token{
							ID:    token.T_STRING,
							Value: []byte("foo"),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  11,
								EndPos:    14,
							},
						},
						Value: []byte("foo"),
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  14,
						EndPos:    15,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestNullsafePropertyFetchInEncapsed(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php "$a?->foo";`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    17,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    17,
				},
				Expr: &ast.ScalarEncapsed{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    16,
					},
					OpenQuoteTkn: &token.Token{
						ID:    token.ID(34),
						Value: []byte("\""),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    7,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_OPEN_TAG,
								Value: []byte("<?php"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  0,
									EndPos:    5,
								},
							},
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    6,
								},
							},
						},
					},
					Parts: []ast.Vertex{
						&ast.ExprNullsafePropertyFetch{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    15,
							},
							Var: &ast.ExprVariable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    9,
								},
								Name: &ast.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  7,
										EndPos:    9,
									},
									IdentifierTkn: &token.Token{
										ID:    token.T_VARIABLE,
										Value: []byte("$a"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  7,
											EndPos:    9,
										},
									},
									Value: []byte("$a"),
								},
							},
							ObjectOperatorTkn: &token.Token{
								ID:    token.T_NULLSAFE_OBJECT_OPERATOR,
								Value: []byte("?->"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    12,
								},
							},
							Prop: &ast.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    15,
								},
								IdentifierTkn: &token.Token{
									ID:    token.T_STRING,
									Value: []byte("foo"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    15,
									},
								},
								Value: []byte("foo"),
							},
						},
					},
					CloseQuoteTkn: &token.Token{
						ID:    token.ID(34),
						Value: []byte("\""),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    16,
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  16,
						EndPos:    17,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestNullsafePropertyFetchForDereferencable(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php (f())?->foo;`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    18,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    18,
				},
				Expr: &ast.ExprNullsafePropertyFetch{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    17,
					},
					Var: &ast.ExprBrackets{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    11,
						},
						OpenParenthesisTkn: &token.Token{
							ID:    token.ID(40),
							Value: []byte("("),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  6,
								EndPos:    7,
							},
							FreeFloating: []*token.Token{
								{
									ID:    token.T_OPEN_TAG,
									Value: []byte("<?php"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  0,
										EndPos:    5,
									},
								},
								{
									ID:    token.T_WHITESPACE,
									Value: []byte(" "),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  5,
										EndPos:    6,
									},
								},
							},
						},
						Expr: &ast.ExprFunctionCall{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  7,
								EndPos:    10,
							},
							Function: &ast.Name{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  7,
									EndPos:    8,
								},
								Parts: []ast.Vertex{
									&ast.NamePart{
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  7,
											EndPos:    8,
										},
										StringTkn: &token.Token{
											ID:    token.T_STRING,
											Value: []byte("f"),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  7,
												EndPos:    8,
											},
										},
										Value: []byte("f"),
									},
								},
							},
							OpenParenthesisTkn: &token.Token{
								ID:    token.ID(40),
								Value: []byte("("),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    9,
								},
							},
							CloseParenthesisTkn: &token.Token{
								ID:    token.ID(41),
								Value: []byte(")"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  9,
									EndPos:    10,
								},
							},
						},
						CloseParenthesisTkn: &token.Token{
							ID:    token.ID(41),
							Value: []byte(")"),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    11,
							},
						},
					},
					ObjectOperatorTkn: &token.Token{
						ID:    token.T_NULLSAFE_OBJECT_OPERATOR,
						Value: []byte("?->"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  11,
							EndPos:    14,
						},
					},
					Prop: &ast.Identifier{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  14,
							EndPos:    17,
						},
						IdentifierTkn: &token.Token{
							ID:    token.T_STRING,
							Value: []byte("foo"),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  14,
								EndPos:    17,
							},
						},
						Value: []byte("foo"),
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  17,
						EndPos:    18,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestExprCast(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
(array)$a;
(boolean)$a;
(bool)$a;
(double)$a;
(float)$a;
(integer)$a;
(int)$a;
(object)$a;
(string)$a;
(binary)$a;
`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   11,
			StartPos:  6,
			EndPos:    120,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    16,
				},
				Expr: &ast.ExprCastArray{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    15,
					},
					CastTkn: &token.Token{
						ID:    token.T_ARRAY_CAST,
						Value: []byte("(array)"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    13,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_OPEN_TAG,
								Value: []byte("<?php"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  0,
									EndPos:    5,
								},
							},
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  5,
									EndPos:    6,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    15,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    15,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  13,
									EndPos:    15,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  15,
						EndPos:    16,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 3,
					EndLine:   3,
					StartPos:  17,
					EndPos:    29,
				},
				Expr: &ast.ExprCastBool{
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  17,
						EndPos:    28,
					},
					CastTkn: &token.Token{
						ID:    token.T_BOOL_CAST,
						Value: []byte("(boolean)"),
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  17,
							EndPos:    26,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  16,
									EndPos:    17,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 3,
							EndLine:   3,
							StartPos:  26,
							EndPos:    28,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  26,
								EndPos:    28,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  26,
									EndPos:    28,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 3,
						EndLine:   3,
						StartPos:  28,
						EndPos:    29,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 4,
					EndLine:   4,
					StartPos:  30,
					EndPos:    39,
				},
				Expr: &ast.ExprCastBool{
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  30,
						EndPos:    38,
					},
					CastTkn: &token.Token{
						ID:    token.T_BOOL_CAST,
						Value: []byte("(bool)"),
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  30,
							EndPos:    36,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  29,
									EndPos:    30,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  36,
							EndPos:    38,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  36,
								EndPos:    38,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  36,
									EndPos:    38,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  38,
						EndPos:    39,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 5,
					EndLine:   5,
					StartPos:  40,
					EndPos:    51,
				},
				Expr: &ast.ExprCastDouble{
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  40,
						EndPos:    50,
					},
					CastTkn: &token.Token{
						ID:    token.T_DOUBLE_CAST,
						Value: []byte("(double)"),
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  40,
							EndPos:    48,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  39,
									EndPos:    40,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  48,
							EndPos:    50,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 5,
								EndLine:   5,
								StartPos:  48,
								EndPos:    50,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  48,
									EndPos:    50,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  50,
						EndPos:    51,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 6,
					EndLine:   6,
					StartPos:  52,
					EndPos:    62,
				},
				Expr: &ast.ExprCastDouble{
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  52,
						EndPos:    61,
					},
					CastTkn: &token.Token{
						ID:    token.T_DOUBLE_CAST,
						Value: []byte("(float)"),
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  52,
							EndPos:    59,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 5,
									EndLine:   5,
									StartPos:  51,
									EndPos:    52,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 6,
							EndLine:   6,
							StartPos:  59,
							EndPos:    61,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 6,
								EndLine:   6,
								StartPos:  59,
								EndPos:    61,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  59,
									EndPos:    61,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 6,
						EndLine:   6,
						StartPos:  61,
						EndPos:    62,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 7,
					EndLine:   7,
					StartPos:  63,
					EndPos:    75,
				},
				Expr: &ast.ExprCastInt{
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  63,
						EndPos:    74,
					},
					CastTkn: &token.Token{
						ID:    token.T_INT_CAST,
						Value: []byte("(integer)"),
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  63,
							EndPos:    72,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 6,
									EndLine:   6,
									StartPos:  62,
									EndPos:    63,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 7,
							EndLine:   7,
							StartPos:  72,
							EndPos:    74,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 7,
								EndLine:   7,
								StartPos:  72,
								EndPos:    74,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  72,
									EndPos:    74,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 7,
						EndLine:   7,
						StartPos:  74,
						EndPos:    75,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 8,
					EndLine:   8,
					StartPos:  76,
					EndPos:    84,
				},
				Expr: &ast.ExprCastInt{
					Position: &position.Position{
						StartLine: 8,
						EndLine:   8,
						StartPos:  76,
						EndPos:    83,
					},
					CastTkn: &token.Token{
						ID:    token.T_INT_CAST,
						Value: []byte("(int)"),
						Position: &position.Position{
							StartLine: 8,
							EndLine:   8,
							StartPos:  76,
							EndPos:    81,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 7,
									EndLine:   7,
									StartPos:  75,
									EndPos:    76,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 8,
							EndLine:   8,
							StartPos:  81,
							EndPos:    83,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 8,
								EndLine:   8,
								StartPos:  81,
								EndPos:    83,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 8,
									EndLine:   8,
									StartPos:  81,
									EndPos:    83,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 8,
						EndLine:   8,
						StartPos:  83,
						EndPos:    84,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 9,
					EndLine:   9,
					StartPos:  85,
					EndPos:    96,
				},
				Expr: &ast.ExprCastObject{
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  85,
						EndPos:    95,
					},
					CastTkn: &token.Token{
						ID:    token.T_OBJECT_CAST,
						Value: []byte("(object)"),
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  85,
							EndPos:    93,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 8,
									EndLine:   8,
									StartPos:  84,
									EndPos:    85,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 9,
							EndLine:   9,
							StartPos:  93,
							EndPos:    95,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 9,
								EndLine:   9,
								StartPos:  93,
								EndPos:    95,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  93,
									EndPos:    95,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 9,
						EndLine:   9,
						StartPos:  95,
						EndPos:    96,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 10,
					EndLine:   10,
					StartPos:  97,
					EndPos:    108,
				},
				Expr: &ast.ExprCastString{
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  97,
						EndPos:    107,
					},
					CastTkn: &token.Token{
						ID:    token.T_STRING_CAST,
						Value: []byte("(string)"),
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  97,
							EndPos:    105,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 9,
									EndLine:   9,
									StartPos:  96,
									EndPos:    97,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 10,
							EndLine:   10,
							StartPos:  105,
							EndPos:    107,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 10,
								EndLine:   10,
								StartPos:  105,
								EndPos:    107,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  105,
									EndPos:    107,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 10,
						EndLine:   10,
						StartPos:  107,
						EndPos:    108,
					},
				},
			},
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 11,
					EndLine:   11,
					StartPos:  109,
					EndPos:    120,
				},
				Expr: &ast.ExprCastString{
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  109,
						EndPos:    119,
					},
					CastTkn: &token.Token{
						ID:    token.T_STRING_CAST,
						Value: []byte("(binary)"),
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  109,
							EndPos:    117,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 10,
									EndLine:   10,
									StartPos:  108,
									EndPos:    109,
								},
							},
						},
					},
					Expr: &ast.ExprVariable{
						Position: &position.Position{
							StartLine: 11,
							EndLine:   11,
							StartPos:  117,
							EndPos:    119,
						},
						Name: &ast.Identifier{
							Position: &position.Position{
								StartLine: 11,
								EndLine:   11,
								StartPos:  117,
								EndPos:    119,
							},
							IdentifierTkn: &token.Token{
								ID:    token.T_VARIABLE,
								Value: []byte("$a"),
								Position: &position.Position{
									StartLine: 11,
									EndLine:   11,
									StartPos:  117,
									EndPos:    119,
								},
							},
							Value: []byte("$a"),
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  119,
						EndPos:    120,
					},
				},
			},
		},
		EndTkn: &token.Token{
			FreeFloating: []*token.Token{
				{
					ID:    token.T_WHITESPACE,
					Value: []byte("\n"),
					Position: &position.Position{
						StartLine: 11,
						EndLine:   11,
						StartPos:  120,
						EndPos:    121,
					},
				},
			},
		},
	}

	suite.Run()
}

func TestRealCastParserError(t *testing.T) {
	suite := tester.NewParserErrorTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php 
(real)$a;   // Error
(float)$a;  // Ok
(double)$a; // Ok
(int)$a;    // Ok
`

	suite.Expected = []*errors.Error{
		{
			Msg: "The (real) cast has been removed, use (float) instead",
			Pos: position.NewPosition(2, 2, 7, 13),
		},
	}

	suite.Run()
}

func TestUnsetCastParserError(t *testing.T) {
	suite := tester.NewParserErrorTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php 
(unset)$a;   // Error
(float)$a;  // Ok
(double)$a; // Ok
(int)$a;    // Ok
`

	suite.Expected = []*errors.Error{
		{
			Msg: "The (unset) cast is no longer supported",
			Pos: position.NewPosition(2, 2, 7, 14),
		},
	}

	suite.Run()
}

func TestFunctionCallWithNamedArgument(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php f(name: $a);`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    18,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    18,
				},
				Expr: &ast.ExprFunctionCall{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    17,
					},
					Function: &ast.Name{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    7,
						},
						Parts: []ast.Vertex{
							&ast.NamePart{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  6,
									EndPos:    7,
								},
								StringTkn: &token.Token{
									ID:    token.T_STRING,
									Value: []byte("f"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  6,
										EndPos:    7,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_OPEN_TAG,
											Value: []byte("<?php"),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  0,
												EndPos:    5,
											},
										},
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  5,
												EndPos:    6,
											},
										},
									},
								},
								Value: []byte("f"),
							},
						},
					},
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    8,
						},
					},
					Args: []ast.Vertex{
						&ast.Argument{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    16,
							},
							Name: &ast.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    12,
								},
								IdentifierTkn: &token.Token{
									ID:    token.T_STRING,
									Value: []byte("name"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    12,
									},
								},
								Value: []byte("name"),
							},
							ColonTkn: &token.Token{
								ID:    token.ID(58),
								Value: []byte(":"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    13,
								},
							},
							Expr: &ast.ExprVariable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    16,
								},
								Name: &ast.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  14,
										EndPos:    16,
									},
									IdentifierTkn: &token.Token{
										ID:    token.T_VARIABLE,
										Value: []byte("$a"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  14,
											EndPos:    16,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte(" "),
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  13,
													EndPos:    14,
												},
											},
										},
									},
									Value: []byte("$a"),
								},
							},
						},
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  16,
							EndPos:    17,
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  17,
						EndPos:    18,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestFunctionCallWithNamedArgumentAndNonNamedArgument(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php f($b, name: $a);`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    22,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    22,
				},
				Expr: &ast.ExprFunctionCall{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    21,
					},
					Function: &ast.Name{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  6,
							EndPos:    7,
						},
						Parts: []ast.Vertex{
							&ast.NamePart{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  6,
									EndPos:    7,
								},
								StringTkn: &token.Token{
									ID:    token.T_STRING,
									Value: []byte("f"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  6,
										EndPos:    7,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_OPEN_TAG,
											Value: []byte("<?php"),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  0,
												EndPos:    5,
											},
										},
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  5,
												EndPos:    6,
											},
										},
									},
								},
								Value: []byte("f"),
							},
						},
					},
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  7,
							EndPos:    8,
						},
					},
					Args: []ast.Vertex{
						&ast.Argument{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    10,
							},
							Expr: &ast.ExprVariable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  8,
									EndPos:    10,
								},
								Name: &ast.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  8,
										EndPos:    10,
									},
									IdentifierTkn: &token.Token{
										ID:    token.T_VARIABLE,
										Value: []byte("$b"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  8,
											EndPos:    10,
										},
									},
									Value: []byte("$b"),
								},
							},
						},
						&ast.Argument{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  12,
								EndPos:    20,
							},
							Name: &ast.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  12,
									EndPos:    16,
								},
								IdentifierTkn: &token.Token{
									ID:    token.T_STRING,
									Value: []byte("name"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  12,
										EndPos:    16,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  11,
												EndPos:    12,
											},
										},
									},
								},
								Value: []byte("name"),
							},
							ColonTkn: &token.Token{
								ID:    token.ID(58),
								Value: []byte(":"),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  16,
									EndPos:    17,
								},
							},
							Expr: &ast.ExprVariable{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  18,
									EndPos:    20,
								},
								Name: &ast.Identifier{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  18,
										EndPos:    20,
									},
									IdentifierTkn: &token.Token{
										ID:    token.T_VARIABLE,
										Value: []byte("$a"),
										Position: &position.Position{
											StartLine: 1,
											EndLine:   1,
											StartPos:  18,
											EndPos:    20,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte(" "),
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  17,
													EndPos:    18,
												},
											},
										},
									},
									Value: []byte("$a"),
								},
							},
						},
					},
					SeparatorTkns: []*token.Token{
						{
							ID:    token.ID(44),
							Value: []byte(","),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    11,
							},
						},
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  20,
							EndPos:    21,
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  21,
						EndPos:    22,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestArgumentsCount(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
foo($a, name: $b, ...$b);
$foo(a: $a, $c, ...$b);
$foo->bar(some: $a, $b, ...$c);
foo::bar($a, b: $b, ...$c);
$foo::bar($b, c: $a, ...$b);
new foo(a: $a, $c, ...$b);
/** anonymous class */
new class (name: $a, $b, ...$c) {};
`

	config := conf.Config{
		Version: &version.Version{Major: 8, Minor: 0},
	}

	root, err := parser.Parse([]byte(suite.Code), config)
	if err != nil {
		t.Fatalf("Error parse: %v", err)
	}

	stmts := root.(*ast.Root).Stmts

	for _, stmt := range stmts {
		stmtExpr, ok := stmt.(*ast.StmtExpression)
		if !ok {
			continue
		}

		switch n := stmtExpr.Expr.(type) {
		case *ast.ExprFunctionCall:
			assert.Assert(t, len(n.Args) == 3)
		case *ast.ExprMethodCall:
			assert.Assert(t, len(n.Args) == 3)
		case *ast.ExprStaticCall:
			assert.Assert(t, len(n.Args) == 3)
		case *ast.ExprNew:
			if class, ok := n.Class.(*ast.StmtClass); ok {
				assert.Assert(t, len(class.Args) == 3)
				continue
			}

			assert.Assert(t, len(n.Args) == 3)
		}
	}
}

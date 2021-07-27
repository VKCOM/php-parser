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

func TestCurlyBracesAccessParserError(t *testing.T) {
	suite := tester.NewParserErrorTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php 
$a[10]; // Ok
$a{10}; // Error

${a}; // Ok
Foo::{a}(); // Ok

new $a[10]; // Ok
new $a{10}; // Error

"hello"[1]; // Ok
"hello"{1}; // Error
`

	suite.Expected = []*errors.Error{
		{
			Msg: "Array and string offset access syntax with curly braces is no longer supported",
			Pos: position.NewPosition(3, 3, 26, 27),
		},
		{
			Msg: "Array and string offset access syntax with curly braces is no longer supported",
			Pos: position.NewPosition(9, 9, 97, 98),
		},
		{
			Msg: "Array and string offset access syntax with curly braces is no longer supported",
			Pos: position.NewPosition(12, 12, 137, 138),
		},
	}

	suite.Run()
}

func TestMatchExpr(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
match (10) {
  10 => 100
};
`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   4,
			StartPos:  6,
			EndPos:    33,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   4,
					StartPos:  6,
					EndPos:    33,
				},
				Expr: &ast.ExprMatch{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   4,
						StartPos:  6,
						EndPos:    32,
					},
					MatchTkn: &token.Token{
						ID:    token.T_MATCH,
						Value: []byte("match"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    11,
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
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  12,
							EndPos:    13,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  11,
									EndPos:    12,
								},
							},
						},
					},
					Expr: &ast.ScalarLnumber{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    15,
						},
						NumberTkn: &token.Token{
							ID:    token.T_LNUMBER,
							Value: []byte("10"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    15,
							},
						},
						Value: []byte("10"),
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  15,
							EndPos:    16,
						},
					},
					OpenCurlyBracketTkn: &token.Token{
						ID:    token.ID(123),
						Value: []byte("{"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  17,
							EndPos:    18,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  16,
									EndPos:    17,
								},
							},
						},
					},
					Arms: []ast.Vertex{
						&ast.MatchArm{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  21,
								EndPos:    30,
							},
							Exprs: []ast.Vertex{
								&ast.ScalarLnumber{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  21,
										EndPos:    23,
									},
									NumberTkn: &token.Token{
										ID:    token.T_LNUMBER,
										Value: []byte("10"),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  21,
											EndPos:    23,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte("\n  "),
												Position: &position.Position{
													StartLine: 2,
													EndLine:   3,
													StartPos:  18,
													EndPos:    21,
												},
											},
										},
									},
									Value: []byte("10"),
								},
							},
							DoubleArrowTkn: &token.Token{
								ID:    token.T_DOUBLE_ARROW,
								Value: []byte("=>"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  24,
									EndPos:    26,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  23,
											EndPos:    24,
										},
									},
								},
							},
							ReturnExpr: &ast.ScalarLnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  27,
									EndPos:    30,
								},
								NumberTkn: &token.Token{
									ID:    token.T_LNUMBER,
									Value: []byte("100"),
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  27,
										EndPos:    30,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  26,
												EndPos:    27,
											},
										},
									},
								},
								Value: []byte("100"),
							},
						},
					},
					CloseCurlyBracketTkn: &token.Token{
						ID:    token.ID(125),
						Value: []byte("}"),
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  31,
							EndPos:    32,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  30,
									EndPos:    31,
								},
							},
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  32,
						EndPos:    33,
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
						StartLine: 4,
						EndLine:   4,
						StartPos:  33,
						EndPos:    34,
					},
				},
			},
		},
	}

	suite.Run()
}

func TestMatchExprWithList(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
match (10) {
  10, 11 => 100
};
`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   4,
			StartPos:  6,
			EndPos:    37,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   4,
					StartPos:  6,
					EndPos:    37,
				},
				Expr: &ast.ExprMatch{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   4,
						StartPos:  6,
						EndPos:    36,
					},
					MatchTkn: &token.Token{
						ID:    token.T_MATCH,
						Value: []byte("match"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    11,
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
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  12,
							EndPos:    13,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  11,
									EndPos:    12,
								},
							},
						},
					},
					Expr: &ast.ScalarLnumber{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    15,
						},
						NumberTkn: &token.Token{
							ID:    token.T_LNUMBER,
							Value: []byte("10"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    15,
							},
						},
						Value: []byte("10"),
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  15,
							EndPos:    16,
						},
					},
					OpenCurlyBracketTkn: &token.Token{
						ID:    token.ID(123),
						Value: []byte("{"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  17,
							EndPos:    18,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  16,
									EndPos:    17,
								},
							},
						},
					},
					Arms: []ast.Vertex{
						&ast.MatchArm{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  21,
								EndPos:    34,
							},
							Exprs: []ast.Vertex{
								&ast.ScalarLnumber{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  21,
										EndPos:    23,
									},
									NumberTkn: &token.Token{
										ID:    token.T_LNUMBER,
										Value: []byte("10"),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  21,
											EndPos:    23,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte("\n  "),
												Position: &position.Position{
													StartLine: 2,
													EndLine:   3,
													StartPos:  18,
													EndPos:    21,
												},
											},
										},
									},
									Value: []byte("10"),
								},
								&ast.ScalarLnumber{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  25,
										EndPos:    27,
									},
									NumberTkn: &token.Token{
										ID:    token.T_LNUMBER,
										Value: []byte("11"),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  25,
											EndPos:    27,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte(" "),
												Position: &position.Position{
													StartLine: 3,
													EndLine:   3,
													StartPos:  24,
													EndPos:    25,
												},
											},
										},
									},
									Value: []byte("11"),
								},
							},
							SeparatorTkns: []*token.Token{
								{
									ID:    token.ID(44),
									Value: []byte(","),
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  23,
										EndPos:    24,
									},
								},
							},
							DoubleArrowTkn: &token.Token{
								ID:    token.T_DOUBLE_ARROW,
								Value: []byte("=>"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  28,
									EndPos:    30,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  27,
											EndPos:    28,
										},
									},
								},
							},
							ReturnExpr: &ast.ScalarLnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  31,
									EndPos:    34,
								},
								NumberTkn: &token.Token{
									ID:    token.T_LNUMBER,
									Value: []byte("100"),
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  31,
										EndPos:    34,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  30,
												EndPos:    31,
											},
										},
									},
								},
								Value: []byte("100"),
							},
						},
					},
					CloseCurlyBracketTkn: &token.Token{
						ID:    token.ID(125),
						Value: []byte("}"),
						Position: &position.Position{
							StartLine: 4,
							EndLine:   4,
							StartPos:  35,
							EndPos:    36,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  34,
									EndPos:    35,
								},
							},
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 4,
						EndLine:   4,
						StartPos:  36,
						EndPos:    37,
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
						StartLine: 4,
						EndLine:   4,
						StartPos:  37,
						EndPos:    38,
					},
				},
			},
		},
	}

	suite.Run()
}

func TestMatchExprWithDefault(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
match (10) {
  10 => 100,
  default => 101,
};
`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   5,
			StartPos:  6,
			EndPos:    52,
		},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   5,
					StartPos:  6,
					EndPos:    52,
				},
				Expr: &ast.ExprMatch{
					Position: &position.Position{
						StartLine: 2,
						EndLine:   5,
						StartPos:  6,
						EndPos:    51,
					},
					MatchTkn: &token.Token{
						ID:    token.T_MATCH,
						Value: []byte("match"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  6,
							EndPos:    11,
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
					OpenParenthesisTkn: &token.Token{
						ID:    token.ID(40),
						Value: []byte("("),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  12,
							EndPos:    13,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  11,
									EndPos:    12,
								},
							},
						},
					},
					Expr: &ast.ScalarLnumber{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    15,
						},
						NumberTkn: &token.Token{
							ID:    token.T_LNUMBER,
							Value: []byte("10"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    15,
							},
						},
						Value: []byte("10"),
					},
					CloseParenthesisTkn: &token.Token{
						ID:    token.ID(41),
						Value: []byte(")"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  15,
							EndPos:    16,
						},
					},
					OpenCurlyBracketTkn: &token.Token{
						ID:    token.ID(123),
						Value: []byte("{"),
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  17,
							EndPos:    18,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  16,
									EndPos:    17,
								},
							},
						},
					},
					Arms: []ast.Vertex{
						&ast.MatchArm{
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  21,
								EndPos:    30,
							},
							Exprs: []ast.Vertex{
								&ast.ScalarLnumber{
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  21,
										EndPos:    23,
									},
									NumberTkn: &token.Token{
										ID:    token.T_LNUMBER,
										Value: []byte("10"),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  21,
											EndPos:    23,
										},
										FreeFloating: []*token.Token{
											{
												ID:    token.T_WHITESPACE,
												Value: []byte("\n  "),
												Position: &position.Position{
													StartLine: 2,
													EndLine:   3,
													StartPos:  18,
													EndPos:    21,
												},
											},
										},
									},
									Value: []byte("10"),
								},
							},
							DoubleArrowTkn: &token.Token{
								ID:    token.T_DOUBLE_ARROW,
								Value: []byte("=>"),
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  24,
									EndPos:    26,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   3,
											StartPos:  23,
											EndPos:    24,
										},
									},
								},
							},
							ReturnExpr: &ast.ScalarLnumber{
								Position: &position.Position{
									StartLine: 3,
									EndLine:   3,
									StartPos:  27,
									EndPos:    30,
								},
								NumberTkn: &token.Token{
									ID:    token.T_LNUMBER,
									Value: []byte("100"),
									Position: &position.Position{
										StartLine: 3,
										EndLine:   3,
										StartPos:  27,
										EndPos:    30,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 3,
												EndLine:   3,
												StartPos:  26,
												EndPos:    27,
											},
										},
									},
								},
								Value: []byte("100"),
							},
						},
						&ast.MatchArm{
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  34,
								EndPos:    48,
							},
							DefaultTkn: &token.Token{
								ID:    token.T_DEFAULT,
								Value: []byte("default"),
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  34,
									EndPos:    41,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_WHITESPACE,
										Value: []byte("\n  "),
										Position: &position.Position{
											StartLine: 3,
											EndLine:   4,
											StartPos:  31,
											EndPos:    34,
										},
									},
								},
							},
							DoubleArrowTkn: &token.Token{
								ID:    token.T_DOUBLE_ARROW,
								Value: []byte("=>"),
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  42,
									EndPos:    44,
								},
								FreeFloating: []*token.Token{
									{
										ID:    token.T_WHITESPACE,
										Value: []byte(" "),
										Position: &position.Position{
											StartLine: 4,
											EndLine:   4,
											StartPos:  41,
											EndPos:    42,
										},
									},
								},
							},
							ReturnExpr: &ast.ScalarLnumber{
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  45,
									EndPos:    48,
								},
								NumberTkn: &token.Token{
									ID:    token.T_LNUMBER,
									Value: []byte("101"),
									Position: &position.Position{
										StartLine: 4,
										EndLine:   4,
										StartPos:  45,
										EndPos:    48,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 4,
												EndLine:   4,
												StartPos:  44,
												EndPos:    45,
											},
										},
									},
								},
								Value: []byte("101"),
							},
						},
					},
					SeparatorTkns: []*token.Token{
						{
							ID:    token.ID(44),
							Value: []byte(","),
							Position: &position.Position{
								StartLine: 3,
								EndLine:   3,
								StartPos:  30,
								EndPos:    31,
							},
						},
						{
							ID:    token.ID(44),
							Value: []byte(","),
							Position: &position.Position{
								StartLine: 4,
								EndLine:   4,
								StartPos:  48,
								EndPos:    49,
							},
						},
					},
					CloseCurlyBracketTkn: &token.Token{
						ID:    token.ID(125),
						Value: []byte("}"),
						Position: &position.Position{
							StartLine: 5,
							EndLine:   5,
							StartPos:  50,
							EndPos:    51,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte("\n"),
								Position: &position.Position{
									StartLine: 4,
									EndLine:   4,
									StartPos:  49,
									EndPos:    50,
								},
							},
						},
					},
				},
				SemiColonTkn: &token.Token{
					ID:    token.ID(59),
					Value: []byte(";"),
					Position: &position.Position{
						StartLine: 5,
						EndLine:   5,
						StartPos:  51,
						EndPos:    52,
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
						StartLine: 5,
						EndLine:   5,
						StartPos:  52,
						EndPos:    53,
					},
				},
			},
		},
	}

	suite.Run()
}

func TestStaticTypeParserError(t *testing.T) {
	suite := tester.NewParserErrorTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php 
function f(static|\Boo $a) {}
function f(static $a) {}
function f(): static|\Boo {} // ok
function f(): static {} // ok

class Foo {
	public static|int $a;
	public static $a;
}
`

	suite.Expected = []*errors.Error{
		{
			Msg: "syntax error",
			Pos: position.NewPosition(2, 2, 18, 24),
		},
		{
			Msg: "syntax error",
			Pos: position.NewPosition(3, 3, 48, 54),
		},
		{
			Msg: "syntax error",
			Pos: position.NewPosition(8, 8, 154, 155),
		},
		{
			Msg: "syntax error",
			Pos: position.NewPosition(10, 10, 182, 183),
		},
	}

	suite.Run()
}

func TestUnionTypeParserError(t *testing.T) {
	suite := tester.NewParserErrorTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php 
function f(?int|string $a) {}
function f(?(int|string) $a) {}
function f(int|string|null $a) {} // ok
`

	suite.Expected = []*errors.Error{
		{
			Msg: "syntax error",
			Pos: position.NewPosition(2, 2, 22, 23),
		},
		{
			Msg: "syntax error",
			Pos: position.NewPosition(3, 3, 49, 50),
		},
		{
			Msg: "syntax error",
			Pos: position.NewPosition(3, 3, 62, 64),
		},
	}

	suite.Run()
}

func TestUnionTypes(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php function f(int|string $a) {}`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  6,
			EndPos:    34,
		},
		Stmts: []ast.Vertex{
			&ast.StmtFunction{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  6,
					EndPos:    34,
				},
				FunctionTkn: &token.Token{
					ID:    token.T_FUNCTION,
					Value: []byte("function"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  6,
						EndPos:    14,
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
				Name: &ast.Identifier{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  15,
						EndPos:    16,
					},
					IdentifierTkn: &token.Token{
						ID:    token.T_STRING,
						Value: []byte("f"),
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  15,
							EndPos:    16,
						},
						FreeFloating: []*token.Token{
							{
								ID:    token.T_WHITESPACE,
								Value: []byte(" "),
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  14,
									EndPos:    15,
								},
							},
						},
					},
					Value: []byte("f"),
				},
				OpenParenthesisTkn: &token.Token{
					ID:    token.ID(40),
					Value: []byte("("),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  16,
						EndPos:    17,
					},
				},
				Params: []ast.Vertex{
					&ast.Parameter{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  17,
							EndPos:    30,
						},
						Type: &ast.Union{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  17,
								EndPos:    27,
							},
							Types: []ast.Vertex{
								&ast.Name{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  17,
										EndPos:    20,
									},
									Parts: []ast.Vertex{
										&ast.NamePart{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  17,
												EndPos:    20,
											},
											StringTkn: &token.Token{
												ID:    token.T_STRING,
												Value: []byte("int"),
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  17,
													EndPos:    20,
												},
											},
											Value: []byte("int"),
										},
									},
								},
								&ast.Name{
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  21,
										EndPos:    27,
									},
									Parts: []ast.Vertex{
										&ast.NamePart{
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  21,
												EndPos:    27,
											},
											StringTkn: &token.Token{
												ID:    token.T_STRING,
												Value: []byte("string"),
												Position: &position.Position{
													StartLine: 1,
													EndLine:   1,
													StartPos:  21,
													EndPos:    27,
												},
											},
											Value: []byte("string"),
										},
									},
								},
							},
							SeparatorTkns: []*token.Token{
								{
									ID:    token.ID(124),
									Value: []byte("|"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  20,
										EndPos:    21,
									},
								},
							},
						},
						Var: &ast.ExprVariable{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  28,
								EndPos:    30,
							},
							Name: &ast.Identifier{
								Position: &position.Position{
									StartLine: 1,
									EndLine:   1,
									StartPos:  28,
									EndPos:    30,
								},
								IdentifierTkn: &token.Token{
									ID:    token.T_VARIABLE,
									Value: []byte("$a"),
									Position: &position.Position{
										StartLine: 1,
										EndLine:   1,
										StartPos:  28,
										EndPos:    30,
									},
									FreeFloating: []*token.Token{
										{
											ID:    token.T_WHITESPACE,
											Value: []byte(" "),
											Position: &position.Position{
												StartLine: 1,
												EndLine:   1,
												StartPos:  27,
												EndPos:    28,
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
						StartPos:  30,
						EndPos:    31,
					},
				},
				OpenCurlyBracketTkn: &token.Token{
					ID:    token.ID(123),
					Value: []byte("{"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  32,
						EndPos:    33,
					},
					FreeFloating: []*token.Token{
						{
							ID:    token.T_WHITESPACE,
							Value: []byte(" "),
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  31,
								EndPos:    32,
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
				CloseCurlyBracketTkn: &token.Token{
					ID:    token.ID(125),
					Value: []byte("}"),
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  33,
						EndPos:    34,
					},
				},
			},
		},
		EndTkn: &token.Token{},
	}

	suite.Run()
}

func TestCatchExprWithoutVariable(t *testing.T) {
	suite := tester.NewParserTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
try {} catch (Exception) {}
`

	suite.Expected = &ast.Root{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  6,
			EndPos:    33,
		},
		Stmts: []ast.Vertex{
			&ast.StmtTry{
				Position: &position.Position{
					StartLine: 2,
					EndLine:   2,
					StartPos:  6,
					EndPos:    33,
				},
				TryTkn: &token.Token{
					ID:    token.T_TRY,
					Value: []byte("try"),
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  6,
						EndPos:    9,
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
				OpenCurlyBracketTkn: &token.Token{
					ID:    token.ID(123),
					Value: []byte("{"),
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  10,
						EndPos:    11,
					},
					FreeFloating: []*token.Token{
						{
							ID:    token.T_WHITESPACE,
							Value: []byte(" "),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  9,
								EndPos:    10,
							},
						},
					},
				},
				Stmts: []ast.Vertex{},
				CloseCurlyBracketTkn: &token.Token{
					ID:    token.ID(125),
					Value: []byte("}"),
					Position: &position.Position{
						StartLine: 2,
						EndLine:   2,
						StartPos:  11,
						EndPos:    12,
					},
				},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Position: &position.Position{
							StartLine: 2,
							EndLine:   2,
							StartPos:  13,
							EndPos:    33,
						},
						CatchTkn: &token.Token{
							ID:    token.T_CATCH,
							Value: []byte("catch"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  13,
								EndPos:    18,
							},
							FreeFloating: []*token.Token{
								{
									ID:    token.T_WHITESPACE,
									Value: []byte(" "),
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  12,
										EndPos:    13,
									},
								},
							},
						},
						OpenParenthesisTkn: &token.Token{
							ID:    token.ID(40),
							Value: []byte("("),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  19,
								EndPos:    20,
							},
							FreeFloating: []*token.Token{
								{
									ID:    token.T_WHITESPACE,
									Value: []byte(" "),
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  18,
										EndPos:    19,
									},
								},
							},
						},
						Types: []ast.Vertex{
							&ast.Name{
								Position: &position.Position{
									StartLine: 2,
									EndLine:   2,
									StartPos:  20,
									EndPos:    29,
								},
								Parts: []ast.Vertex{
									&ast.NamePart{
										Position: &position.Position{
											StartLine: 2,
											EndLine:   2,
											StartPos:  20,
											EndPos:    29,
										},
										StringTkn: &token.Token{
											ID:    token.T_STRING,
											Value: []byte("Exception"),
											Position: &position.Position{
												StartLine: 2,
												EndLine:   2,
												StartPos:  20,
												EndPos:    29,
											},
										},
										Value: []byte("Exception"),
									},
								},
							},
						},
						CloseParenthesisTkn: &token.Token{
							ID:    token.ID(41),
							Value: []byte(")"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  29,
								EndPos:    30,
							},
						},
						OpenCurlyBracketTkn: &token.Token{
							ID:    token.ID(123),
							Value: []byte("{"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  31,
								EndPos:    32,
							},
							FreeFloating: []*token.Token{
								{
									ID:    token.T_WHITESPACE,
									Value: []byte(" "),
									Position: &position.Position{
										StartLine: 2,
										EndLine:   2,
										StartPos:  30,
										EndPos:    31,
									},
								},
							},
						},
						Stmts: []ast.Vertex{},
						CloseCurlyBracketTkn: &token.Token{
							ID:    token.ID(125),
							Value: []byte("}"),
							Position: &position.Position{
								StartLine: 2,
								EndLine:   2,
								StartPos:  32,
								EndPos:    33,
							},
						},
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
						StartLine: 2,
						EndLine:   2,
						StartPos:  33,
						EndPos:    34,
					},
				},
			},
		},
	}

	suite.Run()
}

func TestThrowExprAsStmt(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php throw new InvalidArgumentException();`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtThrow{
			Expr: &ast.ExprNew{
				Class: &ast.Name{
					Parts: []ast.Vertex{
						&ast.NamePart{
							Val: []byte("InvalidArgumentException"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestThrowExpr(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php $a ??= throw new InvalidArgumentException();`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprAssignCoalesce{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Val: []byte("$a"),
					},
				},
				Expr: &ast.ExprThrow{
					Expr: &ast.ExprNew{
						Class: &ast.Name{
							Parts: []ast.Vertex{
								&ast.NamePart{
									Val: []byte("InvalidArgumentException"),
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

func TestThrowExprPrecedence(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php $a ??= throw $this->createNotFoundException();`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprAssignCoalesce{
				Var: &ast.ExprVariable{
					Name: &ast.Identifier{
						Val: []byte("$a"),
					},
				},
				Expr: &ast.ExprThrow{
					Expr: &ast.ExprMethodCall{
						Var: &ast.ExprVariable{
							Name: &ast.Identifier{
								Val: []byte("$this"),
							},
						},
						Method: &ast.Identifier{
							Val: []byte("createNotFoundException"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestArrowFunctionPrecedenceWithOr(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php fn($a, $b) => $a and $b;`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprArrowFunction{
				Params: []ast.Vertex{
					&ast.Parameter{
						Var: &ast.ExprVariable{
							Name: &ast.Identifier{
								Val: []byte("$a"),
							},
						},
					},
					&ast.Parameter{
						Var: &ast.ExprVariable{
							Name: &ast.Identifier{
								Val: []byte("$b"),
							},
						},
					},
				},
				Expr: &ast.ExprBinaryLogicalAnd{
					Left: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$b"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestArrowFunctionPrecedenceWithAnd(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php fn($a, $b) => $a && $b;`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprArrowFunction{
				Params: []ast.Vertex{
					&ast.Parameter{
						Var: &ast.ExprVariable{
							Name: &ast.Identifier{
								Val: []byte("$a"),
							},
						},
					},
					&ast.Parameter{
						Var: &ast.ExprVariable{
							Name: &ast.Identifier{
								Val: []byte("$b"),
							},
						},
					},
				},
				Expr: &ast.ExprBinaryBooleanAnd{
					Left: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$b"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestConcatPrecedenceWithPlus(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php "str" . $a + $b;`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprBinaryConcat{
				Left: &ast.ScalarString{
					Val: []byte("\"str\""),
				},
				Right: &ast.ExprBinaryPlus{
					Left: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$b"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestConcatPrecedenceWithMinus(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php "str" . $a - $b;`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprBinaryConcat{
				Left: &ast.ScalarString{
					Val: []byte("\"str\""),
				},
				Right: &ast.ExprBinaryMinus{
					Left: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$b"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestConcatPrecedenceWithShiftLeft(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php "str" . $a << $b;`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtExpression{
			Expr: &ast.ExprBinaryConcat{
				Left: &ast.ScalarString{
					Val: []byte("\"str\""),
				},
				Right: &ast.ExprBinaryShiftLeft{
					Left: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$a"),
						},
					},
					Right: &ast.ExprVariable{
						Name: &ast.Identifier{
							Val: []byte("$b"),
						},
					},
				},
			},
		},
	},
},`

	suite.Run()
}

func TestPropertiesInConstructor(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
class Point {
  public function __construct(
    public float $x = 0.0,
    public float $y = 0.0,
    public float $z = 0.0,
  ) {}
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtClass{
			Name: &ast.Identifier{
				Val: []byte("Point"),
			},
			Stmts: []ast.Vertex{
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
							Visibility: &ast.Identifier{
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("float"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$x"),
								},
							},
							DefaultValue: &ast.ScalarDnumber{
								Val: []byte("0.0"),
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("float"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$y"),
								},
							},
							DefaultValue: &ast.ScalarDnumber{
								Val: []byte("0.0"),
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										Val: []byte("float"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									Val: []byte("$z"),
								},
							},
							DefaultValue: &ast.ScalarDnumber{
								Val: []byte("0.0"),
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

func TestMixedPropertiesInConstructor(t *testing.T) {
	suite := tester.NewParserDumpTestSuite(t)
	suite.UsePHP8()
	suite.WithTokens()
	suite.Code = `<?php
class Point {
  public function __construct(
    public $x,
    public float $y,
    public float &$z,
    public float $a = 0.0,
    public float &$b = 0.0,
  ) {}
}
`

	suite.Expected = `&ast.Root{
	Stmts: []ast.Vertex{
		&ast.StmtClass{
			ClassTkn: &token.Token{
				ID: token.T_CLASS,
				Val: []byte("class"),
				FreeFloating: []*token.Token{
					{
						ID: token.T_OPEN_TAG,
						Val: []byte("<?php"),
					},
					{
						ID: token.T_WHITESPACE,
						Val: []byte("\n"),
					},
				},
			},
			Name: &ast.Identifier{
				IdentifierTkn: &token.Token{
					ID: token.T_STRING,
					Val: []byte("Point"),
					FreeFloating: []*token.Token{
						{
							ID: token.T_WHITESPACE,
							Val: []byte(" "),
						},
					},
				},
				Val: []byte("Point"),
			},
			OpenCurlyBracketTkn: &token.Token{
				ID: token.ID(123),
				Val: []byte("{"),
				FreeFloating: []*token.Token{
					{
						ID: token.T_WHITESPACE,
						Val: []byte(" "),
					},
				},
			},
			Stmts: []ast.Vertex{
				&ast.StmtClassMethod{
					Modifiers: []ast.Vertex{
						&ast.Identifier{
							IdentifierTkn: &token.Token{
								ID: token.T_PUBLIC,
								Val: []byte("public"),
								FreeFloating: []*token.Token{
									{
										ID: token.T_WHITESPACE,
										Val: []byte("\n  "),
									},
								},
							},
							Val: []byte("public"),
						},
					},
					FunctionTkn: &token.Token{
						ID: token.T_FUNCTION,
						Val: []byte("function"),
						FreeFloating: []*token.Token{
							{
								ID: token.T_WHITESPACE,
								Val: []byte(" "),
							},
						},
					},
					Name: &ast.Identifier{
						IdentifierTkn: &token.Token{
							ID: token.T_STRING,
							Val: []byte("__construct"),
							FreeFloating: []*token.Token{
								{
									ID: token.T_WHITESPACE,
									Val: []byte(" "),
								},
							},
						},
						Val: []byte("__construct"),
					},
					OpenParenthesisTkn: &token.Token{
						ID: token.ID(40),
						Val: []byte("("),
					},
					Params: []ast.Vertex{
						&ast.Parameter{
							Visibility: &ast.Identifier{
								IdentifierTkn: &token.Token{
									ID: token.T_PUBLIC,
									Val: []byte("public"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte("\n    "),
										},
									},
								},
								Val: []byte("public"),
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									IdentifierTkn: &token.Token{
										ID: token.T_VARIABLE,
										Val: []byte("$x"),
										FreeFloating: []*token.Token{
											{
												ID: token.T_WHITESPACE,
												Val: []byte(" "),
											},
										},
									},
									Val: []byte("$x"),
								},
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								IdentifierTkn: &token.Token{
									ID: token.T_PUBLIC,
									Val: []byte("public"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte("\n    "),
										},
									},
								},
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										StringTkn: &token.Token{
											ID: token.T_STRING,
											Val: []byte("float"),
											FreeFloating: []*token.Token{
												{
													ID: token.T_WHITESPACE,
													Val: []byte(" "),
												},
											},
										},
										Val: []byte("float"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									IdentifierTkn: &token.Token{
										ID: token.T_VARIABLE,
										Val: []byte("$y"),
										FreeFloating: []*token.Token{
											{
												ID: token.T_WHITESPACE,
												Val: []byte(" "),
											},
										},
									},
									Val: []byte("$y"),
								},
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								IdentifierTkn: &token.Token{
									ID: token.T_PUBLIC,
									Val: []byte("public"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte("\n    "),
										},
									},
								},
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										StringTkn: &token.Token{
											ID: token.T_STRING,
											Val: []byte("float"),
											FreeFloating: []*token.Token{
												{
													ID: token.T_WHITESPACE,
													Val: []byte(" "),
												},
											},
										},
										Val: []byte("float"),
									},
								},
							},
							AmpersandTkn: &token.Token{
								ID: token.ID(38),
								Val: []byte("&"),
								FreeFloating: []*token.Token{
									{
										ID: token.T_WHITESPACE,
										Val: []byte(" "),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									IdentifierTkn: &token.Token{
										ID: token.T_VARIABLE,
										Val: []byte("$z"),
									},
									Val: []byte("$z"),
								},
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								IdentifierTkn: &token.Token{
									ID: token.T_PUBLIC,
									Val: []byte("public"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte("\n    "),
										},
									},
								},
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										StringTkn: &token.Token{
											ID: token.T_STRING,
											Val: []byte("float"),
											FreeFloating: []*token.Token{
												{
													ID: token.T_WHITESPACE,
													Val: []byte(" "),
												},
											},
										},
										Val: []byte("float"),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									IdentifierTkn: &token.Token{
										ID: token.T_VARIABLE,
										Val: []byte("$a"),
										FreeFloating: []*token.Token{
											{
												ID: token.T_WHITESPACE,
												Val: []byte(" "),
											},
										},
									},
									Val: []byte("$a"),
								},
							},
							EqualTkn: &token.Token{
								ID: token.ID(61),
								Val: []byte("="),
								FreeFloating: []*token.Token{
									{
										ID: token.T_WHITESPACE,
										Val: []byte(" "),
									},
								},
							},
							DefaultValue: &ast.ScalarDnumber{
								NumberTkn: &token.Token{
									ID: token.T_DNUMBER,
									Val: []byte("0.0"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte(" "),
										},
									},
								},
								Val: []byte("0.0"),
							},
						},
						&ast.Parameter{
							Visibility: &ast.Identifier{
								IdentifierTkn: &token.Token{
									ID: token.T_PUBLIC,
									Val: []byte("public"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte("\n    "),
										},
									},
								},
								Val: []byte("public"),
							},
							Type: &ast.Name{
								Parts: []ast.Vertex{
									&ast.NamePart{
										StringTkn: &token.Token{
											ID: token.T_STRING,
											Val: []byte("float"),
											FreeFloating: []*token.Token{
												{
													ID: token.T_WHITESPACE,
													Val: []byte(" "),
												},
											},
										},
										Val: []byte("float"),
									},
								},
							},
							AmpersandTkn: &token.Token{
								ID: token.ID(38),
								Val: []byte("&"),
								FreeFloating: []*token.Token{
									{
										ID: token.T_WHITESPACE,
										Val: []byte(" "),
									},
								},
							},
							Var: &ast.ExprVariable{
								Name: &ast.Identifier{
									IdentifierTkn: &token.Token{
										ID: token.T_VARIABLE,
										Val: []byte("$b"),
									},
									Val: []byte("$b"),
								},
							},
							EqualTkn: &token.Token{
								ID: token.ID(61),
								Val: []byte("="),
								FreeFloating: []*token.Token{
									{
										ID: token.T_WHITESPACE,
										Val: []byte(" "),
									},
								},
							},
							DefaultValue: &ast.ScalarDnumber{
								NumberTkn: &token.Token{
									ID: token.T_DNUMBER,
									Val: []byte("0.0"),
									FreeFloating: []*token.Token{
										{
											ID: token.T_WHITESPACE,
											Val: []byte(" "),
										},
									},
								},
								Val: []byte("0.0"),
							},
						},
					},
					SeparatorTkns: []*token.Token{
						{
							ID: token.ID(44),
							Val: []byte(","),
						},
						{
							ID: token.ID(44),
							Val: []byte(","),
						},
						{
							ID: token.ID(44),
							Val: []byte(","),
						},
						{
							ID: token.ID(44),
							Val: []byte(","),
						},
						{
							ID: token.ID(44),
							Val: []byte(","),
						},
					},
					CloseParenthesisTkn: &token.Token{
						ID: token.ID(41),
						Val: []byte(")"),
						FreeFloating: []*token.Token{
							{
								ID: token.T_WHITESPACE,
								Val: []byte("\n  "),
							},
						},
					},
					Stmt: &ast.StmtStmtList{
						OpenCurlyBracketTkn: &token.Token{
							ID: token.ID(123),
							Val: []byte("{"),
							FreeFloating: []*token.Token{
								{
									ID: token.T_WHITESPACE,
									Val: []byte(" "),
								},
							},
						},
						Stmts: []ast.Vertex{},
						CloseCurlyBracketTkn: &token.Token{
							ID: token.ID(125),
							Val: []byte("}"),
						},
					},
				},
			},
			CloseCurlyBracketTkn: &token.Token{
				ID: token.ID(125),
				Val: []byte("}"),
				FreeFloating: []*token.Token{
					{
						ID: token.T_WHITESPACE,
						Val: []byte("\n"),
					},
				},
			},
		},
	},
	EndTkn: &token.Token{
		FreeFloating: []*token.Token{
			{
				ID: token.T_WHITESPACE,
				Val: []byte("\n"),
			},
		},
	},
},`

	suite.Run()
}

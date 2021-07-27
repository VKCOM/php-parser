package php8

import (
	"github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/pkg/ast"
	position2 "github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

// Builder is responsible for creating nodes inside grammar rules.
// Creating nodes directly in grammar rules is inconvenient, since
// there is no autocompletion and you cannot put breakpoints inside.
type Builder struct {
	Pos *position.Builder
}

// NewBuilder creates a new Builder.
func NewBuilder(pos *position.Builder) *Builder {
	return &Builder{Pos: pos}
}

// NewEmptySeparatedList creates a new empty list.
// Used for places where a delimited list is used.
func (b *Builder) NewEmptySeparatedList() *ParserSeparatedList {
	return &ParserSeparatedList{}
}

// NewSeparatedList creates a new single-element list.
// Used for places where a delimited list is used.
func (b *Builder) NewSeparatedList(node ast.Vertex) *ParserSeparatedList {
	return &ParserSeparatedList{Items: []ast.Vertex{node}}
}

// NewSeparatedListWithTwoElements creates a new two-element list.
// Used for places where a delimited list is used.
func (b *Builder) NewSeparatedListWithTwoElements(node1 ast.Vertex, tkn *token.Token, node2 ast.Vertex) *ParserSeparatedList {
	return &ParserSeparatedList{
		Items:         []ast.Vertex{node1, node2},
		SeparatorTkns: []*token.Token{tkn},
	}
}

// AppendToSeparatedList inserts a new node and/or token into the list.
func (b *Builder) AppendToSeparatedList(list ast.Vertex, tkn *token.Token, node ast.Vertex) *ParserSeparatedList {
	sepList := list.(*ParserSeparatedList)
	if node != nil {
		sepList.Items = append(sepList.Items, node)
	}
	if tkn != nil {
		sepList.SeparatorTkns = append(sepList.SeparatorTkns, tkn)
	}

	return sepList
}

func (b *Builder) NewExpressionStmt(
	Expr ast.Vertex,
	SemiColonTkn *token.Token,
) ast.Vertex {
	if throwExpr, ok := Expr.(*ast.ExprThrow); ok {
		// For backwards-compatibility reasons, convert throw in statement position into
		// StmtThrow rather than StmtExpression(ExprThrow).
		return &ast.StmtThrow{
			Position:     b.Pos.NewTokensPosition(throwExpr.ThrowTkn, SemiColonTkn),
			ThrowTkn:     throwExpr.ThrowTkn,
			Expr:         throwExpr.Expr,
			SemiColonTkn: SemiColonTkn,
		}
	}

	return &ast.StmtExpression{
		Position:     b.Pos.NewNodeTokenPosition(Expr, SemiColonTkn),
		Expr:         Expr,
		SemiColonTkn: SemiColonTkn,
	}
}

func (b *Builder) NewMethodCall(
	Expr ast.Vertex,
	ObjectOperatorTkn *token.Token,
	PropertyName ast.Vertex,
	ArgList ast.Vertex,
) *ast.ExprMethodCall {
	argumentList := ArgList.(*ArgumentList)
	methodCall := &ast.ExprMethodCall{
		Position:            b.Pos.NewNodesPosition(Expr, ArgList),
		Var:                 Expr,
		ObjectOperatorTkn:   ObjectOperatorTkn,
		Method:              PropertyName,
		OpenParenthesisTkn:  argumentList.OpenParenthesisTkn,
		Args:                argumentList.Arguments,
		SeparatorTkns:       argumentList.SeparatorTkns,
		CloseParenthesisTkn: argumentList.CloseParenthesisTkn,
	}

	if brackets, ok := PropertyName.(*ParserBrackets); ok {
		methodCall.OpenCurlyBracketTkn = brackets.OpenBracketTkn
		methodCall.Method = brackets.Child
		methodCall.CloseCurlyBracketTkn = brackets.CloseBracketTkn
	}

	return methodCall
}

func (b *Builder) NewNullsafeMethodCall(
	Expr ast.Vertex,
	ObjectOperatorTkn *token.Token,
	PropertyName ast.Vertex,
	ArgList ast.Vertex,
) *ast.ExprNullsafeMethodCall {
	argumentList := ArgList.(*ArgumentList)
	methodCall := &ast.ExprNullsafeMethodCall{
		Position:            b.Pos.NewNodesPosition(Expr, ArgList),
		Var:                 Expr,
		ObjectOperatorTkn:   ObjectOperatorTkn,
		Method:              PropertyName,
		OpenParenthesisTkn:  argumentList.OpenParenthesisTkn,
		Args:                argumentList.Arguments,
		SeparatorTkns:       argumentList.SeparatorTkns,
		CloseParenthesisTkn: argumentList.CloseParenthesisTkn,
	}

	if brackets, ok := PropertyName.(*ParserBrackets); ok {
		methodCall.OpenCurlyBracketTkn = brackets.OpenBracketTkn
		methodCall.Method = brackets.Child
		methodCall.CloseCurlyBracketTkn = brackets.CloseBracketTkn
	}

	return methodCall
}

func (b *Builder) NewPropertyFetch(
	Expr ast.Vertex,
	ObjectOperatorTkn *token.Token,
	PropertyName ast.Vertex,
) *ast.ExprPropertyFetch {
	propertyFetch := &ast.ExprPropertyFetch{
		Position:          b.Pos.NewNodesPosition(Expr, PropertyName),
		Var:               Expr,
		ObjectOperatorTkn: ObjectOperatorTkn,
		Prop:              PropertyName,
	}

	if brackets, ok := PropertyName.(*ParserBrackets); ok {
		propertyFetch.OpenCurlyBracketTkn = brackets.OpenBracketTkn
		propertyFetch.Prop = brackets.Child
		propertyFetch.CloseCurlyBracketTkn = brackets.CloseBracketTkn
	}

	return propertyFetch
}

func (b *Builder) NewPropertyFetchFromTokens(
	VariableTkn *token.Token,
	ObjectOperatorTkn *token.Token,
	StringTkn *token.Token,
) *ast.ExprPropertyFetch {
	return &ast.ExprPropertyFetch{
		Position: b.Pos.NewTokensPosition(VariableTkn, StringTkn),
		Var: &ast.ExprVariable{
			Position: b.Pos.NewTokenPosition(VariableTkn),
			Name: &ast.Identifier{
				Position:      b.Pos.NewTokenPosition(VariableTkn),
				IdentifierTkn: VariableTkn,
				Value:         VariableTkn.Value,
			},
		},
		ObjectOperatorTkn: ObjectOperatorTkn,
		Prop: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(StringTkn),
			IdentifierTkn: StringTkn,
			Value:         StringTkn.Value,
		},
	}
}

func (b *Builder) NewNullsafePropertyFetch(
	Expr ast.Vertex,
	ObjectOperatorTkn *token.Token,
	PropertyName ast.Vertex,
) *ast.ExprNullsafePropertyFetch {
	propertyFetch := &ast.ExprNullsafePropertyFetch{
		Position:          b.Pos.NewNodesPosition(Expr, PropertyName),
		Var:               Expr,
		ObjectOperatorTkn: ObjectOperatorTkn,
		Prop:              PropertyName,
	}

	if brackets, ok := PropertyName.(*ParserBrackets); ok {
		propertyFetch.OpenCurlyBracketTkn = brackets.OpenBracketTkn
		propertyFetch.Prop = brackets.Child
		propertyFetch.CloseCurlyBracketTkn = brackets.CloseBracketTkn
	}

	return propertyFetch
}

func (b *Builder) NewNullsafePropertyFetchFromTokens(
	VariableTkn *token.Token,
	ObjectOperatorTkn *token.Token,
	StringTkn *token.Token,
) *ast.ExprNullsafePropertyFetch {
	return &ast.ExprNullsafePropertyFetch{
		Position: b.Pos.NewTokensPosition(VariableTkn, StringTkn),
		Var: &ast.ExprVariable{
			Position: b.Pos.NewTokenPosition(VariableTkn),
			Name: &ast.Identifier{
				Position:      b.Pos.NewTokenPosition(VariableTkn),
				IdentifierTkn: VariableTkn,
				Value:         VariableTkn.Value,
			},
		},
		ObjectOperatorTkn: ObjectOperatorTkn,
		Prop: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(StringTkn),
			IdentifierTkn: StringTkn,
			Value:         StringTkn.Value,
		},
	}
}

func (b *Builder) NewArgument(
	Expr ast.Vertex,
) *ast.Argument {
	return &ast.Argument{
		Position: b.Pos.NewNodePosition(Expr),
		Expr:     Expr,
	}
}

func (b *Builder) NewVariadicArgument(
	VariadicTkn *token.Token,
	Expr ast.Vertex,
) *ast.Argument {
	return &ast.Argument{
		Position:    b.Pos.NewTokenNodePosition(VariadicTkn, Expr),
		VariadicTkn: VariadicTkn,
		Expr:        Expr,
	}
}

func (b *Builder) NewNamedArgument(
	Name *token.Token,
	ColonTkn *token.Token,
	Expr ast.Vertex,
) *ast.Argument {
	return &ast.Argument{
		Position: b.Pos.NewTokenNodePosition(Name, Expr),
		Name: &ast.Identifier{
			Position:      Name.Position,
			IdentifierTkn: Name,
			Value:         Name.Value,
		},
		ColonTkn: ColonTkn,
		Expr:     Expr,
	}
}

func (b *Builder) NewMatch(
	MatchTkn *token.Token,
	OpenParenthesisTkn *token.Token,
	Expr ast.Vertex,
	CloseParenthesisTkn *token.Token,
	OpenCurlyBracketTkn *token.Token,
	Cases ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.ExprMatch {
	var arms []ast.Vertex
	var sepTkns []*token.Token
	if Cases != nil {
		cases := Cases.(*ParserSeparatedList)
		arms = cases.Items
		sepTkns = cases.SeparatorTkns
	}

	return &ast.ExprMatch{
		Position:             b.Pos.NewTokensPosition(MatchTkn, CloseCurlyBracketTkn),
		MatchTkn:             MatchTkn,
		OpenParenthesisTkn:   OpenParenthesisTkn,
		Expr:                 Expr,
		CloseParenthesisTkn:  CloseParenthesisTkn,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Arms:                 arms,
		SeparatorTkns:        sepTkns,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}
}

func (b *Builder) NewMatchArm(
	DefaultTkn *token.Token,
	DefaultCommaTkn *token.Token,
	Exprs ast.Vertex,
	DoubleArrowTkn *token.Token,
	ReturnExpr ast.Vertex,
) *ast.MatchArm {
	// Default branch.
	if Exprs == nil {
		return &ast.MatchArm{
			Position:        b.Pos.NewTokenNodePosition(DefaultTkn, ReturnExpr),
			DefaultTkn:      DefaultTkn,
			DefaultCommaTkn: DefaultCommaTkn,
			DoubleArrowTkn:  DoubleArrowTkn,
			ReturnExpr:      ReturnExpr,
		}
	}
	list := Exprs.(*ParserSeparatedList)
	return &ast.MatchArm{
		Position:       b.Pos.NewNodeListNodePosition(list.Items, ReturnExpr),
		DefaultTkn:     nil,
		Exprs:          list.Items,
		SeparatorTkns:  list.SeparatorTkns,
		DoubleArrowTkn: DoubleArrowTkn,
		ReturnExpr:     ReturnExpr,
	}
}

func (b *Builder) NewNameType(
	IdentifierTkn *token.Token,
) *ast.Identifier {
	return &ast.Identifier{
		Position:      b.Pos.NewTokenPosition(IdentifierTkn),
		IdentifierTkn: IdentifierTkn,
		Value:         IdentifierTkn.Value,
	}
}

func (b *Builder) NewNullableType(
	QuestionTkn *token.Token,
	Expr ast.Vertex,
) *ast.Nullable {
	return &ast.Nullable{
		Position:    b.Pos.NewTokenNodePosition(QuestionTkn, Expr),
		QuestionTkn: QuestionTkn,
		Expr:        Expr,
	}
}

func (b *Builder) NewUnionType(
	Types ast.Vertex,
) *ast.Union {
	var types []ast.Vertex
	var sepTkns []*token.Token
	if Types != nil {
		cases := Types.(*ParserSeparatedList)
		types = cases.Items
		sepTkns = cases.SeparatorTkns
	}

	return &ast.Union{
		Position:      b.Pos.NewNodeListPosition(types),
		Types:         types,
		SeparatorTkns: sepTkns,
	}
}

func (b *Builder) NewReturnType(
	ColonTkn *token.Token,
	Type ast.Vertex,
) *ReturnType {
	return &ReturnType{
		ColonTkn: ColonTkn,
		Type:     Type,
	}
}

func (b *Builder) NewTry(
	TryTkn *token.Token,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
	Catches ast.Vertex,
	Finally ast.Vertex,
) *ast.StmtTry {
	var catches []ast.Vertex
	if Catches != nil {
		catchList := Catches.(*ParserSeparatedList)
		catches = catchList.Items
	}

	var pos *position2.Position

	if Finally != nil {
		pos = b.Pos.NewTokenNodePosition(TryTkn, Finally)
	} else {
		pos = b.Pos.NewTokenNodeListPosition(TryTkn, catches)
	}

	return &ast.StmtTry{
		Position:             pos,
		TryTkn:               TryTkn,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
		Catches:              catches,
		Finally:              Finally,
	}
}

func (b *Builder) NewCatch(
	CatchTkn *token.Token,
	OpenParenthesisTkn *token.Token,
	CatchNameList ast.Vertex,
	Variable *token.Token,
	CloseParenthesisTkn *token.Token,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtCatch {
	var types []ast.Vertex
	var sepTkns []*token.Token
	if CatchNameList != nil {
		catchList := CatchNameList.(*ParserSeparatedList)
		types = catchList.Items
		sepTkns = catchList.SeparatorTkns
	}

	var variable ast.Vertex

	if Variable != nil {
		variable = &ast.ExprVariable{
			Position: b.Pos.NewTokenPosition(Variable),
			Name: &ast.Identifier{
				Position:      b.Pos.NewTokenPosition(Variable),
				IdentifierTkn: Variable,
				Value:         Variable.Value,
			},
		}
	}

	return &ast.StmtCatch{
		Position:             b.Pos.NewTokensPosition(CatchTkn, CloseCurlyBracketTkn),
		CatchTkn:             CatchTkn,
		OpenParenthesisTkn:   OpenParenthesisTkn,
		Types:                types,
		SeparatorTkns:        sepTkns,
		Var:                  variable,
		CloseParenthesisTkn:  CloseParenthesisTkn,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}
}

func (b *Builder) NewThrowStmt(
	ThrowTkn *token.Token,
	Expr ast.Vertex,
) *ast.StmtThrow {
	return &ast.StmtThrow{
		Position: b.Pos.NewTokenNodePosition(ThrowTkn, Expr),
		ThrowTkn: ThrowTkn,
		Expr:     Expr,
	}
}

func (b *Builder) NewThrowExpr(
	ThrowTkn *token.Token,
	Expr ast.Vertex,
) *ast.ExprThrow {
	return &ast.ExprThrow{
		Position: b.Pos.NewTokenNodePosition(ThrowTkn, Expr),
		ThrowTkn: ThrowTkn,
		Expr:     Expr,
	}
}

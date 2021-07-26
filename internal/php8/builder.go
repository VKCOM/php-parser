package php8

import (
	"github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/pkg/ast"
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

// NewSeparatedList creates a new single-element list.
// Used for places where a delimited list is used.
func (b *Builder) NewSeparatedList(node ast.Vertex) *ParserSeparatedList {
	return &ParserSeparatedList{Items: []ast.Vertex{node}}
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

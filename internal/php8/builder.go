package php8

import (
	"bytes"

	"github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/pkg/ast"
	position2 "github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

// Builder is responsible for creating nodes inside grammar rules.
// Creating nodes directly in grammar rules is inconvenient, since
// there is no autocompletion and you cannot put breakpoints inside.
type Builder struct {
	Pos    *position.Builder
	Parser *Parser
}

// NewBuilder creates a new Builder.
func NewBuilder(pos *position.Builder, parser *Parser) *Builder {
	return &Builder{
		Pos:    pos,
		Parser: parser,
	}
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

func (b *Builder) NewNonEmptySeparatedList(nodes []ast.Vertex, tkns []*token.Token) *ParserSeparatedList {
	return &ParserSeparatedList{
		Items:         nodes,
		SeparatorTkns: tkns,
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

func (b *Builder) SeparatedListItems(list ast.Vertex) (items []ast.Vertex, sepTkns []*token.Token) {
	if list == nil {
		return nil, nil
	}

	paramsList, ok := list.(*ParserSeparatedList)
	if !ok {
		return nil, nil
	}

	return paramsList.Items, paramsList.SeparatorTkns
}

// namespace\Foo
func (b *Builder) parseNameRelativeToken(t *token.Token) *ast.NameRelative {
	n := &ast.NameRelative{
		Position: b.Pos.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	// namespace token

	p1 := b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 9

	n.NsTkn = &token.Token{
		ID:           token.T_NAMESPACE,
		Value:        v[:9],
		Position:     p1,
		FreeFloating: t.FreeFloating,
	}

	s = s + 9
	v = v[9:]

	// ns separator token

	p1 = b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 1

	n.NsSeparatorTkn = &token.Token{
		ID:       token.T_NS_SEPARATOR,
		Value:    v[:1],
		Position: p1,
	}

	s = s + 1
	v = v[1:]

	// parts

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 = b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := b.Parser.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:       token.T_STRING,
				Value:    v[:i],
				Position: p2,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	// last part

	p1 = b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := b.Parser.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:       token.T_STRING,
			Value:    v,
			Position: p2,
		},
		Value: v,
	})

	return n
}

// \Foo\Boo
func (b *Builder) parseNameFullyQualifiedToken(t *token.Token) *ast.NameFullyQualified {
	n := &ast.NameFullyQualified{
		Position: b.Pos.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	// ns separator token

	p1 := b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + 1

	n.NsSeparatorTkn = &token.Token{
		ID:           token.T_NS_SEPARATOR,
		Value:        v[:1],
		Position:     p1,
		FreeFloating: t.FreeFloating,
	}

	s = s + 1
	v = v[1:]

	// parts

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 = b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := b.Parser.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:       token.T_STRING,
				Value:    v[:i],
				Position: p2,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	// last part

	p1 = b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := b.Parser.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:       token.T_STRING,
			Value:    v,
			Position: p2,
		},
		Value: v,
	})

	return n
}

// Foo\Boo
func (b *Builder) parseNameToken(t *token.Token) *ast.NameRelative {
	n := &ast.NameRelative{
		Position: b.Pos.NewTokenPosition(t),
	}
	s := t.Position.StartPos
	v := t.Value

	if bytes.HasPrefix(bytes.ToLower(v), []byte("namespace")) {
		p1 := b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 9

		n.NsTkn = &token.Token{
			ID:       token.T_NAMESPACE,
			Value:    v[:9],
			Position: p1,
		}

		s = s + 9
		v = v[9:]
	}

	if bytes.HasPrefix(bytes.ToLower(v), []byte("\\")) {
		p1 := b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.NsSeparatorTkn = &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		}

		s = s + 1
		v = v[1:]
	}

	for {
		i := bytes.Index(v, []byte("\\"))
		if i < 0 {
			break
		}

		p1 := b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + i

		p2 := b.Parser.Lexer.positionPool.Get()
		*p2 = *p1

		n.Parts = append(n.Parts, &ast.NamePart{
			Position: p1,
			StringTkn: &token.Token{
				ID:           token.T_STRING,
				Value:        v[:i],
				Position:     p2,
				FreeFloating: t.FreeFloating,
			},
			Value: v[:i],
		})
		t.FreeFloating = nil
		s = s + i
		v = v[i:]

		p1 = b.Parser.Lexer.positionPool.Get()
		p1.StartLine = t.Position.StartLine
		p1.EndLine = t.Position.EndLine
		p1.StartPos = s
		p1.EndPos = s + 1

		n.SeparatorTkns = append(n.SeparatorTkns, &token.Token{
			ID:       token.T_NS_SEPARATOR,
			Value:    v[:1],
			Position: p1,
		})
		s = s + 1
		v = v[1:]
	}

	p1 := b.Parser.Lexer.positionPool.Get()
	p1.StartLine = t.Position.StartLine
	p1.EndLine = t.Position.EndLine
	p1.StartPos = s
	p1.EndPos = s + len(v)
	p2 := b.Parser.Lexer.positionPool.Get()
	*p2 = *p1

	n.Parts = append(n.Parts, &ast.NamePart{
		Position: p1,
		StringTkn: &token.Token{
			ID:           token.T_STRING,
			Value:        v,
			Position:     p2,
			FreeFloating: t.FreeFloating,
		},
		Value: v,
	})

	return n
}

func (b *Builder) NewScalarMagicConstant(
	MagicConstantTkn *token.Token,
) *ast.ScalarMagicConstant {
	return &ast.ScalarMagicConstant{
		Position:      b.Pos.NewTokenPosition(MagicConstantTkn),
		MagicConstTkn: MagicConstantTkn,
		Value:         MagicConstantTkn.Value,
	}
}

func (b *Builder) NewBracket(
	OpenParenthesisTkn *token.Token,
	Expr ast.Vertex,
	CloseParenthesisTkn *token.Token,
) *ast.ExprBrackets {
	return &ast.ExprBrackets{
		Position:            b.Pos.NewTokensPosition(OpenParenthesisTkn, CloseParenthesisTkn),
		OpenParenthesisTkn:  OpenParenthesisTkn,
		Expr:                Expr,
		CloseParenthesisTkn: CloseParenthesisTkn,
	}
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

func (b *Builder) NewIdentifier(
	IdentifierTkn *token.Token,
) ast.Vertex {
	if IdentifierTkn == nil {
		return nil
	}

	return &ast.Identifier{
		Position:      b.Pos.NewTokenPosition(IdentifierTkn),
		IdentifierTkn: IdentifierTkn,
		Value:         IdentifierTkn.Value,
	}
}

func (b *Builder) NewName(
	NameTkn *token.Token,
) *ast.Name {
	return &ast.Name{
		Position: b.Pos.NewTokenPosition(NameTkn),
		Parts: []ast.Vertex{
			&ast.NamePart{
				Position:  b.Pos.NewTokenPosition(NameTkn),
				StringTkn: NameTkn,
				Value:     NameTkn.Value,
			},
		},
	}
}

func (b *Builder) NewNameQualified(
	NameTkn *token.Token,
) *ast.Name {
	name := b.parseNameToken(NameTkn)

	return &ast.Name{
		Position:      b.Pos.NewTokenPosition(NameTkn),
		Parts:         name.Parts,
		SeparatorTkns: name.SeparatorTkns,
	}
}

func (b *Builder) NewNameFullyQualified(
	NameTkn *token.Token,
) *ast.NameFullyQualified {
	return b.parseNameFullyQualifiedToken(NameTkn)
}

func (b *Builder) NewNameRelative(
	NameTkn *token.Token,
) *ast.NameRelative {
	return b.parseNameRelativeToken(NameTkn)
}

func (b *Builder) NewNamePart(
	NameTkn *token.Token,
) *ast.NamePart {
	return &ast.NamePart{
		Position:  b.Pos.NewTokenPosition(NameTkn),
		StringTkn: NameTkn,
		Value:     NameTkn.Value,
	}
}

func (b *Builder) NewConstFetch(
	Name ast.Vertex,
) *ast.ExprConstFetch {
	return &ast.ExprConstFetch{
		Position: b.Pos.NewNodePosition(Name),
		Const:    Name,
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

func (b *Builder) NewParameter(
	AttrGroups []ast.Vertex,
	Visibility ast.Vertex,
	Type ast.Vertex,
	AmpersandTkn *token.Token,
	VariadicTkn *token.Token,
	VarTkn *token.Token,
	EqualTkn *token.Token,
	DefaultValue ast.Vertex,
	WithDefault bool,
) *ast.Parameter {
	var pos *position2.Position

	if WithDefault {
		if AttrGroups != nil {
			pos = b.Pos.NewNodeListNodePosition(AttrGroups, DefaultValue)
		} else if Visibility != nil {
			pos = b.Pos.NewNodesPosition(Visibility, DefaultValue)
		} else if Type != nil {
			pos = b.Pos.NewNodesPosition(Type, DefaultValue)
		} else if AmpersandTkn != nil {
			pos = b.Pos.NewTokenNodePosition(AmpersandTkn, DefaultValue)
		} else if VariadicTkn != nil {
			pos = b.Pos.NewTokenNodePosition(VariadicTkn, DefaultValue)
		} else {
			pos = b.Pos.NewTokenNodePosition(VarTkn, DefaultValue)
		}
	} else {
		if AttrGroups != nil {
			pos = b.Pos.NewNodeListTokenPosition(AttrGroups, VarTkn)
		} else if Visibility != nil {
			pos = b.Pos.NewNodeTokenPosition(Visibility, VarTkn)
		} else if Type != nil {
			pos = b.Pos.NewNodeTokenPosition(Type, VarTkn)
		} else if AmpersandTkn != nil {
			pos = b.Pos.NewTokensPosition(AmpersandTkn, VarTkn)
		} else if VariadicTkn != nil {
			pos = b.Pos.NewTokensPosition(VariadicTkn, VarTkn)
		} else {
			pos = b.Pos.NewTokenPosition(VarTkn)
		}
	}

	return &ast.Parameter{
		Position:     pos,
		AttrGroups:   AttrGroups,
		Visibility:   Visibility,
		Type:         Type,
		AmpersandTkn: AmpersandTkn,
		VariadicTkn:  VariadicTkn,
		Var: &ast.ExprVariable{
			Position: b.Pos.NewTokenPosition(VarTkn),
			Name: &ast.Identifier{
				Position:      b.Pos.NewTokenPosition(VarTkn),
				IdentifierTkn: VarTkn,
				Value:         VarTkn.Value,
			},
		},
		EqualTkn:     EqualTkn,
		DefaultValue: DefaultValue,
	}
}

func (b *Builder) NewAttribute(
	Name ast.Vertex,
	ArgList ast.Vertex,
) *ast.Attribute {
	if ArgList == nil {
		return &ast.Attribute{
			Position: b.Pos.NewNodePosition(Name),
			Name:     Name,
		}
	}

	argList := ArgList.(*ArgumentList)
	return &ast.Attribute{
		Position:            b.Pos.NewNodeTokenPosition(Name, argList.CloseParenthesisTkn),
		Name:                Name,
		OpenParenthesisTkn:  argList.OpenParenthesisTkn,
		Args:                argList.Arguments,
		SeparatorTkns:       argList.SeparatorTkns,
		CloseParenthesisTkn: argList.CloseParenthesisTkn,
	}
}

func (b *Builder) NewAttributeGroup(
	OpenAttributeTkn *token.Token,
	AttrsList ast.Vertex,
	OptionalCommaTkn *token.Token,
	CloseAttributeTkn *token.Token,
) *ast.AttributeGroup {
	attrList := AttrsList.(*ParserSeparatedList)
	if OptionalCommaTkn != nil {
		attrList.SeparatorTkns = append(attrList.SeparatorTkns, OptionalCommaTkn)
	}
	return &ast.AttributeGroup{
		Position:          b.Pos.NewTokensPosition(OpenAttributeTkn, CloseAttributeTkn),
		OpenAttributeTkn:  OpenAttributeTkn,
		Attrs:             attrList.Items,
		SeparatorTkns:     attrList.SeparatorTkns,
		CloseAttributeTkn: CloseAttributeTkn,
	}
}

func (b *Builder) NewClass(
	AttrGroups []ast.Vertex,
	Modifiers []ast.Vertex,
	ClassTkn *token.Token,
	Name *token.Token,

	ExtendsFrom ast.Vertex,
	ImplementsList ast.Vertex,

	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtClass {
	var pos *position2.Position
	if AttrGroups != nil {
		pos = b.Pos.NewNodeListTokenPosition(AttrGroups, CloseCurlyBracketTkn)
	} else if Modifiers != nil {
		pos = b.Pos.NewNodeListTokenPosition(Modifiers, CloseCurlyBracketTkn)
	} else {
		pos = b.Pos.NewTokensPosition(ClassTkn, CloseCurlyBracketTkn)
	}

	var name ast.Vertex

	if Name != nil {
		name = &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(Name),
			IdentifierTkn: Name,
			Value:         Name.Value,
		}
	}

	class := &ast.StmtClass{
		Position:             pos,
		AttrGroups:           AttrGroups,
		Modifiers:            Modifiers,
		ClassTkn:             ClassTkn,
		Name:                 name,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}

	if ExtendsFrom != nil {
		class.ExtendsTkn = ExtendsFrom.(*ast.StmtClass).ExtendsTkn
		class.Extends = ExtendsFrom.(*ast.StmtClass).Extends
	}

	if ImplementsList != nil {
		class.ImplementsTkn = ImplementsList.(*ast.StmtClass).ImplementsTkn
		class.Implements = ImplementsList.(*ast.StmtClass).Implements
		class.ImplementsSeparatorTkns = ImplementsList.(*ast.StmtClass).ImplementsSeparatorTkns
	}

	return class
}

func (b *Builder) NewAnonClass(
	AttrGroups []ast.Vertex,
	ClassTkn *token.Token,

	ArgList ast.Vertex,

	ExtendsFrom ast.Vertex,
	ImplementsList ast.Vertex,

	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtClass {
	class := b.NewClass(
		AttrGroups,
		nil, // Modifiers
		ClassTkn,
		nil, // Name
		ExtendsFrom,
		ImplementsList,
		OpenCurlyBracketTkn,
		Stmts,
		CloseCurlyBracketTkn,
	)

	if ArgList == nil {
		return class
	}

	argList := ArgList.(*ArgumentList)

	class.OpenParenthesisTkn = argList.OpenParenthesisTkn
	class.Args = argList.Arguments
	class.SeparatorTkns = argList.SeparatorTkns
	class.CloseParenthesisTkn = argList.CloseParenthesisTkn

	return class
}

func (b *Builder) NewTrait(
	AttrGroups []ast.Vertex,
	TraitTkn *token.Token,
	Name *token.Token,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtTrait {
	return &ast.StmtTrait{
		Position:   b.Pos.NewOptionalListTokensPosition(AttrGroups, TraitTkn, CloseCurlyBracketTkn),
		AttrGroups: AttrGroups,
		TraitTkn:   TraitTkn,
		Name: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(Name),
			IdentifierTkn: Name,
			Value:         Name.Value,
		},
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}
}

func (b *Builder) NewInterface(
	AttrGroups []ast.Vertex,
	InterfaceTkn *token.Token,
	Name *token.Token,
	ExtendsFrom ast.Vertex,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtInterface {
	iface := &ast.StmtInterface{
		Position:     b.Pos.NewOptionalListTokensPosition(AttrGroups, InterfaceTkn, CloseCurlyBracketTkn),
		AttrGroups:   AttrGroups,
		InterfaceTkn: InterfaceTkn,
		Name: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(Name),
			IdentifierTkn: Name,
			Value:         Name.Value,
		},
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}

	if ExtendsFrom != nil {
		iface.ExtendsTkn = ExtendsFrom.(*ast.StmtInterface).ExtendsTkn
		iface.Extends = ExtendsFrom.(*ast.StmtInterface).Extends
		iface.ExtendsSeparatorTkns = ExtendsFrom.(*ast.StmtInterface).ExtendsSeparatorTkns
	}

	return iface
}

func (b *Builder) NewFunction(
	AttrGroups []ast.Vertex,
	FunctionTkn *token.Token,
	AmpersandTkn *token.Token,
	Name *token.Token,
	OpenParenthesisTkn *token.Token,
	Params ast.Vertex,
	CloseParenthesisTkn *token.Token,
	RetType ast.Vertex,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
) *ast.StmtFunction {
	var params []ast.Vertex
	var sepTkns []*token.Token
	if Params != nil {
		paramsList := Params.(*ParserSeparatedList)
		params = paramsList.Items
		sepTkns = paramsList.SeparatorTkns
	}

	returnType := RetType.(*ReturnType)

	return &ast.StmtFunction{
		Position:     b.Pos.NewOptionalListTokensPosition(AttrGroups, FunctionTkn, CloseCurlyBracketTkn),
		AttrGroups:   AttrGroups,
		FunctionTkn:  FunctionTkn,
		AmpersandTkn: AmpersandTkn,
		Name: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(Name),
			IdentifierTkn: Name,
			Value:         Name.Value,
		},
		OpenParenthesisTkn:   OpenParenthesisTkn,
		Params:               params,
		SeparatorTkns:        sepTkns,
		CloseParenthesisTkn:  CloseParenthesisTkn,
		ColonTkn:             returnType.ColonTkn,
		ReturnType:           returnType.Type,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
	}
}

func (b *Builder) NewPropertyList(
	AttrGroups []ast.Vertex,
	Modifiers []ast.Vertex,
	Type ast.Vertex,
	PropsList ast.Vertex,
	SemiColonTkn *token.Token,
) *ast.StmtPropertyList {
	props, sepTkns := b.SeparatedListItems(PropsList)

	var pos *position2.Position

	if AttrGroups != nil {
		pos = b.Pos.NewNodeListTokenPosition(AttrGroups, SemiColonTkn)
	} else {
		pos = b.Pos.NewNodeListTokenPosition(Modifiers, SemiColonTkn)
	}

	return &ast.StmtPropertyList{
		Position:      pos,
		AttrGroups:    AttrGroups,
		Modifiers:     Modifiers,
		Type:          Type,
		Props:         props,
		SeparatorTkns: sepTkns,
		SemiColonTkn:  SemiColonTkn,
	}
}

func (b *Builder) NewClassConstList(
	AttrGroups []ast.Vertex,
	Modifiers []ast.Vertex,
	ConstTkn *token.Token,
	ConstList ast.Vertex,
	SemiColonTkn *token.Token,
) *ast.StmtClassConstList {
	consts, sepTkns := b.SeparatedListItems(ConstList)

	var pos *position2.Position

	if AttrGroups != nil {
		pos = b.Pos.NewNodeListTokenPosition(AttrGroups, SemiColonTkn)
	} else {
		pos = b.Pos.NewOptionalListTokensPosition(Modifiers, ConstTkn, SemiColonTkn)
	}

	return &ast.StmtClassConstList{
		Position:      pos,
		AttrGroups:    AttrGroups,
		Modifiers:     Modifiers,
		ConstTkn:      ConstTkn,
		Consts:        consts,
		SeparatorTkns: sepTkns,
		SemiColonTkn:  SemiColonTkn,
	}
}

func (b *Builder) NewClassMethod(
	AttrGroups []ast.Vertex,
	Modifiers []ast.Vertex,
	FunctionTkn *token.Token,
	AmpersandTkn *token.Token,
	Name *token.Token,

	OpenParenthesisTkn *token.Token,
	Params ast.Vertex,
	CloseParenthesisTkn *token.Token,

	RetType ast.Vertex,

	Stmt ast.Vertex,
) *ast.StmtClassMethod {
	params, sepTkns := b.SeparatedListItems(Params)
	returnType := RetType.(*ReturnType)

	var pos *position2.Position

	if AttrGroups != nil {
		pos = b.Pos.NewNodeListNodePosition(AttrGroups, Stmt)
	} else if Modifiers != nil {
		pos = b.Pos.NewNodeListNodePosition(Modifiers, Stmt)
	} else {
		pos = b.Pos.NewTokenNodePosition(FunctionTkn, Stmt)
	}

	return &ast.StmtClassMethod{
		Position:     pos,
		AttrGroups:   AttrGroups,
		Modifiers:    Modifiers,
		FunctionTkn:  FunctionTkn,
		AmpersandTkn: AmpersandTkn,
		Name: &ast.Identifier{
			Position:      b.Pos.NewTokenPosition(Name),
			IdentifierTkn: Name,
			Value:         Name.Value,
		},
		OpenParenthesisTkn:  OpenParenthesisTkn,
		Params:              params,
		SeparatorTkns:       sepTkns,
		CloseParenthesisTkn: CloseParenthesisTkn,
		ColonTkn:            returnType.ColonTkn,
		ReturnType:          returnType.Type,
		Stmt:                Stmt,
	}
}

func (b *Builder) NewNamespace(
	NsTkn *token.Token,
	Name ast.Vertex,
	OpenCurlyBracketTkn *token.Token,
	Stmts []ast.Vertex,
	CloseCurlyBracketTkn *token.Token,
	SemiColonTkn *token.Token,
) *ast.StmtNamespace {
	var pos *position2.Position

	if SemiColonTkn != nil {
		pos = b.Pos.NewTokensPosition(NsTkn, SemiColonTkn)
	} else {
		pos = b.Pos.NewTokensPosition(NsTkn, CloseCurlyBracketTkn)
	}

	return &ast.StmtNamespace{
		Position:             pos,
		NsTkn:                NsTkn,
		Name:                 Name,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Stmts:                Stmts,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
		SemiColonTkn:         SemiColonTkn,
	}
}

func (b *Builder) NewUseList(
	UseTkn *token.Token,
	Type *token.Token,
	UsesList ast.Vertex,
	SemiColonTkn *token.Token,
) *ast.StmtUseList {
	uses, sepTkns := b.SeparatedListItems(UsesList)

	return &ast.StmtUseList{
		Position:      b.Pos.NewTokensPosition(UseTkn, SemiColonTkn),
		UseTkn:        UseTkn,
		Type:          b.NewIdentifier(Type),
		Uses:          uses,
		SeparatorTkns: sepTkns,
		SemiColonTkn:  SemiColonTkn,
	}
}

func (b *Builder) NewGroupUseList(
	UseTkn *token.Token,
	Type *token.Token,

	Prefix ast.Vertex,

	NsSeparatorTkn *token.Token,

	OpenCurlyBracketTkn *token.Token,

	UsesList ast.Vertex,

	CloseCurlyBracketTkn *token.Token,
	SemiColonTkn *token.Token,
) *ast.StmtGroupUseList {
	uses, sepTkns := b.SeparatedListItems(UsesList)

	groupUse := &ast.StmtGroupUseList{
		Position:             b.Pos.NewTokensPosition(UseTkn, SemiColonTkn),
		UseTkn:               UseTkn,
		Type:                 b.NewIdentifier(Type),
		Prefix:               Prefix,
		NsSeparatorTkn:       NsSeparatorTkn,
		OpenCurlyBracketTkn:  OpenCurlyBracketTkn,
		Uses:                 uses,
		SeparatorTkns:        sepTkns,
		CloseCurlyBracketTkn: CloseCurlyBracketTkn,
		SemiColonTkn:         SemiColonTkn,
	}

	if n, ok := Prefix.(*ast.NameFullyQualified); ok {
		groupUse.LeadingNsSeparatorTkn = n.NsSeparatorTkn

		prefix := &ast.Name{
			Position:      n.Position,
			Parts:         n.Parts,
			SeparatorTkns: n.SeparatorTkns,
		}
		prefix.Position.StartPos++

		groupUse.Prefix = prefix
	}

	return groupUse
}

func (b *Builder) NewUse(
	Type ast.Vertex,

	UseName ast.Vertex,

	AsTkn *token.Token,
	Alias *token.Token,
) *ast.StmtUse {
	use := &ast.StmtUse{
		Type:  Type,
		Use:   UseName,
		AsTkn: AsTkn,
		Alias: b.NewIdentifier(Alias),
	}

	if Alias != nil {
		use.Position = b.Pos.NewNodeTokenPosition(UseName, Alias)
	} else {
		use.Position = b.Pos.NewNodePosition(UseName)
	}

	if n, ok := UseName.(*ast.NameFullyQualified); ok {
		use.NsSeparatorTkn = n.NsSeparatorTkn

		name := &ast.Name{
			Position:      n.Position,
			Parts:         n.Parts,
			SeparatorTkns: n.SeparatorTkns,
		}
		name.Position.StartPos++

		use.Use = name
	}

	return use
}

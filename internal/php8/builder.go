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

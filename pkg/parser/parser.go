package parser

import (
	"errors"

	"github.com/VKCOM/php-parser/internal/php5"
	"github.com/VKCOM/php-parser/internal/php7"
	"github.com/VKCOM/php-parser/internal/php8"
	"github.com/VKCOM/php-parser/internal/scanner"
	"github.com/VKCOM/php-parser/pkg/ast"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/version"
)

var (
	// ErrVersionOutOfRange is returned if the version is not supported
	ErrVersionOutOfRange = errors.New("the version is out of supported range")

	php5RangeStart = &version.Version{Major: 5}
	php5RangeEnd   = &version.Version{Major: 5, Minor: 6}

	php7RangeStart = &version.Version{Major: 7}
	php7RangeEnd   = &version.Version{Major: 7, Minor: 4}

	php8RangeStart = &version.Version{Major: 8}
	php8RangeEnd   = &version.Version{Major: 8, Minor: 1}
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() ast.Vertex
}

func Parse(src []byte, config conf.Config) (ast.Vertex, error) {
	var parser Parser

	if config.Version == nil {
		config.Version = php7RangeEnd
	}

	if config.Version.InRange(php5RangeStart, php5RangeEnd) {
		lexer := scanner.NewLexer(src, config)
		parser = php5.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	if config.Version.InRange(php7RangeStart, php7RangeEnd) {
		lexer := scanner.NewLexer(src, config)
		parser = php7.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	if config.Version.InRange(php8RangeStart, php8RangeEnd) {
		lexer := php8.NewLexer(src, config)
		parser = php8.NewParser(lexer, config)
		parser.Parse()
		return parser.GetRootNode(), nil
	}

	return nil, ErrVersionOutOfRange
}

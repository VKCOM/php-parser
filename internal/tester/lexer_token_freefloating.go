package tester

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/php8"
	"github.com/VKCOM/php-parser/internal/scanner"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/token"
	"github.com/VKCOM/php-parser/pkg/version"
	"gotest.tools/assert"
)

type Lexer interface {
	Lex() *token.Token
}

type LexerTokenFreeFloatingTestSuite struct {
	t *testing.T

	Code     string
	Expected [][]*token.Token

	Version version.Version
}

func NewLexerTokenFreeFloatingTestSuite(t *testing.T) *LexerTokenFreeFloatingTestSuite {
	return &LexerTokenFreeFloatingTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (l *LexerTokenFreeFloatingTestSuite) UsePHP8() {
	l.Version = version.Version{Major: 8, Minor: 0}
}

func (l *LexerTokenFreeFloatingTestSuite) Run() {
	config := conf.Config{
		Version: &l.Version,
	}

	var lexer Lexer

	if l.Version.Less(&version.Version{Major: 8, Minor: 0}) {
		lexer = scanner.NewLexer([]byte(l.Code), config)
	} else {
		lexer = php8.NewLexer([]byte(l.Code), config)
	}

	for _, expected := range l.Expected {
		tkn := lexer.Lex()
		actual := tkn.FreeFloating
		for _, v := range actual {
			v.Position = nil
		}
		assert.DeepEqual(l.t, expected, actual)
	}
}

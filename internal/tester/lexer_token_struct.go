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

type LexerTokenStructTestSuite struct {
	t *testing.T

	Code     string
	Expected []*token.Token

	Version version.Version

	withPosition     bool
	withFreeFloating bool
}

func NewLexerTokenStructTestSuite(t *testing.T) *LexerTokenStructTestSuite {
	return &LexerTokenStructTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (l *LexerTokenStructTestSuite) UsePHP8() {
	l.Version = version.Version{Major: 8, Minor: 0}
}

func (l *LexerTokenStructTestSuite) WithPosition() {
	l.withPosition = true
}

func (l *LexerTokenStructTestSuite) WithFreeFloating() {
	l.withFreeFloating = true
}

func (l *LexerTokenStructTestSuite) Run() {
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
		actual := lexer.Lex()
		if !l.withPosition {
			actual.Position = nil
		}
		if !l.withFreeFloating {
			actual.FreeFloating = nil
		}
		assert.DeepEqual(l.t, expected, actual)
	}
}

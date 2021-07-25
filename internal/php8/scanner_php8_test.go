package php8_test

import (
	"testing"

	"github.com/z7zmey/php-parser/internal/tester"
	"github.com/z7zmey/php-parser/pkg/token"
)

func TestNullsafeMethodCallTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php $a?->foo();"
	suite.Expected = []string{
		"$a",
		"?->",
		"foo",
		"(",
		")",
		";",
	}
	suite.Run()
}

func TestNullsafePropertyFetchTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php $a?->prop;"
	suite.Expected = []string{
		"$a",
		"?->",
		"prop",
		";",
	}
	suite.Run()
}

func TestNullsafePropertyFetchInStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStringTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php \"$a?->prop\";"
	suite.Expected = []string{
		"\"",
		"$a",
		"?->",
		"prop",
		"\"",
		";",
	}
	suite.Run()
}

func TestNullsafeMethodCallTokensFreeFloating(t *testing.T) {
	suite := tester.NewLexerTokenFreeFloatingTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
	$a ?-> bar ( '' ) ;`

	suite.Expected = [][]*token.Token{
		{
			{
				ID:    token.T_OPEN_TAG,
				Value: []byte("<?php"),
			},
			{
				ID:    token.T_WHITESPACE,
				Value: []byte("\n\t"),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
		{
			{
				ID:    token.T_WHITESPACE,
				Value: []byte(" "),
			},
		},
	}

	suite.Run()
}

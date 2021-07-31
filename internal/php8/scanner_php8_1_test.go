package php8_test

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/tester"
	"github.com/VKCOM/php-parser/pkg/token"
)

func TestReadonlyTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php readonly public $a;"
	suite.Expected = []*token.Token{
		{
			ID:    token.T_READONLY,
			Value: []byte("readonly"),
		},
		{
			ID:    token.T_PUBLIC,
			Value: []byte("public"),
		},
		{
			ID:    token.T_VARIABLE,
			Value: []byte("$a"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

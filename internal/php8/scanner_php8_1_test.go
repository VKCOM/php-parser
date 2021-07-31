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

func TestNumberTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
0x10;
0X10;
0b10;
0B10;
0o10;
0O10;
`
	suite.Expected = []*token.Token{
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0x10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0X10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0b10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0B10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0o10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
		{
			ID:    token.T_LNUMBER,
			Value: []byte("0O10"),
		},
		{
			ID:    ';',
			Value: []byte(";"),
		},
	}
	suite.Run()
}

func TestNumberStringTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = `<?php
"$a[0o10]"
"$a[0O10]"
`
	suite.Expected = []*token.Token{
		{
			ID:    '"',
			Value: []byte("\""),
		},
		{
			ID:    token.T_VARIABLE,
			Value: []byte("$a"),
		},
		{
			ID:    '[',
			Value: []byte("["),
		},
		{
			ID:    token.T_NUM_STRING,
			Value: []byte("0o10"),
		},
		{
			ID:    ']',
			Value: []byte("]"),
		},
		{
			ID:    '"',
			Value: []byte("\""),
		},
		{
			ID:    '"',
			Value: []byte("\""),
		},
		{
			ID:    token.T_VARIABLE,
			Value: []byte("$a"),
		},
		{
			ID:    '[',
			Value: []byte("["),
		},
		{
			ID:    token.T_NUM_STRING,
			Value: []byte("0O10"),
		},
		{
			ID:    ']',
			Value: []byte("]"),
		},
		{
			ID:    '"',
			Value: []byte("\""),
		},
	}
	suite.Run()
}

func TestEnumTokens(t *testing.T) {
	suite := tester.NewLexerTokenStructTestSuite(t)
	suite.UsePHP8()
	suite.Code = "<?php enum A {}"
	suite.Expected = []*token.Token{
		{
			ID:    token.T_ENUM,
			Value: []byte("enum"),
		},
		{
			ID:    token.T_STRING,
			Value: []byte("A"),
		},
		{
			ID:    '{',
			Value: []byte("{"),
		},
		{
			ID:    '}',
			Value: []byte("}"),
		},
	}
	suite.Run()
}

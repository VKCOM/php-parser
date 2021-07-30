package tester

import (
	"testing"

	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/errors"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"gotest.tools/assert"
)

type ParserErrorTestSuite struct {
	t *testing.T

	Code     string
	Expected []*errors.Error

	Version version.Version
}

func NewParserErrorTestSuite(t *testing.T) *ParserErrorTestSuite {
	return &ParserErrorTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (p *ParserErrorTestSuite) UsePHP8() {
	p.Version = version.Version{Major: 8, Minor: 0}
}

func (p *ParserErrorTestSuite) Run() {
	config := conf.Config{
		Version: &p.Version,
	}

	var errs []*errors.Error

	config.ErrorHandlerFunc = func(e *errors.Error) {
		errs = append(errs, e)
	}

	_, err := parser.Parse([]byte(p.Code), config)
	if err != nil {
		p.t.Fatalf("Error parse: %v", err)
	}
	assert.DeepEqual(p.t, p.Expected, errs)
}

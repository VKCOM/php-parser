package tester

import (
	"testing"

	"github.com/VKCOM/php-parser/pkg/ast"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"gotest.tools/assert"
)

type ParserTestSuite struct {
	t *testing.T

	Code     string
	Expected ast.Vertex

	Version version.Version
}

func NewParserTestSuite(t *testing.T) *ParserTestSuite {
	return &ParserTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (p *ParserTestSuite) UsePHP8() {
	p.Version = version.Version{Major: 8, Minor: 0}
}

func (p *ParserTestSuite) Run() {
	config := conf.Config{
		Version: &p.Version,
	}

	actual, err := parser.Parse([]byte(p.Code), config)
	if err != nil {
		p.t.Fatalf("Error parse: %v", err)
	}
	assert.DeepEqual(p.t, p.Expected, actual)
}

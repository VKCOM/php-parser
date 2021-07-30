package tester

import (
	"bytes"
	"testing"

	"github.com/VKCOM/php-parser/pkg/ast"
	"github.com/VKCOM/php-parser/pkg/conf"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"github.com/VKCOM/php-parser/pkg/visitor/printer"
	"gotest.tools/assert"
)

type ParserPrintTestSuite struct {
	t       *testing.T
	Version version.Version
}

func NewParserPrintTestSuite(t *testing.T) *ParserPrintTestSuite {
	return &ParserPrintTestSuite{
		t: t,
		Version: version.Version{
			Major: 7,
			Minor: 4,
		},
	}
}

func (p *ParserPrintTestSuite) UsePHP8() *ParserPrintTestSuite {
	p.Version = version.Version{Major: 8, Minor: 0}
	return p
}

func (p *ParserPrintTestSuite) Run(code string) {
	actual := p.print(p.parse(code))
	assert.DeepEqual(p.t, code, actual)
}

func (p *ParserPrintTestSuite) parse(src string) ast.Vertex {
	config := conf.Config{
		Version: &p.Version,
	}

	root, err := parser.Parse([]byte(src), config)
	if err != nil {
		p.t.Fatal(err)
	}

	return root
}

func (p *ParserPrintTestSuite) print(n ast.Vertex) string {
	o := bytes.NewBufferString("")

	pr := printer.NewPrinter(o)
	n.Accept(pr)

	return o.String()
}

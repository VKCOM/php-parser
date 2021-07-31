package printer_test

import (
	"testing"

	"github.com/VKCOM/php-parser/internal/tester"
)

func TestParseAndPrintReadonlyModifierPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
class Foo {
	readonly string $a;
	private readonly string $a;
	private string $a;
	private readonly $a = 100;

	public function __construct(
		readonly string $a,
		private readonly string $a,
		private string $a,
		private readonly $a = 100,
	) {}
}
`)
}

func TestNeverTypePHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
function f(): never {}
`)
}

func TestNumbersPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
echo 0x10;
echo 0X10;
echo 0b10;
echo 0B10;
echo 0o10;
echo 0O10;
`)
}

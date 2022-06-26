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

func TestEnumPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
enum A {
  case B;
  case B = 100;
  case C = "aa";
  case 
	D;
}

enum A: int {
  case B;
  case B = 100;

  #[Attribute]
  case C = 100;

  #[Attribute1]
  #[Attribute2]
  case D;
}

enum A implements B {
  case C;
  public function f(): string {}
}

enum A implements B, C, D {
  case E;
  public function f(): string {}
}
`)
}

func TestFinalConstantPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
class Foo {
    public const X = "foo";
    final public const X = "foo";
    final private const X = "foo";
}
`)
}

func TestFirstClassCallableSyntaxPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
$fn = id(...);
stdClass::doesNotExist(...);
(new stdClass)->doesNotExist(...);
Test::privateMethod(...);
Test::instanceMethod(...);
$fn = (new A)->$name2(...);
false && strlen(...);
$foo?->foo->bar(...);
$foo?->foo($foo->bar(...));
$foo = $closure->__invoke(...);

// #[Attribute(...)] // not working
// class Foo {}

// new Foo(...); // not working
`)
}

func TestIntersectionTypesSyntaxPHP81(t *testing.T) {
	tester.NewParserPrintTestSuite(t).UsePHP8().Run(`<?php
class Test {
    public A&B $prop;
}

function test(A&B $a): A&B {}
`)
}
